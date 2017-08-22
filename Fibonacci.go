package main

import "fmt"

func fb(n int) []int{
	ch := make(chan int)

	go func() {
		x ,y := 0,1
		for i:= 0;i<n;i++ {
			ch <- x
			x, y = y, x + y

		}
		close(ch)

	}()
	results := make([]int,0)
	for i := range ch {
		results = append(results, i)
	}
	return results
}

func main()  {
	results := fb(10)
	fmt.Println(results)

}
