# ConsentimientoDeInformacionSobreCuentasDePagoAIS
(*ConsentimientoDeInformacionSobreCuentasDePagoAIS*)

### Available Operations

* [PostConsents](#postconsents) - Solicitud de consentimiento AIS

## PostConsents

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
    res, err := s.ConsentimientoDeInformacionSobreCuentasDePagoAIS.PostConsents(ctx, operations.PostConsentsRequest{
        Digest: "Account Sleek",
        PSUAccept: tjsolerpsd2.String("Maryland"),
        PSUAcceptCharset: tjsolerpsd2.String("Metal male"),
        PSUAcceptEncoding: tjsolerpsd2.String("silent"),
        PSUAcceptLanguage: tjsolerpsd2.String("withdrawal"),
        PSUCorporateID: tjsolerpsd2.String("Soap Direct"),
        PSUCorporateIDType: tjsolerpsd2.String("Director bandwidth"),
        PSUDeviceID: tjsolerpsd2.String("Shoes"),
        PSUGeoLocation: tjsolerpsd2.String("nervous Blues Borders"),
        PSUHTTPMethod: tjsolerpsd2.String("Berkshire RSS"),
        PsuID: tjsolerpsd2.String("Trial male"),
        PSUIDType: tjsolerpsd2.String("Blues"),
        PSUIPAddress: "Kyat Van Lek",
        PSUIPPort: tjsolerpsd2.String("biology"),
        PSUUserAgent: tjsolerpsd2.String("Bicycle Bedfordshire Metal"),
        RequestBody: []byte("X)68BkYEQ1"),
        Signature: "Elmo Borders Hoffman",
        TPPBrandLoggingInformation: tjsolerpsd2.String("lavender"),
        TPPNokRedirectURI: tjsolerpsd2.String("Representative even"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("fuchsia gold"),
        TPPSignatureCertificate: "round",
        XRequestID: "Northwest Tactics",
        Aspsp: "National",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PostConsentsRequest](../../models/operations/postconsentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PostConsentsResponse](../../models/operations/postconsentsresponse.md), error**

