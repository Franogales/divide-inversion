package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerRun() {
	/* Uso la configuracion por default de gin */
	router := gin.Default()
	/* Utilizo el sistema de logs de gin con un formato diferente*/
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	/* Se instancea el sistema de recovery custom en caso de panic y returna 500 en caso de uno */
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, struct {
				ErrorName string `json:"error"`
			}{ErrorName: err})
		}

		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	credit := router.Group("/credit-asigment")
	{
		credit.GET("/panic", func(c *gin.Context) {
			// panic de prueba
			panic("foo")
		})

		credit.POST("/", func(c *gin.Context) {
			test := 0
			for i := 0; i < 3; i++ {
				test += i
			}
			c.String(http.StatusOK, "ohai")
		})
	}

	router.Run()
}
