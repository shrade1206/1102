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

type MyError struct {
	Time   time.Time `json:"time"`
	Reason string    `json:"reason"`
}

var (
	newImg   []byte
	data     string
	t1       = time.Now().Year()
	t2       = time.Now().Month()
	t3       = time.Now().Day()
	t4       = time.Now().Hour()
	t5       = time.Now().Minute()
	timedata = time.Date(t1, t2, t3, t4, t5, 0, 0, time.Local)

	Err = MyError{
		Time:   timedata,
		Reason: "連線正常，但Cliend Data有誤",
	}

	ServerErr = MyError{
		Time:   timedata,
		Reason: "連線正常，但Server異常",
	}
)

func (e MyError) Error() string {
	return fmt.Sprintf("At %v, %s", e.Time, e.Reason)
}

func main() {
	err := InitMysql()
	if err != nil {
		log.Println("InitMysql() Run Error")
		return
	}
	defer DB.Close()
	r := gin.New()
	//---------------------------------------
	r.GET("/take", func(c *gin.Context) {
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			log.Printf("VideoCaptureDevice Error %s:", err.Error())
			c.JSON(http.StatusOK, Err)
			return
		}
		time.Sleep(time.Second)
		img := gocv.NewMat()
		defer img.Close()

		webcam.Read(&img)
		defer webcam.Close()
		buf, err := gocv.IMEncode(".jpg", img)
		if err != nil {
			log.Printf("Gocv Encode Error %s:", err.Error())
			c.JSON(http.StatusOK, ServerErr)
			return
		}
		defer buf.Close()
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
			c.JSON(http.StatusOK, Err)
		}
		p := client.Picture
		if p != nil {
			pic = Pic{Name: client.Name, Picture: []byte(client.Picture)}
			err = DB.Create(&pic).Error
			if err != nil {
				log.Printf("Data Error : %s", err.Error())
				c.JSON(http.StatusOK, ServerErr)
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"name": pic.Name,
				})
			}
		}
	})
	//----------------------------
	r.GET("/get", func(c *gin.Context) {
		var pics []Pic
		err := DB.Find(&pics).Error
		if err != nil {
			log.Printf("Find All Error : %s", err.Error())
			c.JSON(http.StatusOK, ServerErr)
			return
		} else {
			c.JSON(http.StatusOK, pics)
		}
	})
	//----------------------------
	// r.GET("/getpage", func(c *gin.Context) {
	// 	DDB := DB
	// 	var pages []Pic
	// 	page, err := strconv.Atoi(c.Query("page"))
	// 	if err != nil {
	// 		log.Printf("Page Error %s: ", err.Error())
	// 		return
	// 	}
	// 	pageSize := 4
	// 	if page > 0 && pageSize > 0 {
	// 		DDB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	// 	}
	// 	err = DDB.Find(&pages).Error
	// 	if err != nil {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"Get Error": err.Error(),
	// 		})
	// 		return
	// 	} else {
	// 		c.JSON(http.StatusOK, pages)
	// 	}
	// })
	//----------------------------
	r.POST("/getWho", func(c *gin.Context) {
		var pic []Pic
		var client Client
		err = c.BindJSON(&client)
		if err != nil {
			log.Printf("BindJSON %s: ", err.Error())
			c.JSON(http.StatusOK, Err)
			return
		}
		a := "%" + client.Name + "%"
		err := DB.Where("name like ?", a).Find(&pic).Error
		if err != nil {
			c.JSON(http.StatusOK, Err)
			return
		} else {
			c.JSON(http.StatusOK, pic)
		}
	})
	//---------------------------------------
	r.PUT("/put/:id", func(c *gin.Context) {
		var pic Pic
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, Err)
			return
		}
		err = DB.Where("id =?", id).First(&pic).Error
		if err != nil {
			c.JSON(http.StatusOK, Err)
			return
		}
		err = c.BindJSON(&pic)
		if err != nil {
			log.Printf("BindJson Error : %s", err.Error())
			c.JSON(http.StatusOK, Err)
			return
		}
		err = DB.Save(&pic).Error
		if err != nil {
			c.JSON(http.StatusOK, Err)
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
			c.JSON(http.StatusOK, Err)
			return
		}
		err := DB.Where("id = ?", id).Delete(Pic{}).Error
		if err != nil {
			c.JSON(http.StatusOK, Err)
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
			fmt.Println("Route ok")
		}
	})
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("8080 err : ", err.Error())
		return
	}
}
