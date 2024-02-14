package main

import (

	"gocrud/controllers"
	"gocrud/initializers"

	"github.com/gin-gonic/gin"

)

func init (){
	initializers.LoadEnvVaribale()
	initializers.ConnectToDB()
}


func main(){
	r := gin.Default()
	r.POST("/posts", controllers.UsersCreate)
	r.GET("/posts", controllers.UserIndex)
	r.PUT("/posts/:id", controllers.UserUpdate)
	r.DELETE("/posts/:id", controllers.DeleteUsers)
	r.Run()


}