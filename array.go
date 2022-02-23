// package main

// import "fmt"

// func main() {
// 	var a [5]int
// 	fmt.Println("emp:", a)

// 	a[4] = 100
// 	fmt.Println("set:", a)
// 	fmt.Println("get:", a[4])

// 	fmt.Println("len:", len(a))

// 	b := [5]int{1, 2, 3, 4,5}
// 	fmt.Println("dcl: ", b)

// 	// [2]jumlah array [3]isi dalam array
// 	var twoD [2][3]int
// 	for i := 0; i < 2;i++ {
// 		for j := 0; j < 3;j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)

// 	var fruits = [4]string{"apple","grape","orange","melon"}

// 	// for i := 0; i < len(fruits); i++ {
// 	// 	fmt.Printf("Element %d : %s\n", i,fruits[i])
// 	// }

// 	for _, fruit := range fruits {
// 		fmt.Printf("Element : %s\n",fruit)
// 	}
// }