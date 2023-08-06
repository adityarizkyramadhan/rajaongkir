package main

import (
	"fmt"
	"os"

	"github.com/adityarizkyramadhan/rajaongkir"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
		return
	}
	roStarter := rajaongkir.New(os.Getenv("RAJA_ONGKIR_KEY")).Starter()
	provinces, err := roStarter.GetProvince()
	fmt.Println("Province : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(provinces)
	cities, err := roStarter.GetCityByIDProvince(provinces[10].ProvinceID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("City : ")
	fmt.Println(cities)

	costs, err := roStarter.GetCost(rajaongkir.CostInput{
		Origin:      "179",
		Destination: "343",
		Weight:      1700,
		Courier:     "jne",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(costs)

}
