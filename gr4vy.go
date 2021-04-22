package gr4vy

import (
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
	"net/http"
	"runtime"
	"strings"
	. "github.com/steve-gr4vy/gr4vy/api"
)

const VERSION = "0.0.1"

type Gr4vyClient struct {
	gr4vyId string
	privateKey string
	baseUrl string
	Debug bool
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

func GetClient(c *Gr4vyClient) (*Client, error) {

	if c.Debug {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
      	fmt.Println("Gr4vy - Request - " + lastString(strings.Split(funcName, ".")))
    }

	authRequestEditorFn, err := authentication(c.privateKey, c.Debug)

    if err != nil {
    	return nil, err
    }


	client, clientErr := NewClient(c.BaseUrl(), 
		[]ClientOption{
			WithRequestEditorFn(authRequestEditorFn),
		}...,
	)

	if clientErr != nil {
		return nil, clientErr
	}

	return client, nil
}

func (c *Gr4vyClient) GetEmbedToken(params map[string]string) (string, error) {
	scopes := []string{"embed"}
	return getToken(c.privateKey, scopes, params)
}

func (c *Gr4vyClient) HandleResponse(response *http.Response, error error) (*http.Response, error) {
	if c.Debug {
		var sb strings.Builder
		sb.WriteString("Gr4vy - Response - ")

		if error != nil {
			sb.WriteString("error: " + error.Error())
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
	return response, error
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

func GetKeyFromFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
