package entidades

import "time"


const VALORTOTALARRIENDO = 1000000

// Mes mapa de la relación número del mes con sus siglas en inglés
var Mes = map[int]string{
	1: "jan",
	2: "feb",
	3: "mar",
	4: "apr",
	5: "may",
	6: "jun",
	7: "jul",
	8: "aug",
	9: "sep",
	10: "oct",
	11: "nov",
	12: "dic",
}

//Pago estructura de Pago
type Pago struct {

	DocumentoIdentificacionArrendatario int
	CodigoInmueble 						string
	ValorPagado 						int
	FechaPago 							time.Time
}

//CopiaPago copia de la estructura copia pero con los datos tipo string que recibe del json
type CopiaPago struct {

	DocumentoIdentificacionArrendatario string			`json:"documentoIdentificacionArrendatario"`
	CodigoInmueble 						string			`json:"codigoInmueble"`
	ValorPagado 						string			`json:"valorPagado"`
	FechaPago 							string			`json:"fechaPago"`
}

//RepositorioPago interface para que el repositorio implemente los métodos Crear y ListarTodos los pagos
type RepositorioPago interface {
	Crear(pago Pago) (int,error)
	ListarTodos() ([]*Pago,error)
}
