package payment

import (
	"errors"

	"github.com/stripe/stripe-go/v74/client"
)

type StripeService struct {
	stripeClient *client.API
}

func NewStripeService(apiKey string) (*StripeService, error) {
	if apiKey == "" {
		return nil, errors.New("api key is missing")
	}

	stripeClient := client.New(apiKey, nil)

	return &StripeService{
		stripeClient: stripeClient,
	}, nil
}
