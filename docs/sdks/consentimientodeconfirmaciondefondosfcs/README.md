# ConsentimientoDeConfirmacionDeFondosFCS
(*ConsentimientoDeConfirmacionDeFondosFCS*)

### Available Operations

* [PostConsentsConfirmationOfFunds](#postconsentsconfirmationoffunds) - Solicitud de consentimiento FCS

## PostConsentsConfirmationOfFunds

Con este servicio, un TPP a través del HUB puede solicitar un consentimiento para acceder a las cuentas del PSU. Esta solicitud puede ser sobre unas cuentas indicadas o no.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ConsentimientoDeConfirmacionDeFondosFCS.PostConsentsConfirmationOfFunds(ctx, operations.PostConsentsConfirmationOfFundsRequest{
        Digest: "string",
        PSUIPAddress: "string",
        RequestBody: []byte("0x2c4f80e8fF"),
        Signature: "string",
        TPPSignatureCertificate: "string",
        XRequestID: "string",
        Aspsp: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PostConsentsConfirmationOfFundsRequest](../../pkg/models/operations/postconsentsconfirmationoffundsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PostConsentsConfirmationOfFundsResponse](../../pkg/models/operations/postconsentsconfirmationoffundsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
