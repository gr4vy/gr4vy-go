package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/steve-gr4vy/gr4vy/api"
)

type Gr4vyListPaymentOptionsParams ListPaymentOptionsParams

func (c *Gr4vyClient) ListPaymentOptions(params Gr4vyListPaymentOptionsParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListPaymentOptionsParams = ListPaymentOptionsParams(params)
    return c.HandleResponse(client.ListPaymentOptions(context.TODO(), &p))
}
