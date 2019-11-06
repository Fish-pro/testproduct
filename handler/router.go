package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/testproduct/context"
	"github.com/leyle/userandrole/api"
)

func PRouter(ctx *context.Context, g *gin.RouterGroup) {
	pR := g.Group("/product", func(c *gin.Context) {
		api.Auth(c)
	})
	{
		// 新建 product
		pR.POST("/product", func(c *gin.Context) {
			CreateProductHandler(ctx, c)
		})

		// 修改 product
		pR.PUT("/product/:id", func(c *gin.Context) {
			DummyHandler(ctx, c)
		})

		// 上线
		pR.POST("/product/:id/online", func(c *gin.Context) {
			DummyHandler(ctx, c)
		})

	}

	noAuth := g.Group("/product")
	{
		// 明细
		noAuth.GET("/product/:id", func(c *gin.Context) {
			GetProductInfoHandler(ctx, c)
		})

		// 搜索
		noAuth.GET("/products", func(c *gin.Context) {
			QueryProductHandler(ctx, c)
		})

	}
}

func DummyHandler(ctx *context.Context, c *gin.Context) {

}