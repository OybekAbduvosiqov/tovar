package api

import (
	"github.com/OybekAbduvosiqov/tovar/storage"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "github.com/OybekAbduvosiqov/tovar/api/docs"
	"github.com/OybekAbduvosiqov/tovar/api/handler"
)

func NewApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(storage)

	r.POST("/tovar", handlerV1.CreateTovar)
	r.GET("/tovar/:id", handlerV1.GetByIDTovar)
	r.GET("/tovar", handlerV1.GetListTovar)
	r.PUT("/tovar/:id", handlerV1.UpdateTovar)
	r.DELETE("/tovar/:id", handlerV1.DeleteTovar)

	// r.POST("/category", handlerV1.CreateCategory)
	// r.GET("/category/:id", handlerV1.GetByIdCategory)
	// r.GET("/category", handlerV1.GetListCategory)
	// r.DELETE("/category/:id", handlerV1.DeleteCategory)
	// r.PUT("/category/:id", handlerV1.UpdateCategory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
