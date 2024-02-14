package controllers

import (
	"gocrud/initializers"
	"gocrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UsersCreate(c *gin.Context){
	// get data of request

	var body struct {
		Name string 
		Email string 
	}

	c.Bind(&body)

	//create a post 

	post := models.Users{Name:body.Name,Email:body.Email}
	result := initializers.DB.Create(&post)
	if result.Error != nil{
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"post":post,
	})
}


func UserIndex(c *gin.Context){
		//get the post 
		var posts []models.Users
		initializers.DB.Find(&posts)

		//Respond with them 
		c.JSON(http.StatusOK, gin.H{
			"post":posts,
		})


}

func UserUpdate(c*gin.Context){
	//get the id off url 
	id :=c.Param("id")

	//get the data req body 

	var body struct {
		Name string 
		Email string
	}
	c.Bind(&body)

	//find the post were updating 

	var post models.Users
	initializers.DB.First(&post,id)

	//update it 
	initializers.DB.Model(&post).Updates(models.Users{Name:body.Name,Email:body.Email})

	//respond with it 
	c.JSON(http.StatusOK,gin.H{
		"post":post,
	})
}

func DeleteUsers(c *gin.Context){
	//get the post 
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Users{},id)

	//respond with them 
	c.Status(200)
}

