package mocks

import (
	"github.com/stretchr/testify/mock"
	"pagosdearrendamiento/arquitectura/limpia/centro/dominio/entidades"
)

// mock de RepositorioPago
type RepositorioPago struct {
	mock.Mock
}

// mock de Crear
func (repo RepositorioPago) Crear(pago entidades.Pago) (int,error) {
	args := repo.Called(pago)
	args.Error(0)
	return 0, nil
}

// mock de Listar
func (repo RepositorioPago) ListarTodos() ([]*entidades.Pago, error){
	args := repo.Called()
	pago, ok := args.Get(0).([]*entidades.Pago)
	if !ok {
		return nil, args.Error(1)
	}
	return pago, nil
}

