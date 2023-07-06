package gr4vy

import (
	"context"
	"net/http"

	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyPaymentServiceRequestFields PaymentServiceRequestFields
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
func (c *Gr4vyClient) ListPaymentServicesContext(ctx context.Context, limit *int32) (*Gr4vyPaymentServices, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
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
func (c *Gr4vyClient) GetPaymentServiceContext(ctx context.Context, payment_service_id string) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.GetPaymentService(auth, payment_service_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentService = Gr4vyPaymentService(response)
    return &r, http, err
}

func (c *Gr4vyClient) AddPaymentService(body Gr4vyPaymentServiceRequest, fields []Gr4vyPaymentServiceRequestFields) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.NewPaymentService(auth)

    var b PaymentServiceRequest = PaymentServiceRequest(body)
	var f []PaymentServiceRequestFields
    for _, element := range fields {
        var c PaymentServiceRequestFields = PaymentServiceRequestFields(element)
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
func (c *Gr4vyClient) AddPaymentServiceContext(ctx context.Context, body Gr4vyPaymentServiceRequest, fields []Gr4vyPaymentServiceRequestFields) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.NewPaymentService(auth)

    var b PaymentServiceRequest = PaymentServiceRequest(body)
    var f []PaymentServiceRequestFields
    for _, element := range fields {
        var c PaymentServiceRequestFields = PaymentServiceRequestFields(element)
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
func (c *Gr4vyClient) UpdatePaymentServiceContext(ctx context.Context, payment_service_id string, body Gr4vyPaymentServiceUpdate) (*Gr4vyPaymentService, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
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
func (c *Gr4vyClient) DeletePaymentServiceContext(ctx context.Context, payment_service_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.PaymentServicesApi.DeletePaymentService(auth, payment_service_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}