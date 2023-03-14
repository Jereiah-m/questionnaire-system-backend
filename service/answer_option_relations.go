package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/models"
	"time"
)

// 用户针对问题id答题一组一组分别存储
func AddAnswerOptionRelationsList(c *gin.Context) {
	//谷歌人机验证
	recaptchaToken := c.GetHeader("recaptchaToken")
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
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	if ip == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户ip不能为空",
		})
		return
	}

	var answerSurvey models.AnswerInfo
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	//通过问卷id查询问卷状态是否为0
	surveyInfo, err := models.GetSurveyInfoById(survey_id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询失败" + err.Error(),
		})
		return
	}
	if surveyInfo.Status == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问卷已结束",
		})
		return
	}
	answerSurvey = models.AnswerInfo{
		Id:         helper.GetUUID(),
		Survey_Id:  survey_id,
		UserIp:     ip,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	var answerOptionRelationsList []models.AnswerOptionRelations // 一组一组的答案
	if err := c.ShouldBindJSON(&answerOptionRelationsList); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	for _, answerOptionRelations := range answerOptionRelationsList { // 遍历一组一组的答案
		if answerOptionRelations.Question_Id == "" || fmt.Sprint(answerOptionRelations.Content) == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "参数不能为空",
			})
			return
		}
		answerOptionRelations = models.AnswerOptionRelations{
			Id:          helper.GetUUID(),
			Answer_Id:   answerSurvey.Id,
			Question_Id: answerOptionRelations.Question_Id,
			Content:     answerOptionRelations.Content,
		}
		err := models.AddAnswerOptionRelations(&answerOptionRelations) // 一组一组的存储
		if err != nil {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": "答题失败" + err.Error(),
			})
			return
		}
	}
	if err := models.AddAnswerSurvey(&answerSurvey); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "新增答卷失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "新增答卷成功",
		"data":    answerSurvey,
	})
}

// 依据答卷id查询答题详情
func GetAnswerOptionRelationsList(c *gin.Context) {
	answer_id := c.Param("answer_id")
	if answer_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	answerOptionRelationsList, err := models.GetAnswerOptionRelationsByAIdentity(answer_id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询成功",
		"data": gin.H{
			"list": answerOptionRelationsList,
		},
	})
}
