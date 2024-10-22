package routes

import (
	"bookstore-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	bookController := &controllers.BookController{DB: db}
	authorController := &controllers.AuthorController{DB: db}
	publisherController := &controllers.PublisherController{DB: db}

	api := r.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("", bookController.GetBooks)
			books.POST("", bookController.CreateBook)
			books.GET("/:id", bookController.GetBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}

		authors := api.Group("/authors")
		{
			authors.GET("", authorController.GetAuthors)
			authors.POST("", authorController.CreateAuthor)
			authors.GET("/:id", authorController.GetAuthor)
			authors.PUT("/:id", authorController.UpdateAuthor)
			authors.DELETE("/:id", authorController.DeleteAuthor)
		}

		publishers := api.Group("/publishers")
		{
			publishers.GET("", publisherController.GetPublishers)
			publishers.POST("", publisherController.CreatePublisher)
			publishers.GET("/:id", publisherController.GetPublisher)
			publishers.PUT("/:id", publisherController.UpdatePublisher)
			publishers.DELETE("/:id", publisherController.DeletePublisher)
		}
	}
}