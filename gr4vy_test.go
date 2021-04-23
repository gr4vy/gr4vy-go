package gr4vy

import (
	"encoding/json"
	"testing"
	"strconv"
)

const keyPath = "./test/gr4vy-private-key-kH4ndJM8QHyfYGpKdgfAhq2j9RySIOWa3kY_urrI5PI.pem"

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
	client := NewGr4vyClient("demo", string(key))
	
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
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyAddBuyer
	helper := string("Jane Smith")
	req.DisplayName = &helper
	response, err := client.AddBuyer(req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var p Gr4vyBuyer
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	buyerId = *p.Id
	embed := map[string]string{"amount": "200", "currency": "USD", "buyer_id": *p.Id}

	client = NewGr4vyClient("demo", string(key))
	
	_, err = client.GetEmbedToken(embed)

	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	//Clean up
	client = NewGr4vyClient("demo", string(key))
	
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
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListBuyersParams
	helper := int32(5)
    params.Limit = &helper  
	_, err = client.ListBuyers(params)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestAddBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyAddBuyer
	helper := string("Jane Smith")
	req.DisplayName = &helper
	response, err := client.AddBuyer(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyBuyer
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	buyerId = *p.Id
	t.Log("Set buyerId: " + buyerId)
}
func TestGetBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.GetBuyer(buyerId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	var p Gr4vyBuyer
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Retrieved buyer: " + *p.Id)
}
func TestUpdateBuyer(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyUpdateBuyer
	helper := string("Jane Smith")
	req.DisplayName = &helper
	_, err = client.UpdateBuyer(buyerId, req)
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
	client := NewGr4vyClient("demo", string(key))
	
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
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListPaymentMethodsParams
	_, err = client.ListPaymentMethods(params)
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
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyStorePaymentMethod
	req.Method = "card"
	req.Number = "4111111111111111"
	req.ExpirationDate = "12/24"
	req.SecurityCode = "123"
	response, err := client.StorePaymentMethod(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyCard
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	paymentMethodId = *p.Id
	t.Log("Set paymentMethodId: " + paymentMethodId)
}
func TestGetPaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.GetPaymentMethod(paymentMethodId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	var p Gr4vyCard
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Retrieved paymentMethod: " + *p.Id)
}
func TestDeletePaymentMethod(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
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
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListPaymentOptionsParams
	_, err = client.ListPaymentOptions(params)	
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
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListPaymentServiceDefinitionsParams
	_, err = client.ListPaymentServiceDefinitions(params)	
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
	client := NewGr4vyClient("demo", string(key))
	
	_, err = client.GetPaymentServiceDefinition("braintree-card")
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
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListPaymentServicesParams
	response, err := client.ListPaymentServices(params)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyPaymentServices
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	
	paymentServiceId = *(*p.Items)[0].Id
	t.Log("Set paymentServiceId: " + paymentServiceId)
}
func TestGetPaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	_, err = client.GetPaymentService(paymentServiceId)
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
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyAddPaymentService
	helper := string("stripe-card")
	req.PaymentServiceDefinitionId = &helper
	helper1 := string("Test")
	req.DisplayName = &helper1
	helper2 := []string{"GBP"}
	req.AcceptedCurrencies = &helper2
	helper3 := []string{"GB"}
	req.AcceptedCountries = &helper3

	fields := &[]struct {
		Key string `json:"key"`
		Value string `json:"value"`
	}{
		{Key: "secret_key", Value: "abc"},
	}
	req.Fields = fields
	response, err := client.AddPaymentService(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyPaymentService
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	
	paymentServiceIdDelete = *p.Id
	t.Log("Set paymentServiceIdDelete: " + paymentServiceIdDelete)
}
func TestUpdatePaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyUpdatePaymentService
	helper := string("Braintree UK")
	req.DisplayName = &helper
	response, err := client.UpdatePaymentService(paymentServiceIdDelete, req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyPaymentService
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Updated paymentService: " + *p.Id)
}
func TestDeletePaymentService(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	_, err = client.DeletePaymentService(paymentServiceIdDelete)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Deleted paymentService: " + paymentServiceIdDelete)
}

func TestListCardRules(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListCardsRulesParams
	_, err = client.ListCardsRules(params)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
}
func TestAddCardRule(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	var req Gr4vyAddCardRule
	helper := true
	req.Active = &helper
	var conditions []interface{}
	conditions = append(conditions, Gr4vyCardRuleNumberCondition{Match: "number", Key: "amount", Operator: ">", Value: 100})
	req.Conditions = conditions
	req.PaymentServiceIds = []string{paymentServiceId}
	response, err := client.AddCardRule(req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	var p Gr4vyCardRule
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	cardRuleId = *p.Id
	t.Log("Set cardRuleId: " + cardRuleId)
}
func TestGetCardRule(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.GetCardRule(cardRuleId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	var p Gr4vyCardRule
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log("Retrieved cardRule: " + *p.Id)
}
func TestUpdateCardRule(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyUpdateCardRule
	helper := true
	req.Active = &helper
	_, err = client.UpdateCardRule(cardRuleId, req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Updated cardRule: " + cardRuleId)
}
func TestDeleteCardRule(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	_, err = client.DeleteCardRule(cardRuleId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	t.Log("Deleted cardRule: " + cardRuleId)
}

func TestListTransactions(t *testing.T) {
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var params Gr4vyListTransactionsParams
	_, err = client.ListTransactions(params)
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
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyAuthorizeNewTransaction
	req.Amount = 12.99
	req.Currency = "USD"
	var paymentMethod Gr4vyPaymentMethod
	paymentMethod.Method = "card"
	paymentMethod.Number = "4111111111111111"
	paymentMethod.ExpirationDate = "12/24"
	paymentMethod.SecurityCode = "123"
	response, err := client.AuthorizeNewTransaction(req, paymentMethod)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyTransaction
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	
	transactionId = *p.Id
	t.Log("Set transactionId: " + transactionId)
}
func TestGetTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.GetTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}

	var p Gr4vyTransaction
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Log("Retrieved transaction: " + *p.Id)
}
func TestCaptureTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	var req Gr4vyCaptureTransaction
	req.Amount = 12.99
	req.Currency = "USD"
	response, err := client.CaptureTransaction(transactionId, req)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	if (response.StatusCode != 200) {
		t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	}
}
func TestAuthorizeTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.AuthorizeTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	if (response.StatusCode != 200) {
		t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	}
}
func TestRefundTransaction(t *testing.T) {
	// t.Skip("skipping test for now")
	key, err := GetKeyFromFile(keyPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	client := NewGr4vyClient("demo", string(key))
	
	response, err := client.RefundTransaction(transactionId)
	if err != nil {
		t.Errorf(err.Error())
		return;
	}
	if (response.StatusCode != 200) {
		t.Errorf("expected StatusCode 200: received: " + strconv.Itoa(response.StatusCode))
	}
}