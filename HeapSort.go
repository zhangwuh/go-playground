package main

import "fmt"

var ARRAY = []int{4,2,9,1,7,12,6}
var result = []int{}
func main()  {
	for i:=0;i<len(ARRAY);i++ {
		pickBiggest(ARRAY[i:])
	}
	fmt.Println(result)
}

func pickBiggest(array []int) {
	fmt.Println(array)
	var length = len(array)
	for i := parent(length - 1);i>=0;i-- {
		siftDown(array,i)
	}
	result = append(result,array[0])
}

func parent(i int) int {
	return (i - 1)/2
}

func swap(array []int,i int, j int)  {
	array[i],array[j] = array[j], array[i]
}

func siftDown(array []int,node int)  {
	val := array[node]
	left := left(node,len(array))
	right := right(node,len(array))
	if(left != -1 && val < array[left]) {
		swap(array,node,left)
		siftDown(array,left)
	}
	val = array[node]
	if(right != -1 && val < array[right]){
		swap(array,node,right)
		siftDown(array,right)
	}
}

func left(node int,length int) int {
	left := node << 1 + 1
	fmt.Printf("left:%v\n",left)
	if(left >= length){
		return -1
	}
	return left
}

func right(node int,length int) int {
	right := (node+1) << 1
	fmt.Printf("right:%v\n",right)
	if(right >= length) {
		return -1
	}
	return right
}

