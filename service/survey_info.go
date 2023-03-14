package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/models"
	"time"
)

// 获取用户的所有问卷
func GetSurveyInfo(c *gin.Context) {
	creator_id := c.Param("creator_id")
	if creator_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoByCID(creator_id) //根据用户id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	data, err := models.GetSurveyInfoByCIdentity(creator_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 获取用户未开始问卷
func GetSurveyInfoW(c *gin.Context) {
	creator_id := c.Param("creator_id")
	if creator_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoByCID(creator_id) //根据用户id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	data, err := models.GetSurveyInfoByCIdentityW(creator_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 获取用户正在进行问卷
func GetSurveyInfoG(c *gin.Context) {
	creator_id := c.Param("creator_id")
	if creator_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoByCID(creator_id) //根据用户id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	data, err := models.GetSurveyInfoByCIdentitySWE(creator_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 获取用户已完成问卷
func GetSurveyInfoE(c *gin.Context) {
	creator_id := c.Param("creator_id")
	if creator_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoByCID(creator_id) //根据用户id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	data, err := models.GetSurveyInfoByCIdentityE(creator_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 获取用户禁用问卷
func GetSurveyInfo0(c *gin.Context) {
	creator_id := c.Param("creator_id")
	if creator_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoByCID(creator_id) //根据用户id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	data, err := models.GetSurveyInfoByCIdentityS(creator_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": gin.H{
			"list": data,
		},
	})
}

// 用户新增问卷
func AddSurveyInfo(c *gin.Context) {
	var survey models.SurveyInfo
	if err := c.ShouldBindJSON(&survey); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if survey.Name == "" || survey.Creator_Id == "" || survey.Survey_Description == "" || survey.Start_Time == "" || survey.End_Time == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}

	surveyInfo := &models.SurveyInfo{
		Id:                 helper.GetUUID(),
		Name:               survey.Name,
		Creator_Id:         survey.Creator_Id,
		Survey_Description: survey.Survey_Description,
		Start_Time:         survey.Start_Time,
		End_Time:           survey.End_Time,
		Status:             1,
		Create_Date:        time.Now().Format("2006-01-02 15:04:05"),
	}

	err := models.AddSurveyInfo(surveyInfo)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据添加异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":               http.StatusOK,
		"message":            "添加成功",
		"_id":                surveyInfo.Id,
		"name":               surveyInfo.Name,
		"start_time":         surveyInfo.Start_Time,
		"end_time":           surveyInfo.End_Time,
		"survey_description": surveyInfo.Survey_Description,
	})
}

// 用户修改问卷
func UpdateSurveyInfo(c *gin.Context) {
	token := c.GetHeader("token")
	var survey models.SurveyInfo
	if err := c.ShouldBindJSON(&survey); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	if survey.Id == "" || survey.Name == "" || survey.Survey_Description == "" || survey.Start_Time == "" || survey.End_Time == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	//根据问卷id获取该问卷的用户id
	creator_id, err := models.GetSurveyInfoByCid(survey.Id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据查询异常" + err.Error(),
		})
		return
	}
	//解析token以获取用户id
	claims, err := helper.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":     -1,
			"message:": "token解析失败",
		})
		return
	}
	//判断用户id是否一致
	if creator_id != claims {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户id不一致",
		})
		return
	}
	count, error := models.GetSurveyInfoBySID(survey.Id) //根据问卷id查询问卷
	if error != nil {
		log.Printf("[DB EROR] :%v\n", error)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "问卷不存在",
		})
		return
	}
	err = models.UpdateSurveyInfoById(&survey)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据修改异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":               http.StatusOK,
		"message":            "修改成功",
		"_id":                survey.Id,
		"name":               survey.Name,
		"start_time":         survey.Start_Time,
		"end_time":           survey.End_Time,
		"survey_description": survey.Survey_Description,
	})
}

// 通过问卷id修改问卷的状态为0
func UpdateSurveyStatusW(c *gin.Context) {
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoBySID(survey_id) //根据问卷id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "问卷不存在",
		})
		return
	}
	if err := models.UpdateSurveyInfoStatusById0(survey_id); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "禁用问卷失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "禁用问卷成功",
	})
}

// 通过问卷id修改问卷的状态为1
func UpdateSurveyStatus1(c *gin.Context) {
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "问题id不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoBySID(survey_id) //根据问卷id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "问卷不存在",
		})
		return
	}
	if err := models.UpdateSurveyInfoStatusById1(survey_id); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "启动问卷失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "启动问卷成功",
	})
}

// 通过问卷id修改问卷的状态为11
func UpdateSurveyStatusD(c *gin.Context) {
	var survey models.SurveyInfo
	if err := c.ShouldBindJSON(&survey); err != nil {
		log.Printf("[ERROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	if survey.Id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoBySID(survey.Id) //根据问卷id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "问卷不存在",
		})
		return
	}
	if err := models.UpdateSurveyInfoStatusByIdD(&survey); err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "删除问卷失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除问卷成功",
	})
}

// 通过问卷id查询问卷
func GetSurveyInfoById(c *gin.Context) {
	survey_id := c.Param("survey_id")
	if survey_id == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数不能为空",
		})
		return
	}
	count, err := models.GetSurveyInfoBySID(survey_id) //根据问卷id查询问卷
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "问卷不存在",
		})
		return
	}
	survey, err := models.GetSurveyInfoById(survey_id)
	if err != nil {
		log.Printf("[DB EROR] :%v\n", err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":               200,
		"message":            "查询成功",
		"_id":                survey.Id,
		"name":               survey.Name,
		"start_time":         survey.Start_Time,
		"end_time":           survey.End_Time,
		"status":             survey.Status,
		"survey_description": survey.Survey_Description,
	})
}

// 不断地执行这个函数，每隔一段时间就会执行一次
func UpdateSurveyStatus() func(c *gin.Context) {
	return func(c *gin.Context) {
		for {
			err := models.UpdateSurveyInfoStatus()
			if err != nil {
				log.Printf("[DB EROR] :%v\n", err)
				return
			}
			err1 := models.UpdateSurveyInfoStatus1()
			if err1 != nil {
				log.Printf("[DB EROR] :%v\n", err1)
				return
			}
			time.Sleep(time.Second * 10)
		}
	}
}
