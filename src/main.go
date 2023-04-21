package main

import "fmt"

func main() {
	//Declaracion de Constantes que jamas va a cambiar con el tiempo
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(pi, pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area Cuadrado
	const BaseCuadrado = 10
	areaCuadrado := BaseCuadrado * BaseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("suma", result)

	//resta
	result = y - x
	fmt.Println("resta", result)

	//multiplicacion
	result = x * y
	fmt.Println("multiplicacion", result)

	//division
	result = y / x
	fmt.Println("Divicion", result)

	//modulo si es 0 es exacta
	result = y % x
	fmt.Println("modulo", result)

	//Printf %s enteros %d numericos %v si no sabes el dato
	nombre := "Plazi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)

	//Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)

	fmt.Println(message)

	fmt.Printf("nombre : %T \n ", nombre)
	fmt.Printf("Cursos : %T \n ", cursos)
}
