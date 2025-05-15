package todo

import (
	"fmt"
	"log"
	"slices"
)

func Add(tarefa string, arquivo string) error {
	lista, err := JsonToSlice(arquivo)
	if err != nil {
		return fmt.Errorf("Erro ao converter o json para slice: %w", err)
	}

	// Adicionar novo item
	lista = append(lista, tarefa)

	//Codificar a slice atualizada
	err = SliceToJson(lista, arquivo)
	if err != nil {
		return fmt.Errorf("Erro ao converter a lice para json: %w", err)
	}
	return nil
}

func List(arquivo string) error {
	var lista []string
	var err error
	if lista, err = JsonToSlice(arquivo); err != nil {
		return fmt.Errorf("Erro ao converter json para uma slice: %w", err)
	}

	PrintList("TO-DO", lista)
	return nil
}

func Finished(index int, arquivoLista, arquivoListaFinalizada string) error {
	lista, err := JsonToSlice(arquivoLista)
	if err != nil {
		return fmt.Errorf("Erro ao converter json para uma slice: %w", err)
	}

	if len(lista) < 1 {
		return fmt.Errorf("Não a tarefas para serem concluidas")
	}

	if index > len(lista) {
		return fmt.Errorf("Passe um index valido!")
	}

	listaFin, err := JsonToSlice(arquivoListaFinalizada)
	if err != nil {
		return fmt.Errorf("Erro ao converter o json para uma slice: %w", err)
	}

	// 14/05/25 parei aqui
	listaFin = append(listaFin, lista[index-1])
	err = SliceToJson(listaFin, FINALIZADAS)
	if err != nil {
		return fmt.Errorf("Erro ao converter a slice para um json: %w", err)
	}

	lista = slices.Delete(lista, index-1, index)
	err = SliceToJson(lista, LISTA)
	if err != nil {
		return fmt.Errorf("Erro ao converter a slice para um json: %w", err)
	}

	return nil
}

func ListFin(arquivo string) {
	var lista []string
	var err error
	if lista, err = JsonToSlice(arquivo); err != nil {
		log.Fatal(err)
	}

	PrintList("FINALIZADAS", lista)
}

func Cancel(index int, arquivo string) {
	var lista []string
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
