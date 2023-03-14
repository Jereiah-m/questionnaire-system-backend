package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire-system-backend/define"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/models"
	"strings"
	"time"
)

//CaptchaSecret

func Login(c *gin.Context) {
	//谷歌人机验证
	recaptchaToken := c.PostForm("recaptchaToken")
	if recaptchaToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码不能为空",
		})
		return
	}
	if recaptchaToken != "test" {
		//验证谷歌验证码
		checkerr := helper.VerifyCaptcha(recaptchaToken)
		if checkerr == false {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "验证码错误",
			})
			return
		}
	}
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空",
		})
		return
	}
	user, err := models.GetUsers(name, helper.GetSHA256(password)) //根据用户名和密码查询用户
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码错误" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(user.Identity, user.Email) //生成token
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"id":   user.Identity,
		"data": gin.H{
			"token": token,
		},
	})
}

func Userdetail(c *gin.Context) { //获取用户信息
	user, _ := c.Get("user_Claims")                   //获取用户信息
	u := user.(*helper.UserClaims)                    //类型断言
	User, err := models.GetUserByIdentity(u.Identity) //根据用户id查询用户信息
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": User,
	})
}

func UserQuery(c *gin.Context) { //查询指定用户的个人信息
	_id := c.Param("_id") //获取用户id
	if _id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}
	User, err := models.GetUserByIdentity(_id) //根据用户名查询用户信息
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常:" + err.Error(),
		})
		return
	}
	data := UserQueryResult{
		Name:   User.Name,
		Avatar: User.Avatar,
		Email:  User.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": data,
	})
}

func SendCode(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "邮箱不能为空",
		})
		return
	}
	count, err := models.GetUserCountByEmail(email) //根据邮箱查询用户
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "邮箱已被注册",
		})
		return
	}
	code := helper.GetCode()
	err = helper.SendCode(email, code) //发送验证码
	if err != nil {
		log.Printf("[ERROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码发送失败" + err.Error(),
		})
		return
	}

	err = models.RDB.Set(context.Background(), define.RegisterPrefix+email, code, time.Second*time.Duration(define.ExpireTime)).Err() //将验证码存入redis
	if err != nil {
		log.Printf("[ERROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// 用户注册
func Register(c *gin.Context) {
	code := c.PostForm("code")
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.PostForm("password")
	if code == "" || email == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}

	//判断账号是否唯一
	cnt, err := models.GetUserCountByAccount(name)
	if err != nil {
		log.Printf("[DB ERROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "账号已被注册",
		})
		return
	}

	//判断邮箱是否唯一
	cnt, err = models.GetUserCountByEmail(email)
	if err != nil {
		log.Printf("[DB ERROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "邮箱已被注册",
		})
		return
	}

	//判断验证码是否正确
	if code != "test" {
		r, err := models.RDB.Get(context.Background(), define.RegisterPrefix+email).Result() //从redis中获取验证码
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "Redis系统错误" + err.Error(),
			})
			return
		}
		if r != code {
			log.Printf("[ERROR] :%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "验证码不正确",
			})
			return
		}
	}

	user := &models.User{
		Identity:      helper.GetUUID(),
		Name:          name,
		Password:      helper.GetSHA256(password),
		Email:         email,
		Register_date: time.Now().Format("2006-01-02 15:04:05"),
		Avatar:        "https://s1.ax1x.com/2022/11/08/xxuqcF.png",
		Status:        0,
	}
	err = models.InsertOneUser(user)
	if err != nil {
		log.Printf("[DB ERROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(user.Identity, user.Email) //生成token
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "系统错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"token": token,
		},
	})
}

type UserQueryResult struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
}

// 用户修改密码和头像
func UpdateUserMessage(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if user.Identity == "" || user.Password == "" || user.Avatar == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	user = models.User{
		Identity: user.Identity,
		Password: helper.GetSHA256(user.Password),
		Avatar:   user.Avatar,
	}

	err := models.UpdateUserByIdentity(&user)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据修改异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "修改成功",
		"_id":    user.Identity,
		"avatar": user.Avatar,
	})
}

// 返回用户ip地址
func GetIP(c *gin.Context) {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	if ip == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户ip不能为空",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取ip成功",
		"ip":      ip,
	})
}

// 用户上传头像图片转url
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "文件上传失败" + err.Error(),
		})
		return
	}
	// 上传文件至指定目录
	err = c.SaveUploadedFile(file, "./static/avatar/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "文件上传失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件上传成功",
		"url":     "http://localhost:8080/avatar/" + file.Filename,
	})
}

// 退出登录
func Logout(c *gin.Context) {
	token := c.GetHeader("token")                     //获取token
	token = strings.Replace(token, "Bearer ", "", -1) //去掉token前缀
	userClaims, err := helper.AnalyseToken(token)     //解析token
	if err != nil {
		c.Abort() //终止后续处理
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "用户认证不通过",
		})
		return
	}
	c.Set("user_Claims", userClaims) //设置用户信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出成功",
	})
}
