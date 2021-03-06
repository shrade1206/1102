package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

type Pic struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture []byte `json:"picture"`
}

type Client struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture []byte `json:"picture"`
}

// type Error struct {
// 	err string
// }

var (
	newImg []byte
	data   string
)

func main() {
	err := InitMysql()
	if err != nil {
		log.Printf("initMysql() invalid : %s", err.Error())
		return
	}
	defer DB.Close()
	r := gin.New()
	//---------------------------------------
	r.GET("/take", func(c *gin.Context) {
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
		img := gocv.NewMat()
		defer img.Close()

		webcam.Read(&img)
		defer webcam.Close()
		buf, err := gocv.IMEncode(".jpg", img)
		if err != nil {
			log.Fatal(err)
		}
		defer buf.Close() //nolint
		newImg = buf.GetBytes()
		data = base64.StdEncoding.EncodeToString(newImg)
		c.JSON(http.StatusOK, data)
	})
	//---------------------------------------
	r.POST("/save", func(c *gin.Context) {
		var client Client
		var pic Pic
		err := c.BindJSON(&client)
		if err != nil {
			log.Printf("BindJson Error : %s", err.Error())
			return
		}
		// err = DB.Create(&pic).Error
		// if err != nil {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"Create Error": err.Error()})
		// 	return
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"name": pic.Name,
		// 	})
		// }

		pic = Pic{Name: client.Name, Picture: []byte(client.Picture)}
		err = DB.Debug().Create(&pic).Error
		if err != nil {
			log.Printf("Data Error : %s", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": pic.Name,
			})
		}

	})
	//----------------------------
	r.GET("/get", func(c *gin.Context) {
		var pics []Pic
		err := DB.Find(&pics).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Get Error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, pics)
		}
	})
	//----------------------------
	r.GET("/get", func(c *gin.Context) {
		var pics []Pic
		var client Client
		err = c.BindJSON(&client)
		if err != nil {
			log.Printf("BindJSON %s: ", err.Error())
			return
		}
		err := DB.Where("name like ?", "%"+client.Name+"%").Find(&pics).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Get Error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, pics)
		}
	})
	//---------------------------------------
	r.PUT("/put/:id", func(c *gin.Context) {
		var pic Pic
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"Put Error": "ID??????",
			})
			return
		}
		err = DB.Where("id =?", id).First(&pic).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Put Error": err.Error(),
			})
			return
		}
		err = c.BindJSON(&pic)
		if err != nil {
			log.Printf("BindJson Error : %s", err.Error())
			return
		}
		err = DB.Save(&pic).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Save Error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"id":      pic.Id,
				"name":    pic.Name,
				"picture": pic.Picture,
			})
		}
	})
	//---------------------------------------
	r.DELETE("/del/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"Del Error": "ID??????",
			})
			return
		}
		err := DB.Where("id = ?", id).Delete(Pic{}).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Del Error": err.Error(),
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
		//??????path?????????????????????"/"
		if strings.HasPrefix(path, "/") {
			fmt.Println("Route ok")
		}
	})
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("8080 err : ", err.Error())
	}
}
