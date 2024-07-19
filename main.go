package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TitoMuller/githubapi.git/github"
)

func main() {
	// Extrair dados de um usuario. Parametro = username
	usernames := []string{
		"TitoMuller", "rbmuller", "octocat", "mojombo", "teste", "ErickWendel", "ratrock", "ComunidadeDevSpace",
	}
	users := github.ExtractMultipleUsers(usernames)
	saveJSON("user_data.json", users)
	convertJSONToCSV("user_data.json", "user_data.csv", github.ConvertUsersToCSV)

	// Extrair dados de organizacao. Parametro = orgname
	organizations := []string{
		"github", "valor-labs", "oldschoolvalue", "facebook", "ComunidadeDevSpace",
	}
	orgs := github.ExtractMultipleOrgs(organizations)
	saveJSON("org_data.json", orgs)
	convertJSONToCSV("org_data.json", "org_data.csv", github.ConvertOrganizationsToCSV)

	// Extrair dados de repositorio, parametro = username, repo
	repositories := []github.RepositoryParameters{
		{Owner: "TitoMuller", Repo: "golangulator"},
		{Owner: "rbmuller", Repo: "devTools"},
		{Owner: "ErickWendel", Repo: "fingerpose"},
		{Owner: "ErickWendel", Repo: "myownnode"},
		{Owner: "TitoMuller", Repo: "CalculadoraIMC"},
		{Owner: "TitoMuller", Repo: "calculatorJS"},
		{Owner: "rbmuller", Repo: "awsGENAI"},
	}
	repos := github.ExtractMultipleRepos(repositories)
	saveJSON("repo_data.json", repos)
	convertJSONToCSV("repo_data.json", "repo_data.csv", github.ConvertRepositoriesToCSV)

	fmt.Println("Dados extraídos e salvos localmente.")
}

// Função para salvar dados em um arquivo JSON localmente
func saveJSON(filename string, data interface{}) {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar dados:", err)
		return
	}
	err = os.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}

	fmt.Printf("Dados salvos em %s\n", filename)
}

// Funcao para converter arquivos JSON em arquivos CSV
func convertJSONToCSV(jsonFilename, csvFilename string, convertFunc func([]byte, *csv.Writer) error) {
	jsonFile, err := os.ReadFile(jsonFilename)
	if err != nil {
		fmt.Println("Erro ao ler arquivo JSON:", err)
		return
	}

	csvFile, err := os.Create(csvFilename)
	if err != nil {
		fmt.Println("Erro ao criar arquivo CSV:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	err = convertFunc(jsonFile, writer)
	if err != nil {
		fmt.Println("Erro ao converter dados:", err)
		return
	}

	fmt.Printf("Dados convertidos e salvos em %s\n", csvFilename)
}
