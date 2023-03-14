package main

import (
	_ "golang.org/x/crypto/bcrypt"
	"math/rand"
	"questionnaire-system-backend/router"
	"time"
)

func main() {
	//r := gin.Default()
	//r.Run() // listen and serve on 0.0.0.0:8080
	e := router.Router()
	//初始化UUID
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	e.Run(":8080")
}
