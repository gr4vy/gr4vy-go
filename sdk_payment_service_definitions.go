package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyListPaymentServiceDefinitionsParams ListPaymentServiceDefinitionsParams
type Gr4vyPaymentServiceDefinition PaymentServiceDefinition

func (c *Gr4vyClient) ListPaymentServiceDefinitions(params Gr4vyListPaymentServiceDefinitionsParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListPaymentServiceDefinitionsParams = ListPaymentServiceDefinitionsParams(params)
    return c.HandleResponse(client.ListPaymentServiceDefinitions(context.TODO(), &p))
}

func (c *Gr4vyClient) GetPaymentServiceDefinition(payment_service_definition_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetPaymentServiceDefinition(context.TODO(), payment_service_definition_id))
}