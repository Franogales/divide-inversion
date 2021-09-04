package api

import (
	"net/http"

	"github.com/Franogales/divide-inversion/second/credito"
	"github.com/gin-gonic/gin"
)

type data struct {
	Investment int32 `form:"investment" json:"investment"  binding:"required"`
}
type CreditService struct {
	credit credito.CreditType
}

/* Funcion para assignar una inversion desde el api rest */
func Assign(c *gin.Context) {
	service := CreditService{credit: *credito.New()}
	var json data
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad parameter, expected 'investment' type integer"})
		return
	}
	invest := int32(json.Investment)
	/* aqui pongo un mensaje solo por practica para saber que esta pasando, no se me va bien solo negar peticiones :) */
	_, _, _, err := service.credit.Assign(invest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, service.credit)
}
