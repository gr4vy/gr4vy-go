package gr4vy

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	. "github.com/gr4vy/gr4vy-go/api"
)

const VERSION = "0.33.0"

type Gr4vyClient struct {
	gr4vyId           string
	privateKey        string
	baseUrl           string
	Debug             bool
	accessToken       string
	environment       string
	merchantAccountId string
}

type EmbedParams struct {
	Amount            int32                  `json:"amount"`
	Currency          string                 `json:"currency"`
	BuyerID           string                 `json:"buyer_id"`
	Metadata          map[string]string      `json:"metadata"`
	ConnectionOptions map[string]interface{} `json:"connection_options"`
	MerchantAccountId string                 `json:"merchant_account_id"`
}

func NewGr4vyClient(gr4vy_id string, private_key string, environment string) *Gr4vyClient {
	client := Gr4vyClient{
		gr4vyId:           gr4vy_id,
		privateKey:        private_key,
		Debug:             false,
		environment:       environment,
		merchantAccountId: "default",
	}
	return &client
}

func NewGr4vyClientWithMid(gr4vy_id string, private_key string, environment string, merchant_account_id string) *Gr4vyClient {
	client := Gr4vyClient{
		gr4vyId:           gr4vy_id,
		privateKey:        private_key,
		Debug:             false,
		environment:       environment,
		merchantAccountId: merchant_account_id,
	}
	return &client
}

func NewGr4vyClientWithBaseUrl(base_url string, private_key string) *Gr4vyClient {
	client := Gr4vyClient{
		privateKey:        private_key,
		baseUrl:           base_url,
		Debug:             false,
		environment:       "sandbox",
		merchantAccountId: "default",
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
	config.AddDefaultHeader("X-GR4VY-MERCHANT-ACCOUNT-ID", c.merchantAccountId)

	client := NewAPIClient(config)

	return client, nil
}

func (c *Gr4vyClient) GetEmbedToken(params EmbedParams, checkout_session_id string) (string, error) {
	if params.MerchantAccountId == "" {
		params.MerchantAccountId = c.merchantAccountId
	}
	return getEmbedToken(c.privateKey, params, checkout_session_id)
}

func (c *Gr4vyClient) GetEmbedTokenWithCheckoutSession(params EmbedParams) (string, CheckoutSession, error) {
	req := Gr4vyCheckoutSessionCreateRequest{}
	checkoutSession, _, _ := c.AddCheckoutSession(req)
	token, err := getEmbedToken(c.privateKey, params, *checkoutSession.Id)
	return token, CheckoutSession(*checkoutSession), err
}

func (c *Gr4vyClient) GetToken() (string, error) {
	scopes := []string{"*.read", "*.write"}
	return getToken(c.privateKey, scopes)
}

// VerifyWebhook verifies the webhook signature and timestamp.
func VerifyWebhook(secret string, payload string, signatureHeader *string, timestampHeader *string, timestampTolerance int) error {
	if signatureHeader == nil || timestampHeader == nil {
		return errors.New("missing header values")
	}

	timestamp, err := strconv.Atoi(*timestampHeader)
	if err != nil {
		return errors.New("invalid header timestamp")
	}

	signatures := strings.Split(*signatureHeader, ",")
	expectedSignature := computeHMAC(secret, strconv.Itoa(timestamp)+"."+payload)

	if !contains(signatures, expectedSignature) {
		return errors.New("no matching signature found")
	}

	if timestampTolerance > 0 && timestamp < int(time.Now().Unix())-timestampTolerance {
		return errors.New("timestamp too old")
	}

	return nil
}

// computeHMAC computes the HMAC-SHA256 signature.
func computeHMAC(secret string, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
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
				response.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				sb.WriteString("error: " + err.Error())
			} else {
				sb.WriteString(string(body))
				response.Body = io.NopCloser(bytes.NewBuffer(body))
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
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func String(v string) NullableString {
	return *NewNullableString(&v)
}

func Gr4vyNullableInt32(v int32) NullableInt32 {
	return *NewNullableInt32(&v)
}

func Gr4vyNullableString(v string) NullableString {
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
