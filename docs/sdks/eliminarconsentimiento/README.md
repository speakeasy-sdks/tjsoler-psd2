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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.EliminarConsentimiento.DeleteConsentID(ctx, operations.DeleteConsentIDRequest{
        Digest: "delectus",
        PSUAccept: tjsolerpsd2.String("eum"),
        PSUAcceptCharset: tjsolerpsd2.String("non"),
        PSUAcceptEncoding: tjsolerpsd2.String("eligendi"),
        PSUAcceptLanguage: tjsolerpsd2.String("sint"),
        PSUDeviceID: tjsolerpsd2.String("aliquid"),
        PSUGeoLocation: tjsolerpsd2.String("provident"),
        PSUHTTPMethod: tjsolerpsd2.String("necessitatibus"),
        PSUIPAddress: tjsolerpsd2.String("sint"),
        PSUIPPort: tjsolerpsd2.String("officia"),
        PSUUserAgent: tjsolerpsd2.String("dolor"),
        Signature: "debitis",
        TPPNokRedirectURI: tjsolerpsd2.String("a"),
        TPPRedirectPreferred: tjsolerpsd2.String("dolorum"),
        TPPRedirectURI: tjsolerpsd2.String("in"),
        TPPSignatureCertificate: "in",
        XRequestID: "illum",
        Aspsp: "maiores",
        ConsentID: "rerum",
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
        Digest: "dicta",
        PSUAccept: tjsolerpsd2.String("magnam"),
        PSUAcceptCharset: tjsolerpsd2.String("cumque"),
        PSUAcceptEncoding: tjsolerpsd2.String("facere"),
        PSUAcceptLanguage: tjsolerpsd2.String("ea"),
        PSUDeviceID: tjsolerpsd2.String("aliquid"),
        PSUGeoLocation: tjsolerpsd2.String("laborum"),
        PSUHTTPMethod: tjsolerpsd2.String("accusamus"),
        PSUIPAddress: tjsolerpsd2.String("non"),
        PSUIPPort: tjsolerpsd2.String("occaecati"),
        PSUUserAgent: tjsolerpsd2.String("enim"),
        Signature: "accusamus",
        TPPNokRedirectURI: tjsolerpsd2.String("delectus"),
        TPPRedirectPreferred: tjsolerpsd2.String("quidem"),
        TPPRedirectURI: tjsolerpsd2.String("provident"),
        TPPSignatureCertificate: "nam",
        XRequestID: "id",
        Aspsp: "blanditiis",
        ConsentID: "deleniti",
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

