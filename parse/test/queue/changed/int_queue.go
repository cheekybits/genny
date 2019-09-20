// Code generated https://github.com/cheekybits/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package changed

// IntQueue is a queue of Ints.
type IntQueue struct {
	items []int
}

func NewIntQueue() *IntQueue {
	return &IntQueue{items: make([]int, 0)}
}
func (q *IntQueue) Push(item int) {
	q.items = append(q.items, item)
}
func (q *IntQueue) Pop() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
