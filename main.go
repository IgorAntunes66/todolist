package main

import (
	"fmt"
	"os"
	"strconv"
	"todolist/todo"
)

const LISTA = "lista.json"
const FINALIZADAS = "listaFinalizadas.json"

func main() {
	todo.
	switch os.Args[1] {
	case "add":
		if len(os.Args) == 3 {
			if !isLetter(os.Args[2]) {
				fmt.Println("A tarefa precisa ser apenas letras.")
				os.Exit(1)
			}
			add(os.Args[2], LISTA)
		} else {
			err := fmt.Errorf("uso: %s add <tarefa>", os.Args[0])
			fmt.Println(err)
		}
	case "list":
		if len(os.Args) == 2 {
			list(LISTA)
		} else {
			fmt.Println("Para utilizar a função \"list\" não deve ser passado nada antes e nem apos do comando! Ex: ./todolist.go list")
			os.Exit(1)
		}
	case "concluir":
		if len(os.Args) != 3 {
			err := fmt.Errorf("uso: %s concluir <indice>", os.Args[0])
			fmt.Println(err)
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			err := fmt.Errorf("indice invalido '%s' : %w", os.Args[2], err)
			fmt.Println(err)
		}

		if err := finished(index, LISTA, FINALIZADAS); err != nil {
			err = fmt.Errorf("falha ao concluir tarefa: %w", err)
			fmt.Println(err)
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
