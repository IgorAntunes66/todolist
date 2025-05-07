package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	if os.Args[1] == "add" {
		add(os.Args[2], "lista.json")
	}

	switch os.Args[1] {
	case "add":
		add(os.Args[2], "lista.json")
	case "list":
		list("lista.json")
	}
}

func add(tarefa string, arquivo string) error {
	var lista []string

	if _, err := os.Stat(arquivo); err == nil {
		//Arquivo existe: Lê o conteudo
		conteudo, err := os.ReadFile(arquivo)
		if err != nil {
			return err
		}
		// Decodificar o json para a slice
		if err := json.Unmarshal(conteudo, &lista); err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		// Arquivo não existe: cria uma slice vazia
		lista = []string{}
	} else {
		// Outro erro (ex: permissão)
		return err
	}

	// Adicionar novo item
	lista = append(lista, tarefa)

	fmt.Println(lista)

	//Codificar a slice atualizada
	novoConteudo, err := json.MarshalIndent(lista, "", " ")
	if err != nil {
		return err
	}

	//Escrever no arquivo
	return os.WriteFile(arquivo, novoConteudo, 0644)
}

func list(arquivo string) {
	var lista []string
	if _, err := os.Stat(arquivo); err == nil {
		conteudo, err := os.ReadFile(arquivo)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(conteudo, &lista); err != nil {
			log.Fatal(err)
		} else if os.IsNotExist(err) {
			fmt.Printf("Nenhum item adicionado na lista!")
			log.Fatal(err)
		}
	}

	fmt.Println("------ TO-DO ------")
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}
