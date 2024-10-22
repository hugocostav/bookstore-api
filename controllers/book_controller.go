package controllers

import (
	"bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func (bc *BookController) GetBooks(c *gin.Context) {
	var books []models.Book
	bc.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bc.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func (bc *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bc.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bc.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bc.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bc.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	bc.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Livro excluído com sucesso"})
}