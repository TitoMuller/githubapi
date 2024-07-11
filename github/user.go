package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
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
