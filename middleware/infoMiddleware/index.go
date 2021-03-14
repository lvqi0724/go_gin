package infoMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PrintInfo() gin.HandlerFunc{
	return func(context *gin.Context) {
		fmt.Println(context.Request.UserAgent())
		fmt.Println(context.FullPath())
		context.String(http.StatusOK,"2222")
		context.Abort()
	}
}

func TestMiddleware () gin.HandlerFunc{
	return func(context *gin.Context) {
		fmt.Println("单例middleware")

		context.Next()

		fmt.Println(context.Writer.Status())
	}
}