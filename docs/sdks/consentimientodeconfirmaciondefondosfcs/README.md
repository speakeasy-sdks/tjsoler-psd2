# ConsentimientoDeConfirmacionDeFondosFCS

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
        Digest: "maiores",
        PSUAccept: tjsolerpsd2.String("dicta"),
        PSUAcceptCharset: tjsolerpsd2.String("corporis"),
        PSUAcceptEncoding: tjsolerpsd2.String("dolore"),
        PSUAcceptLanguage: tjsolerpsd2.String("iusto"),
        PSUCorporateID: tjsolerpsd2.String("dicta"),
        PSUCorporateIDType: tjsolerpsd2.String("harum"),
        PSUDeviceID: tjsolerpsd2.String("enim"),
        PSUGeoLocation: tjsolerpsd2.String("accusamus"),
        PSUHTTPMethod: tjsolerpsd2.String("commodi"),
        PsuID: tjsolerpsd2.String("repudiandae"),
        PSUIDType: tjsolerpsd2.String("quae"),
        PSUIPAddress: "ipsum",
        PSUIPPort: tjsolerpsd2.String("quidem"),
        PSUUserAgent: tjsolerpsd2.String("molestias"),
        RequestBody: []byte("excepturi"),
        Signature: "pariatur",
        TPPBrandLoggingInformation: tjsolerpsd2.String("modi"),
        TPPNokRedirectURI: tjsolerpsd2.String("praesentium"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("rem"),
        TPPSignatureCertificate: "voluptates",
        XRequestID: "quasi",
        Aspsp: "repudiandae",
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

