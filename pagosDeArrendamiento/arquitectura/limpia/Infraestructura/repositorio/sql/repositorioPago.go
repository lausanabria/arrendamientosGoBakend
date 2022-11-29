package sql

import (
	"database/sql"
	"log"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type RepositorioPago struct {
	db *sql.DB
}

//Crear crea el pago en la base de datos
func (repo RepositorioPago) Crear(pago entidades.Pago) (int, error) {
	repo.db = ConexionBD()
	insertarRegistros, err := repo.db.Prepare("INSERT INTO pagos values (" +
		strconv.Itoa(pago.DocumentoIdentificacionArrendatario) + ",'" +
		pago.CodigoInmueble + "'," +
		strconv.Itoa(pago.ValorPagado) +
		",'" + pago.FechaPago.String() + "')")
	if err != nil {

		repo.db.Close()
		panic(err.Error())
		return -1, err
	}
	_, err = insertarRegistros.Exec()
	me, _ := err.(*mysql.MySQLError)
	if me != nil {
		if me.Number == 1062 { // ya existe el registro en la base de datos, entonces actualizar el valor pagado
			valorPagado, err := repo.obtenerValorPagadoPorIDArrendatarioYCodigoInmueble(pago)
			if err != nil {
				return -1, err
			}
			pago.ValorPagado = valorPagado + pago.ValorPagado
			errActualizar := repo.actualizarPago(pago)
			if errActualizar != nil {
				return -1, errActualizar
			}
		}
	}
	repo.db.Close()
	faltaPorPagar := entidades.VALORTOTALARRIENDO - pago.ValorPagado
	if faltaPorPagar > 0 {
		return faltaPorPagar, nil
	}
	return 0, nil
}

//ListarTodos consulta todos los pagos de la base de datos
func (repo RepositorioPago) ListarTodos() ([]*entidades.Pago, error) {
	repo.db = ConexionBD()
	filas, err := repo.db.Query("SELECT * FROM pagos")
	if err != nil {
		return nil, err
	}
	pagos := []*entidades.Pago{}

	for filas.Next() {
		pago := entidades.Pago{}
		var fecha string
		if err := filas.Scan(&pago.DocumentoIdentificacionArrendatario, &pago.CodigoInmueble, &pago.ValorPagado, &fecha); err != nil {
			log.Fatal(err)
		}
		pago.FechaPago = repo.stringAFecha(fecha)
		pagos = append(pagos, &pago)
	}
	repo.db.Close()
	return pagos, nil
}

//obtenerValorPagadoPorIDArrendatarioYCodigoInmueble obtiene el valor pagado de un registro ya existente
func (repo RepositorioPago) obtenerValorPagadoPorIDArrendatarioYCodigoInmueble(pago entidades.Pago) (int, error) {
	repo.db = ConexionBD()
	filas, err := repo.db.Query("SELECT valorPagado FROM pagos WHERE documentoIdentificacionArrendatario=" +
		strconv.Itoa(pago.DocumentoIdentificacionArrendatario) + " and codigoInmueble='" +
		pago.CodigoInmueble + "'")
	if err != nil {
		return -1, err
	}
	var valorPagado int = 0
	for filas.Next() {
		if err := filas.Scan(&valorPagado); err != nil {
			return -1, err
		}
		break
	}
	repo.db.Close()
	return valorPagado, nil

}

//actualizarPago actualiza el pago en la base de datos
func (repo RepositorioPago) actualizarPago(pago entidades.Pago) error {
	_, err := repo.db.Exec(`UPDATE pagos SET valorPagado = ? WHERE documentoIdentificacionArrendatario = ? and codigoInmueble = ?`, pago.ValorPagado, pago.DocumentoIdentificacionArrendatario, pago.CodigoInmueble)
	if err != nil {
		return err
	}
	return nil
}

//stringAFecha convierte un string a tipo fecha
func (repo RepositorioPago) stringAFecha(fechaPago string) time.Time {

	fechaPorPartes := strings.Split(fechaPago, "-")
	mes, _ := strconv.Atoi(fechaPorPartes[1])
	fecha := fechaPorPartes[2] + " " + entidades.Mes[mes] + " " + fechaPorPartes[0]

	formato := "2 Jan 2006"
	marcaDeFecha, _ := time.Parse(formato, fecha)
	return marcaDeFecha
}
