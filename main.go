package main

import (
	"log"

	"github.com/jhoguer/Bases-de-datos-conGo/pkg/product"
	"github.com/jhoguer/Bases-de-datos-conGo/storage"
)

func main() {
	storage.NewPostgresDB()

	// Creamos una struct llamada storageProduct qu tiene implementada
	// el metodo Migrate de la interface Storage de Product
	storageProduct := storage.NewPsqlProduct(storage.Pool())

	// Creamos una struct serviceProduct que ya implemento el
	// metodo Migrate
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrated(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
