package gr4vy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	. "github.com/gr4vy/gr4vy-go/api"
)

const VERSION = "0.24.0"

type Gr4vyClient struct {
	gr4vyId     string
	privateKey  string
	baseUrl     string
	Debug       bool
	accessToken string
	environment string
	merchantAccountId string
}

type EmbedParams struct {
	Amount   int32             `json:"amount"`
	Currency string            `json:"currency"`
	BuyerID  string            `json:"buyer_id"`
	Metadata map[string]string `json:"metadata"`
}

func NewGr4vyClient(gr4vy_id string, private_key string, environment string) *Gr4vyClient {
	client := Gr4vyClient{
		gr4vyId:     gr4vy_id,
		privateKey:  private_key,
		Debug:       false,
		environment: environment,
		merchantId: "default",
	}
	return &client
}

func NewGr4vyClientWithMid(gr4vy_id string, private_key string, environment string, merchant_id string) *Gr4vyClient {
	client := Gr4vyClient{
		gr4vyId:     gr4vy_id,
		privateKey:  private_key,
		Debug:       false,
		environment: environment,
		merchantId: merchant_id,
	}
	return &client
}

func NewGr4vyClientWithBaseUrl(base_url string, private_key string) *Gr4vyClient {
	client := Gr4vyClient{
		privateKey:  private_key,
		baseUrl:     base_url,
		Debug:       false,
		environment: "sandbox",
		merchantId: "default",
	}
	return &client
}

func GetClient(c *Gr4vyClient) (*APIClient, error) {

	if c.Debug {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Println("Gr4vy - Request - " + lastString(strings.Split(funcName, ".")))
	}

	bearerToken, err := authentication(c.privateKey, c.Debug)
	if err != nil {
		return nil, err
	}
	c.accessToken = bearerToken

	config := NewConfiguration()
	config.Servers[0].URL = c.BaseUrl()
	config.AddDefaultHeader("X-GR4VY-MERCHANT-ACCOUNT-ID", c.merchantId)

	client := NewAPIClient(config)

	return client, nil
}

func (c *Gr4vyClient) GetEmbedToken(params EmbedParams) (string, error) {
	return getEmbedToken(c.privateKey, params)
}

func (c *Gr4vyClient) GetToken() (string, error) {
	scopes := []string{"*.read", "*.write"}
	return getToken(c.privateKey, scopes)
}

func (c *Gr4vyClient) HandleResponse(response *http.Response, error error) {
	if c.Debug {
		var sb strings.Builder
		sb.WriteString("Gr4vy - Response - ")

		if error != nil {
			sb.WriteString("error: " + error.Error())
			body, err := io.ReadAll(response.Body)
			if err == nil {
				sb.WriteString(string(body))
				response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				sb.WriteString("error: " + err.Error())
			} else {
				sb.WriteString(string(body))
				response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		fmt.Println(sb.String())
	}
}

func (c *Gr4vyClient) BaseUrl() string {
	if c.baseUrl != "" {
		return c.baseUrl
	}

	var baseUrl string
	if c.environment == "sandbox" {
		baseUrl = fmt.Sprintf("https://api.%v.%v.gr4vy.app", c.environment, c.gr4vyId)
	} else {
		baseUrl = fmt.Sprintf("https://api.%v.gr4vy.app", c.gr4vyId)
	}

	return baseUrl
}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}

func GetKeyFromFile(fileName string) (string, error) {
	value, exists := os.LookupEnv("PRIVATE_KEY")
	if exists {
		return string(value), nil
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func String(v string) NullableString {
	return *NewNullableString(&v)
}

func StringPtr(v string) *string {
	return &v
}

func Int32(v int32) *int32 {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}
