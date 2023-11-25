package model

import (
	"Nuevo_go/JCesarBat/Nuevo_go/db"
	"gorm.io/gorm"
)

type Comida struct {
	gorm.Model
	Nombre       string `json:"Nombre"`
	Ingrediente1 string `json:"Ingrediente1"`
	Ingrediente2 string `json:"Ingrediente2"`
}

func Reed(ID uint) (Comida, error) {
	var comida Comida

	result := db.DB.Where(ID).Find(&comida)

	if result.Error != nil {

		return comida, result.Error
	}

	return comida, nil
}

func Save(comida Comida) error {
	var guardar Comida
	guardar = comida
	result := db.DB.Create(&guardar)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
func Listar() ([]Comida, error) {
	var comidas []Comida
	result := db.DB.Find(&comidas)
	if result.Error != nil {
		return comidas, result.Error
	}
	return comidas, nil
}

func Update(ID uint, comida Comida) error {

	err := db.DB.Model(&Comida{}).Where(ID).Updates(Comida{
		Nombre:       comida.Nombre,
		Ingrediente1: comida.Ingrediente1,
		Ingrediente2: comida.Ingrediente2,
	})
	if err != nil {
		return err.Error
	}
	return nil
}
