package main

import "fmt"

func changeNum(num *int) { // reciving address of num which is integer
	*num = 11
}

func changeList(list *[]int) { // reciving address of list which is slice of integer
	*list = append(*list, 4)
}

func main() {
	num := 10
	changeNum(&num) // passing address of the num variable
	fmt.Println(num)

	list := []int{1, 2, 3}
	changeList(&list) // passing address of the list slice
	fmt.Println(list)

}
