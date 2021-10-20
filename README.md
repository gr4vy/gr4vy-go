# Gr4vy SDK for Go

Gr4vy provides any of your payment integrations through one unified API. For
more details, visit [gr4vy.com](https://gr4vy.com).

## Installation

To add Gr4vy to your project, add the `github.com/gr4vy/gr4vy-go` package to
your project.

```sh
go get github.com/gr4vy/gr4vy-go
```

Add import:

```golang
import "github.com/gr4vy/gr4vy-go"
```

## Getting Started

To make your first API call, you will need to [request](https://gr4vy.com) a
Gr4vy instance to be set up. Please contact our sales team for a demo.

Once you have been set up with a Gr4vy account you will need to head over to the
**Integrations** panel and generate a private key. We recommend storing this key
in a secure location but in this code sample we simply read the file from disk.

```golang
package main

import (
  "io"
  "io/ioutil"
  "fmt"
  "github.com/gr4vy/gr4vy-go"
)

func main() {
  key, err := gr4vy.GetKeyFromFile(PRIVATE_KEY_FILENAME)
  if err != nil {
    fmt.Println(err)
    return
  }
  client := gr4vy.NewGr4vyClient("demo", key)
  client.Debug = true

  var response *gr4vy.Gr4vyBuyers
  response, _, err = client.ListBuyers(gr4vy.Int32(2))
  if err != nil {
    fmt.Println(err.Error())
    return;
  }
  fmt.Printf("%+v\n", (*response.Items)[0].GetId())
}
```

## Gr4vy Embed

To create a token for Gr4vy Embed, call the `client.GetEmbedToken(embed)`
function with the amount, currency, and optional buyer information for Gr4vy
Embed.

```golang
embed := gr4vy.EmbedParams{
  Amount:   200,
  Currency: "USD",
  BuyerID:  "d757c76a-cbd7-4b56-95a3-40125b51b29c",
}
token, err = client.GetEmbedToken(embed)
```

You can now pass this token to your frontend where it can be used to
authenticate Gr4vy Embed.

The `buyer_id` and/or `buyer_external_identifier` fields can be used to allow
the token to pull in previously stored payment methods for a user. A buyer
needs to be created before it can be used in this way.

```golang
  key, err := gr4vy.GetKeyFromFile(PRIVATE_KEY)
  if err != nil {
    fmt.Println(err)
    return
  }
  client := gr4vy.NewGr4vyClient("demo", key)
  client.Debug = true
  req := gr4vy.Gr4vyBuyerRequest{
    DisplayName: gr4vy.String("Jane Smith"),
  }
  var response *gr4vy.Gr4vyBuyer
  response, _, err = client.AddBuyer(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  embed := gr4vy.EmbedParams{
    Amount:   200,
    Currency: "USD",
    BuyerID:  (*response.Id),
  }
  client = gr4vy.NewGr4vyClient("demo", key)
  client.Debug = true
  var responseStr string
  responseStr, err = client.GetEmbedToken(embed)

  if err != nil {
    fmt.Println(err)
    return;
  }

  fmt.Println("embed token: " + responseStr)
```

## Initialization

The client can be initialized with the Gr4vy ID (`gr4vyId`) and the private key
string.

```golang
  client := gr4vy.NewGr4vyClient("acme", key)
```

Alternatively, instead of the `gr4vyId` it can be initialized with the `baseUrl`
of the server to use directly.

```golang
  client := gr4vy.NewGr4vyClientWithBaseUrl("https://api.acme.gr4vy.app", key)
```

Your API private key can be created in your admin panel on the **Integrations**
tab.


## Making API calls

This library conveniently maps every API path to a seperate function. For
example, `GET /buyers?limit=100` would be:

```golang
  response, _, error := client.ListBuyers(2)
```

To create, the API requires a request object for that resource that is conventiently
named `Gr4vy<Resource>Request`.  To update, the API requires a request object
for that resource that is named `Gr4vy<Resource>Update`.

For example, to create a buyer you will need to pass a `Gr4vyBuyerRequest` object to
the `AddBuyer` method.

```golang
  req := gr4vy.Gr4vyBuyerRequest{
    DisplayName: gr4vy.String("Jane Smith"),
  }
  response, _, error := client.AddBuyer(req)
```

So to update a buyer you will need to pass in the `Gr4vyBuyerUpdate` to the
`UpdateBuyer` method.

```golang
  req := gr4vy.Gr4vyBuyerUpdate{
    DisplayName: gr4vy.String("Janet Smith"),
  }
  response, err := client.UpdateBuyer(buyerId, req)
```

## Response

Every resolved API call returns the requested resource, a `*http.Response`
object from the "net/http" package and an `error` object.


```golang
  var response *gr4vy.Gr4vyBuyers
  response, http, err = client.ListBuyers(gr4vy.Int32(2))
  if err != nil {
    fmt.Println(err.Error())
    return;
  }
```

## Logging & Debugging

The SDK makes it easy possible to the requests and responses to the console.

```golang
  client := gr4vy.NewGr4vyClient("YOUR_GR4VY_ID", key)
  client.Debug = true
```

This will output the request parameters and response to the console as follows.

```sh
Gr4vy - Request - ListBuyers
Gr4vy - Response - {"items":[{"id":"b8433347-a16f-46b5-958f-d681876546a6","type":"buyer","display_name":"Jane Smith","external_identifier":null,"created_at":"2021-04-22T06:51:16.910297+00:00","updated_at":"2021-04-22T07:18:49.816242+00:00"}],"limit":1,"next_cursor":"fAA0YjY5NmU2My00NzY5LTQ2OGMtOTEyNC0xODVjMDdjZTY5MzEAMjAyMS0wNC0yMlQwNjozNTowNy4yNTMxMDY","previous_cursor":null}
```

## Development

### Adding new APIs

To add new APIs, run the following command to update the models and APIs based
on the API spec.

```sh
./openapi-generator-generate.sh
```

Next, update `sdk_<object_name>.go` to bind any new APIs or remove any APIs that are no
longer available.

Run the tests to ensure the changes do not break any existing tests.

```sh
go test -v
```

### Publishing

Create a PR for your change to the master branch of github.com/gr4vy/gr4vy-go

## License

This library is released under the [MIT License](LICENSE).
