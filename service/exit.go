package service

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Exit(c *gin.Context) {
	defer ExitFunc()
}

func ExitFunc() {
	os.Exit(0) //退出
}
