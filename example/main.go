package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/deitrix/borger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbPass = "1234"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "car_company"
)

type Car struct {
	ID          string
	Make        string
	Model       string
	Description string
}

type Price struct {
	CarID   string
	Mileage int
	Term    int
	Price   int
}

var Cars = borger.Table[Car]{
	Name: "cars",
	Columns: func(car *Car) map[string]any {
		return map[string]any{
			"id":          &car.ID,
			"make":        &car.Make,
			"model":       &car.Model,
			"description": &car.Description,
		}
	},
}

var Prices = borger.Table[Price]{
	Name: "pricing",
	Columns: func(price *Price) map[string]any {
		return map[string]any{
			"car_id":  &price.CarID,
			"mileage": &price.Mileage,
			"term":    &price.Term,
			"price":   &price.Price,
		}
	},
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	borger.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	cars, err := Cars.Select("id", "make", "model")
	if err != nil {
		log.Fatalf("failed to select cars: %v", err)
	}

	for _, car := range cars {
		fmt.Printf("%+v\n", car)
	}
}
