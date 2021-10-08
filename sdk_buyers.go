package gr4vy

import (
	"net/http"
	"context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyBuyerRequest BuyerRequest
type Gr4vyBuyerUpdate BuyerUpdate
type Gr4vyBuyer Buyer
type Gr4vyBuyers Buyers

func (c *Gr4vyClient) ListBuyers(limit *int32) (*Gr4vyBuyers, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.BuyersApi.ListBuyers(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyers = Gr4vyBuyers(response)
    return &r, http, err
}
func (c *Gr4vyClient) ListBuyersContext(ctx context.Context, limit *int32) (*Gr4vyBuyers, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.BuyersApi.ListBuyers(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyers = Gr4vyBuyers(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetBuyer(buyer_id string) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.BuyersApi.GetBuyer(auth, buyer_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}
func (c *Gr4vyClient) GetBuyerContext(ctx context.Context, buyer_id string) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.BuyersApi.GetBuyer(auth, buyer_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}

func (c *Gr4vyClient) AddBuyer(body Gr4vyBuyerRequest) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.BuyersApi.AddBuyer(auth)

    var b BuyerRequest = BuyerRequest(body)
    response, http, err := p.BuyerRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}
func (c *Gr4vyClient) AddBuyerContext(ctx context.Context, body Gr4vyBuyerRequest) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.BuyersApi.AddBuyer(auth)

    var b BuyerRequest = BuyerRequest(body)
    response, http, err := p.BuyerRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}

func (c *Gr4vyClient) UpdateBuyer(buyer_id string, body Gr4vyBuyerUpdate) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.BuyersApi.UpdateBuyer(auth, buyer_id)

    var b BuyerUpdate = BuyerUpdate(body)
    response, http, err := p.BuyerUpdate(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}
func (c *Gr4vyClient) UpdateBuyerContext(ctx context.Context, buyer_id string, body Gr4vyBuyerUpdate) (*Gr4vyBuyer, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.BuyersApi.UpdateBuyer(auth, buyer_id)

    var b BuyerUpdate = BuyerUpdate(body)
    response, http, err := p.BuyerUpdate(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyBuyer = Gr4vyBuyer(response)
    return &r, http, err
}

func (c *Gr4vyClient) DeleteBuyer(buyer_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.BuyersApi.DeleteBuyer(auth, buyer_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}
func (c *Gr4vyClient) DeleteBuyerContext(ctx context.Context, buyer_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.BuyersApi.DeleteBuyer(auth, buyer_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}