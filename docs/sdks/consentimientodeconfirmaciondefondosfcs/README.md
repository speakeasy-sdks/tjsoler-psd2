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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ConsentimientoDeConfirmacionDeFondosFCS.PostConsentsConfirmationOfFunds(ctx, operations.PostConsentsConfirmationOfFundsRequest{
        Digest: "id",
        PSUAccept: psd2cajarural.String("possimus"),
        PSUAcceptCharset: psd2cajarural.String("aut"),
        PSUAcceptEncoding: psd2cajarural.String("quasi"),
        PSUAcceptLanguage: psd2cajarural.String("error"),
        PSUCorporateID: psd2cajarural.String("temporibus"),
        PSUCorporateIDType: psd2cajarural.String("laborum"),
        PSUDeviceID: psd2cajarural.String("quasi"),
        PSUGeoLocation: psd2cajarural.String("reiciendis"),
        PSUHTTPMethod: psd2cajarural.String("voluptatibus"),
        PsuID: psd2cajarural.String("vero"),
        PSUIDType: psd2cajarural.String("nihil"),
        PSUIPAddress: "praesentium",
        PSUIPPort: psd2cajarural.String("voluptatibus"),
        PSUUserAgent: psd2cajarural.String("ipsa"),
        RequestBody: []byte("omnis"),
        Signature: "voluptate",
        TPPBrandLoggingInformation: psd2cajarural.String("cum"),
        TPPNokRedirectURI: psd2cajarural.String("perferendis"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("doloremque"),
        TPPSignatureCertificate: "reprehenderit",
        XRequestID: "ut",
        Aspsp: "maiores",
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

