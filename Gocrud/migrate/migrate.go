package main

import ("gocrud/initializers"
		"gocrud/models"
)

func init(){
	initializers.LoadEnvVaribale()
	initializers.ConnectToDB()
}


func main(){
	initializers.DB.AutoMigrate(&models.Users{})

}