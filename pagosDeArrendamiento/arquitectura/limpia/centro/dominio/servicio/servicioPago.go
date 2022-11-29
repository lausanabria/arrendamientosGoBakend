package servicio

import (
	"net/http"
	"pagosdearrendamiento/arquitectura/limpia/Infraestructura/repositorio/sql"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/mensajes"
	"strconv"
)

//Pago interface para implementar los métodos Crear y ListarTodos los pagos
type Pago interface {
	Crear(pago entidades.CopiaPago) mensajes.Respuesta
	ListarTodos() ([]*entidades.Pago, mensajes.Respuesta)
}

//ServicioPago estructura con el objeto de tipo repositorioPago
type ServicioPago struct {
	repositorioPago entidades.RepositorioPago
}

//Crear valida que todos los campos del json estén bien definidos y en sus rangos establecidos y luego crea el pago a través del repositorio
func (p ServicioPago) Crear(jsonPago entidades.CopiaPago) mensajes.Respuesta {

	//validaciones de tipos de dato

	documento, convertidoOK := validaDocumentoIdentificacionArrendatario(jsonPago.DocumentoIdentificacionArrendatario)
	if !convertidoOK {
		return mensajes.GetRespuesta(http.StatusConflict, "El documento de identificación debe ser un valor numérico")
	}

	valor, convertidoOK := validaValorPagado(jsonPago.ValorPagado)
	if !convertidoOK {
		return mensajes.GetRespuesta(http.StatusConflict, "El valor pagado debe ser un valor numérico")
	}

	fecha, convertidoOK := validaFechaPago(jsonPago.FechaPago)
	if !convertidoOK {
		return mensajes.GetRespuesta(http.StatusBadRequest, "Formato de fecha incorrecto")
	}

	pago := entidades.Pago{}
	pago.DocumentoIdentificacionArrendatario = documento
	pago.CodigoInmueble = jsonPago.CodigoInmueble
	pago.ValorPagado = valor
	pago.FechaPago = fecha

	// validación de reglas de negocio

	valido := validaFechaImpar(pago)
	if !valido {
		return mensajes.GetRespuesta(http.StatusBadRequest, "lo siento pero no se puede recibir el pago por decreto de administración")
	}

	valido = validaRangoDeValorPagado(pago)
	if !valido {
		return mensajes.GetRespuesta(http.StatusConflict, "Valor pagado no permitido, debe estar entre 1 y 1000000")
	}

	p.repositorioPago = sql.RepositorioPago{}

	total, valor := validaPagoTotal(pago)
	faltaPorPagar, err := p.repositorioPago.Crear(pago)

	if total {
		if err == nil {
			return mensajes.GetRespuesta(http.StatusOK, "gracias por pagar todo tu arriendo")
		}
		return mensajes.GetRespuesta(http.StatusInternalServerError, "error al intentar insertar el pago en la base de datos")
	} else {
		if err == nil && faltaPorPagar > 0 {
			return mensajes.GetRespuesta(http.StatusOK, "gracias por tu abono, sin embargo recuerda que te hace falta pagar "+strconv.Itoa(faltaPorPagar))
		}
		if faltaPorPagar == 0 {
			return mensajes.GetRespuesta(http.StatusOK, "gracias por pagar todo tu arriendo")
		}
		return mensajes.GetRespuesta(http.StatusInternalServerError, "error al intentar insertar el pago en la base de datos")

	}

}

//ListarTodos devuelve una lista de todos los pagos existentes en la base de datos
func (p ServicioPago) ListarTodos() ([]*entidades.Pago, mensajes.Respuesta) {
	p.repositorioPago = sql.RepositorioPago{}
	pagos, err := p.repositorioPago.ListarTodos()
	if err != nil {
		return pagos, mensajes.GetRespuesta(http.StatusInternalServerError, "Error al obtener los pagos de la base de datos")
	}
	return pagos, mensajes.GetRespuesta(http.StatusOK, "pagos listados")
}
