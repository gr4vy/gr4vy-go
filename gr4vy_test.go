package gr4vy

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const keyPath = "./private_key.pem"
const gr4vyId = "spider"
const environment = "sandbox"

var buyerId string
var paymentMethodId string
var cardRuleId string
var paymentServiceId string
var paymentServiceIdDelete string
var transactionId string

func TestEmbedClaimsSerialization(t *testing.T) {
	embed := EmbedParams{
		Amount:   200,
		Currency: "USD",
		BuyerID:  "d757c76a-cbd7-4b56-95a3-40125b51b29c",
	}

	serialized := getEmbedClaims(embed)

	if _, present := serialized["buyer_id"]; !present {
		t.Error("EmbedParams serialization is not snake_casing keys")
	}
}

func TestEmbedClaimsWithMerchantAccountId(t *testing.T) {
	embed := EmbedParams{
		Amount:            200,
		Currency:          "USD",
		BuyerID:           "d757c76a-cbd7-4b56-95a3-40125b51b29c",
		MerchantAccountId: "plantly-uk",
	}

	serialized := getEmbedClaims(embed)

	if _, present := serialized["buyer_id"]; !present {
		t.Error("EmbedParams serialization is not snake_casing keys")
	}
}
func TestEmbedToken(t *testing.T) {

	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	embed := EmbedParams{
		Amount:   200,
		Currency: "USD",
		BuyerID:  "d757c76a-cbd7-4b56-95a3-40125b51b29c",
		Metadata: map[string]string{"key": "value"},
	}

	_, err = client.GetEmbedToken(embed, "")

	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestEmbedTokenWithConnectionOptions(t *testing.T) {

	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	embed := EmbedParams{
		Amount:   200,
		Currency: "USD",
		BuyerID:  "d757c76a-cbd7-4b56-95a3-40125b51b29c",
		Metadata: map[string]string{"key": "value"},
		ConnectionOptions: map[string]interface{}{
			"stripe-card": map[string]interface{}{
				"stripe_connect": map[string]string{"key": "value"},
			},
		},
	}

	_, err = client.GetEmbedToken(embed, "")

	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestAddBuyerAndEmbed(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyBuyerRequest{
		DisplayName: String("Jane Smith"),
	}
	var response *Gr4vyBuyer
	response, _, err = client.AddBuyer(req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	buyerId = *response.Id

	embed := EmbedParams{
		Amount:   200,
		Currency: "USD",
		BuyerID:  buyerId,
		Metadata: map[string]string{"key": "value"},
	}

	client = NewGr4vyClient(gr4vyId, key, environment)

	_, err = client.GetEmbedToken(embed, "")

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//Clean up
	client = NewGr4vyClient(gr4vyId, key, environment)

	_, err = client.DeleteBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestListBuyers(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyBuyers
	response, _, err = client.ListBuyers(Int32(5))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("%+v\n", *response)
}

func TestListBuyersForMid(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClientWithMid(gr4vyId, key, environment, "default")

	var response *Gr4vyBuyers
	response, _, err = client.ListBuyers(Int32(5))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("%+v\n", *response)
}
func TestListBuyersContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyBuyers
	response, _, err = client.ListBuyersContext(context.Background(), Int32(5))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("%+v\n", *response)
}
func TestAddBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyBuyerRequest{
		DisplayName: String("Jane Smith"),
	}
	var response *Gr4vyBuyer
	response, _, err = client.AddBuyer(req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	buyerId = *response.Id
	t.Log("Set buyerId: " + buyerId)
}
func TestGetBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)
	var response *Gr4vyBuyer
	response, _, err = client.GetBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Retrieved buyer: " + *response.Id)
}
func TestUpdateBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyBuyerUpdate{
		DisplayName: String("Janet Smith"),
	}

	_, _, err = client.UpdateBuyer(buyerId, req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Updated buyer: " + buyerId)
}
func TestDeleteBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, err = client.DeleteBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Deleted buyer: " + buyerId)
}

