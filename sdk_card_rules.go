package gr4vy

import (
	"net/http"
	"golang.org/x/net/context"
	. "github.com/gr4vy/gr4vy-go/api"
)

type Gr4vyCardRuleCondition CardRuleCondition
type Gr4vyCardRuleRequest CardRuleRequest
type Gr4vyCardRuleUpdate CardRuleUpdate
type Gr4vyCardRule CardRule
type Gr4vyCardRules CardRules

func (c *Gr4vyClient) ListCardsRules(limit *int32) (*Gr4vyCardRules, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }

    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CardRulesApi.ListCardsRules(auth)
    if (limit != nil) {
        p.Limit(*limit)
    }
    
    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCardRules = Gr4vyCardRules(response)
    return &r, http, err
}

func (c *Gr4vyClient) GetCardRule(card_rule_id string) (*Gr4vyCardRule, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CardRulesApi.GetCardRule(auth, card_rule_id)

    response, http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCardRule = Gr4vyCardRule(response)
    return &r, http, err
}

func (c *Gr4vyClient) AddCardRule(body Gr4vyCardRuleRequest, conditions []Gr4vyCardRuleCondition) (*Gr4vyCardRule, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CardRulesApi.AddCardRule(auth)

    var b CardRuleRequest = CardRuleRequest(body)
    var cs []CardRuleCondition
    for _, element := range conditions {
        var c CardRuleCondition = CardRuleCondition(element)
        cs = append(cs, c)
    }
    b.Conditions = []CardRule{{Conditions: &cs}}
    response, http, err := p.CardRuleRequest(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCardRule = Gr4vyCardRule(response)
    return &r, http, err
}

func (c *Gr4vyClient) UpdateCardRule(card_rule_id string, body Gr4vyCardRuleUpdate) (*Gr4vyCardRule, *http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CardRulesApi.UpdateCardRule(auth, card_rule_id)

    var b CardRuleUpdate = CardRuleUpdate(body)
    response, http, err := p.CardRuleUpdate(b).Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return nil, http, err
    }
    var r Gr4vyCardRule = Gr4vyCardRule(response)
    return &r, http, err
}

func (c *Gr4vyClient) DeleteCardRule(card_rule_id string) (*http.Response, error) {
    client, err := GetClient(c)
    if err != nil {
        return nil, err
    }
    auth := context.WithValue(context.Background(), ContextAccessToken, c.accessToken)
    p := client.CardRulesApi.DeleteCardRule(auth, card_rule_id)

    http, err := p.Execute()
    c.HandleResponse(http, err)
    if (err != nil) {
        return http, err
    }
    return http, err
}