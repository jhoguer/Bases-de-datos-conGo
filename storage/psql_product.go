package storage

import (
	"database/sql"
	"fmt"

	"github.com/jhoguer/Bases-de-datos-conGo/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
		)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) 
												VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at
												FROM products`
)

// psqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de producto ejecutada correctamente")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)

	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)

	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")
	return nil
}

// GetAll implementa la interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProducts)
	if err != nil {
		// Retorna el valor cero de un slice que es nil y el err
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	// var ms []product.Models
	for rows.Next() {
		m := &product.Model{}
		observationNull := sql.NullString{}
		updatedAtNull := sql.NullTime{}

		err := rows.Scan(
			&m.ID,
			&m.Name,
			&observationNull,
			&m.Price,
			&m.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		m.Observations = observationNull.String
		m.UpdatedAt = updatedAtNull.Time
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}
