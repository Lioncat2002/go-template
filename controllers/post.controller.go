package controllers

import (
	"backend/go-catbox"
	"backend/models"
	"backend/services"
	"backend/utils/token"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostdata struct {
	Name string `json:"name" binding:"required"`
	Text string `json:"desc" binding:"required"`
}

type PostData struct {
	PostID string `json:"post_id" binding:"required"`
}

type RatePostData struct {
	Rate   float32 `json:"rate" binding:"required"`
	PostID string  `json:"post_id" binding:"required"`
}

func SearchPost(c *gin.Context) {
	name := c.Param("name")
	post := models.Post{}
	if err := services.DB.Where("name = ?", name).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   post,
	})
}

func AddPostUrl(c *gin.Context) {
	f, _ := c.FormFile("file")
	file, _ := f.Open()
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	if err := services.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	url, err := catbox.New(nil).Upload(buf.Bytes(), string(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id = c.Param("id")
	post := models.Post{}
	if err := services.DB.Where("id = ?", id).Find(&post).Update("url", url).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   post,
	})
}

func CreatePost(c *gin.Context) {
	var postData CreatePostdata
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	if err := services.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := models.Post{}
	post.UserID = id
	post.Text = postData.Text
	if err := services.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
	})
}

func GetOnePost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	if err := services.DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   post,
	})
}

func AllPosts(c *gin.Context) {
	var posts []models.Post
	if err := services.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   posts,
	})
}
