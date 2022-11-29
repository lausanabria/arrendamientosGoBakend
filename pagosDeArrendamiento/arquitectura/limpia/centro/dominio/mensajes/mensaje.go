package mensajes

//Respuesta estructura que contiene los campos para mostrar una respuesta tipo JSON
type Respuesta struct {
	Codigo int
	Mensaje string
}

//GetRespuesta devuelve una estructura tipo Respuesta dado un código y un mensaje
func GetRespuesta(codigo int, mensaje string) Respuesta {

	msj := Respuesta{}
	msj.Codigo = codigo
	msj.Mensaje = mensaje

	return msj
}
