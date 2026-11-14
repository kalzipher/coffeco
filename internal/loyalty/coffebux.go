package loyalty

import (
	"uuid"

	coffeco "github.com/kalzipher/coffeco/internal"
	"github.com/kalzipher/coffeco/internal/store"
)

type CoffeBux struct {
	ID                                     uuid.UUID
	Store                                  store.Store
	CoffeLover                             coffeco.CoffeLover
	FreeDrinksAvailable                    int
	RemainingDrinksPurchasesUntilFreeDrink int
}
