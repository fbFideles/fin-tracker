package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type database struct {
	User *string `json:"user"`
	Name *string `json:"name"`
	SSL  *string `json:"ssl"`
}

// Config é uma estrutura que define
// o retorno de um arquivo de configuração
type Config struct {
	Database database `json:"database"`
}

// GetConfig é uma função
func GetConfig() (configuracao Config) {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(byteValue, &configuracao)

	return
}
