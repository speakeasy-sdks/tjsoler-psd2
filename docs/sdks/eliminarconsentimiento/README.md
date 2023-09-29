# EliminarConsentimiento
(*EliminarConsentimiento*)

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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.EliminarConsentimiento.DeleteConsentID(ctx, operations.DeleteConsentIDRequest{
        Digest: "Bicycle",
        PSUAccept: tjsolerpsd2.String("sans Cambridgeshire"),
        PSUAcceptCharset: tjsolerpsd2.String("doubtfully"),
        PSUAcceptEncoding: tjsolerpsd2.String("Cambridgeshire"),
        PSUAcceptLanguage: tjsolerpsd2.String("Radon"),
        PSUDeviceID: tjsolerpsd2.String("Bike conciliate"),
        PSUGeoLocation: tjsolerpsd2.String("Fish Erie"),
        PSUHTTPMethod: tjsolerpsd2.String("online"),
        PSUIPAddress: tjsolerpsd2.String("Stream Touring architecture"),
        PSUIPPort: tjsolerpsd2.String("Architect"),
        PSUUserAgent: tjsolerpsd2.String("grey"),
        Signature: "Southeast indexing Integrated",
        TPPNokRedirectURI: tjsolerpsd2.String("however invoice"),
        TPPRedirectPreferred: tjsolerpsd2.String("withdrawal payment"),
        TPPRedirectURI: tjsolerpsd2.String("illum"),
        TPPSignatureCertificate: "Alabama bankbook unwilling",
        XRequestID: "Analyst navigate",
        Aspsp: "array Kids female",
        ConsentID: "Maserati Folk Cambridgeshire",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.EliminarConsentimiento.DeleteConsentsConfirmationOfFunds(ctx, operations.DeleteConsentsConfirmationOfFundsRequest{
        Digest: "deposit granular Money",
        PSUAccept: tjsolerpsd2.String("copying withdrawal farad"),
        PSUAcceptCharset: tjsolerpsd2.String("Security Sausages Books"),
        PSUAcceptEncoding: tjsolerpsd2.String("coffee Concrete West"),
        PSUAcceptLanguage: tjsolerpsd2.String("calculating"),
        PSUDeviceID: tjsolerpsd2.String("Gloves ratione"),
        PSUGeoLocation: tjsolerpsd2.String("Lanka from pink"),
        PSUHTTPMethod: tjsolerpsd2.String("Tuckahoe Buckinghamshire"),
        PSUIPAddress: tjsolerpsd2.String("hmph"),
        PSUIPPort: tjsolerpsd2.String("Bacon anti Shoes"),
        PSUUserAgent: tjsolerpsd2.String("innovate Exclusive"),
        Signature: "Gasoline",
        TPPNokRedirectURI: tjsolerpsd2.String("Electronic Cyrus"),
        TPPRedirectPreferred: tjsolerpsd2.String("Automotive models"),
        TPPRedirectURI: tjsolerpsd2.String("Passenger scandalise Tools"),
        TPPSignatureCertificate: "to sponge",
        XRequestID: "psst",
        Aspsp: "East River experiences",
        ConsentID: "Astatine",
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

