package dynamic_array

import "fmt"

// INTERFACE
// Define a estrutura de um Array Dinâmico genérico.
// - dynamic_array.New[T](capacity) -> cria array dinâmico com capacidade inicial
// - dynamic_array.PushBack(T)
// - dynamic_array.Get(index) -> (T, error)
// - dynamic_array.Set(index, T) -> error
// - dynamic_array.Remove(index) -> error
// - dynamic_array.Size() -> int
// - dynamic_array.Capacity() -> int
// - dynamic_array.IsEmpty() -> bool

// DynamicArray representa um array que cresce e encolhe sua capacidade de forma autônoma.
type DynamicArray[T any] struct {
	data     []T // Usamos slices em Go de forma burra para simular a alocação de blocos na memória.
	size     int // Elementos armazenados (tamanho lógico)
	capacity int // Capacidade máxima ocupada na memória (tamanho físico)
}

// New cria um novo DynamicArray com uma capacidade inicial pré-estabelecida.
func New[T any](initialCapacity int) *DynamicArray[T] {
	if initialCapacity <= 0 {
		initialCapacity = 2 // Capacidade padrão mínima para facilitar redimensionamento
	}
	return &DynamicArray[T]{
		data:     make([]T, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// CORE LOGIC

// resize é um método não-exportado (private) usado para expandir ou encolher a estrutura subjacente.
func (da *DynamicArray[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	// Copiamos os dados no custo de O(n) durante esse pico isolado.
	for i := 0; i < da.size; i++ {
		newData[i] = da.data[i]
	}
	da.data = newData
	da.capacity = newCapacity
}

// PushBack insere um novo elemento no final do array.
// Complexidade de Tempo: O(1) amortizado, O(n) pior-caso se precisar redimensionar
// Complexidade de Espaço: O(1) auxiliar (constante), ignorando os O(n) na realocação
func (da *DynamicArray[T]) PushBack(value T) {
	// EDGE CASE: Limite da capacidade atingido?
	// Se size alcaçar capacity, o array encheu, então dobramos ele de tamanho:
	if da.size == da.capacity {
		da.resize(da.capacity * 2)
	}
	// Com o espaço a disposição garantido, guardamos o valor O(1)
	da.data[da.size] = value
	da.size++
}

// Get acessa um elemento pelo índice sem risco de boundary violantion se usando as convenções corretas.
// O(1)
func (da *DynamicArray[T]) Get(index int) (T, error) {
	var zero T // Instancia o valor padrão zeroado desse Generic (nil, 0, "", struct{})
	if index < 0 || index >= da.size {
		return zero, fmt.Errorf("index out of bounds: %d", index)
	}
	return da.data[index], nil
}

// Set altera um elemento no índice passado.
// O(1)
func (da *DynamicArray[T]) Set(index int, value T) error {
	if index < 0 || index >= da.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	da.data[index] = value
	return nil
}

// Remove exclui o elemento de um dado index. Se the array ficar muito vazio e grande em capacity, ele encolherá.
// O(n) pela varredura em deslocar todos posteriores 1 casa à esquerda.
func (da *DynamicArray[T]) Remove(index int) error {
	if index < 0 || index >= da.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	// Deslocamento de arranjo = O(n) a partir do index, até o final da cadeia.
	for i := index; i < da.size-1; i++ {
		da.data[i] = da.data[i+1]
	}

	// Decrementa e evita memory leak limpando o ref value da velha extremidade se fosse um ponteiro.
	var zero T
	da.data[da.size-1] = zero
	da.size--

	// SHRINKING:
	// Pra diminuir memory leaks teóricos e gastar menos RAM, se tamanho for 1/4 da capacidade, cortamos por metade.
	// EDGE CASE: nunca baixar demais ex: min de 2.
	if da.capacity > 2 && da.size <= da.capacity/4 {
		da.resize(da.capacity / 2)
	}

	return nil
}

// INVARIANCES AND UTILS: Métodos auxiliares O(1) puros de inspeção ("getters").
func (da *DynamicArray[T]) Size() int {
	return da.size
}

func (da *DynamicArray[T]) Capacity() int {
	return da.capacity
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.size == 0
}
