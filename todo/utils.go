package todo

import "fmt"

func IsLetter(texto string) bool {
	for i := range len(texto) {
		if int(texto[i]) >= 65 && int(texto[i]) <= 90 || int(texto[i]) >= 97 && int(texto[i]) <= 122 || int(texto[1]) == 32 {
			continue
		} else {
			return false
		}
	}
	return true
}

func PrintList(texto string, lista []string) {
	fmt.Printf("------ %s ------\n", texto)
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}
