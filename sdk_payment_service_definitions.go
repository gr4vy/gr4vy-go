package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyPaymentServiceDefinition PaymentServiceDefinition
type Gr4vyPaymentServiceDefinitions PaymentServiceDefinitions

func (c *Gr4vyClient) ListPaymentServiceDefinitions(limit *int32) (*Gr4vyPaymentServiceDefinitions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServiceDefinitionsApi.ListPaymentServiceDefinitions(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentServiceDefinitions = Gr4vyPaymentServiceDefinitions(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetPaymentServiceDefinition(payment_service_definition_id string) (*Gr4vyPaymentServiceDefinition, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServiceDefinitionsApi.GetPaymentServiceDefinition(auth, payment_service_definition_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentServiceDefinition = Gr4vyPaymentServiceDefinition(response)
    return &r, http, err
}