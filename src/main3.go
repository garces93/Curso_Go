package main

import "fmt"

func main() {
	/*modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("es par", modulo)
	default:
		fmt.Println("es impar", modulo)

	}

	// sin condicion

	value := 200
	switch {
	case value > 100:
		fmt.Println("es mayor a 100")
	case value < 0:
		fmt.Println("es menor a 0")
	default:
		fmt.Println("no condicional")

	}
	*/

	//vamos a ver los KeyWords
	// defer va a ejecutar la ultima funcion antes de cualquier cosa
	// se puede utilizar para cerrar una base de datos o archivos

	/*defer fmt.Println("hola")
	fmt.Println("mundo")

	//continue y break
	// se utiliza cuando es un error que no afecta para nada el codigo

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break

		}
	}
	*/
	//Arry
	//len dice cuantos elementos hay en un array cap me dice la
	//capacidad maxima, son imutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice igual que un array pero sin declarar la cantidad son mutables
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el Slice 4: inclucivo :4 exclisivo
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append fila
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append columna
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
