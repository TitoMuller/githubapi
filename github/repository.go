package github

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct pra poder fazer o loop no main.go
type RepositoryParameters struct {
	Owner string
	Repo  string
}

// Estrutura para dados de um repositório
type Repository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FullName    string `json:"full_name"`
	HtmlUrl     string `json:"html_url"`
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

func ConvertRepositoriesToCSV(jsonData []byte, writer *csv.Writer) error {
	var repos []*Repository
	err := json.Unmarshal(jsonData, &repos)
	if err != nil {
		return fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	writer.Write([]string{"ID", "Name", "Description", "Full name", "URL"})

	for _, repo := range repos {
		writer.Write([]string{fmt.Sprint(repo.ID), repo.Name, repo.Description, repo.FullName, repo.HtmlUrl})
	}

	return nil
}
