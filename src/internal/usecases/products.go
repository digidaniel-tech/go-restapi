package usecases

import "github.com/digidaniel-tech/go-restapi/internal/models"

func GetProduct(id string) (models.Product, error) {
	return models.Product{ID: 1, Name: "Product 1", Price: 1.00}, nil
}

func GetProducts() ([]models.Product, error) {
	category := models.Category{ID: 1, Name: "Root"}

	productList := []models.Product{
		{ID: 1, Name: "Product 1", Price: 1.00, Category: category},
		{ID: 2, Name: "Product 2", Price: 2.00, Category: category},
		{ID: 3, Name: "Product 3", Price: 3.00, Category: category},
	}

	return productList, nil
}
