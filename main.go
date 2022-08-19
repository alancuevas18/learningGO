package main

import (
	"fmt"
	"log"
	"strconv"
)

// functions
func funcion(message string) {
	fmt.Println(message)
}
func returnValue(a int) int {
	return a + 2
}

func multipleArguments(a int) (c, b int) {
	return a, a * 2
}
func main() {
	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	//Variables
	base := 12 //: para declarar e asignar
	var altura int = 14
	var area int
	area = base * altura
	fmt.Println(area)

	//Zero Values valores nuevos tienen un valor
	var a int     //0
	var b float64 //0
	var c string  //sting vacio
	var d bool    //false
	fmt.Println(a, b, c, d)

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//opeardores
	num1 := 10
	num2 := 15

	//suma
	result := num1 + num2
	fmt.Println("Suma ", result)

	//Resta
	result = num1 - num2
	fmt.Println("Resta ", result)

	//Multiplicacion
	result = num1 * num2
	fmt.Println("Multiplicacion ", result)

	//Division
	result = num1 / num2
	fmt.Println("Division ", result)

	//Increment
	num1++
	fmt.Println(num1)

	//funciones
	funcion("Impimir desde funcion")
	fmt.Println(returnValue(4))

	return1, _ := multipleArguments(4)
	fmt.Println(return1)
	fmt.Println(multipleArguments(4))

	//Ciclos
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Condicionales
	value1 := 1

	if value1 == 1 {
		fmt.Println("Es uno")
	}

	//Convertir texto a numero strconv return value, error
	value, error := strconv.Atoi("43")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Value : ", value)
}
