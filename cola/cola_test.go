package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

var INT_PRUEBA = []int{3, 746321, 12, 1000}
var FLOAT_PRUEBA = []float64{3.23, 4.23412, 123.1000}
var STRING_PRUEBA = []string{"Algoritmos 2", "Messi rve", "Stephen Curry"}

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola no esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestOrdenFifoInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	testOrdenFifoAny(t, cola, INT_PRUEBA)

	//Comprobacion que sea igual a una cola recien creada
	TestColaVacia(t)
}

func TestOrdenFifoString(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	testOrdenFifoAny(t, cola, STRING_PRUEBA)
	//Comprobacion que sea igual a una cola recien creada
	TestColaVacia(t)

}

func TestOrdenFifoFloat64(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()

	testOrdenFifoAny(t, cola, FLOAT_PRUEBA)
	//Comprobacion que sea igual a una cola recien creada
	TestColaVacia(t)
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= INT_PRUEBA[3]; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= INT_PRUEBA[3]; i++ {
		require.Equal(t, i, cola.VerPrimero())
		cola.Desencolar()
	}
	TestColaVacia(t)
}

func testOrdenFifoAny[T any](t *testing.T, cola TDACola.Cola[T], lista []T) {

	for i := 0; i <= len(lista)-1; i++ {
		cola.Encolar(lista[i])
		require.Equal(t, lista[0], cola.VerPrimero())

	}
	for i := 0; i <= len(lista)-1; i++ {
		require.Equal(t, lista[i], cola.VerPrimero())
		require.Equal(t, lista[i], cola.Desencolar())
	}

	//Comprobacion que sea igual a una cola recien creada
	TestColaVacia(t)
}
