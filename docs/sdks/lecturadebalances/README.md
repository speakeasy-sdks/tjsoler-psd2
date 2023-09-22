# LecturaDeBalances

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
        ConsentID: "suscipit",
        Digest: "deserunt",
        PSUAccept: tjsolerpsd2.String("provident"),
        PSUAcceptCharset: tjsolerpsd2.String("minima"),
        PSUAcceptEncoding: tjsolerpsd2.String("repellendus"),
        PSUAcceptLanguage: tjsolerpsd2.String("totam"),
        PSUDeviceID: tjsolerpsd2.String("similique"),
        PSUGeoLocation: tjsolerpsd2.String("alias"),
        PSUHTTPMethod: tjsolerpsd2.String("at"),
        PSUIPAddress: tjsolerpsd2.String("quaerat"),
        PSUIPPort: tjsolerpsd2.String("tempora"),
        PSUUserAgent: tjsolerpsd2.String("vel"),
        Signature: "quod",
        TPPSignatureCertificate: "officiis",
        XRequestID: "qui",
        AccountID: "dolorum",
        Aspsp: "a",
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

