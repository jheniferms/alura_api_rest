package models

import "github.com/jheniferms/alura_api_rest/database"

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade

func GetAll() []Personalidade {
	var personalidades []Personalidade

	database.DB.Find(&personalidades)

	return personalidades
}

func Get(id string) Personalidade {
	var personalidade Personalidade

	database.DB.First(&personalidade, id)

	return personalidade
}

func Create(novaPersonalidade Personalidade) Personalidade {
	database.DB.Create(&novaPersonalidade)

	return novaPersonalidade
}

func Delete(id string) Personalidade {
	var personalidade Personalidade

	database.DB.Delete(&personalidade, id)

	return personalidade
}

func Update(personalidade Personalidade) {
	database.DB.Save(&personalidade)
}
