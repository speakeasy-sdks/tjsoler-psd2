# ConsentimientoDeInformacionSobreCuentasDePagoAIS

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ConsentimientoDeInformacionSobreCuentasDePagoAIS.PostConsents(ctx, operations.PostConsentsRequest{
        Digest: "dicta",
        PSUAccept: psd2cajarural.String("corporis"),
        PSUAcceptCharset: psd2cajarural.String("dolore"),
        PSUAcceptEncoding: psd2cajarural.String("iusto"),
        PSUAcceptLanguage: psd2cajarural.String("dicta"),
        PSUCorporateID: psd2cajarural.String("harum"),
        PSUCorporateIDType: psd2cajarural.String("enim"),
        PSUDeviceID: psd2cajarural.String("accusamus"),
        PSUGeoLocation: psd2cajarural.String("commodi"),
        PSUHTTPMethod: psd2cajarural.String("repudiandae"),
        PsuID: psd2cajarural.String("quae"),
        PSUIDType: psd2cajarural.String("ipsum"),
        PSUIPAddress: "quidem",
        PSUIPPort: psd2cajarural.String("molestias"),
        PSUUserAgent: psd2cajarural.String("excepturi"),
        RequestBody: []byte("pariatur"),
        Signature: "modi",
        TPPBrandLoggingInformation: psd2cajarural.String("praesentium"),
        TPPNokRedirectURI: psd2cajarural.String("rem"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("voluptates"),
        TPPSignatureCertificate: "quasi",
        XRequestID: "repudiandae",
        Aspsp: "sint",
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

