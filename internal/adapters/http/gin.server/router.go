package gin_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"ml-x-men/internal/application"
	"net/http"
)

type RouterHandler struct {
	ucHandler application.Handler
	Logger    application.Logger
}

func NewRouter() RouterHandler {
	return RouterHandler{}
}

func NewRouterWithLogger(i application.Handler, logger application.Logger) RouterHandler {
	return RouterHandler{
		ucHandler: i,
		Logger:    logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.Use(rH.errorCatcher())

	rH.mutantRoutes(api)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}

func (rH RouterHandler) mutantRoutes(api *gin.RouterGroup) {
	mutantGroup := api.Group("/mutant")
	mutantGroup.POST("", rH.mutantPost)
}

func (rH RouterHandler) errorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() > 399 {
			c.Render(
				c.Writer.Status(),
				render.Data{
					ContentType: "application/json; charset=utf-8",
					Data:        []byte(`{"error": "There is an error! Please check the http status code" }}`),
				},
			)
		}
	}
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
