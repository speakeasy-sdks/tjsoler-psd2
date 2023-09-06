# ObtenerEstadoDeConsentimiento

### Available Operations

* [GetConsentStatus](#getconsentstatus) - Estado de consentimiento AIS
* [GetConsentsConfirmationOfFunds](#getconsentsconfirmationoffunds) - Estado de consentimiento FCS

## GetConsentStatus

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
    res, err := s.ObtenerEstadoDeConsentimiento.GetConsentStatus(ctx, operations.GetConsentStatusRequest{
        Digest: "porro",
        PSUAccept: psd2cajarural.String("maiores"),
        PSUAcceptCharset: psd2cajarural.String("doloribus"),
        PSUAcceptEncoding: psd2cajarural.String("iusto"),
        PSUAcceptLanguage: psd2cajarural.String("eligendi"),
        PSUDeviceID: psd2cajarural.String("ducimus"),
        PSUGeoLocation: psd2cajarural.String("alias"),
        PSUHTTPMethod: psd2cajarural.String("officia"),
        PSUIPAddress: psd2cajarural.String("tempora"),
        PSUIPPort: psd2cajarural.String("ipsam"),
        PSUUserAgent: psd2cajarural.String("ea"),
        Signature: "aspernatur",
        TPPNokRedirectURI: psd2cajarural.String("vel"),
        TPPRedirectPreferred: psd2cajarural.String("possimus"),
        TPPRedirectURI: psd2cajarural.String("magnam"),
        TPPSignatureCertificate: "ratione",
        XRequestID: "ex",
        Aspsp: "laudantium",
        ConsentID: "dicta",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetConsentStatusRequest](../../models/operations/getconsentstatusrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetConsentStatusResponse](../../models/operations/getconsentstatusresponse.md), error**


## GetConsentsConfirmationOfFunds

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
    res, err := s.ObtenerEstadoDeConsentimiento.GetConsentsConfirmationOfFunds(ctx, operations.GetConsentsConfirmationOfFundsRequest{
        Digest: "dolor",
        PSUAccept: psd2cajarural.String("maiores"),
        PSUAcceptCharset: psd2cajarural.String("quasi"),
        PSUAcceptEncoding: psd2cajarural.String("ex"),
        PSUAcceptLanguage: psd2cajarural.String("nulla"),
        PSUDeviceID: psd2cajarural.String("excepturi"),
        PSUGeoLocation: psd2cajarural.String("voluptatibus"),
        PSUHTTPMethod: psd2cajarural.String("nostrum"),
        PSUIPAddress: psd2cajarural.String("sapiente"),
        PSUIPPort: psd2cajarural.String("quisquam"),
        PSUUserAgent: psd2cajarural.String("saepe"),
        Signature: "ea",
        TPPNokRedirectURI: psd2cajarural.String("impedit"),
        TPPRedirectPreferred: psd2cajarural.String("corporis"),
        TPPRedirectURI: psd2cajarural.String("veniam"),
        TPPSignatureCertificate: "aliquid",
        XRequestID: "inventore",
        Aspsp: "magnam",
        ConsentID: "ea",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetConsentsConfirmationOfFundsRequest](../../models/operations/getconsentsconfirmationoffundsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetConsentsConfirmationOfFundsResponse](../../models/operations/getconsentsconfirmationoffundsresponse.md), error**

