package helpers

import (
	"log"

	validator "github.com/go-playground/validator/v10"
)

func IsValid(req interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	log.Println(req)
	if err != nil {
		return false, err
	}
	return true, nil
}
