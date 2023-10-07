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
        Digest: "neque",
        PSUIPAddress: "ability Strategist Fundamental",
        RequestBody: []byte("n4:!0Qz2eo"),
        Signature: "engage transmitter Blues",
        TPPSignatureCertificate: "Sri",
        XRequestID: "payment",
        Aspsp: "Bellingham",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PostConsentsConfirmationOfFundsRequest](../../models/operations/postconsentsconfirmationoffundsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PostConsentsConfirmationOfFundsResponse](../../models/operations/postconsentsconfirmationoffundsresponse.md), error**

