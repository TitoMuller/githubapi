package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Estrutura para dados de uma organização
type Organization struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func ExtractOrgData(orgname string) *Organization {
	orgURL := fmt.Sprintf("https://api.github.com/orgs/%s", orgname)
	resp, err := http.Get(orgURL)
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
