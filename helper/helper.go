package helper

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"questionnaire-system-backend/define"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GetSha256
// 生成SHA256
func GetSHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

var myKey = []byte("MHQ")

// GenerateToken
// 生成 token
func GenerateToken(identity string, email string) (string, error) { // 生成token
	UserClaim := &UserClaims{
		Identity:       identity,             // 这里的 user.ID 是 uint 类型，不是 int
		Email:          email,                // 标准声明
		StandardClaims: jwt.StandardClaims{}, // 标准声明
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	//设置token过期时间
	UserClaim.ExpiresAt = GetExpireTime()
	//token过期后重新发送token
	UserClaim.NotBefore = time.Now().Unix()
	//每次请求都会刷新token过期时间
	UserClaim.IssuedAt = time.Now().Unix()
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// token过期时间
func GetExpireTime() int64 {
	return time.Now().Add(time.Hour * 1).Unix()
}

// token过期后重新发送token
func RefreshToken(tokenString string) (string, error) {
	userClaim, err := AnalyseToken(tokenString)
	if err != nil {
		return "", err
	}
	tokenString, err = GenerateToken(userClaim.Identity, userClaim.Email)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// 解析token得到原来的数据
func ParseToken(tokenString string) (string, error) {
	claims := new(UserClaims)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return "", err
	}
	return claims.Identity, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <925714253@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "925714253@qq.com", define.MailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

// 生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// 生成唯一码
func GetUUID() string {
	uuid := make([]byte, 16)
	rand.Read(uuid)
	return fmt.Sprintf("%x", uuid)
}

// 使用Google reCAPTCHA进行人机验证
func VerifyCaptcha(captcha string) bool {
	//请求Google reCAPTCHA API
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{"secret": {define.CaptchaSecret}, "response": {captcha}})
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	//解析响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false
	}
	//返回验证结果
	return result["success"].(bool)
}

// 获取用户ip
func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-Ip")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func JsonToStruct(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
