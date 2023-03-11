package main

import (
	"github.com/gin-gonic/gin"
	"pra_go/routes"
)




func main()	 {
	router := gin.Default()
    router.LoadHTMLGlob("./files/templates/*.html")
    router.Static("/static", "./files/static")

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogin)
	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.Login)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)
	var port string = ":8080"//ポート番号の宣言
	println("\n your localhost => http://localhost"+port+"\n")

    router.Run(port)
}













/*
func sum(srcs ...int) int {
	//可変長引数はようそがスライスで返される
	num := 0
	for _, src := range srcs {
		num += src
	}
	return num
}
*/
