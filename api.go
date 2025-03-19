package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Diccionario de sistemas averiados y sus códigos
var systemCodes = map[string]string{
	"navigation":       "NAV-01",
	"communications":   "COM-02",
	"life_support":     "LIFE-03",
	"engines":          "ENG-04",
	"deflector_shield": "SHLD-05",
}

// Lista de sistemas para seleccionar aleatoriamente
var systems = []string{"navigation", "communications", "life_support", "engines", "deflector_shield"}
var damagedSystem string

func main() {
	r := gin.Default()

	// Endpoint /status
	r.GET("/status", func(c *gin.Context) {
		damagedSystem = systems[rand.Intn(len(systems))]
		c.JSON(http.StatusOK, gin.H{"damaged_system": damagedSystem})
	})

	// Endpoint /repair-bay
	r.GET("/repair-bay", func(c *gin.Context) {
		systemCode := systemCodes[damagedSystem]
		htmlContent := "<!DOCTYPE html><html><head><title>Repair</title></head><body><div class='anchor-point'>" + systemCode + "</div></body></html>"
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// Endpoint /teapot
	r.POST("/teapot", func(c *gin.Context) {
		c.Status(http.StatusTeapot)
	})

	// Iniciar servidor en el puerto 8000
	r.Run(":3000")
}
