package service

import (
	"log"
	"time"

	"github.com/GabriellGds/petshop-dti/internal/domain"
	"github.com/GabriellGds/petshop-dti/internal/model"
	
)


func ListPetshops(req *model.Request) (*domain.Petshop, float64, error) {
	date, err := time.Parse("2006/01/02", req.Date)
	if err != nil {
		return nil, 0, &model.ErrorResponse{Message: "The date format must follow the pattern MM-DD-YYYY."}
	}
	isWeekend := date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
	
	log.Print(date)
	log.Print(isWeekend)

	var bestPetshop domain.Petshop
	var bestPrice float64 = 1e9

	for _, petshop := range domain.Petshops {
		var price float64
		if isWeekend {
			price = float64(req.SmallDogs) * petshop.WeekendPriceSmallDog + float64(req.BigDogs) * petshop.WeekendPriceBigDog
		} else {
			price = float64(req.SmallDogs) * petshop.WeekdayPriceSmallDog + float64(req.BigDogs) * petshop.WeekdayPriceBigDog
		}

		if price < bestPrice || (price == bestPrice && petshop.Distance < bestPetshop.Distance) {
			bestPetshop = petshop
			bestPrice = price
		}
	}

	return &bestPetshop, bestPrice, nil
}