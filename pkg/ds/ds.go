package ds

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(i string) {
	q.elements = append(q.elements, i)
}

func (q *Queue) Dequeue() {
	q.elements = q.elements[1:]
}

func (q *Queue) GetAll() []string {
	return q.elements
}

func (q *Queue) Clear() {
	q.elements = []string{}
}
