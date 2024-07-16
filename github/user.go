package github

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	HtmlUrl   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

// Função para extrair dados de um usuário
// username string = parametro       -      *User = retorna um ponteiro para uma struct User
func ExtractUserData(username string) *User {
	userURL := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(userURL)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return nil
	}

	defer resp.Body.Close()

	// Cria uma variavel do tipo User
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
