package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Food struct
type Food struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
}

// Food slices with their ID, name, country, type, and price.
var foods = []Food{
	// Africa
	{ID: "1", Name: "Couscous", Country: "Algeria", Type: "Main Course", Price: 14.50},
	{ID: "2", Name: "Koshari", Country: "Egypt", Type: "Main Course", Price: 8.00},
	{ID: "3", Name: "Doro Wat", Country: "Ethiopia", Type: "Stew", Price: 16.00},
	{ID: "4", Name: "Jollof Rice", Country: "Ghana", Type: "Main Course", Price: 12.50},
	{ID: "5", Name: "Ugali with Sukuma Wiki", Country: "Kenya", Type: "Main Course", Price: 9.50},
	{ID: "6", Name: "Tagine", Country: "Morocco", Type: "Stew", Price: 18.00},
	{ID: "7", Name: "Jollof Rice", Country: "Nigeria", Type: "Main Course", Price: 13.00},
	{ID: "8", Name: "Bobotie", Country: "South Africa", Type: "Casserole", Price: 15.50},
	// Antarctica (Symbolic)
	{ID: "9", Name: "Pemmican", Country: "Antarctica (Historical)", Type: "Survival Food", Price: 25.00},
	// Asia
	{ID: "10", Name: "Peking Duck", Country: "China", Type: "Main Course", Price: 45.00},
	{ID: "11", Name: "Biryani", Country: "India", Type: "Main Course", Price: 14.00},
	{ID: "12", Name: "Nasi Goreng", Country: "Indonesia", Type: "Main Course", Price: 7.50},
	{ID: "13", Name: "Sushi", Country: "Japan", Type: "Main Course", Price: 22.00},
	{ID: "14", Name: "Adobo", Country: "Philippines", Type: "Main Course", Price: 11.50},
	{ID: "15", Name: "Kimchi", Country: "South Korea", Type: "Side Dish", Price: 5.00},
	{ID: "16", Name: "Pad Thai", Country: "Thailand", Type: "Noodles", Price: 10.00},
	{ID: "17", Name: "Pho", Country: "Vietnam", Type: "Soup", Price: 9.00},
	// Europe
	{ID: "18", Name: "Pot-au-feu", Country: "France", Type: "Stew", Price: 25.00},
	{ID: "19", Name: "Sauerbraten", Country: "Germany", Type: "Main Course", Price: 19.50},
	{ID: "20", Name: "Moussaka", Country: "Greece", Type: "Casserole", Price: 16.50},
	{ID: "21", Name: "Goulash", Country: "Hungary", Type: "Stew", Price: 15.00},
	{ID: "22", Name: "Ragu alla Bolognese", Country: "Italy", Type: "Pasta", Price: 17.00},
	{ID: "23", Name: "Pierogi", Country: "Poland", Type: "Dumpling", Price: 12.00},
	{ID: "24", Name: "Paella", Country: "Spain", Type: "Main Course", Price: 24.00},
	{ID: "25", Name: "Fish and Chips", Country: "United Kingdom", Type: "Main Course", Price: 13.50},
	// North America
	{ID: "26", Name: "Poutine", Country: "Canada", Type: "Side Dish", Price: 9.50},
	{ID: "27", Name: "Ropa Vieja", Country: "Cuba", Type: "Main Course", Price: 18.50},
	{ID: "28", Name: "Ackee and Saltfish", Country: "Jamaica", Type: "Main Course", Price: 20.00},
	{ID: "29", Name: "Mole Poblano", Country: "Mexico", Type: "Main Course", Price: 19.00},
	{ID: "30", Name: "Hamburger", Country: "United States", Type: "Main Course", Price: 11.00},
	// Oceania
	{ID: "31", Name: "Meat Pie", Country: "Australia", Type: "Snack", Price: 6.50},
	{ID: "32", Name: "Kokoda", Country: "Fiji", Type: "Appetizer", Price: 14.00},
	{ID: "33", Name: "Hāngi", Country: "New Zealand", Type: "Main Course", Price: 30.00},
	{ID: "34", Name: "Mumu", Country: "Papua New Guinea", Type: "Main Course", Price: 28.00},
	{ID: "35", Name: "Palusami", Country: "Samoa", Type: "Side Dish", Price: 8.50},
	// South America
	{ID: "36", Name: "Asado", Country: "Argentina", Type: "Main Course", Price: 35.00},
	{ID: "37", Name: "Feijoada", Country: "Brazil", Type: "Stew", Price: 21.00},
	{ID: "38", Name: "Pastel de Choclo", Country: "Chile", Type: "Casserole", Price: 16.00},
	{ID: "39", Name: "Bandeja Paisa", Country: "Colombia", Type: "Main Course", Price: 20.00},
	{ID: "40", Name: "Ceviche", Country: "Peru", Type: "Appetizer", Price: 17.50},
	{ID: "41", Name: "Arepa", Country: "Venezuela", Type: "Snack", Price: 7.00},
}

// APIs function.

func getFoodType(c *gin.Context) {
	typeQuery := c.Query("type")
	if typeQuery != "" {
		filter := []Food{}
		for _, food := range foods {
			if food.Type == typeQuery {
				filter = append(filter, food)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func getFoodPrice(c *gin.Context) {
	priceQuery := c.Query("price")
	if priceQuery != "" {
		filter := []Food{}
		for _, food := range foods {
			if fmt.Sprint(food.Price) == priceQuery {
				filter = append(filter, food)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func getFoodName(c *gin.Context) {
	nameQuery := c.Query("name")
	if nameQuery != "" {
		filter := []Food{}
		for _, food := range foods {
			if food.Name == nameQuery {
				filter = append(filter, food)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func getFoodID(c *gin.Context) {
	idQuery := c.Query("id")
	if idQuery != "" {
		filter := []Food{}
		for _, food := range foods {
			if food.ID == idQuery {
				filter = append(filter, food)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func getFoodCountry(c *gin.Context) {
	countryQuery := c.Query("country")
	if countryQuery != "" {
		filter := []Food{}
		for _, food := range foods {
			if food.Country == countryQuery {
				filter = append(filter, food)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, foods)
}

// Calling APIs by path.

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Working Normally"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/foods/id", getFoodID)
		api.GET("/foods/country", getFoodCountry)
		api.GET("/foods/name", getFoodName)
		api.GET("/foods/price", getFoodPrice)
		api.GET("/foods/type", getFoodType)
	}

	fmt.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
