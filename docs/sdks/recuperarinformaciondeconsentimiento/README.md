# RecuperarInformacionDeConsentimiento
(*RecuperarInformacionDeConsentimiento*)

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
        Digest: "Ocala nasalise Calcium",
        PSUAccept: tjsolerpsd2.String("reprehenderit for"),
        PSUAcceptCharset: tjsolerpsd2.String("huzzah"),
        PSUAcceptEncoding: tjsolerpsd2.String("Loan"),
        PSUAcceptLanguage: tjsolerpsd2.String("online Marketing"),
        PSUDeviceID: tjsolerpsd2.String("barring Assistant utilize"),
        PSUGeoLocation: tjsolerpsd2.String("woot"),
        PSUHTTPMethod: tjsolerpsd2.String("behind"),
        PSUIPAddress: tjsolerpsd2.String("programming Dodge"),
        PSUIPPort: tjsolerpsd2.String("Trans with Classical"),
        PSUUserAgent: tjsolerpsd2.String("Southwest Northeast"),
        Signature: "dicta",
        TPPNokRedirectURI: tjsolerpsd2.String("Bacon logistical Infrastructure"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Kia Small"),
        TPPSignatureCertificate: "false",
        XRequestID: "infomediaries minus focus",
        Aspsp: "impactful",
        ConsentID: "Soap Wagon",
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
        Digest: "Representative",
        PSUAccept: tjsolerpsd2.String("payment"),
        PSUAcceptCharset: tjsolerpsd2.String("functionalities Forward"),
        PSUAcceptEncoding: tjsolerpsd2.String("Hybrid kelvin"),
        PSUAcceptLanguage: tjsolerpsd2.String("Northeast partially"),
        PSUDeviceID: tjsolerpsd2.String("sparkling copy Planner"),
        PSUGeoLocation: tjsolerpsd2.String("why Tools Seaborgium"),
        PSUHTTPMethod: tjsolerpsd2.String("cyan Uranium"),
        PSUIPAddress: tjsolerpsd2.String("models"),
        PSUIPPort: tjsolerpsd2.String("Convertible Central"),
        PSUUserAgent: tjsolerpsd2.String("Diesel"),
        Signature: "Sulfur Ames",
        TPPNokRedirectURI: tjsolerpsd2.String("Planner Pants"),
        TPPRedirectPreferred: tjsolerpsd2.String("extend"),
        TPPRedirectURI: tjsolerpsd2.String("wireless whoa visualize"),
        TPPSignatureCertificate: "Planner Saint",
        XRequestID: "Card Classical",
        Aspsp: "calculate",
        ConsentID: "North sed",
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

