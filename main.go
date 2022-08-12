package main

import "fmt"

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
}
