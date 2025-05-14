package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

const LISTA = "lista.json"
const FINALIZADAS = "listaFinalizadas.json"

func main() {
	switch os.Args[1] {
	case "add":
		if len(os.Args) == 3 {
			if !isLetter(os.Args[2]) {
				fmt.Println("A tarefa precisa ser apenas letras.")
				os.Exit(1)
			}
			add(os.Args[2], LISTA)
		} else {
			fmt.Println("Para utilizar a função \"add\" deve ser passado apenas uma tarefa para ser adicionada e sem espaços! Ex: ./todolist.go add \"tarefa\"")
			os.Exit(1)
		}
	case "list":
		if len(os.Args) == 2 {
			list(LISTA)
		} else {
			fmt.Println("Para utilizar a função \"list\" não deve ser passado nada antes e nem apos do comando! Ex: ./todolist.go list")
			os.Exit(1)
		}
	case "concluir":
		if len(os.Args) == 3 {
			index, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Erro ao converter parametro")
				os.Exit(1)
			}
			finished(index, "lista.json", "listaFinalizadas.json")
		} else {
			fmt.Println("Para utilizar a função \"concluir\" deve ser passado apenas um index da tarefa a ser concluida e sem espaços! Ex: ./todolist.go concluir \"index\"")
			os.Exit(1)
		}
	case "finalizadas":
		if len(os.Args) == 2 {
			listFin(FINALIZADAS)
		} else {
			fmt.Println("Para utilizar a função \"finalizadas\" não deve ser passado nada antes e nem apos do comando! Ex: ./todolist.go list")
			os.Exit(1)
		}
	case "remover":
		if len(os.Args) == 3 {
			index, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Erro ao converter parametro")
				return
			}
			cancel(index, LISTA)
		} else {
			fmt.Println("Para utilizar a função \"remover\" deve ser passado apenas um index da tarefa a ser concluida e sem espaços! Ex: ./todolist.go remover \"index\"")
			os.Exit(1)
		}
	}
}

func add(tarefa string, arquivo string) error {
	lista, err := jsonToSlice(arquivo)
	if err != nil {
		return err
	}

	// Adicionar novo item
	lista = append(lista, tarefa)

	//Codificar a slice atualizada
	return sliceToJson(lista, arquivo)
}

func list(arquivo string) {
	var lista []string
	var err error
	if lista, err = jsonToSlice(arquivo); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------ TO-DO ------")
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}

func finished(index int, arquivoLista, arquivoListaFinalizada string) {
	lista, err := jsonToSlice(arquivoLista)
	if err != nil {
		log.Fatal(err)
	}

	if len(lista) < 1 {
		log.Fatal("Nenhum item na lista de tarefa.")
	}

	if index < 1 || index > len(lista) {
		log.Fatal("Passe um index valido!")
	}

	listaFin, err := jsonToSlice(arquivoListaFinalizada)
	if err != nil {
		log.Fatal(err)
	}

	// 14/05/25 parei aqui
	listaFin = append(listaFin, lista[index-1])

	novaLista, err := json.MarshalIndent(listaFin, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(arquivoListaFinalizada, novaLista, 0644)

	lista = slices.Delete(lista, index-1, index)
	novoConteudo, err := json.MarshalIndent(lista, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(arquivoLista, novoConteudo, 0644)
}

func listFin(arquivo string) {
	var lista []string
	var err error
	if lista, err = jsonToSlice(arquivo); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------ FINALIZADAS ------")
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}

func cancel(index int, arquivo string) {
	var lista []string
	var err error
	lista, err = jsonToSlice(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	if len(lista) < 1 {
		log.Fatal("Nenhum item na lista de tarefa.")
	}

	lista = slices.Delete(lista, index-1, index)
	novoConteudo, err := json.MarshalIndent(lista, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(arquivo, novoConteudo, 0644)
}

func isLetter(texto string) bool {
	for i := range len(texto) {
		if int(texto[i]) >= 65 && int(texto[i]) <= 90 || int(texto[i]) >= 97 && int(texto[i]) <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}

func jsonToSlice(arquivo string) ([]string, error) {
	var lista []string
	if _, err := os.Stat(arquivo); err == nil {
		//Arquivo existe: Lê o conteudo
		conteudo, err := os.ReadFile(arquivo)
		if err != nil {
			return nil, err
		}

		//Decodificar o json para a slice
		if err := json.Unmarshal(conteudo, &lista); err != nil {
			return nil, err
		}
	} else if os.IsNotExist(err) {
		lista = []string{}
	} else {
		return nil, err
	}

	return lista, nil
}

func sliceToJson(slice []string, arquivo string) error {
	// Codificar a slice atualizada
	novoConteudo, err := json.MarshalIndent(slice, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(arquivo, novoConteudo, 0644)
}
