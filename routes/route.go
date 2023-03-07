package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
//ルータから呼び出されるハンドラ
func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "sign_up.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code":"page not found", "message":"Page Not Found"})
}

