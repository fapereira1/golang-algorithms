package dynamic_array_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dsa_array "github.com/fapereira1/golang-algorithms/internal/dsa/dynamic_array"
)

func TestDynamicArray_PushBackAndGet(t *testing.T) {
	da := dsa_array.New[int](2)

	assert.True(t, da.IsEmpty())
	assert.Equal(t, 2, da.Capacity())

	// Push items
	da.PushBack(10)
	da.PushBack(20)

	assert.Equal(t, 2, da.Size())
	assert.Equal(t, 2, da.Capacity())

	// Push more, trigger resize
	da.PushBack(30)
	assert.Equal(t, 3, da.Size())
	assert.Equal(t, 4, da.Capacity()) // 2 * 2 = 4

	// Verify values
	val, err := da.Get(0)
	require.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = da.Get(2)
	require.NoError(t, err)
	assert.Equal(t, 30, val)
}

func TestDynamicArray_Set(t *testing.T) {
	da := dsa_array.New[string](5)
	da.PushBack("hello")
	da.PushBack("world")

	err := da.Set(1, "golang")
	require.NoError(t, err)

	val, err := da.Get(1)
	require.NoError(t, err)
	assert.Equal(t, "golang", val)

	// Out of bounds set
	err = da.Set(5, "out")
	assert.Error(t, err)
}

func TestDynamicArray_GetOutOfBounds(t *testing.T) {
	da := dsa_array.New[int](2)
	da.PushBack(1)

	// Negative index
	_, err := da.Get(-1)
	assert.Error(t, err)

	// Zero-indexed element (size is 1)
	_, err = da.Get(1)
	assert.Error(t, err)
}

func TestDynamicArray_RemoveAndShrink(t *testing.T) {
	da := dsa_array.New[int](2)
	
	// Add 9 elements. Capacity will scale: 2 -> 4 -> 8 -> 16
	for i := 0; i < 9; i++ {
		da.PushBack(i)
	}

	assert.Equal(t, 9, da.Size())
	assert.Equal(t, 16, da.Capacity())

	// Remove element at index 4 (value '4')
	err := da.Remove(4)
	require.NoError(t, err)

	// Value at index 4 should now be 5
	val, _ := da.Get(4)
	assert.Equal(t, 5, val)

	// Remove elements until size becomes <= capacity/4 (16 / 4 = 4)
	// We currently have 8 elements, capacity 16.
	// Remove 5 more elements so size = 3, which is <= 4.
	for i := 0; i < 5; i++ {
		da.Remove(0)
	}

	assert.Equal(t, 3, da.Size())
	// Expected new capacity: 16 / 2 = 8
	assert.Equal(t, 8, da.Capacity())
}

func TestDynamicArray_RemoveOutOfBounds(t *testing.T) {
	da := dsa_array.New[int](1)
	err := da.Remove(0)
	assert.Error(t, err)

	da.PushBack(1)
	err = da.Remove(1) // out of bounds
	assert.Error(t, err)
}

func TestDynamicArray_Strings(t *testing.T) {
	da := dsa_array.New[string](0) // Test behavior with 0 capacity input
	assert.Equal(t, 2, da.Capacity()) // Default minimum applies
	
	da.PushBack("test")
	val, err := da.Get(0)
	require.NoError(t, err)
	assert.Equal(t, "test", val)
}
