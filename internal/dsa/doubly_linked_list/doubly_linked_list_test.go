package doubly_linked_list

import "testing"

func TestDoublyLinkedList_PrependAppend(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Prepend(0)
	list.Append(2)

	expected := []int{0, 1, 2}
	vals := list.Values()
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	for i, v := range expected {
		if vals[i] != v {
			t.Errorf("Expected %d, got %d", v, vals[i])
		}
	}
}

func TestDoublyLinkedList_InsertAt(t *testing.T) {
	list := New[string]()
	list.Append("b")
	list.InsertAt("a", 0) // Prepend
	list.InsertAt("c", 2) // Append
	list.InsertAt("b2", 2)

	expected := []string{"a", "b", "b2", "c"}
	vals := list.Values()
	for i, v := range expected {
		if vals[i] != v {
			t.Errorf("Expected %s at idx %d, got %s", v, i, vals[i])
		}
	}
    // Verify Prev links are correct by walking backwards
    curr := list.tail
    for i := len(expected)-1; i >= 0; i-- {
        if curr.Value != expected[i] {
            t.Errorf("Reverse link broke")
        }
        curr = curr.Prev
    }
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	list := New[int]()
	list.Append(0)
    list.Append(1)
    list.Append(2)
    list.Append(3)

	list.RemoveAt(2) // removes 2
	expected := []int{0, 1, 3}
	vals := list.Values()
	for i, v := range vals {
		if expected[i] != v {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}

	list.RemoveAt(0) // removes 0
	if list.head.Value != 1 {
		t.Errorf("Expected new head to be 1, got %v", list.head.Value)
	}
	
	list.RemoveAt(1) // removes 3 (now tail)
	if list.tail.Value != 1 {
		t.Errorf("Expected new tail to be 1, got %v", list.tail.Value)
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	list := New[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	
	list.Remove(20)
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	
	err := list.Remove(100)
	if err != ErrValueNotFound {
		t.Errorf("Expected ErrValueNotFound, got %v", err)
	}
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	
	expected := []int{3, 2, 1}
	vals := list.Values()
	for i, v := range expected {
		if vals[i] != v {
			t.Errorf("Expected %d, got %d", v, vals[i])
		}
	}
	if list.head.Value != 3 || list.tail.Value != 1 {
		t.Errorf("Head or tail not matched after reverse")
	}
    
    // verify prev pointers are correct
    curr := list.tail
    for i := len(expected)-1; i >= 0; i-- {
        if curr.Value != expected[i] {
            t.Errorf("Reverse link broke after Reverse()")
        }
        curr = curr.Prev
    }
}
