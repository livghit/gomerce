package main

import (
	"encoding/json"
	"log"
	"os"

	shopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_key := os.Getenv("API_KEY")
	api_secret := os.Getenv("API_SECRET")
	redirect_url := os.Getenv("REDIRECT_URL")
	admin_token := os.Getenv("ADMIN_TOKEN")
	shop_name := os.Getenv("SHOP_NAME")

	app := shopify.App{
		ApiKey:      api_key,
		ApiSecret:   api_secret,
		RedirectUrl: redirect_url,
		Scope:       "read_products , read_customers",
	}

	client := shopify.NewClient(app, shop_name, admin_token)

	// Fetch the number of products.
	products, err := client.Customer.List(nil)
	if err != nil {
		log.Printf("Error happended while fatching products:  %d", err)
	}

	log.Printf(ToJson(products))
}

func ToJson(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
