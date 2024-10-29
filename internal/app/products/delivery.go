package products

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetProducts(g *gin.Context)
	GetProduct(g *gin.Context)
	UpdateProduct(g *gin.Context)
	DeleteProduct(g *gin.Context)
}
