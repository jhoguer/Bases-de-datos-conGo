package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jhoguer/Bases-de-datos-conGo/pkg/product"
	"github.com/jhoguer/Bases-de-datos-conGo/storage"
)

func main() {
	storage.NewPostgresDB()

	// Creamos una struct llamada storageProduct qu tiene implementada
	// el metodo Migrate de la interface Storage de Product
	// storageProduct := storage.NewPsqlProduct(storage.Pool())

	// // Creamos una struct serviceProduct que ya implemento el
	// // metodo Migrate
	// serviceProduct := product.NewService(storageProduct)

	// if err := serviceProduct.Migrated(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// // Migracion de tabla invoice_headers
	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	// if err := serviceInvoiceHeader.Migrated(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }

	// // Migracion de tabla invoice_items
	// storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	// serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	// if err := serviceInvoiceItem.Migrated(); err != nil {
	// 	log.Fatalf("invoiceItem.Migrate: %v", err)
	// }

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m := &product.Model{
	// 	Name:         "Curso de db con Go",
	// 	Price:        80,
	// 	Observations: "On fire",
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// err := serviceProduct.Create(m)
	// if err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v\n", m)

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// ms, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }

	// fmt.Println(ms)

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetById(10)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}

}
