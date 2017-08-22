package main

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right,ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int,10)
	ch2 := make(chan int,10)
	r1 := make([]int, 10)
	r2 := make([]int, 10)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i := 0 ; i < 10;i++ {
		select {
		case v1 := <-ch1:
			r1 = append(r1, v1)
		case v2 := <-ch2:
			r2 = append(r2, v2)
		}
	}
	for i := range r1{
		if(r1[i] != r2[i]){
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1),tree.New(2)))
}
