package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyListPaymentMethodsParams ListPaymentMethodsParams
type Gr4vyStorePaymentMethod StorePaymentMethodJSONRequestBody
type Gr4vyPaymentMethod CardRequest


func (c *Gr4vyClient) ListPaymentMethods(params Gr4vyListPaymentMethodsParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
	var p ListPaymentMethodsParams = ListPaymentMethodsParams(params)
    return c.HandleResponse(client.ListPaymentMethods(context.TODO(), &p))
}

func (c *Gr4vyClient) GetPaymentMethod(payment_method_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetPaymentMethod(context.TODO(), payment_method_id))
}

func (c *Gr4vyClient) StorePaymentMethod(body Gr4vyStorePaymentMethod) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    var b StorePaymentMethodJSONRequestBody = StorePaymentMethodJSONRequestBody(body)
    return c.HandleResponse(client.StorePaymentMethod(context.TODO(), b))
}

func (c *Gr4vyClient) DeletePaymentMethod(payment_method_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.DeletePaymentMethod(context.TODO(), payment_method_id))
}