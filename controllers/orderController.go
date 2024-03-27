package controllers

import (
	"assignment_2_3/database"
	"assignment_2_3/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var (
		newOrder models.Order
	)

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Creating Order Data: %v", err),
		})
		return
	}

	newOrder.OrderedAt = time.Now()

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Creating Order Data: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"result": newOrder,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()

	var (
		orders []models.Order
	)

	if err := db.Model(&models.Order{}).Preload("Items").Order("order_id asc").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Getting Order Data: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var (
		orderId      string
		updatedOrder models.Order
		countItem    int
	)

	orderId = ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.Model(&models.Order{}).Where("order_id = ?", orderId).Updates(updatedOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Updating Order Data: %v", err.Error()),
		})
		return
	}

	countItem = 0
	for _, item := range updatedOrder.Items {
		if err := db.First(&models.Item{}, "order_id = ? and item_id = ?", orderId, item.ItemId).Error; err == nil {
			db.Model(&models.Item{}).Where("item_id = ?", item.ItemId).Updates(item)
			countItem += 1
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Updated (%v Items Updated)", orderId, countItem),
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	var (
		orderId string
	)

	orderId = ctx.Param("orderId")

	if err := db.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Order: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Deleted", orderId),
	})
}
