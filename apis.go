package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserGet(c *gin.Context){
	userId := c.Param("id")
	c.String(http.StatusOK, "Hello %s", userId)
}

func UserRegister(c *gin.Context){

}

func UserLogin(c *gin.Context){

}

func UrlGetList(c *gin.Context){

}

func UrlAdd(c *gin.Context){

}

func UrlRemove(c *gin.Context){

}

func UrlGetDetail(c *gin.Context){

}