package payment

type Means string

const (
	MEANS_CASH     Means = "Cash"
	MEANS_CARD     Means = "Card"
	MEANS_COFFEBUX Means = "CoffeBux"
)

type CardDetails struct {
	cardToken string
}
