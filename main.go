package main

import "fmt"

func main() {
 testQueque()
}

func testQueque(){
	queue := NewQueue()
	queue.Enquque(100)
	fmt.Println("queque:", queue)
	queue.Enquque(25).Enquque(77)
	fmt.Println("queque", queue)
	fmt.Println("is queque empty?", queue.IsEmpty())

	var result, _ = queue.Peek()
	fmt.Println("the next item in the queue:",result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeque()
	fmt.Println(result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeque()
	fmt.Println(result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeque()
	fmt.Println(result)

}   

// declares type of Queue
type Queue struct {
	data []int
}

// returns new  Queue
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// checks if Queue empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// returns next element of Queue
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// adds new element to Queue
func (q *Queue) Enquque(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// returns current element and removes it
func (q *Queue) Dequeque() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queque is empty")

	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
