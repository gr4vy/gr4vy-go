package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/steve-gr4vy/gr4vy/api"
)

type Gr4vyListTransactionsParams ListTransactionsParams
type Gr4vyAuthorizeNewTransaction AuthorizeNewTransactionJSONRequestBody
type Gr4vyCaptureTransaction CaptureTransactionJSONRequestBody
// type Gr4vyUpdateBuyer UpdateBuyerJSONRequestBody
// type Gr4vyBuyer Buyer

func (c *Gr4vyClient) ListTransactions(params Gr4vyListTransactionsParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListTransactionsParams = ListTransactionsParams(params)
	
    return c.HandleResponse(client.ListTransactions(context.TODO(), &p))
}

func (c *Gr4vyClient) GetTransaction(transaction_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetTransaction(context.TODO(), transaction_id))
}

func (c *Gr4vyClient) AuthorizeNewTransaction(body Gr4vyAuthorizeNewTransaction, pm Gr4vyPaymentMethod) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var b AuthorizeNewTransactionJSONRequestBody = AuthorizeNewTransactionJSONRequestBody(body)
    var p CardRequest = CardRequest(pm)
    b.PaymentMethod = p
    return c.HandleResponse(client.AuthorizeNewTransaction(context.TODO(), b))
}

func (c *Gr4vyClient) CaptureTransaction(transaction_id string, body Gr4vyCaptureTransaction) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    var b CaptureTransactionJSONRequestBody = CaptureTransactionJSONRequestBody(body)
    return c.HandleResponse(client.CaptureTransaction(context.TODO(), transaction_id, b))
}

func (c *Gr4vyClient) AuthorizeTransaction(transaction_id string,) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    return c.HandleResponse(client.AuthorizeTransaction(context.TODO(), transaction_id))
}

func (c *Gr4vyClient) RefundTransaction(transaction_id string,) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    return c.HandleResponse(client.RefundTransaction(context.TODO(), transaction_id))
}
