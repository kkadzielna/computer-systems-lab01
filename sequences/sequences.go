package main
import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	for i := range slice {
		slice = append(slice, slice[i])
	}
}

func mapSlice(f func(a int) int, slice []int) {
	for i := range slice {
		slice[i] = f (slice[i])
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i := range array {
		array[i] = f (array[i])
	}
}

func main() {
	intsSlice := make([]int, 5)
	intsSlice[0] = 1
	intsSlice[1] = 2
	intsSlice[2] = 3
	intsSlice[3] = 4
	intsSlice[4] = 5
	intsArray := [3]int{1, 2, 3}
	

	mapSlice(addOne,intsSlice)
	fmt.Println("mapSlice: ", intsSlice )
	mapArray(addOne,intsArray) //doesn't work idk why :((
	newSlice := intsSlice[1:3]
	fmt.Println("newSlice: ", newSlice )
	double(newSlice) 
	fmt.Println("doubleSlice: ", intsSlice)
	fmt.Println("mapArray: ", intsArray )
	
}