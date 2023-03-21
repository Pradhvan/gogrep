package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue("bar")
	elements := queue.GetAll()
	assert.Equal(t, "bar", elements[0])
}

func TestDequeue(t *testing.T) {
	queue := Queue{elements: []string{"foobar"}}
	queue.Dequeue()
	assert.Equal(t, []string{}, queue.GetAll())

}

func TestDequeueError(t *testing.T) {
	queue := Queue{}
	errstr := "elements slice bounds out of range [1:0]"
	assert.Contains(t, errstr, queue.Dequeue().Error())
}

func TestGetAll(t *testing.T) {
	queue := Queue{elements: []string{"foobar", "jhondoe"}}
	assert.Equal(t, 2, len(queue.GetAll()))
}

func TestClear(t *testing.T) {
	queue := Queue{elements: []string{"foobar", "jhondoe"}}
	queue.Clear()
	assert.Equal(t, []string{}, queue.GetAll())
}
