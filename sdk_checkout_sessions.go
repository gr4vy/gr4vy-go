package gr4vy

import (
	"context"
	"net/http"

	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyCheckoutSessionRequest CheckoutSessionRequest
type Gr4vyCheckoutSession CheckoutSession
type Gr4vyCheckoutSessionCreateRequest CheckoutSessionCreateRequest
type Gr4vyCheckoutSessionUpdateRequest CheckoutSessionUpdateRequest
type Gr4vyCartItem = CartItem

func (c *Gr4vyClient) GetCheckoutSession(checkout_session_id string) (*Gr4vyCheckoutSession, *http.Response, error) {
	client, err := GetClient(c)
	if err != nil {
		return nil, nil, err
	}
	auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
	p := client.CheckoutSessionsApi.GetCheckoutSession(auth, checkout_session_id)

	response, http, err := p.Execute()
	c.HandleResponse(http, err)
	if err != nil {
		return nil, http, err
	}
	var r Gr4vyCheckoutSession = Gr4vyCheckoutSession(response)
	return &r, http, err
}

func (c *Gr4vyClient) AddCheckoutSession(body Gr4vyCheckoutSessionCreateRequest) (*Gr4vyCheckoutSession, *http.Response, error) {
	client, err := GetClient(c)
	if err != nil {
		return nil, nil, err
	}
	auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
	p := client.CheckoutSessionsApi.NewCheckoutSession(auth)
	var b CheckoutSessionCreateRequest = CheckoutSessionCreateRequest(body)

	response, http, err := p.CheckoutSessionCreateRequest(b).Execute()
	c.HandleResponse(http, err)
	if err != nil {
		return nil, http, err
	}
	var r Gr4vyCheckoutSession = Gr4vyCheckoutSession(response)
	return &r, http, err
}

func (c *Gr4vyClient) UpdateCheckoutSession(checkoutSessionId string, body Gr4vyCheckoutSessionUpdateRequest) (*Gr4vyCheckoutSession, *http.Response, error) {
	client, err := GetClient(c)
	if err != nil {
		return nil, nil, err
	}
	auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
	p := client.CheckoutSessionsApi.UpdateCheckoutSession(auth, checkoutSessionId)
	var b CheckoutSessionUpdateRequest = CheckoutSessionUpdateRequest(body)

	response, http, err := p.CheckoutSessionUpdateRequest(b).Execute()
	c.HandleResponse(http, err)
	if err != nil {
		return nil, http, err
	}
	var r Gr4vyCheckoutSession = Gr4vyCheckoutSession(response)
	return &r, http, err
}

func (c *Gr4vyClient) DeleteCheckoutSession(checkoutSessionId string) (*http.Response, error) {
	client, err := GetClient(c)
	if err != nil {
		return nil, err
	}
	auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
	p := client.CheckoutSessionsApi.DeleteCheckoutSession(auth, checkoutSessionId)

	http, err := p.Execute()
	c.HandleResponse(http, err)
	if err != nil {
		return http, err
	}
	return http, err
}
