package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Estrutura para dados de um repositório
type Repository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ExtractRepoData(owner, repo string) *Repository {
	repoURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	resp, err := http.Get(repoURL)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return nil
	}
	defer resp.Body.Close()

	var repoData Repository
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return nil
	}

	err = json.Unmarshal(body, &repoData)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return nil
	}

	return &repoData
}
