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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.LecturaDeBalances.GetAccountBalances(ctx, operations.GetAccountBalancesRequest{
        ConsentID: "7890-asdf-4321",
        Digest: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "TestTPPCertificate",
        XRequestID: "<value>",
        AccountID: "<value>",
        Aspsp: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAccountBalancesRequest](../../pkg/models/operations/getaccountbalancesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAccountBalancesResponse](../../pkg/models/operations/getaccountbalancesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
