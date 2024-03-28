package storage

import (
	"marketplace/internal/storage/postgresql"

	"github.com/jmoiron/sqlx"
)

type Storages struct {
	Listing IListingStorage
	Auth    IAuthStorage
}

func NewPostgresStorages(db *sqlx.DB) *Storages {
	return &Storages{
		Auth: postgresql.NewAuthStorage(db),
		// Listing: postgresql.NewListingStorage(db),
	}
}
