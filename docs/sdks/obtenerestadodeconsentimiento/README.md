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
        Digest: "ducimus",
        PSUAccept: psd2cajarural.String("alias"),
        PSUAcceptCharset: psd2cajarural.String("officia"),
        PSUAcceptEncoding: psd2cajarural.String("tempora"),
        PSUAcceptLanguage: psd2cajarural.String("ipsam"),
        PSUDeviceID: psd2cajarural.String("ea"),
        PSUGeoLocation: psd2cajarural.String("aspernatur"),
        PSUHTTPMethod: psd2cajarural.String("vel"),
        PSUIPAddress: psd2cajarural.String("possimus"),
        PSUIPPort: psd2cajarural.String("magnam"),
        PSUUserAgent: psd2cajarural.String("ratione"),
        Signature: "ex",
        TPPNokRedirectURI: psd2cajarural.String("laudantium"),
        TPPRedirectPreferred: psd2cajarural.String("dicta"),
        TPPRedirectURI: psd2cajarural.String("dolor"),
        TPPSignatureCertificate: "maiores",
        XRequestID: "quasi",
        Aspsp: "ex",
        ConsentID: "nulla",
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
        Digest: "excepturi",
        PSUAccept: psd2cajarural.String("voluptatibus"),
        PSUAcceptCharset: psd2cajarural.String("nostrum"),
        PSUAcceptEncoding: psd2cajarural.String("sapiente"),
        PSUAcceptLanguage: psd2cajarural.String("quisquam"),
        PSUDeviceID: psd2cajarural.String("saepe"),
        PSUGeoLocation: psd2cajarural.String("ea"),
        PSUHTTPMethod: psd2cajarural.String("impedit"),
        PSUIPAddress: psd2cajarural.String("corporis"),
        PSUIPPort: psd2cajarural.String("veniam"),
        PSUUserAgent: psd2cajarural.String("aliquid"),
        Signature: "inventore",
        TPPNokRedirectURI: psd2cajarural.String("magnam"),
        TPPRedirectPreferred: psd2cajarural.String("ea"),
        TPPRedirectURI: psd2cajarural.String("quo"),
        TPPSignatureCertificate: "consectetur",
        XRequestID: "recusandae",
        Aspsp: "aspernatur",
        ConsentID: "minima",
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

