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
        Digest: "accusamus",
        PSUAccept: psd2cajarural.String("veritatis"),
        PSUAcceptCharset: psd2cajarural.String("esse"),
        PSUAcceptEncoding: psd2cajarural.String("quod"),
        PSUAcceptLanguage: psd2cajarural.String("nam"),
        PSUDeviceID: psd2cajarural.String("vero"),
        PSUGeoLocation: psd2cajarural.String("aliquid"),
        PSUHTTPMethod: psd2cajarural.String("quasi"),
        PSUIPAddress: psd2cajarural.String("saepe"),
        PSUIPPort: psd2cajarural.String("vel"),
        PSUUserAgent: psd2cajarural.String("harum"),
        Signature: "molestiae",
        TPPNokRedirectURI: psd2cajarural.String("rerum"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("occaecati"),
        TPPSignatureCertificate: "minima",
        XRequestID: "distinctio",
        Aspsp: "eligendi",
        ConsentID: "sit",
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
        Digest: "culpa",
        PSUAccept: psd2cajarural.String("tempore"),
        PSUAcceptCharset: psd2cajarural.String("adipisci"),
        PSUAcceptEncoding: psd2cajarural.String("cumque"),
        PSUAcceptLanguage: psd2cajarural.String("consequuntur"),
        PSUDeviceID: psd2cajarural.String("consequatur"),
        PSUGeoLocation: psd2cajarural.String("minus"),
        PSUHTTPMethod: psd2cajarural.String("quaerat"),
        PSUIPAddress: psd2cajarural.String("sapiente"),
        PSUIPPort: psd2cajarural.String("consectetur"),
        PSUUserAgent: psd2cajarural.String("esse"),
        Signature: "blanditiis",
        TPPNokRedirectURI: psd2cajarural.String("provident"),
        TPPRedirectPreferred: psd2cajarural.String("a"),
        TPPRedirectURI: psd2cajarural.String("nulla"),
        TPPSignatureCertificate: "quas",
        XRequestID: "esse",
        Aspsp: "quasi",
        ConsentID: "a",
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

