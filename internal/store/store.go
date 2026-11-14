package store

import (
	"uuid"

	coffeco "github.com/kalzipher/coffeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeco.Product
}
