package ds

import "fmt"

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(i string) {
	q.elements = append(q.elements, i)
}

func (q *Queue) Dequeue() error {
	if len(q.elements) > 0 {
		q.elements = q.elements[1:]
		return nil
	}
	return fmt.Errorf("elements slice bounds out of range [1:0]")
}

func (q *Queue) GetAll() []string {
	return q.elements
}

func (q *Queue) Clear() {
	q.elements = []string{}
}
