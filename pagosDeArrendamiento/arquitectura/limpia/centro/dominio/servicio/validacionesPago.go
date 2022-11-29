package servicio

import (
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
	"strconv"
	"strings"
	"time"
)

//validaDocumentoIdentificacionArrendatario valida que el documento de identificación del arrendatario sea numérico entero
func validaDocumentoIdentificacionArrendatario(documentoIdentificacionArrendatario string) (int,bool){
	documento , error := strconv.Atoi(documentoIdentificacionArrendatario)
	if error == nil{
		return documento,true
	}
	return -1,false
}

//validaValorPagado valida que el valor pagado sea numérico entero
func validaValorPagado(valorPagado string) (int,bool){
	valor, err := strconv.Atoi(valorPagado)
	if err== nil {
		return valor,true
	}
	return -1 , false
}

//validaRangoDeValorPagado valida que el pago se encuentre entre 1 y 1000000
func validaRangoDeValorPagado(pago entidades.Pago) bool{
	if(pago.ValorPagado >=1 && pago.ValorPagado<=1000000){
		return true
	}
	return false
}

//validaFechaPago valida que la fecha de pago sea una fecha válida y que esté en el formato DD/MM/AAAA
func validaFechaPago(fechaPago string) (time.Time,bool){

	fechaPorPartes :=  strings.Split(fechaPago, "/")
	mes,err := strconv.Atoi(fechaPorPartes[1])
	if(err!= nil){
		return time.Now(),false
	}
	fecha := fechaPorPartes[0]+ " " + entidades.Mes[mes]+" "+fechaPorPartes[2]

	formato := "2 Jan 2006"
	marcaDeFecha, err := time.Parse(formato, fecha)

	if err != nil {
		return marcaDeFecha,false
	}
	return marcaDeFecha,true
}

//validaFechaImpar valida que el pago se realice en una fecha impar
func validaFechaImpar(pago entidades.Pago) bool{

	_, _, dia := pago.FechaPago.Date()
	if dia % 2 != 0 {
		return true
	}
	return false
}

//valida si el pago que está realizando el arrendatario es el pago total de 1000000
func validaPagoTotal(pago entidades.Pago) (bool,int){
	if(pago.ValorPagado == entidades.VALORTOTALARRIENDO){
		return true,-1
	}
	return false,entidades.VALORTOTALARRIENDO-pago.ValorPagado
}