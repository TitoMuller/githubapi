package github

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	HtmlUrl   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

// Função para extrair dados de um usuário
func ExtractUserData(username string) *User {
	userURL := fmt.Sprintf("https://api.github.com/users/%s", username)
	req, err := http.NewRequest("GET", userURL, nil)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
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
		fmt.Println("Erro ao fazer requisicao:", err)
		return nil
	}
	defer resp.Body.Close()

	var userData User
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return nil
	}

	err = json.Unmarshal(body, &userData)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return nil
	}

	return &userData
}

// Extrair multiplos users
func ExtractMultipleUsers(usernames []string) []*User {
	var users []*User
	for _, username := range usernames {
		userData := ExtractUserData(username)
		if userData != nil {
			users = append(users, userData)
		}
	}
	return users
}

// Converte os dados para CSV
func ConvertUsersToCSV(jsonData []byte, writer *csv.Writer) error {
	var users []*User
	err := json.Unmarshal(jsonData, &users)
	if err != nil {
		return fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	writer.Write([]string{"Login", "ID", "URL", "Avatar URL"})

	for _, user := range users {
		writer.Write([]string{user.Login, fmt.Sprint(user.ID), user.HtmlUrl, user.AvatarURL})
	}

	return nil
}
