package main

import (
	"fmt"
	"os"
	"strconv"
	"todolist/todo"
)

func main() {
	switch os.Args[1] {
	case "add":
		if len(os.Args) == 3 {
			if !todo.IsLetter(os.Args[2]) {
				fmt.Println("A tarefa precisa ser apenas letras.")
				os.Exit(1)
			}
			todo.Add(os.Args[2], todo.LISTA)
		} else {
			err := fmt.Errorf("uso: %s add <tarefa>", os.Args[0])
			fmt.Println(err)
		}
	case "list":
		if len(os.Args) == 2 {
			todo.List(todo.LISTA)
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

		if err := todo.Finished(index, todo.LISTA, todo.FINALIZADAS); err != nil {
			err = fmt.Errorf("falha ao concluir tarefa: %w", err)
			fmt.Println(err)
		}
	case "finalizadas":
		if len(os.Args) == 2 {
			todo.ListFin(todo.FINALIZADAS)
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
			todo.Cancel(index, todo.LISTA)
		} else {
			fmt.Println("Para utilizar a função \"remover\" deve ser passado apenas um index da tarefa a ser concluida e sem espaços! Ex: ./todolist.go remover \"index\"")
			os.Exit(1)
		}
	}
}
