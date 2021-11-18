package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todos []Todo

func main() {
	err := InitMysql()
	if err != nil {
		log.Printf("initMysql() invalid : %s", err.Error())
		return
	}
	defer DB.Close()
	r := gin.Default()
	//---------------------------------------
	r.POST("/insert", func(c *gin.Context) {
		var todo Todo
		err := c.BindJSON(&todo)
		if err != nil {
			log.Printf("BindJson Error : %s", err.Error())
			return
		}
		err = DB.Create(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Create Error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Title": todo.Title,
			})
		}
	})
	//----------------------------
	r.GET("/get", func(c *gin.Context) {

		err := DB.Find(&todos).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Find Error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, todos)
		}
	})
	//---------------------------------------
	r.GET("/getpage", func(c *gin.Context) {
		DDB := DB
		var pages []Todo
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Printf("Page Error %s: ", err.Error())
			return
		}
		pageSize := 4
		if page > 0 && pageSize > 0 {
			DDB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
		}
		err = DDB.Find(&pages).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, pages)
		}
	})
	//---------------------------------------
	r.PUT("/put/:id", func(c *gin.Context) {
		var todo Todo
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"Put Error": "ID無效",
			})
			return
		}
		err = DB.Where("id =?", id).First(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Put Error": err.Error(),
			})
			return
		}
		err = c.BindJSON(&todo)
		if err != nil {
			log.Printf("BindJson Error : %s", err.Error())
			return
		}
		err = DB.Save(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Save Error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"id":     todo.Id,
				"title":  todo.Title,
				"status": todo.Status,
			})
		}
	})
	//---------------------------------------
	r.DELETE("/del/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		log.Printf("id: %v\n", id)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"err": nil,
			})
			return
		}
		err := DB.Where("id = ?", id).Delete(Todo{}).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				id: "Deleted ok",
			})
		}
	})
	//---------------------------------------
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./public"))), func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		//檢查path的開頭使是否為"/"
		if strings.HasPrefix(path, "/") {
			fmt.Println("ok")
		}
	})
	err = r.Run(":8082")
	if err != nil {
		log.Fatal("8082 err : ", err.Error())
	}
}
