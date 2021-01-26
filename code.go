package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GenericInterface interface {
	DeliversTo(string) bool
}

type Address struct {
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	FirstLine  string `json:"first_line"`
	SecondLine string `json:"second_line"`
}

type Seller struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func (seller Seller) DeliversTo(city string) bool {
	return city == seller.Address.City
}

func HasAnySellersFromCity(sellers []Seller, city string) {
	city = city

	for i := range sellers {
		if sellers[i].Address.City == city {
			fmt.Printf("Found seller %s in %s city", sellers[i].Name, city)
		}
		break
	}
}

type Product struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Seller      Seller `json:"seller"`
}

func (product Product) DeliversTo(city string) bool {
	return product.Seller.DeliversTo(city)
}

func NewProduct(name string, price int, description string, seller Seller) Product {
	return Product{
		Name:        name,
		Price:       price,
		Description: description,
		Seller:      seller,
	}
}

func (product Product) Update(updatedProduct Product) {
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Description = updatedProduct.Description
	product.Seller = updatedProduct.Seller
}

func LoadProducts(jsonPath string) ([]Product, error) {
	productBytes, err := ioutil.ReadFile(jsonPath)
	products := []Product{}
	err = json.Unmarshal(productBytes, &products)

	if err != nil {
		fmt.Println(err)
		return products, err
	}

	return products, nil
}

func WriteProducts(productsSold []Product, productsLeft []Product, jsonPath string) error {
	allProducts := []Product{}

	allProducts = append(allProducts, productsSold...)

	for i := range productsLeft {
		productsLeft = append(allProducts, productsLeft[i])
	}

	fmt.Println(allProducts[:])

	return nil
}
