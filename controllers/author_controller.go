package controllers

import (
	"bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthorController struct {
	DB *gorm.DB
}

func (ac *AuthorController) GetAuthors(c *gin.Context) {
	var authors []models.Author
	ac.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func (ac *AuthorController) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ac.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

func (ac *AuthorController) GetAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := ac.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Autor não encontrado"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (ac *AuthorController) UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := ac.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Autor não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ac.DB.Save(&author)
	c.JSON(http.StatusOK, author)
}

func (ac *AuthorController) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := ac.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Autor não encontrado"})
		return
	}

	ac.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Autor excluído com sucesso"})
}