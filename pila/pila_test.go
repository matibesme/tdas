package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

var INT_PRUEBA = []int{3, 746321, 12, 1000}
var FLOAT_PRUEBA = []float64{3.23, 4.23412, 123.1000}

var STRING_PRUEBA = []string{"Algoritmos 2", "Messi rve", "Stephen Curry"}

// Se pueda crear una Pila vacía, y esta se comporta como tal
// Condición de borde: la acción de esta_vacía en una pila recién creada es verdadero.
// Condición de borde: las acciones de desapilar y ver_tope en una pila recién creada son inválidas.
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila no esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

// Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de pila (que esta es LIFO).
// Probar con elementos diferentes, y ver que salgan en el orden deseado.
func TestOrdenLifoInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= len(INT_PRUEBA)-1; i++ {
		pila.Apilar(INT_PRUEBA[i])
		require.Equal(t, INT_PRUEBA[i], pila.VerTope())

	}
	for i := len(INT_PRUEBA) - 1; i >= 0; i-- {
		require.Equal(t, INT_PRUEBA[i], pila.VerTope())
		require.Equal(t, INT_PRUEBA[i], pila.Desapilar())
	}

	//Comprobacion que sea igual a una pila recien creada
	TestPilaVacia(t)
}

func TestOrdenLifoString(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	for i := 0; i <= len(STRING_PRUEBA)-1; i++ {
		pila.Apilar(STRING_PRUEBA[i])
		require.Equal(t, STRING_PRUEBA[i], pila.VerTope())

	}
	for i := len(STRING_PRUEBA) - 1; i >= 0; i-- {
		require.Equal(t, STRING_PRUEBA[i], pila.VerTope())
		require.Equal(t, STRING_PRUEBA[i], pila.Desapilar())
	}

	//Comprobacion que sea igual a una pila recien creada
	TestPilaVacia(t)

}

func TestOrdenLifoFloat64(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	for i := 0; i <= len(FLOAT_PRUEBA)-1; i++ {
		pila.Apilar(FLOAT_PRUEBA[i])
		require.Equal(t, FLOAT_PRUEBA[i], pila.VerTope())

	}
	for i := len(FLOAT_PRUEBA) - 1; i >= 0; i-- {
		require.Equal(t, FLOAT_PRUEBA[i], pila.VerTope())
		require.Equal(t, FLOAT_PRUEBA[i], pila.Desapilar())
	}

	//Comprobacion que sea igual a una pila recien creada
	TestPilaVacia(t)

}

// Prueba de volumen: hacer crecer la pila, y desapilar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante. Validar que se cumpla siempre que el tope de la pila sea el correcto paso a paso, y que el nuevo tope después de cada desapilar también sea el correcto.
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= INT_PRUEBA[3]; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())

	}
	for i := INT_PRUEBA[3]; i >= 1; i-- {
		require.Equal(t, i, pila.VerTope())
		pila.Desapilar()
	}
	TestPilaVacia(t)
}
