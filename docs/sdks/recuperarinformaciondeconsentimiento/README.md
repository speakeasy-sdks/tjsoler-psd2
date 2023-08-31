# RecuperarInformacionDeConsentimiento

### Available Operations

* [GetConsentIDDetails](#getconsentiddetails) - Recuperar información de consentimiento AIS
* [GetConsentsConfirmationOfFundsInfo](#getconsentsconfirmationoffundsinfo) - Recuperar información de consentimiento FCS

## GetConsentIDDetails

Este servicio permite al TPP conocer el estado de una solicitud de consentimiento iniciada previamente.

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
    res, err := s.RecuperarInformacionDeConsentimiento.GetConsentIDDetails(ctx, operations.GetConsentIDDetailsRequest{
        Digest: "vero",
        PSUAccept: psd2cajarural.String("aliquid"),
        PSUAcceptCharset: psd2cajarural.String("quasi"),
        PSUAcceptEncoding: psd2cajarural.String("saepe"),
        PSUAcceptLanguage: psd2cajarural.String("vel"),
        PSUDeviceID: psd2cajarural.String("harum"),
        PSUGeoLocation: psd2cajarural.String("molestiae"),
        PSUHTTPMethod: psd2cajarural.String("rerum"),
        PSUIPAddress: psd2cajarural.String("occaecati"),
        PSUIPPort: psd2cajarural.String("minima"),
        PSUUserAgent: psd2cajarural.String("distinctio"),
        Signature: "eligendi",
        TPPNokRedirectURI: psd2cajarural.String("sit"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("culpa"),
        TPPSignatureCertificate: "tempore",
        XRequestID: "adipisci",
        Aspsp: "cumque",
        ConsentID: "consequuntur",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetConsentIDDetailsRequest](../../models/operations/getconsentiddetailsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetConsentIDDetailsResponse](../../models/operations/getconsentiddetailsresponse.md), error**


## GetConsentsConfirmationOfFundsInfo

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
    res, err := s.RecuperarInformacionDeConsentimiento.GetConsentsConfirmationOfFundsInfo(ctx, operations.GetConsentsConfirmationOfFundsInfoRequest{
        Digest: "consequatur",
        PSUAccept: psd2cajarural.String("minus"),
        PSUAcceptCharset: psd2cajarural.String("quaerat"),
        PSUAcceptEncoding: psd2cajarural.String("sapiente"),
        PSUAcceptLanguage: psd2cajarural.String("consectetur"),
        PSUDeviceID: psd2cajarural.String("esse"),
        PSUGeoLocation: psd2cajarural.String("blanditiis"),
        PSUHTTPMethod: psd2cajarural.String("provident"),
        PSUIPAddress: psd2cajarural.String("a"),
        PSUIPPort: psd2cajarural.String("nulla"),
        PSUUserAgent: psd2cajarural.String("quas"),
        Signature: "esse",
        TPPNokRedirectURI: psd2cajarural.String("quasi"),
        TPPRedirectPreferred: psd2cajarural.String("a"),
        TPPRedirectURI: psd2cajarural.String("error"),
        TPPSignatureCertificate: "sint",
        XRequestID: "pariatur",
        Aspsp: "possimus",
        ConsentID: "quia",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetConsentsConfirmationOfFundsInfoRequest](../../models/operations/getconsentsconfirmationoffundsinforequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetConsentsConfirmationOfFundsInfoResponse](../../models/operations/getconsentsconfirmationoffundsinforesponse.md), error**

