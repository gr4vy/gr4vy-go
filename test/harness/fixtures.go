package harness

import (
	"fmt"
	"sync/atomic"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
)

// The deterministic mock-card connector approves this card with a future
// expiry and CVV 123 (see TESTING.md).
const (
	ApprovingCardNumber = "4111111111111111"
	CardExpirationDate  = "12/35"
	CardSecurityCode    = "123"
	MissingID           = "11111111-1111-1111-1111-111111111111"
)

var counter uint64

// UniqueID returns a stable-prefix, monotonically-unique identifier so tests
// don't collide on external identifiers / display names within a shard.
func UniqueID(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, atomic.AddUint64(&counter, 1))
}

// ApprovingCard is the card-payment-method input the mock connector approves.
func ApprovingCard() components.CardPaymentMethodCreate {
	return components.CardPaymentMethodCreate{
		Number:         ApprovingCardNumber,
		ExpirationDate: CardExpirationDate,
		SecurityCode:   gr4vygo.String(CardSecurityCode),
	}
}

// SampleAddress is a complete, valid GB address.
func SampleAddress() *components.Address {
	return &components.Address{
		City:              gr4vygo.String("London"),
		Country:           gr4vygo.String("GB"),
		PostalCode:        gr4vygo.String("EC1A 1BB"),
		State:             gr4vygo.String("England"),
		HouseNumberOrName: gr4vygo.String("10"),
		Line1:             gr4vygo.String("10 Downing Street"),
		Organization:      gr4vygo.String("Gr4vy"),
	}
}

// SampleBillingDetails is a complete buyer billing-details payload.
func SampleBillingDetails() *components.BillingDetails {
	return &components.BillingDetails{
		FirstName:    gr4vygo.String("Jane"),
		LastName:     gr4vygo.String("Doe"),
		EmailAddress: gr4vygo.String("jane.doe@example.com"),
		PhoneNumber:  gr4vygo.String("+447700900000"),
		Address:      SampleAddress(),
	}
}

// SampleCartItem is one valid line item.
func SampleCartItem() components.CartItem {
	return components.CartItem{
		Name:       "Joust Duffle Bag",
		Quantity:   1,
		UnitAmount: 1299,
		Categories: []string{"Gear", "Bags"},
	}
}
