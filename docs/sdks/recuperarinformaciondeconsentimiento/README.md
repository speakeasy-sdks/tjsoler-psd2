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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.RecuperarInformacionDeConsentimiento.GetConsentIDDetails(ctx, operations.GetConsentIDDetailsRequest{
        Digest: "consequuntur",
        PSUAccept: tjsolerpsd2.String("consequatur"),
        PSUAcceptCharset: tjsolerpsd2.String("minus"),
        PSUAcceptEncoding: tjsolerpsd2.String("quaerat"),
        PSUAcceptLanguage: tjsolerpsd2.String("sapiente"),
        PSUDeviceID: tjsolerpsd2.String("consectetur"),
        PSUGeoLocation: tjsolerpsd2.String("esse"),
        PSUHTTPMethod: tjsolerpsd2.String("blanditiis"),
        PSUIPAddress: tjsolerpsd2.String("provident"),
        PSUIPPort: tjsolerpsd2.String("a"),
        PSUUserAgent: tjsolerpsd2.String("nulla"),
        Signature: "quas",
        TPPNokRedirectURI: tjsolerpsd2.String("esse"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("quasi"),
        TPPSignatureCertificate: "a",
        XRequestID: "error",
        Aspsp: "sint",
        ConsentID: "pariatur",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.RecuperarInformacionDeConsentimiento.GetConsentsConfirmationOfFundsInfo(ctx, operations.GetConsentsConfirmationOfFundsInfoRequest{
        Digest: "possimus",
        PSUAccept: tjsolerpsd2.String("quia"),
        PSUAcceptCharset: tjsolerpsd2.String("eveniet"),
        PSUAcceptEncoding: tjsolerpsd2.String("asperiores"),
        PSUAcceptLanguage: tjsolerpsd2.String("facere"),
        PSUDeviceID: tjsolerpsd2.String("veritatis"),
        PSUGeoLocation: tjsolerpsd2.String("consequuntur"),
        PSUHTTPMethod: tjsolerpsd2.String("quasi"),
        PSUIPAddress: tjsolerpsd2.String("similique"),
        PSUIPPort: tjsolerpsd2.String("culpa"),
        PSUUserAgent: tjsolerpsd2.String("aliquid"),
        Signature: "tenetur",
        TPPNokRedirectURI: tjsolerpsd2.String("quae"),
        TPPRedirectPreferred: tjsolerpsd2.String("earum"),
        TPPRedirectURI: tjsolerpsd2.String("vel"),
        TPPSignatureCertificate: "in",
        XRequestID: "eius",
        Aspsp: "libero",
        ConsentID: "illum",
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

