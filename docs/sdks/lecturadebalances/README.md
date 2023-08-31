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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.LecturaDeBalances.GetAccountBalances(ctx, operations.GetAccountBalancesRequest{
        ConsentID: "omnis",
        Digest: "necessitatibus",
        PSUAccept: psd2cajarural.String("distinctio"),
        PSUAcceptCharset: psd2cajarural.String("asperiores"),
        PSUAcceptEncoding: psd2cajarural.String("nihil"),
        PSUAcceptLanguage: psd2cajarural.String("ipsum"),
        PSUDeviceID: psd2cajarural.String("voluptate"),
        PSUGeoLocation: psd2cajarural.String("id"),
        PSUHTTPMethod: psd2cajarural.String("saepe"),
        PSUIPAddress: psd2cajarural.String("eius"),
        PSUIPPort: psd2cajarural.String("aspernatur"),
        PSUUserAgent: psd2cajarural.String("perferendis"),
        Signature: "amet",
        TPPSignatureCertificate: "optio",
        XRequestID: "accusamus",
        AccountID: "ad",
        Aspsp: "saepe",
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

