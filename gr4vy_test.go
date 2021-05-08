package gr4vy

import (
	"testing"
)

// const keyPath = "./test/gr4vy-private-key-kH4ndJM8QHyfYGpKdgfAhq2j9RySIOWa3kY_urrI5PI.pem"
const keyPath = "./test/dev-key.pem"
// const gr4vyId = "demo"
const gr4vyId = "spider"

var buyerId string
var paymentMethodId string
var cardRuleId string
var paymentServiceId string
var paymentServiceIdDelete string
var transactionId string

func TestEmbedToken(t *testing.T) {
	
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	embed := map[string]string{"amount": "200", "currency": "USD", "buyer_id": "d757c76a-cbd7-4b56-95a3-40125b51b29c"}
	_, err = client.GetEmbedToken(embed)

	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestAddBuyerAndEmbed(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
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
	embed := map[string]string{"amount": "200", "currency": "USD", "buyer_id": buyerId}

	client = NewGr4vyClient(gr4vyId, key)
	
	_, err = client.GetEmbedToken(embed)

	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	//Clean up
	client = NewGr4vyClient(gr4vyId, key)
	
	_, err = client.DeleteBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}

func TestListBuyers(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	var response *Gr4vyBuyers
	response, _, err = client.ListBuyers(Int32(5))
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Logf("%+v\n", *response)
}
func TestAddBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyBuyerRequest{
		DisplayName: String("Jane Smith"),
	}
	var response *Gr4vyBuyer
	response, _, err = client.AddBuyer(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
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
	client := NewGr4vyClient(gr4vyId, key)
	var response *Gr4vyBuyer
	response, _, err = client.GetBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Retrieved buyer: " + *response.Id)
}
func TestUpdateBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyBuyerUpdate{
		DisplayName: String("Janet Smith"),
	}

	_, _, err = client.UpdateBuyer(buyerId, req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Updated buyer: " + buyerId)
}
func TestDeleteBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, err = client.DeleteBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Deleted buyer: " + buyerId)
}

func TestListPaymentMethods(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, _, err = client.ListPaymentMethods(nil)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestStorePaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyPaymentMethodRequest{
		Method: "card",
		Number: "4111111111111111",
		ExpirationDate: "12/24",
		SecurityCode: "123"	,
	}

	var response *Gr4vyCard
	response, _, err = client.StorePaymentMethod(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
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
	client := NewGr4vyClient(gr4vyId, key)
	
	var response *Gr4vyPaymentMethod
	response, _, err = client.GetPaymentMethod(paymentMethodId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	
	t.Log("Retrieved paymentMethod: " + *response.Id)
}
func TestDeletePaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, err = client.DeletePaymentMethod(paymentMethodId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Deleted paymentMethod: " + paymentMethodId)
}

func TestListPaymentOptions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, _, err = client.ListPaymentOptions()	
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}

func TestListPaymentServiceDefinitions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, _, err = client.ListPaymentServiceDefinitions(nil)	
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestGetPaymentServiceDefinition(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	_, _, err = client.GetPaymentServiceDefinition("braintree-card")
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}

func TestListPaymentServices(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	var response *Gr4vyPaymentServices
	response, _, err = client.ListPaymentServices(nil)
	if err != nil {
		t.Errorf(err.Error())
		return;
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
	client := NewGr4vyClient(gr4vyId, key)

	_, _, err = client.GetPaymentService(paymentServiceId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestAddPaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyPaymentServiceRequest{
		PaymentServiceDefinitionId: "stripe-card",
		DisplayName: "Test",
		AcceptedCurrencies: []string{"GBP"},
		AcceptedCountries: []string{"GB"},
	}
	fields := []Gr4vyPaymentServiceUpdateFields{
		{Key: "secret_key", Value: "abc"},
	}
	var response *Gr4vyPaymentService
	response, _, err = client.AddPaymentService(req, fields)
	if err != nil {
		t.Errorf(err.Error())
		return;
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
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyPaymentServiceUpdate{
		DisplayName: StringPtr("Braintree UK"),
	}
	var response *Gr4vyPaymentService
	response, _, err = client.UpdatePaymentService(paymentServiceIdDelete, req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	t.Log("Updated paymentService: " + *response.Id)
}
func TestDeletePaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, err = client.DeletePaymentService(paymentServiceIdDelete)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Deleted paymentService: " + paymentServiceIdDelete)
}

func TestListCardRules(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// _, _, err = client.ListCardsRules(nil)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
}
func TestAddCardRule(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	// var req Gr4vyAddCardRule
	// helper := true
	// req.Active = &helper
	// var conditions []interface{}
	// conditions = append(conditions, Gr4vyCardRuleCondition{Match: "number", Key: "amount", Operator: ">", Value: 100})
	// req.Conditions = conditions
	// req.PaymentServiceIds = []string{paymentServiceId}
	// response, err := client.AddCardRule(req)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// var p Gr4vyCardRule
	// defer response.Body.Close()
	// err = json.NewDecoder(response.Body).Decode(&p)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// cardRuleId = *p.Id
	// t.Log("Set cardRuleId: " + cardRuleId)
}
func TestGetCardRule(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// response, err := client.GetCardRule(cardRuleId)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// var p Gr4vyCardRule
	// defer response.Body.Close()
	// err = json.NewDecoder(response.Body).Decode(&p)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// t.Log("Retrieved cardRule: " + *p.Id)
}
func TestUpdateCardRule(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// var req Gr4vyUpdateCardRule
	// helper := true
	// req.Active = &helper
	// _, err = client.UpdateCardRule(cardRuleId, req)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// t.Log("Updated cardRule: " + cardRuleId)
}
func TestDeleteCardRule(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// _, err = client.DeleteCardRule(cardRuleId)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// t.Log("Deleted cardRule: " + cardRuleId)
}

func TestListTransactions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	_, _, err = client.ListTransactions(nil)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestAuthorizeNewTransaction(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient(gr4vyId, key)
	
	req := Gr4vyTransactionRequest{
		Amount: 12.99,
		Currency: "USD",
	}

	paymentMethod := Gr4vyTransactionPaymentMethodRequest{
		Method: "card",
		Number: StringPtr("4111111111111111"),
		ExpirationDate: StringPtr("12/24"),
		SecurityCode: StringPtr("123"),
	}

	var response *Gr4vyTransaction
	response, _, err = client.AuthorizeNewTransaction(req, paymentMethod)
	if err != nil {
		t.Errorf(err.Error())
		return;
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
	client := NewGr4vyClient(gr4vyId, key)
	
	var response *Gr4vyTransaction
	response, _, err = client.GetTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	t.Log("Retrieved transaction: " + *response.Id)
}
func TestCaptureTransaction(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// var req Gr4vyCaptureTransaction
	// req.Amount = 12.99
	// req.Currency = "USD"
	// response, err := client.CaptureTransaction(transactionId, req)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// if (response.StatusCode != 200) {
	// 	t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	// }
}
func TestAuthorizeTransaction(t *testing.T) {
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
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
	t.Skip("skipping test for now")
	// key, err := GetKeyFromFile(keyPath)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return
	// }
	// client := NewGr4vyClient(gr4vyId, key)
	
	// response, err := client.RefundTransaction(transactionId)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// 	return;
	// }
	// if (response.StatusCode != 200) {
	// 	t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	// }
}