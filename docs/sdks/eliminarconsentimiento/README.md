# EliminarConsentimiento

### Available Operations

* [DeleteConsentID](#deleteconsentid) - Eliminar consentimiento AIS
* [DeleteConsentsConfirmationOfFunds](#deleteconsentsconfirmationoffunds) - Eliminar consentimiento FCS

## DeleteConsentID

Este servicio permite al TPP eliminar una solicitud de consentimiento

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
    res, err := s.EliminarConsentimiento.DeleteConsentID(ctx, operations.DeleteConsentIDRequest{
        Digest: "explicabo",
        PSUAccept: psd2cajarural.String("deserunt"),
        PSUAcceptCharset: psd2cajarural.String("distinctio"),
        PSUAcceptEncoding: psd2cajarural.String("quibusdam"),
        PSUAcceptLanguage: psd2cajarural.String("labore"),
        PSUDeviceID: psd2cajarural.String("modi"),
        PSUGeoLocation: psd2cajarural.String("qui"),
        PSUHTTPMethod: psd2cajarural.String("aliquid"),
        PSUIPAddress: psd2cajarural.String("cupiditate"),
        PSUIPPort: psd2cajarural.String("quos"),
        PSUUserAgent: psd2cajarural.String("perferendis"),
        Signature: "magni",
        TPPNokRedirectURI: psd2cajarural.String("assumenda"),
        TPPRedirectPreferred: psd2cajarural.String("ipsam"),
        TPPRedirectURI: psd2cajarural.String("alias"),
        TPPSignatureCertificate: "fugit",
        XRequestID: "dolorum",
        Aspsp: "excepturi",
        ConsentID: "tempora",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteConsentIDRequest](../../models/operations/deleteconsentidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteConsentIDResponse](../../models/operations/deleteconsentidresponse.md), error**


## DeleteConsentsConfirmationOfFunds

Este servicio permite al TPP, a través del Hub, conocer el estado en el que se encuentra un recurso de consentimiento de confirmación de fondos en el ASPSP.

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
    res, err := s.EliminarConsentimiento.DeleteConsentsConfirmationOfFunds(ctx, operations.DeleteConsentsConfirmationOfFundsRequest{
        Digest: "facilis",
        PSUAccept: psd2cajarural.String("tempore"),
        PSUAcceptCharset: psd2cajarural.String("labore"),
        PSUAcceptEncoding: psd2cajarural.String("delectus"),
        PSUAcceptLanguage: psd2cajarural.String("eum"),
        PSUDeviceID: psd2cajarural.String("non"),
        PSUGeoLocation: psd2cajarural.String("eligendi"),
        PSUHTTPMethod: psd2cajarural.String("sint"),
        PSUIPAddress: psd2cajarural.String("aliquid"),
        PSUIPPort: psd2cajarural.String("provident"),
        PSUUserAgent: psd2cajarural.String("necessitatibus"),
        Signature: "sint",
        TPPNokRedirectURI: psd2cajarural.String("officia"),
        TPPRedirectPreferred: psd2cajarural.String("dolor"),
        TPPRedirectURI: psd2cajarural.String("debitis"),
        TPPSignatureCertificate: "a",
        XRequestID: "dolorum",
        Aspsp: "in",
        ConsentID: "in",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteConsentsConfirmationOfFundsRequest](../../models/operations/deleteconsentsconfirmationoffundsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.DeleteConsentsConfirmationOfFundsResponse](../../models/operations/deleteconsentsconfirmationoffundsresponse.md), error**

