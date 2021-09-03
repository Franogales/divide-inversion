package api

import (
	"net/http"

	"github.com/Franogales/divide-inversion/second/credito"
	"github.com/gin-gonic/gin"
)

type data struct {
	Investment int32 `form:"investment" json:"investment"  binding:"required"`
}

func Assign(c *gin.Context) {
	credito := credito.New()
	var json data
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad parameter, expected 'investment' type integer"})
		return
	}
	invest := int32(json.Investment)
	_, _, _, err := credito.Assign(invest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// c.Writer.WriteHeader(400)
		return
	}
	c.JSON(200, credito)
}
