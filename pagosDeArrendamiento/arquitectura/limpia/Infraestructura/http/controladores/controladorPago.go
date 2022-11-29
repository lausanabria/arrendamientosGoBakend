package controladores

import (
	"net/http"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/servicio"

	"github.com/gin-gonic/gin"
)

type ControladorPago struct {
	servicioPago servicio.Pago
}

func (p ControladorPago) Crear(c *gin.Context) {

	var datosPago entidades.CopiaPago
	error := c.BindJSON(&datosPago)
	if error != nil {
		panic(error)
	}

	if error != nil {
		panic("Error al leer los parámetros")
	}

	payment := entidades.Pago{}

	err := c.ShouldBind(&payment)
	if err != nil {
		panic("Error binding")
	}

	p.servicioPago = servicio.ServicioPago{}

	mensaje := p.servicioPago.Crear(datosPago)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(mensaje.Codigo, gin.H{
		"respuesta": mensaje.Mensaje,
	})

}

func (p ControladorPago) Get(c *gin.Context) {

	p.servicioPago = servicio.ServicioPago{}
	Pagos, _ := p.servicioPago.ListarTodos()
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{"results": Pagos})
}
