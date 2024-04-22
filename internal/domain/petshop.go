package domain

type Petshop struct {
	Name string
	Distance float64
	WeekdayPriceSmallDog float64
	WeekdayPriceBigDog float64
	WeekendPriceSmallDog float64
	WeekendPriceBigDog float64
}

var Petshops []Petshop = []Petshop{
	{
		Name: "Meu Canino Feliz", 
		Distance: 2, 
		WeekdayPriceSmallDog: 20, 
		WeekdayPriceBigDog: 40,
		WeekendPriceSmallDog: 20 * 1.2,
		WeekendPriceBigDog: 40 * 1.2,
	},
	{
		Name: "Vai Rex", 
		Distance: 1.7, 
		WeekdayPriceSmallDog: 15, 
		WeekdayPriceBigDog: 50, 
		WeekendPriceSmallDog: 20, 
		WeekendPriceBigDog: 55,
	},
	{
		Name: "ChowChawgas",
		Distance: 0.8, 
		WeekdayPriceSmallDog: 30, 
		WeekdayPriceBigDog: 45, 
		WeekendPriceSmallDog: 30, 
		WeekendPriceBigDog: 45,
	},
}