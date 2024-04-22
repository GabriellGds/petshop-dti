package model

import "github.com/go-playground/validator/v10"

type Request struct {
	Date      string `json:"date" validate:"required"`
	SmallDogs int    `json:"smallDogs"`
	BigDogs   int    `json:"bigDogs"`
}

func (r *Request) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	
	return nil
}

type Response struct {
	Petshop    string  `json:"petshop"`
	TotalPrice float64 `json:"totalPrice"`
}


type ErrorResponse struct {
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}