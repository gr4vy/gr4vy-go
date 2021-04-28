package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyPaymentServiceUpdateFields PaymentServiceUpdateFields
type Gr4vyPaymentServiceRequest PaymentServiceRequest
type Gr4vyPaymentServiceUpdate PaymentServiceUpdate
type Gr4vyPaymentService PaymentService
type Gr4vyPaymentServices PaymentServices

func (c *Gr4vyClient) ListPaymentServices(limit *int32) (*Gr4vyPaymentServices, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.ListPaymentServices(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentServices = Gr4vyPaymentServices(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetPaymentService(payment_service_id string) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.GetPaymentService(auth, payment_service_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentService = Gr4vyPaymentService(response)
    return &r, http, err
}

func (c *Gr4vyClient) AddPaymentService(body Gr4vyPaymentServiceRequest, fields []Gr4vyPaymentServiceUpdateFields) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.AddPaymentService(auth)

    var b PaymentServiceRequest = PaymentServiceRequest(body)
	var f []PaymentServiceUpdateFields
    for _, element := range fields {
        var c PaymentServiceUpdateFields = PaymentServiceUpdateFields(element)
        f = append(f, c)
    }
    b.Fields = f
    response, http, err := p.PaymentServiceRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentService = Gr4vyPaymentService(response)
    return &r, http, err
}

func (c *Gr4vyClient) UpdatePaymentService(payment_service_id string, body Gr4vyPaymentServiceUpdate) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.UpdatePaymentService(auth, payment_service_id)

    var b PaymentServiceUpdate = PaymentServiceUpdate(body)
    response, http, err := p.PaymentServiceUpdate(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentService = Gr4vyPaymentService(response)
    return &r, http, err
}

func (c *Gr4vyClient) DeletePaymentService(payment_service_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.DeletePaymentService(auth, payment_service_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}