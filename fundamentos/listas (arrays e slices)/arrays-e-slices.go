package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	var slice = []int{10} // nao tem tamanho fixo
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array2)) // [5]string

	slice = append(slice, 40)
	fmt.Println(slice) // [10 40]

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)
}
