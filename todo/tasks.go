package todo

import (
	"fmt"
	"log"
	"slices"
)

func Add(tarefa string, arquivo string) error {
	lista, err := JsonToSlice(arquivo)
	if err != nil {
		return fmt.Errorf("erro ao converter o json para slice: %w", err)
	}

	task := Task{
		ID:        len(lista) + 1,
		Tarefa:    tarefa,
		Completed: false,
	}

	// Adicionar novo item
	lista = append(lista, task)

	//Codificar a slice atualizada
	err = SliceToJson(lista, arquivo)
	if err != nil {
		return fmt.Errorf("erro ao converter a slice para json: %w", err)
	}
	return nil
}

func List(arquivo string) error {
	var lista []Task
	var err error
	if lista, err = JsonToSlice(arquivo); err != nil {
		return fmt.Errorf("erro ao converter json para uma slice: %w", err)
	}

	PrintList("TO-DO", lista)
	return nil
}

func Finished(ID int, arquivoLista string) error {
	lista, err := JsonToSlice(arquivoLista)
	if err != nil {
		return fmt.Errorf("erro ao converter json para slice: %w", err)
	}

	for i := 0; i < len(lista); i++ {
		if lista[i].ID == ID {
			lista[i].Completed = true
			return nil
		}
	}

	return fmt.Errorf("erro ao concluir tarefa")
}

func ListFin(arquivo string) {
	var lista []Task
	var err error
	if lista, err = JsonToSlice(arquivo); err != nil {
		log.Fatal(err)
	}

	PrintList("FINALIZADAS", lista)
}

func Cancel(index int, arquivo string) {
	var lista []Task
	var err error
	lista, err = JsonToSlice(arquivo)

	if index < 1 || index > len(lista) {
		log.Fatal("Passe um index valido!")
	}

	if err != nil {
		log.Fatal(err)
	}

	if len(lista) < 1 {
		log.Fatal("Nenhum item na lista de tarefa.")
	}

	lista = slices.Delete(lista, index-1, index)
	SliceToJson(lista, LISTA)
}
