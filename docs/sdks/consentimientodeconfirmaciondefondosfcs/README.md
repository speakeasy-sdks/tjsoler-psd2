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
        PSUAccept: tjsolerpsd2.String("ability Strategist Fundamental"),
        PSUAcceptCharset: tjsolerpsd2.String("Ball North parsing"),
        PSUAcceptEncoding: tjsolerpsd2.String("Fermium"),
        PSUAcceptLanguage: tjsolerpsd2.String("engage transmitter Blues"),
        PSUCorporateID: tjsolerpsd2.String("Sri"),
        PSUCorporateIDType: tjsolerpsd2.String("payment"),
        PSUDeviceID: tjsolerpsd2.String("Bellingham"),
        PSUGeoLocation: tjsolerpsd2.String("Connecticut Road blue"),
        PSUHTTPMethod: tjsolerpsd2.String("Games state male"),
        PsuID: tjsolerpsd2.String("protocol"),
        PSUIDType: tjsolerpsd2.String("hack Louisiana Dollar"),
        PSUIPAddress: "kelvin Southeast",
        PSUIPPort: tjsolerpsd2.String("Aston"),
        PSUUserAgent: tjsolerpsd2.String("Coordinator"),
        RequestBody: []byte("@:.pXsCO\b"),
        Signature: "lime Turkish Bicycle",
        TPPBrandLoggingInformation: tjsolerpsd2.String("copy Louisiana Electric"),
        TPPNokRedirectURI: tjsolerpsd2.String("whenever deposit"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("nervously convergence empower"),
        TPPSignatureCertificate: "Money",
        XRequestID: "invoice Northwest",
        Aspsp: "California Pants",
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

