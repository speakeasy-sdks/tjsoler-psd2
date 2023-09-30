# LecturaDeBalances
(*LecturaDeBalances*)

### Available Operations

* [GetAccountBalances](#getaccountbalances) - Lectura de balances de una cuenta

## GetAccountBalances

Este servicio permite obtener los balances de una cuenta determinada por su identificador. Como requisito, se asume que el PSU ha dado su consentimiento para este acceso y ha sido almacenado por el ASPSP.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.LecturaDeBalances.GetAccountBalances(ctx, operations.GetAccountBalancesRequest{
        ConsentID: "ampere copying",
        Digest: "Venezuela Chevrolet taxicab",
        PSUAccept: tjsolerpsd2.String("Hybrid voluptatum benchmark"),
        PSUAcceptCharset: tjsolerpsd2.String("repudiandae monitor Russian"),
        PSUAcceptEncoding: tjsolerpsd2.String("East consequently International"),
        PSUAcceptLanguage: tjsolerpsd2.String("Accounts Dubnium lux"),
        PSUDeviceID: tjsolerpsd2.String("project Hip"),
        PSUGeoLocation: tjsolerpsd2.String("Pat"),
        PSUHTTPMethod: tjsolerpsd2.String("fairly Extended"),
        PSUIPAddress: tjsolerpsd2.String("Euclid"),
        PSUIPPort: tjsolerpsd2.String("Intelligent maroon transmit"),
        PSUUserAgent: tjsolerpsd2.String("encryption"),
        Signature: "meh functionalities Compatible",
        TPPSignatureCertificate: "Principal withdrawal kilogram",
        XRequestID: "Southwest",
        AccountID: "structure unexpectedly white",
        Aspsp: "commotion Mills french",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAccountBalancesRequest](../../models/operations/getaccountbalancesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetAccountBalancesResponse](../../models/operations/getaccountbalancesresponse.md), error**

