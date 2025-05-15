package todo

import (
	"encoding/json"
	"os"
)

func JsonToSlice(arquivo string) ([]string, error) {
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

func SliceToJson(slice []string, arquivo string) error {
	// Codificar a slice atualizada
	novoConteudo, err := json.MarshalIndent(slice, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(arquivo, novoConteudo, 0644)
}
