package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyPaymentMethodRequest CardRequest
type Gr4vyCard Card
type Gr4vyPaymentMethod PaymentMethod
type Gr4vyPaymentMethods PaymentMethods

func (c *Gr4vyClient) ListPaymentMethods(limit *int32) (*Gr4vyPaymentMethods, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentMethodsApi.ListPaymentMethods(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentMethods = Gr4vyPaymentMethods(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetPaymentMethod(payment_method_id string) (*Gr4vyPaymentMethod, *http.Response, error) {
   client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentMethodsApi.GetPaymentMethod(auth, payment_method_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentMethod = Gr4vyPaymentMethod(response)
    return &r, http, err
}

func (c *Gr4vyClient) StorePaymentMethod(body Gr4vyPaymentMethodRequest) (*Gr4vyCard, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentMethodsApi.StorePaymentMethod(auth)

    var b CardRequest = CardRequest(body)
    response, http, err := p.CardRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCard = Gr4vyCard(response)
    return &r, http, err
}

func (c *Gr4vyClient) DeletePaymentMethod(payment_method_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentMethodsApi.DeletePaymentMethod(auth, payment_method_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}