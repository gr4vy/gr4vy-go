package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyListCardsRulesParams ListCardsRulesParams
type Gr4vyAddCardRule AddCardRuleJSONRequestBody
type Gr4vyUpdateCardRule UpdateCardRuleJSONRequestBody
type Gr4vyCardRuleNumberCondition CardRuleNumberCondition
type Gr4vyCardRuleTextCondition CardRuleTextCondition

func (c *Gr4vyClient) ListCardsRules(params Gr4vyListCardsRulesParams) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
    	return nil, err
    }
    var p ListCardsRulesParams = ListCardsRulesParams(params)
    return c.HandleResponse(client.ListCardsRules(context.TODO(), &p))
}

func (c *Gr4vyClient) GetCardRule(card_rule_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.GetCardRule(context.TODO(), card_rule_id))
}
func (c *Gr4vyClient) AddCardRule(body Gr4vyAddCardRule) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    var b AddCardRuleJSONRequestBody = AddCardRuleJSONRequestBody(body)
    return c.HandleResponse(client.AddCardRule(context.TODO(), b))
}

func (c *Gr4vyClient) UpdateCardRule(card_rule_id string, body Gr4vyUpdateCardRule) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    var b UpdateCardRuleJSONRequestBody = UpdateCardRuleJSONRequestBody(body)
    return c.HandleResponse(client.UpdateCardRule(context.TODO(), card_rule_id, b))
}

func (c *Gr4vyClient) DeleteCardRule(card_rule_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
      return nil, err
    }
    return c.HandleResponse(client.DeleteCardRule(context.TODO(), card_rule_id))
}