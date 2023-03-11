package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup\n___________________________________")
	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordconfirmaiton := ctx.PostForm("passwordconfirmaiton")

	println("username : "+username)
	println("emailaddress : "+email)
	println("password : "+password)
	println("passwordconfirmation : "+passwordconfirmaiton)
}

func UserLogin(ctx *gin.Context) {
	println("post/login\n________________________________")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	println("username : "+username)
	println("password : "+password)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}