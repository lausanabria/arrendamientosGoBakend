package rutas

import (
	"net/http"
	"pagosdearrendamiento/arquitectura/limpia/Infraestructura/http/controladores"

	"github.com/gin-gonic/gin"
)

// Pago rutas de pagos de arrendamiento
func Pago() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Conexión establecida puerto 8084",
		})
	})

	controladorPago := &controladores.ControladorPago{}
	r.POST("/api/pagos", controladorPago.Crear)
	r.GET("/api/pagos", controladorPago.Get)

	r.Run(":8084")

}
