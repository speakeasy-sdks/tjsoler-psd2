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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ConsentimientoDeInformacionSobreCuentasDePagoAIS.PostConsents(ctx, operations.PostConsentsRequest{
        Digest: "sint",
        PSUAccept: tjsolerpsd2.String("veritatis"),
        PSUAcceptCharset: tjsolerpsd2.String("itaque"),
        PSUAcceptEncoding: tjsolerpsd2.String("incidunt"),
        PSUAcceptLanguage: tjsolerpsd2.String("enim"),
        PSUCorporateID: tjsolerpsd2.String("consequatur"),
        PSUCorporateIDType: tjsolerpsd2.String("est"),
        PSUDeviceID: tjsolerpsd2.String("quibusdam"),
        PSUGeoLocation: tjsolerpsd2.String("explicabo"),
        PSUHTTPMethod: tjsolerpsd2.String("deserunt"),
        PsuID: tjsolerpsd2.String("distinctio"),
        PSUIDType: tjsolerpsd2.String("quibusdam"),
        PSUIPAddress: "labore",
        PSUIPPort: tjsolerpsd2.String("modi"),
        PSUUserAgent: tjsolerpsd2.String("qui"),
        RequestBody: []byte("aliquid"),
        Signature: "cupiditate",
        TPPBrandLoggingInformation: tjsolerpsd2.String("quos"),
        TPPNokRedirectURI: tjsolerpsd2.String("perferendis"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("magni"),
        TPPSignatureCertificate: "assumenda",
        XRequestID: "ipsam",
        Aspsp: "alias",
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

