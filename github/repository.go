package github

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Struct pra poder fazer o loop
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

// Função para extrair dados de um repositorio
func ExtractRepoData(owner, repo string) *Repository {
	repoURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequest("GET", repoURL, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return nil
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Token não encontrado")
		return nil
	}

	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
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

// Extrair multiplos repostirios
func ExtractMultipleRepos(repositories []RepositoryParameters) []*Repository {
	var repos []*Repository
	for _, repository := range repositories {
		repoData := ExtractRepoData(repository.Owner, repository.Repo)
		if repoData != nil {
			repos = append(repos, repoData)
		}
	}
	return repos
}

// Converte os dados para CSV
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
