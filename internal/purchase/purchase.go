package purchase

import (
	"context"
	"errors"
	"time"

	"uuid"

	"github.com/Rhymond/go-money"
	coffeco "github.com/kalzipher/coffeco/internal"
	"github.com/kalzipher/coffeco/internal/payment"
	"github.com/kalzipher/coffeco/internal/store"
)

type Purchase struct {
	id                uuid.UUID
	Store             store.Store
	ProductToPurchase []coffeco.Product
	total             money.Money
	PaymentMeans      payment.Means
	timeOfPurchase    time.Time
	CardToken         *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductToPurchase) == 0 {
		return errors.New("purchase must consist of at least one product")
	}

	p.total = *money.New(0, "USD")
	for _, product := range p.ProductToPurchase {
		newTotal, _ := p.total.Add(&product.BasePrice)
		p.total = *newTotal
	}

	if p.total.IsZero() {
		return errors.New("purchase total must be greater than zero")
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}

type CardChargeService interface {
	ChargeCard(ctx context.Context, cardToken string, amount money.Money) error
}

type Service struct {
	CardChargeService  CardChargeService
	purchaseRepository Repository
}

func (s Service) MakePurchase(ctx context.Context, purchase *Purchase) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.CardChargeService.ChargeCard(ctx, *purchase.CardToken, purchase.total); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		// Todo: For the reader to add :)
	case payment.MEANS_COFFEBUX:
		// Todo: For the reader to add :)
	default:
		return errors.New("unsupported payment method")
	}

	if err := s.purchaseRepository.Store(ctx, purchase); err != nil {
		return errors.New("failed to store purchase")
	}

	return nil
}
