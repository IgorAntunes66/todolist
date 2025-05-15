package todo

import (
	"fmt"
	"log"
	"slices"
)

func add(tarefa string, arquivo string) error {
	lista, err := jsonToSlice(arquivo)
	if err != nil {
		return fmt.Errorf("Erro ao converter o json para slice: %w", err)
	}

	// Adicionar novo item
	lista = append(lista, tarefa)

	//Codificar a slice atualizada
	err = sliceToJson(lista, arquivo)
	if err != nil {
		return fmt.Errorf("Erro ao converter a lice para json: %w", err)
	}
	return nil
}

func list(arquivo string) error {
	var lista []string
	var err error
	if lista, err = jsonToSlice(arquivo); err != nil {
		return fmt.Errorf("Erro ao converter json para uma slice: %w", err)
	}

	printList("TO-DO", lista)
	return nil
}

func finished(index int, arquivoLista, arquivoListaFinalizada string) error {
	lista, err := jsonToSlice(arquivoLista)
	if err != nil {
		return fmt.Errorf("Erro ao converter json para uma slice: %w", err)
	}

	if len(lista) < 1 {
		return fmt.Errorf("Não a tarefas para serem concluidas")
	}

	if index > len(lista) {
		return fmt.Errorf("Passe um index valido!")
	}

	listaFin, err := jsonToSlice(arquivoListaFinalizada)
	if err != nil {
		return fmt.Errorf("Erro ao converter o json para uma slice: %w", err)
	}

	// 14/05/25 parei aqui
	listaFin = append(listaFin, lista[index-1])
	err = sliceToJson(listaFin, FINALIZADAS)
	if err != nil {
		return fmt.Errorf("Erro ao converter a slice para um json: %w", err)
	}

	lista = slices.Delete(lista, index-1, index)
	err = sliceToJson(lista, LISTA)
	if err != nil {
		return fmt.Errorf("Erro ao converter a slice para um json: %w", err)
	}

	return nil
}

func printList(texto string, lista []string) {
	fmt.Printf("------ %s ------\n", texto)
	for i := 0; i < len(lista); i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}

func listFin(arquivo string) {
	var lista []string
	var err error
	if lista, err = jsonToSlice(arquivo); err != nil {
		log.Fatal(err)
	}

	printList("FINALIZADAS", lista)
}

func cancel(index int, arquivo string) {
	var lista []string
	var err error
	lista, err = jsonToSlice(arquivo)

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
	sliceToJson(lista, LISTA)
}
