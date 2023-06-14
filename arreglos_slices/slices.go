package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{5, 8, 7, 9, 3, 4, 56, 76, 45, 1}

func MuestraSlice() {
	fmt.Println(tablaS)
	//Create partial portions slices from array
	porcion01 := arreglo[3:]  //Desde la posicion 3 al final
	porcion02 := arreglo[:5]  //Desde el inicio hasta la posicon 5
	porcion03 := arreglo[6:7] //De la posicion 6 a la 7

	fmt.Println(porcion01)
	fmt.Println(porcion02)
	fmt.Println(porcion03)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	//We declarate a empty slice
	nums := make([]int, 0, 0)
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Printf("Large %d - Increse Capacity %d \n", len(nums), cap(nums))
	}

	fmt.Printf("Slice : Elementos: Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

}
