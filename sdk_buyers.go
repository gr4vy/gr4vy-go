package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/steve-gr4vy/gr4vy/api"
)

type Gr4vyListBuyersParams ListBuyersParams
type Gr4vyAddBuyer AddBuyerJSONRequestBody
type Gr4vyUpdateBuyer UpdateBuyerJSONRequestBody
type Gr4vyBuyer Buyer

func (c *Gr4vyClient) ListBuyers(params Gr4vyListBuyersParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListBuyersParams = ListBuyersParams(params)
	
    return c.HandleResponse(client.ListBuyers(context.TODO(), &p))
}

func (c *Gr4vyClient) GetBuyer(buyer_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetBuyer(context.TODO(), buyer_id))
}

func (c *Gr4vyClient) AddBuyer(body Gr4vyAddBuyer) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var b AddBuyerJSONRequestBody = AddBuyerJSONRequestBody(body)
    return c.HandleResponse(client.AddBuyer(context.TODO(), b))
}

func (c *Gr4vyClient) UpdateBuyer(buyer_id string, body Gr4vyUpdateBuyer) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    var b UpdateBuyerJSONRequestBody = UpdateBuyerJSONRequestBody(body)
    return c.HandleResponse(client.UpdateBuyer(context.TODO(), buyer_id, b))
}

func (c *Gr4vyClient) DeleteBuyer(buyer_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.DeleteBuyer(context.TODO(), buyer_id))
}