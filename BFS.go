package main

import "fmt"

const n  = 5

var dimension = [5][5]int{ {0,0,0,1,0},
	{1,0,0,0,1},
	{1,0,0,1,0},
	{0,1,0,0,0},
	{0,0,0,0,0} }


type Node struct {
	x,y int
	step int
	prev *Node
}

func New(x,y,step int) Node {
	return Node{x:x,y:y,step:step}
}

type LinkedQueue struct {
	head int
	tail int
	list []Node
}

func (queue *LinkedQueue) push(node Node)  {
	queue.tail++
	queue.list = append(queue.list,node)
}

func (queue *LinkedQueue) popHead()  {
	queue.head++
}
var queue LinkedQueue

func bfs() {

}

func main() {
	node := New(1,2,3)
	node1 := New(4,5,6)
	queue.push(node)
	queue.push(node1)
	fmt.Println(queue)
	fmt.Println(queue.tail)
	queue.popHead()
	fmt.Println(queue.head)
	fmt.Println(queue.tail)
	fmt.Println(queue)
}