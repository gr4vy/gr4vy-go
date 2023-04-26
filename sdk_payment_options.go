package gr4vy

import (
	"net/http"
	"context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyPaymentOptions PaymentOptions
type Gr4vyPaymentOptionsRequest PaymentOptionsRequest

func (c *Gr4vyClient) ListPaymentOptions() (*Gr4vyPaymentOptions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentOptionsApi.ListPaymentOptions(auth)
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentOptions = Gr4vyPaymentOptions(response)
    return &r, http, err
}
func (c *Gr4vyClient) ListPaymentOptionsContext(ctx context.Context) (*Gr4vyPaymentOptions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.PaymentOptionsApi.ListPaymentOptions(auth)
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentOptions = Gr4vyPaymentOptions(response)
    return &r, http, err
}
func (c *Gr4vyClient) PostListPaymentOptions(body Gr4vyPaymentOptionsRequest) (*Gr4vyPaymentOptions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.PaymentOptionsApi.PostListPaymentOptions(auth)

    var b PaymentOptionsRequest = PaymentOptionsRequest(body)
    response, http, err := p.PaymentOptionsRequest(b).Execute()

    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentOptions = Gr4vyPaymentOptions(response)
    return &r, http, err
}
func (c *Gr4vyClient) PostListPaymentOptionsContext(ctx context.Context, body Gr4vyPaymentOptionsRequest) (*Gr4vyPaymentOptions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.PaymentOptionsApi.PostListPaymentOptions(auth)

    var b PaymentOptionsRequest = PaymentOptionsRequest(body)
    response, http, err := p.PaymentOptionsRequest(b).Execute()

    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyPaymentOptions = Gr4vyPaymentOptions(response)
    return &r, http, err
}
