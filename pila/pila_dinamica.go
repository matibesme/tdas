package pila

const CAPACIDAD_INICIAL = 5
const FACTOR_CAMBIO_CAPACIDAD = 2
const FACTOR_REDUCCION_CAPACIDAD = 4

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {

	nueva_pila := pilaDinamica[T]{
		datos:    make([]T, CAPACIDAD_INICIAL),
		cantidad: 0,
	}
	return &nueva_pila

}

func (pila *pilaDinamica[T]) EstaVacia() bool {

	return pila.cantidad == 0

}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	pila.cantidad++
	if pila.cantidad == cap(pila.datos) {
		nueva_capacidad := cap(pila.datos) * FACTOR_CAMBIO_CAPACIDAD
		pila.redimensionarPila(nueva_capacidad)
	}

	pila.datos[pila.cantidad-1] = dato

}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.verificarNoVacia()
	pila.cantidad--
	if pila.cantidad <= cap(pila.datos)/FACTOR_REDUCCION_CAPACIDAD {
		nueva_capacidad := cap(pila.datos) / FACTOR_CAMBIO_CAPACIDAD
		pila.redimensionarPila(nueva_capacidad)
	}

	return pila.datos[pila.cantidad]

}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.verificarNoVacia()
	return pila.datos[pila.cantidad-1]
}


func (pila *pilaDinamica[T]) redimensionarPila(nueva_capacidad int) {
	nuevos_datos := make([]T, nueva_capacidad)
	copy(nuevos_datos, pila.datos)
	pila.datos = nuevos_datos

}

func (pila *pilaDinamica[T]) verificarNoVacia() {

	if pila.EstaVacia() {
		panic("La pila esta vacia")

	}

}
