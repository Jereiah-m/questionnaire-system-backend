package router

import (
	"github.com/gin-gonic/gin"
	"questionnaire-system-backend/helper"
	"questionnaire-system-backend/middlewares"
	"questionnaire-system-backend/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(helper.Cors())

	go func() {
		//r.Use(service.UpdateSurveyStatus())                                        //定时任务
		r.POST("/api/register", service.Register)                                 //注册
		r.POST("/api/login", service.Login)                                       //登录
		r.GET("/api/user/QuestionInfo/:survey_id", service.GetQuestionBySurveyId) //依据问卷id获取问卷的问题
		r.GET("/api/user/SurveyInfoT/:survey_id", service.GetSurveyInfoById)      //依据问卷id查询问卷
		auth := r.Group("/api/user", middlewares.AuthCheck())                     //登录验证//跨域
		auth.GET("/query/:_id", service.UserQuery)                                //获取用户个人信息
		auth.GET("/SurveyInfo/:creator_id", service.GetSurveyInfo)                //依据用户id获取用户的调查问卷
		auth.GET("/SurveyInfoW/:creator_id", service.GetSurveyInfoW)              //依据用户id获取用户的未开始调查问卷
		auth.GET("/SurveyInfoG/:creator_id", service.GetSurveyInfoG)              //依据用户id获取用户的正在进行调查问卷
		auth.GET("/SurveyInfoE/:creator_id", service.GetSurveyInfoE)              //依据用户id获取用户的已完结调查问卷
		auth.GET("/SurveyInfo0/:creator_id", service.GetSurveyInfo0)              //依据用户id获取用户的禁用问卷
		auth.PUT("/updateMessage", service.UpdateUserMessage)                     //更新用户个人信息
		auth.POST("/addSurveyInfo", service.AddSurveyInfo)                        //添加问卷
		auth.PUT("/updateSurveyInfo", service.UpdateSurveyInfo)                   //修改问卷

		auth.GET("/AnswerInfo/:survey_id", service.GetAnswerSurvey)                //依据问卷id获取问卷的所有答卷
		auth.GET("/AnswerSurvey/:answer_id", service.GetAnswerOptionRelationsList) //依据答卷id获取答卷的所有问题和选项
		auth.PUT("/updateSurveyStatus0/:survey_id", service.UpdateSurveyStatusW)   //修改问卷状态为0
		auth.PUT("/updateSurveyStatusD", service.UpdateSurveyStatusD)              //修改问卷状态为11
		auth.PUT("/updateSurveyStatus1/:survey_id", service.UpdateSurveyStatus1)   //修改问卷状态为1

		auth.GET("/QuestionOption", service.GetQuestionOptionById) //依据问题id查询问题
		auth.GET("Count", service.GetQuestionAnswerCount)
		auth.GET("GetQID", service.GetQuestionIdByQId)
		auth.POST("GetAnswerCountByQId", service.GetAnswerCountByQId) //依据问题id查询答案数量

		auth.POST("/addQuestion", service.AddQuestion)              //添加问题
		auth.PUT("/updateQuestion", service.UpdateQuestion)         //修改问题
		auth.DELETE("/deleteQuestion/:_id", service.DeleteQuestion) //删除问题

		auth.POST("/loginOut", service.Logout) //退出登录

		r.POST("/api/answer/:survey_id", service.AddAnswerOptionRelationsList) //提交问卷答案
		r.POST("/api/returnIp", service.GetIP)                                 //获取IP
		r.POST("/api/answerSurvey", service.AnswerSurvey)                      //点击答题
		r.POST("/api/exit", service.Exit)                                      //退出
		r.GET("/api/user/detail", service.Userdetail)                          //获取用户详情信息
		r.POST("/api/send/code/:email", service.SendCode)                      //发送验证码
	}()
	return r
}
