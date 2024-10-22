package controllers

import (
	"bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PublisherController struct {
	DB *gorm.DB
}

func (pc *PublisherController) GetPublishers(c *gin.Context) {
	var publishers []models.Publisher
	pc.DB.Find(&publishers)
	c.JSON(http.StatusOK, publishers)
}

func (pc *PublisherController) CreatePublisher(c *gin.Context) {
	var publisher models.Publisher
	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Create(&publisher)
	c.JSON(http.StatusCreated, publisher)
}

func (pc *PublisherController) GetPublisher(c *gin.Context) {
	id := c.Param("id")
	var publisher models.Publisher

	if err := pc.DB.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Editora não encontrada"})
		return
	}

	c.JSON(http.StatusOK, publisher)
}

func (pc *PublisherController) UpdatePublisher(c *gin.Context) {
	id := c.Param("id")
	var publisher models.Publisher

	if err := pc.DB.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Editora não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&publisher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Save(&publisher)
	c.JSON(http.StatusOK, publisher)
}

func (pc *PublisherController) DeletePublisher(c *gin.Context) {
	id := c.Param("id")
	var publisher models.Publisher

	if err := pc.DB.First(&publisher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Editora não encontrada"})
		return
	}

	pc.DB.Delete(&publisher)
	c.JSON(http.StatusOK, gin.H{"message": "Editora excluída com sucesso"})
}