package main

import (
	"fmt"
	"log"
	"strconv"
)

//Solo existe el ciclo For recomendacion es poder sacar de un ciclo

func main() {

	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println((i))
	}

	//for While
	counter := 10

	for counter < 20 {

		fmt.Println(counter)
		counter++

	}

	//for forever

	/*counterForever := 0
	for {

		fmt.Println(counterForever)
		counterForever++

	}*/

	valor1 := 1
	// valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	//Convertir texyo a numero
	valu, err := strconv.Atoi("asd")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("value:", valu)

}
