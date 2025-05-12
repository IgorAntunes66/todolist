package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	switch os.Args[1] {
	case "add":
		if len(os.Args) == 3 {
			add(os.Args[2], "lista.json")
		} else {
			fmt.Println("Para utilizar a função \"add\" deve ser passado apenas uma tarefa para ser adicionada e sem espaços! Ex: ./todolist.go add \"tarefa\"")
			os.Exit(1)
		}
	case "list":
		if len(os.Args) == 2 {
			list("lista.json")
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
			listFin("listaFinalizadas.json")
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
			cancel(index, "lista.json")
		} else {
			fmt.Println("Para utilizar a função \"remover\" deve ser passado apenas um index da tarefa a ser concluida e sem espaços! Ex: ./todolist.go remover \"index\"")
			os.Exit(1)
		}
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

func finished(index int, arquivoLista, arquivoListaFinalizada string) {
	var lista, listaFin []string
	if _, err := os.Stat(arquivoLista); err == nil {
		conteudo, err := os.ReadFile(arquivoLista)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(conteudo, &lista); err != nil {
			log.Fatal(err)
		} else if os.IsNotExist(err) {
			fmt.Println("Lista de tarefas não existe!")
			log.Fatal(err)
		}
	}
	fmt.Println(lista)

	if len(lista) < 1 {
		log.Fatal("Nenhum item na lista de tarefa.")
	}

	if index < 1 || index > len(lista) {
		log.Fatal("Passe um index valido!")
	}

	if _, err := os.Stat(arquivoListaFinalizada); err == nil {
		conteudoFin, err := os.ReadFile(arquivoListaFinalizada)

		if err != nil {
			log.Fatal("Erro ao ler arquivo da lista finalizada.")
		}

		if err := json.Unmarshal(conteudoFin, &listaFin); err != nil {
			log.Fatal(err)
		}
	}

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
	if _, err := os.Stat(arquivo); err == nil {
		conteudo, err := os.ReadFile(arquivo)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(conteudo, &lista); err != nil {
			log.Fatal(err)
		} else if os.IsNotExist(err) {
			fmt.Printf("Lista ainda não existe!")
			log.Fatal(err)
		}
	}

	fmt.Println("------ FINALIZADAS ------")
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}

func cancel(index int, arquivo string) {
	var lista []string
	if _, err := os.Stat(arquivo); err == nil {
		conteudo, err := os.ReadFile(arquivo)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(conteudo, &lista); err != nil {
			log.Fatal(err)
		} else if os.IsNotExist(err) {
			fmt.Println("Nenhum item adicionado na lista!")
			log.Fatal(err)
		}
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
