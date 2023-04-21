package main

import "fmt"

//ToLower pone todo en minuscula
//ToUpper pone todo en mayuscula
/*func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])

	}

	if text == textReverse {
		fmt.Println("es palindromo")
	} else {
		fmt.Println("no es un palindromo")
	}

}*/

type car struct {
	brand string
	year  int
}

func main() {
	/*slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)

	}
	//ama
	//amor a roma

	isPalindromo("Ama")
	*/
	/*m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	//recirer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["jose"]
	fmt.Println(value, ok)
	*/
	mycar := car{brand: "ford", year: 2023}
	fmt.Println(mycar)

	//Otra manera
	var othercar car
	othercar.brand = "ferrari"
	fmt.Println(othercar)

}
