package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TitoMuller/githubapi.git/github"
)

func main() {
	// Criando um SLICE de strings chamado usernames
	usernames := []string{"TitoMuller", "rbmuller", "octocat", "mojombo"}

	// Cria um SLICE de "User" (estudar pointers)
	var users []*github.User

	// Cria um loop que itera sobre cada elemento do SLICE usernames
	// _ é usado como um placeholder para o índice (estudar)
	// Pra cada iteração, chama a funcao ExcractUserData e armazena o resultado na variavel userData
	// Se esse valor for diferente de NIL, adiciona o valor de userData ao SLICE users
	for _, username := range usernames {
		userData := github.ExtractUserData(username)
		if userData != nil {
			users = append(users, userData)
		}
	}

	// Chama a funcao saveJSON, criando o nome de arquivo definido na chamado
	// Passa como data o conteudo de users, que é o resultado do loop for
	saveJSON("user_data.json", users)

	// Extrair dados de organizacao, parametro = orgname
	orgData := github.ExtractOrgData("github")
	saveJSON("org_data.json", orgData)

	// Extrair dados de repositorio, parametro = username, repo
	repoData := github.ExtractRepoData("TitoMuller", "golangulator")
	saveJSON("repo_data.json", repoData)

	fmt.Println("Dados extraídos e salvos localmente.")
}

// Função para salvar dados em um arquivo JSON localmente
func saveJSON(filename string, data interface{}) {

	// dataBytes é o nome da variavel que armazena os dados em formato JSON
	// err é a variavel que armazenara qualquer erro que possa ocorrer
	// MarshalIndent é uma funcao do package encoding/json que serializa e formata os dados JSON
	// data sao os dados a serem formatados
	dataBytes, err := json.MarshalIndent(data, "", "  ")

	// Verifica se houve um erro na formatacao
	if err != nil {
		fmt.Println("Erro ao serializar dados:", err)
		return
	}

	// Funcao que escreve os dados no arquivo
	// WriteFile é a funcao do package "os" que escreve os dados
	// filename é o nome do arquivo
	// dataBytes, nome da variavel que armazenou os dados
	// 0644 representa as permissoes de leitura, escrita e execucao (estudar)
	err = os.WriteFile(filename, dataBytes, 0644)

	// verifica se houve erro na escrita
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}

	// log pra conferir se deu certo
	fmt.Printf("Dados salvos em %s\n", filename)
}
