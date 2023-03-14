package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/models"
	"time"
)

// 用户答题
func AnswerSurvey(c *gin.Context) {
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
	if err := c.ShouldBindJSON(&answerSurvey); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if answerSurvey.Survey_Id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	answerSurvey = models.AnswerInfo{
		Id:         helper.GetUUID(),
		Survey_Id:  answerSurvey.Survey_Id,
		UserIp:     ip,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
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

// 依据问卷id所有答卷
func GetAnswerSurvey(c *gin.Context) {
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问卷id不能为空",
		})
		return
	}

	data, err := models.GetAnswerInfoBySIdentity(survey_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}
