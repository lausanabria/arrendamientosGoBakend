package sql

import (
	"database/sql"
	"fmt"
)

//ConexionBD realiza la conexión a la base de datos
func ConexionBD() (conexion *sql.DB) {

	controlador := "mysql"
	usuario := "root"
	clave := ""
	nombreBD := "pruebaceiba"
	TnsBD := "tcp"
	conexion, err := sql.Open(controlador, usuario+":"+clave+"@"+TnsBD+"/"+nombreBD)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexión establecida a la base de datos")
	}

	return conexion
}
