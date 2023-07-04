package gr4vy

import (
	"context"
	"net/http"

	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyCheckoutSessionRequest CheckoutSessionRequest
type Gr4vyCheckoutSession CheckoutSession

func (c *Gr4vyClient) GetCheckoutSession(checkout_session_id string) (*Gr4vyCheckoutSession, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CheckoutSessionsApi.GetCheckoutSession(auth, checkout_session_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCheckoutSession = Gr4vyCheckoutSession(response)
    return &r, http, err
}

func (c *Gr4vyClient) AddCheckoutSession() (*Gr4vyCheckoutSession, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CheckoutSessionsApi.NewCheckoutSession(auth)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCheckoutSession = Gr4vyCheckoutSession(response)
    return &r, http, err
}