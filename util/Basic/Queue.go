package basic

// Queue is an interface
// that refers to a FIFO data structure
type Queue interface {
	Push(x interface{})
	Pop() interface{}
	Peek() interface{}
}

// ArrayQueue is the Impl of Queue Interface using array
type ArrayQueue struct {
	data []interface{}
}

// Push is a impl of Queue based on array
func (a *ArrayQueue) Push(x interface{}) {
	a.data = append(a.data, x)
}

// Pop is a impl of Queue based on array
func (a *ArrayQueue) Pop() interface{} {
	x := a.data[0]
	if len(a.data) > 0 {
		a.data = a.data[1:]
	}
	return x
}

// Peek is a impl of Queue based on array
func (a *ArrayQueue) Peek() interface{} {
	if len(a.data) > 0 {
		return a.data[0]
	}
	return nil
}
