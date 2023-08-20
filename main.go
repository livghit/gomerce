package main

import (
	"encoding/json"
	"log"

	shopify "github.com/bold-commerce/go-shopify/v3"
)

func main() {
	app := shopify.App{
		ApiKey:      "load from env",
		ApiSecret:   "load from env",
		RedirectUrl: "load from env",
		Scope:       "read_products",
	}

	client := shopify.NewClient(app, "load from env", "load from env")

	// Fetch the number of products.
	products, err := client.Product.List(nil)
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
