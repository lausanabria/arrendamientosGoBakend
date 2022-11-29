package services

import (
	"pagosdearrendamiento/arquitectura/limpia/Infraestructura/repositorio/mocks"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
	"testing"
	"time"
)

func PruebaCrearPago(t *testing.T) {

	var repo mocks.RepositorioPago

	pago := entidades.Pago{}
	pago.DocumentoIdentificacionArrendatario = 1036946622
	pago.CodigoInmueble = "8870"
	pago.ValorPagado = 1000000

	fecha := "25 sep 2020"
	formato := "2 Jan 2006"
	marcaDeFecha, _ := time.Parse(formato, fecha)

	pago.FechaPago = marcaDeFecha

	repo.On("Crear", pago).Return(0,nil)

	repo.AssertCalled(t, "Crear", 0, nil)

}
