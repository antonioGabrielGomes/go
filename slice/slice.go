package slice

import "fmt"

func Slice() int {
	slice := []int{3, 2, 6, 4, 8}
	index := 2

	sliceInit := slice[:index]
	sliceEnd := slice[index+1:]

	fmt.Println(sliceInit)
	fmt.Println(sliceEnd)

	return 1
}