func TestListPaymentMethods(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentMethods(nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestListPaymentMethodsContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentMethodsContext(context.Background(), nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestStorePaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyPaymentMethodRequest{
		Method:         "card",
		Number:         StringPtr("4111111111111111"),
		ExpirationDate: StringPtr("12/25"),
	}

	var response *Gr4vyPaymentMethod
	response, _, err = client.StorePaymentMethod(req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	paymentMethodId = *response.Id
	t.Log("Set paymentMethodId: " + paymentMethodId)
}
func TestGetPaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyPaymentMethod
	response, _, err = client.GetPaymentMethod(paymentMethodId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Retrieved paymentMethod: " + *response.Id)
}
func TestDeletePaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, err = client.DeletePaymentMethod(paymentMethodId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Deleted paymentMethod: " + paymentMethodId)
}

func TestListPaymentOptions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentOptions()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestListPaymentOptionsContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentOptionsContext(context.Background())
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestPostListPaymentOptions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	// cartItems := []CartItem{
	// 	{Name: "Aloe", UnitAmount: 1000, Quantity: 1, Sku: String("aloe")},
	// }

	req := Gr4vyPaymentOptionsRequest{
		Country:  String("US"),
		Currency: String("USD"),
		Amount:   Gr4vyNullableInt32(1000),
		Metadata: map[string]string{"TypeOfPayment": "purchase", "Carbon_FootPrint": "10"},
		// CartItems:	cartItems,
		Locale: String("en"),
	}

	var response *Gr4vyPaymentOptions
	response, _, err = client.PostListPaymentOptions(req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var paymentOptionMethod = (*response.Items)[0].GetMethod()
	t.Log("Retrieved payment options: " + paymentOptionMethod)
}
func TestPostListPaymentOptionsContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	// cartItems := []CartItem{
	// 	{Name: "Aloe", UnitAmount: 1000, Quantity: 1, Sku: String("aloe")},
	// }

	req := Gr4vyPaymentOptionsRequest{
		Country:  String("US"),
		Currency: String("USD"),
		Amount:   Gr4vyNullableInt32(1000),
		Metadata: map[string]string{"TypeOfPayment": "purchase", "Carbon_FootPrint": "10"},
		// CartItems:	cartItems,
		Locale: String("en"),
	}

	var response *Gr4vyPaymentOptions
	response, _, err = client.PostListPaymentOptionsContext(context.Background(), req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var paymentOptionMethod = (*response.Items)[1].GetMethod()
	t.Log("Retrieved payment options: " + paymentOptionMethod)
}

func TestListPaymentServiceDefinitions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentServiceDefinitions(nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestListPaymentServiceDefinitionsContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListPaymentServiceDefinitionsContext(context.Background(), nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestGetPaymentServiceDefinition(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)
	_, _, err = client.GetPaymentServiceDefinition("braintree-card")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestListPaymentServices(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyPaymentServices
	response, _, err = client.ListPaymentServices(nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	paymentServiceId = (*response.Items)[0].GetId()
	t.Log("Set paymentServiceId: " + paymentServiceId)
}
func TestListPaymentServicesContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyPaymentServices
	response, _, err = client.ListPaymentServicesContext(context.Background(), nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	paymentServiceId = (*response.Items)[0].GetId()
	t.Log("Set paymentServiceId: " + paymentServiceId)
}
func TestGetPaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.GetPaymentService(paymentServiceId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestAddPaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyPaymentServiceRequest{
		PaymentServiceDefinitionId: "stripe-card",
		DisplayName:                "Test",
		AcceptedCurrencies:         []string{"GBP"},
		AcceptedCountries:          []string{"GB"},
	}
	fields := []Gr4vyPaymentServiceRequestFields{
		{Key: "secret_key", Value: "abc"},
	}
	var response *Gr4vyPaymentService
	response, _, err = client.AddPaymentService(req, fields)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	paymentServiceIdDelete = *response.Id
	t.Log("Set paymentServiceIdDelete: " + paymentServiceIdDelete)
}
func TestUpdatePaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyPaymentServiceUpdate{
		DisplayName: StringPtr("Braintree UK"),
	}
	var response *Gr4vyPaymentService
	response, _, err = client.UpdatePaymentService(paymentServiceIdDelete, req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Updated paymentService: " + *response.Id)
}
func TestDeletePaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, err = client.DeletePaymentService(paymentServiceIdDelete)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Deleted paymentService: " + paymentServiceIdDelete)
}

func TestListTransactions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListTransactions(nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestListTransactionsContext(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	_, _, err = client.ListTransactionsContext(context.Background(), nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
func TestAuthorizeNewTransaction(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyTransactionRequest{
		Amount:   1299,
		Currency: "USD",
	}

	paymentMethod := Gr4vyTransactionPaymentMethodRequest{
		Method:         "card",
		Number:         Gr4vyNullableString("4111111111111111"),
		ExpirationDate: Gr4vyNullableString("12/25"),
		SecurityCode:   Gr4vyNullableString("123"),
	}

	id := uuid.New()

	var response *Gr4vyTransaction
	var body *http.Response
	response, body, err = client.AuthorizeNewTransactionWithIdempotencyKey(req, paymentMethod, id.String())
	if err != nil {
		t.Errorf(err.Error())
		fmt.Printf("%+v\n", response)
		fmt.Printf("%+v\n", body)
		return
	}

	transactionId = *response.Id
	t.Log("Set transactionId: " + transactionId)
}
func TestGetTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	var response *Gr4vyTransaction
	response, _, err = client.GetTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Retrieved transaction: " + *response.Id)
}
func TestCaptureTransaction(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyTransactionCaptureRequest{
		Amount: Int32(1299),
	}
	response, _, err := client.CaptureTransaction(transactionId, req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Capture status: " + *response.Status)
}
func TestAuthorizeTransaction(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key, environment)

	// response, err := client.AuthorizeTransaction(transactionId)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// if (response.StatusCode != 200) {
	// 	t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	// }
}
func TestRefundTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyTransactionRefundRequest{}

	body, response, err := client.RefundTransaction(transactionId, req)
	if err != nil {
		t.Errorf(err.Error())
		fmt.Printf("%+v\n", response)
		fmt.Printf("%+v\n", body)
		return
	}
	if response.StatusCode != 201 {
		t.Errorf("expected StatusCode 201: received: " + strconv.Itoa(response.StatusCode))
	}
}
func TestAuthorizeSecondTransaction(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	req := Gr4vyTransactionRequest{
		Amount:   1299,
		Currency: "USD",
	}

	paymentMethod := Gr4vyTransactionPaymentMethodRequest{
		Method:         "card",
		Number:         Gr4vyNullableString("4111111111111111"),
		ExpirationDate: Gr4vyNullableString("12/25"),
		SecurityCode:   Gr4vyNullableString("123"),
	}

	id := uuid.New()

	var response *Gr4vyTransaction
	var body *http.Response
	response, body, err = client.AuthorizeNewTransactionWithIdempotencyKey(req, paymentMethod, id.String())
	if err != nil {
		t.Errorf(err.Error())
		fmt.Printf("%+v\n", response)
		fmt.Printf("%+v\n", body)
		return
	}

	transactionId = *response.Id
	t.Log("Set transactionId: " + transactionId)
}
func TestVoidTransaction(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key, environment)

	body, response, err := client.VoidTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		fmt.Printf("%+v\n", response)
		fmt.Printf("%+v\n", body)
		return
	}
	if response.StatusCode != 200 {
		t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	}
}

func TestEmbedTokenWithCheckoutSession(t *testing.T) {
	// fetch/decode the private key
	key, err := GetKeyFromFile(keyPath)
	block, _ := pem.Decode([]byte(key))

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	ecPrivateKey, _ := privateKey.(*ecdsa.PrivateKey)

	if err != nil {
		t.Errorf(err.Error())
		fmt.Println("Failed to parse private key:", err)
		return
	}

	// create a client
	client := NewGr4vyClient(gr4vyId, key, environment)

	// create a checkout session
	req := Gr4vyCheckoutSessionCreateRequest{
		Metadata: map[string]string{"foo": "bar"},
	}
	checkoutSession, _, err := client.AddCheckoutSession(req)

	if err != nil {
		t.Errorf(err.Error())
	}

	// create an embed token
	embed := EmbedParams{
		Amount:            200,
		Currency:          "USD",
		BuyerID:           "d757c76a-cbd7-4b56-95a3-40125b51b29c",
		MerchantAccountId: "plantly-uk",
	}
	token, err := client.GetEmbedToken(embed, *checkoutSession.Id)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// verify/parse the token
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is ES256
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the ECDSA private key for verification
		return ecPrivateKey.Public(), nil
	})

	if err != nil {
		fmt.Println("Failed to parse token:", err)
		t.Errorf(err.Error())
		return
	}

	// check the token is valid
	if !result.Valid {
		t.Errorf("token is not valid")
		return
	}

	claims, ok := result.Claims.(jwt.MapClaims)

	if !ok {
		t.Errorf("invalid token claims")
		return
	}

	if claims["checkout_session_id"] != *checkoutSession.Id {
		t.Errorf("checkout session id does not match %v %v", claims["checkout_session_id"], *checkoutSession.Id)
		return
	}
}

func TestVerifyWebhook_HappyCase(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "1744018920"
	payload := "payload"

	err := VerifyWebhook(secret, payload, &signatureHeader, &timestampHeader, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestVerifyWebhook_OldTimestamp(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "1744018920"
	payload := "payload"

	err := VerifyWebhook(secret, payload, &signatureHeader, &timestampHeader, 60)
	if err == nil || err.Error() != "timestamp too old" {
		t.Errorf("expected 'timestamp too old' error, got: %v", err)
	}
}

func TestVerifyWebhook_WrongSignature(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "other"
	timestampHeader := "1744018920"
	payload := "payload"

	err := VerifyWebhook(secret, payload, &signatureHeader, &timestampHeader, 0)
	if err == nil || err.Error() != "no matching signature found" {
		t.Errorf("expected 'no matching signature found' error, got: %v", err)
	}
}

func TestVerifyWebhook_InvalidTimestamp(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	signatureHeader := "78aca0c78005107a654a957b8566fa6e0e5e06aea92d7da72a6da9e5a690d013,other"
	timestampHeader := "wrong"
	payload := "payload"

	err := VerifyWebhook(secret, payload, &signatureHeader, &timestampHeader, 0)
	if err == nil || err.Error() != "invalid header timestamp" {
		t.Errorf("expected 'invalid header timestamp' error, got: %v", err)
	}
}

func TestVerifyWebhook_MissingHeader(t *testing.T) {
	secret := "Ik4L-8FH0ihWczctcIPXZRR_8F0fPNgmhEfVBbZ3zNwqQVa1Or4tBz4Pgw2eNaVDod7H56Y268h_wohEUaWbUg"
	timestampHeader := "1744018920"
	payload := "payload"

	err := VerifyWebhook(secret, payload, nil, &timestampHeader, 0)
	if err == nil || err.Error() != "missing header values" {
		t.Errorf("expected 'missing header values' error, got: %v", err)
	}
}
