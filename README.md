# Gr4vy SDK for Go

<!-- [![Module Version](https://badge.fury.io/js/%40gr4vy%2Fnode.svg)][npm] -->

Gr4vy provides any of your payment integrations through one unified API. For more details, visit [gr4vy.com](https://gr4vy.com).

## Installation

To add Gr4vy to your project, add the `github.com/gr4vy/gr4vy-go` package to your project.

```sh
go get github.com/gr4vy/gr4vy-go
```

Add import:

```sh
import (
	"github.com/gr4vy/gr4vy-go"
)
```

## Getting Started

To make your first API call, you will need to [request](https://gr4vy.com) a Gr4vy instance to be set up. Please contact our sales team for a demo.

Once you have been set up with a Gr4vy account you will need to head over to the **Integrations** panel and generate a private key. We recommend storing this key in a secure location but in this code sample we simply read the file from disk.

```golang
package main

import (
  "io"
  "io/ioutil"
  "fmt"
  "github.com/gr4vy/gr4vy-go"
)

func main() {
  key, err := ioutil.ReadFile("privateKey.pem")
  if err != nil {
    fmt.Println(err)
    return
  }
  client := gr4vy.NewGr4vyClient("YOUR_GR4VY_ID", string(key))
  var params ListBuyersParams
  response, error := client.ListBuyers(params)
  
  if error != nil {
    fmt.Println(error)
    return
  }
  defer response.Body.Close()
  body, err := io.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(string(body))
  }
}
```

## Gr4vy Embed

To create a token for Gr4vy Embed, call the `client.GetEmbedToken(embed)` function with the amount, currency, and optional buyer information for Gr4vy Embed.

```golang
  embed := map[string]string{"amount": "200", "currency": "USD", "buyer_id": "d757c76a-cbd7-4b56-95a3-40125b51b29c"}
  client := gr4vy.NewGr4vyClient("YOUR_GR4VY_ID", string(key))
  responseStr, err := client.GetEmbedToken(embed)

```

You can now pass this token to your frontend where it can be used to authenticate Gr4vy Embed.

The `buyer_id` and/or `buyer_external_identifier` fields can be used to allow the token to pull in previously stored payment methods for a user. A buyer needs to be created before it can be used in this way.

```
  var req AddBuyerJSONRequestBody
  helper := string("Jane Smith")
  req.DisplayName = &helper
  response, err := client.AddBuyer(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  var p Buyer
  defer response.Body.Close()
  err = json.NewDecoder(response.Body).Decode(&p)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  embed := map[string]string{"amount": "200", "currency": "USD", "buyer_id": *p.Id}
  
  client = gr4vy.NewGr4vyClient("YOUR_GR4VY_ID", string(key))
  var responseStr string
  responseStr, err = client.GetEmbedToken(embed)
  
  if err != nil {
    fmt.Println(err)
    return;
  }
  
  fmt.Println("embed token: " + responseStr)
```

## Initialization

The client can be initialized with the Gr4vy ID (`gr4vyId`) and the private key.

```golang
  client := gr4vy.NewGr4vyClient("acme", string(key))
```

Alternatively, instead of the `gr4vyId` it can be initialized with the `baseUrl` of the server to use directly.

```golang
  client := gr4vy.NewGr4vyClientWithBaseUrl("https://api.acme.gr4vy.app", string(key))
```

Your API private key can be created in your admin panel on the **Integrations** tab.


## Making API calls

This library conveniently maps every API path to a seperate function. For example, `GET /buyers?limit=100` would be:

```golang
  var params ListBuyersParams
  helper := int32(5)
  params.Limit = &helper  
  response, error := client.ListBuyers(params)
```

To create or update a resource, the API requires a request object for that
resource that is conventiently named `<Resource>JsonRequestBody`.

For example, to create a buyer you will need to pass a `AddBuyerJsonRequestBody` object to
the `AddBuyer` method.

```golang
  var req AddBuyerJSONRequestBody
  helper := string("Jane Smith")
  req.DisplayName = &helper
  response, error := client.AddBuyer(req)
```

Similarly, to update a buyer you will need to pass in the `UpdateBuyerJsonRequestBody` to the `UpdateBuyer` method.

```golang
  var req UpdateBuyerJSONRequestBody
  helper := string("Jane Smith")
  req.DisplayName = &helper
  response, error := client.UpdateBuyer(buyer.id, req)
```

## Response 

Every resolved API call returns both a `*http.Response` object from the "net/http" 
package and an `error` object. where the response object contains a `Body` attribute 
with the parsed JSON body other attributes such as the HTTP response status code.


```golang
  response, error := client.ListBuyers(params)
  if error != nil {
    fmt.Println(error)
    return
  }

  defer response.Body.Close()
  body, err := io.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(string(body))
  }
```

## Logging & Debugging

The SDK makes it easy possible to the requests and responses to the console.

```golang
  client := gr4vy.NewGr4vyClient("YOUR_GR4VY_ID", string(key))
  client.Debug = true
```

This will output the request parameters and response to the console as follows.

```sh
  Gr4vy - Request - ListBuyers
  Gr4vy - Response - {"items":[{"id":"b8433347-a16f-46b5-958f-d681876546a6","type":"buyer","display_name":"Jane Smith","external_identifier":null,"created_at":"2021-04-22T06:51:16.910297+00:00","updated_at":"2021-04-22T07:18:49.816242+00:00"}],"limit":1,"next_cursor":"fAA0YjY5NmU2My00NzY5LTQ2OGMtOTEyNC0xODVjMDdjZTY5MzEAMjAyMS0wNC0yMlQwNjozNTowNy4yNTMxMDY","previous_cursor":null}
```

## Available APIs

* ListBuyers
* ListBuyer

## Development

### Adding new APIs

To add new APIs, run the following command to update the models and APIs based
on the API spec.

```sh
  go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
  
  download https://raw.githubusercontent.com/gr4vy/gr4vy-openapi/sdks/openapi.v1.json
  
  oapi-codegen gr4vy-openapi.v1.json > api/gr4vy.gen.go
```

```sh
  ./openapi-generator-generate.sh
```

Next, update `sdk_<object_name>.go` to bind any new APIs or remove any APIs that are no
longer available.

### Publishing

Create a PR for your change to the master branch of github.com/gr4vy/gr4vy-go

## License

This library is released under the [MIT License](LICENSE).