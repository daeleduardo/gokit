package controllers

import (
	"context"
	"gokit/config"
	"gokit/models"
	db "gokit/repository"
	"gokit/services"
	"net/http"

	"gokit/lib"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
type Payload struct {
	Id   int    `form:"id" json:"id" binding:"omitempty"`
	Name string `form:"name" json:"name" binding:"omitempty"`
}

func Default(c *gin.Context) {

	JWT, err := services.GenerateJWT("email@string.com", "username")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "Error to generate token"})
		return
	}

	//Insert the token on cache just for example
	db.RedisConn().SetEX(context.Background(), lib.GetMd5(JWT), JWT, time.Duration(config.GetConfig().JwtTimeout))
	Data, err := db.ReadAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Erro to connect"})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
		"token": JWT,
		"users": Data,
	})
}

func GetToken(c *gin.Context) {

	payload := &Payload{}
	if err := c.ShouldBind(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Invalid Payload"})
		return
	}

	if db.Create(models.User{Id: payload.Id, Name: payload.Name}).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Error on Create Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "OK"})
}

func Create(c *gin.Context) {

	payload := &Payload{}
	if err := c.ShouldBind(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Invalid Payload"})
		return
	}

	if db.Create(models.User{Id: payload.Id, Name: payload.Name}).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Error on Create"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "OK"})
}

// func ReadAll(c *gin.Context) {

// 	result, err := db.ReadAll()

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"response": "NOK"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"response": result})

// }

// func ReadOne(c *gin.Context) {

// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"response": "Erro to get Id"})
// 		return
// 	}
// 	result, err := db.ReadOne(models.User{Id: id})

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"response": "Erro to get result by Id"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"response": result})

// }

func Update(c *gin.Context) {

	payload := &Payload{}
	if err := c.ShouldBind(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Invalid Payload"})
		return
	}

	if db.Update(models.User{Id: payload.Id, Name: payload.Name}).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Error on Update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "OK"})
}

func Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Invalid Parameter"})
		return
	}

	if db.Delete(models.User{Id: id}).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "Error on Delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "OK"})

}
