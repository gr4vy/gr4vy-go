package gr4vy

import (
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
	"net/http"
	"runtime"
	"strings"
	. "github.com/gr4vy/gr4vy-go/api"
)

const VERSION = "0.1.0"

type Gr4vyClient struct {
	gr4vyId string
	privateKey string
	baseUrl string
	Debug bool
	accessToken string
}

func NewGr4vyClient(gr4vy_id string, private_key string) (*Gr4vyClient) {
	client := Gr4vyClient{
		gr4vyId: gr4vy_id,
		privateKey: private_key,
		Debug: false,
	}
	return &client
}

func NewGr4vyClientWithBaseUrl(base_url string, private_key string) (*Gr4vyClient) {
	client := Gr4vyClient{
		privateKey: private_key,
		baseUrl: base_url,
		Debug: false,
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

	client := NewAPIClient(config)

	return client, nil
}

func (c *Gr4vyClient) GetEmbedToken(params map[string]interface{}) (string, error) {
	scopes := []string{"embed"}
	return getToken(c.privateKey, scopes, params)
}

func (c *Gr4vyClient) GetToken() (string, error) {
	scopes := []string{"*.read", "*.write"}
	return getToken(c.privateKey, scopes, nil)
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
   return fmt.Sprintf("https://api.%v.gr4vy.app", c.gr4vyId)
}

func lastString(ss []string) string {
    return ss[len(ss)-1]
}

func GetKeyFromFile(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func String(v string) (NullableString) {
	return *NewNullableString(&v)
}

func StringPtr(v string) (*string) {
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

