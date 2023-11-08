package cola

/* Definición del struct pila proporcionado por la cátedra. */

type nodo[T any] struct {
	dato    T
	proximo *nodo[T]
}

type colaEnlazada[T any] struct {
	primer_nodo *nodo[T]
	ultimo_nodo *nodo[T]
}

func crearNodo[T any](datos T) *nodo[T] {

	nuevo_nodo := nodo[T]{dato: datos}

	return &nuevo_nodo

}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primer_nodo == nil
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevo_nodo := crearNodo(dato)
	if cola.EstaVacia() {
		cola.primer_nodo = nuevo_nodo
	} else {
		cola.ultimo_nodo.proximo = nuevo_nodo
	}
	cola.ultimo_nodo = nuevo_nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {

	cola.verificarNoVacia()
	dato_a_desencolar := cola.primer_nodo.dato
	cola.primer_nodo = cola.primer_nodo.proximo
	if cola.EstaVacia() {
		cola.ultimo_nodo = nil
	}
	return dato_a_desencolar

}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.verificarNoVacia()
	return cola.primer_nodo.dato

}

func (cola *colaEnlazada[T]) verificarNoVacia() {

	if cola.EstaVacia() {
		panic("La cola esta vacia")

	}

}
