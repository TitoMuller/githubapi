package github

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Estrutura para dados de uma organização
type Organization struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Repos       string `json:"repos_url"`
}

// Função para extrair dados de uma organizacao
func ExtractOrgData(orgname string) *Organization {
	orgURL := fmt.Sprintf("https://api.github.com/orgs/%s", orgname)
	req, err := http.NewRequest("GET", orgURL, nil)
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

	var orgData Organization
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return nil
	}

	err = json.Unmarshal(body, &orgData)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return nil
	}

	return &orgData
}

// Funcao para extrair multiplas organizacoes
func ExtractMultipleOrgs(organizations []string) []*Organization {
	var orgs []*Organization
	for _, organization := range organizations {
		orgData := ExtractOrgData(organization)
		if orgData != nil {
			orgs = append(orgs, orgData)
		}
	}
	return orgs
}

// Converte os dados para CSV
func ConvertOrganizationsToCSV(jsonData []byte, writer *csv.Writer) error {
	var orgs []*Organization
	err := json.Unmarshal(jsonData, &orgs)
	if err != nil {
		return fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	writer.Write([]string{"Login", "ID", "Description", "Repos"})

	for _, org := range orgs {
		writer.Write([]string{org.Login, fmt.Sprint(org.ID), org.Description, org.Repos})
	}

	return nil
}
