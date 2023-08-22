package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Itemdata struct {
	AuthorID    uint    `json:"author_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"desc"`
	Price       float32 `json:"price" binding:"required"`
	Tag         string  `json:"tag" binding:"required"`
}

type BuyItemData struct {
	UserID uint `json:"user_id" binding:"required"`
	ItemID uint `json:"item_id" binding:"required"`
}

func BuyItem(c *gin.Context) {
	var buyItemData BuyItemData
	if err := c.ShouldBindJSON(&buyItemData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{}
	if err := services.DB.Where("id = ?", buyItemData.UserID).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	item := models.Item{}
	if err := services.DB.Where("id = ?", buyItemData.ItemID).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	coins := user.Coins - item.Price
	if coins < 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "insufficient balance",
		})
		return
	}
	if err := services.DB.Where("id = ?", buyItemData.UserID).Find(&user).Update("coins", coins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := services.DB.Where("id = ?", buyItemData.ItemID).Find(&item).Update("current_owner_id", buyItemData.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   item,
	})
}

func CreateItem(c *gin.Context) {
	var itemData Itemdata
	if err := c.ShouldBindJSON(&itemData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	item := models.Item{}
	item.UserID = itemData.AuthorID
	item.Name = itemData.Name
	item.Description = itemData.Description
	item.Tag = itemData.Tag
	item.CurrentOwnerID = itemData.AuthorID
	item.Price = itemData.Price
	if err := services.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	//services.DB.Debug().Model(&models.User{}).Related()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   item,
	})
}

func GetOneItem(c *gin.Context) {
	id := c.Param("id")
	item := models.Item{}
	if err := services.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   item,
	})
}

func AllItems(c *gin.Context) {
	var items []models.Item
	if err := services.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   items,
	})
}
