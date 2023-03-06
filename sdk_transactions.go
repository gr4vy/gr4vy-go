package gr4vy

import (
	"net/http"
	"context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyTransactionPaymentMethodRequest TransactionPaymentMethodRequest
type Gr4vyTransactionRequest TransactionRequest
type Gr4vyTransactionCaptureRequest TransactionCaptureRequest
type Gr4vyTransactionRefundRequest TransactionRefundRequest
type Gr4vyTransactions Transactions
type Gr4vyTransaction Transaction
type Gr4vyRefund Refund


func (c *Gr4vyClient) ListTransactions(limit *int32) (*Gr4vyTransactions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.ListTransactions(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransactions = Gr4vyTransactions(response)
    return &r, http, err
}
func (c *Gr4vyClient) ListTransactionsContext(ctx context.Context, limit *int32) (*Gr4vyTransactions, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.ListTransactions(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransactions = Gr4vyTransactions(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetTransaction(transaction_id string) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.GetTransaction(auth, transaction_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}
func (c *Gr4vyClient) GetTransactionContext(ctx context.Context, transaction_id string) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.GetTransaction(auth, transaction_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}

func (c *Gr4vyClient) AuthorizeNewTransaction(body Gr4vyTransactionRequest, pm Gr4vyTransactionPaymentMethodRequest) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.AuthorizeNewTransaction(auth)

    var b TransactionRequest = TransactionRequest(body)
    var tpm TransactionPaymentMethodRequest = TransactionPaymentMethodRequest(pm)
    b.PaymentMethod = tpm
    response, http, err := p.TransactionRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}
func (c *Gr4vyClient) AuthorizeNewTransactionContext(ctx context.Context, body Gr4vyTransactionRequest, pm Gr4vyTransactionPaymentMethodRequest) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.AuthorizeNewTransaction(auth)

    var b TransactionRequest = TransactionRequest(body)
    var tpm TransactionPaymentMethodRequest = TransactionPaymentMethodRequest(pm)
    b.PaymentMethod = tpm
    response, http, err := p.TransactionRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}
func (c *Gr4vyClient) AuthorizeNewTransactionWithIdempotencyKey(body Gr4vyTransactionRequest, pm Gr4vyTransactionPaymentMethodRequest, ik string) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.AuthorizeNewTransaction(auth)

    p = p.IdempotencyKey(ik)
    
    var b TransactionRequest = TransactionRequest(body)
    var tpm TransactionPaymentMethodRequest = TransactionPaymentMethodRequest(pm)
    b.PaymentMethod = tpm
    response, http, err := p.TransactionRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}
func (c *Gr4vyClient) AuthorizeNewTransactionContextWithIdempotencyKey(ctx context.Context, body Gr4vyTransactionRequest, pm Gr4vyTransactionPaymentMethodRequest, ik string) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.AuthorizeNewTransaction(auth)
    
    p = p.IdempotencyKey(ik)

    var b TransactionRequest = TransactionRequest(body)
    var tpm TransactionPaymentMethodRequest = TransactionPaymentMethodRequest(pm)
    b.PaymentMethod = tpm
    response, http, err := p.TransactionRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}

func (c *Gr4vyClient) CaptureTransaction(transaction_id string, body Gr4vyTransactionCaptureRequest) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.CaptureTransaction(auth, transaction_id)

    var b TransactionCaptureRequest = TransactionCaptureRequest(body)
    response, http, err := p.TransactionCaptureRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}
func (c *Gr4vyClient) CaptureTransactionContext(ctx context.Context, transaction_id string, body Gr4vyTransactionCaptureRequest) (*Gr4vyTransaction, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.CaptureTransaction(auth, transaction_id)

    var b TransactionCaptureRequest = TransactionCaptureRequest(body)
    response, http, err := p.TransactionCaptureRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyTransaction = Gr4vyTransaction(response)
    return &r, http, err
}

func (c *Gr4vyClient) RefundTransaction(transaction_id string, body Gr4vyTransactionRefundRequest) (*Gr4vyRefund, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.RefundTransaction(auth, transaction_id)

    var b TransactionRefundRequest = TransactionRefundRequest(body)
    response, http, err := p.TransactionRefundRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyRefund = Gr4vyRefund(response)
    return &r, http, err
}
func (c *Gr4vyClient) RefundTransactionContext(ctx context.Context, transaction_id string, body Gr4vyTransactionRefundRequest) (*Gr4vyRefund, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(ctx, ContextAccessToken, c.accessToken)
    p := client.TransactionsApi.RefundTransaction(auth, transaction_id)

    var b TransactionRefundRequest = TransactionRefundRequest(body)
    response, http, err := p.TransactionRefundRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyRefund = Gr4vyRefund(response)
    return &r, http, err
}
