package storage

import (
	"database/sql"

	"github.com/jhoguer/Bases-de-datos-conGo/pkg/invoice"
	"github.com/jhoguer/Bases-de-datos-conGo/pkg/invoiceheader"
	"github.com/jhoguer/Bases-de-datos-conGo/pkg/invoiceitem"
)

// PsqlInvoice used for work with postgres -invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice return a new pointer of PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Store
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
