package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/models"
)

// 用户通过问卷id查询问卷的问题
func GetQuestionBySurveyId(c *gin.Context) {
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问卷id不能为空",
		})
		return
	}

	data, err := models.GetQuestionInfoBySIdentity(survey_id)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询问题失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询问题成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 用户新增问题
func AddQuestion(c *gin.Context) {
	var question models.QuestionInfo
	if err := c.ShouldBindJSON(&question); err != nil {
		log.Printf("[ERROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if question.Title == "" || question.Survey_Id == "" || fmt.Sprint(question.Options) == "" || fmt.Sprint(question.Type) == "" || fmt.Sprint(question.Flag) == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	question = models.QuestionInfo{
		Id:        helper.GetUUID(),
		Title:     question.Title,
		Survey_Id: question.Survey_Id,
		Options:   question.Options,
		Type:      question.Type,
		Flag:      question.Flag,
	}

	if err := models.AddQuestion(&question); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "新增问题失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "新增问题成功",
		"data":    question,
	})
}

// 用户修改问题
func UpdateQuestion(c *gin.Context) {
	var question models.QuestionInfo
	if err := c.ShouldBindJSON(&question); err != nil {
		log.Printf("[ERROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if question.Id == "" || question.Title == "" || fmt.Sprint(question.Options) == "" || fmt.Sprint(question.Type) == "" || fmt.Sprint(question.Flag) == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	if err := models.UpdateQuestionInfoById(&question); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "修改问题失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改问题成功",
		"_id":     question.Id,
		"title":   question.Title,
		"type":    question.Type,
		"options": question.Options,
		"flag":    question.Flag,
	})
}

// 用户删除问题
func DeleteQuestion(c *gin.Context) {
	_id := c.Param("_id")
	if _id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}
	if err := models.DeleteQuestionInfoById(_id); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "删除问题失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除问题成功",
	})
}

// 根据问题id查询问题选项
func GetQuestionOptionById(c *gin.Context) {
	_id := c.PostForm("_id")
	if _id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}

	data, err := models.GetQuestionOptionsById(_id)
	//将data存成一个数组
	var options []string
	for _, v := range data {
		options = append(options, v)
	}

	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询问题选项失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询问题选项成功",
		"data": gin.H{
			"list": options[1],
		},
	})
}

// 根据问题id统计答案内容的数量
func GetQuestionAnswerCount(c *gin.Context) {
	_id := c.PostForm("_id")
	if _id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}
	content, err := models.GetQuestionOptionsById(_id)
	//将data存成一个数组
	var options []string
	for _, v := range content {
		options = append(options, v)
	}
	////循环遍历数组，统计每个选项的数量
	//var count []int
	//for _, v := range options[1] {
	//	num, err := models.GetQuestionAnswerCount(_id, string(v))
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"code":    -1,
	//			"message": "查询问题选项失败" + err.Error(),
	//		})
	//		return
	//	}
	//	count = append(count, num)
	//}

	//data, err := models.GetAnswerOptionRelationsByQIdentityAndContent(_id, content)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询问题答案失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询问题答案成功",
		"data":    gin.H{
			//"count": data,
		},
	})
}

// 通过问卷id得到问题id数组
func GetQuestionIdByQId(c *gin.Context) {
	survey_id := c.PostForm("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问卷id不能为空",
		})
		return
	}
	data, err := models.GetQuestionIdBySIdentity(survey_id)
	var questionId []string
	for _, v := range data {
		questionId = append(questionId, v)
	}
	//循环在控制台打印问题id数组
	for _, v := range questionId {
		log.Println(v)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询问题失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询问题成功",
		"data": gin.H{
			"list": questionId,
		},
	})
}

// 通过问题id统计答案内容的数量
func GetAnswerCountByQId(c *gin.Context) {
	_id := c.PostForm("_id")
	if _id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}
	content, err := models.GetQuestionOptionsById(_id)
	//将data存成一个数组
	var options []string
	for _, v := range content {
		options = append(options, v)
	}
	data, err := models.GetAnswerOptionRelationsByQIdentityAndContents(_id, options)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询问题答案失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "查询问题各选项数量成功",
		"data": gin.H{
			"list": data,
		},
	})
}
