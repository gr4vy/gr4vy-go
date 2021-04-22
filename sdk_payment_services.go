package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyListPaymentServicesParams ListPaymentServicesParams
type Gr4vyAddPaymentService AddPaymentServiceJSONRequestBody
type Gr4vyUpdatePaymentService UpdatePaymentServiceJSONRequestBody
type Gr4vyPaymentService PaymentService

func (c *Gr4vyClient) ListPaymentServices(params Gr4vyListPaymentServicesParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListPaymentServicesParams = ListPaymentServicesParams(params)
    return c.HandleResponse(client.ListPaymentServices(context.TODO(), &p))
}

func (c *Gr4vyClient) GetPaymentService(payment_service_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetPaymentService(context.TODO(), payment_service_id))
}

func (c *Gr4vyClient) AddPaymentService(body Gr4vyAddPaymentService) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    var b AddPaymentServiceJSONRequestBody = AddPaymentServiceJSONRequestBody(body)
    return c.HandleResponse(client.AddPaymentService(context.TODO(), b))
}

func (c *Gr4vyClient) UpdatePaymentService(payment_service_id string, body Gr4vyUpdatePaymentService) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    var b UpdatePaymentServiceJSONRequestBody = UpdatePaymentServiceJSONRequestBody(body)
    return c.HandleResponse(client.UpdatePaymentService(context.TODO(), payment_service_id, b))
}

func (c *Gr4vyClient) DeletePaymentService(payment_service_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.DeletePaymentService(context.TODO(), payment_service_id))
}