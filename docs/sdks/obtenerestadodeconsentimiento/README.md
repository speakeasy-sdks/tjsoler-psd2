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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoDeConsentimiento.GetConsentStatus(ctx, operations.GetConsentStatusRequest{
        Digest: "ex",
        PSUAccept: tjsolerpsd2.String("nulla"),
        PSUAcceptCharset: tjsolerpsd2.String("excepturi"),
        PSUAcceptEncoding: tjsolerpsd2.String("voluptatibus"),
        PSUAcceptLanguage: tjsolerpsd2.String("nostrum"),
        PSUDeviceID: tjsolerpsd2.String("sapiente"),
        PSUGeoLocation: tjsolerpsd2.String("quisquam"),
        PSUHTTPMethod: tjsolerpsd2.String("saepe"),
        PSUIPAddress: tjsolerpsd2.String("ea"),
        PSUIPPort: tjsolerpsd2.String("impedit"),
        PSUUserAgent: tjsolerpsd2.String("corporis"),
        Signature: "veniam",
        TPPNokRedirectURI: tjsolerpsd2.String("aliquid"),
        TPPRedirectPreferred: tjsolerpsd2.String("inventore"),
        TPPRedirectURI: tjsolerpsd2.String("magnam"),
        TPPSignatureCertificate: "ea",
        XRequestID: "quo",
        Aspsp: "consectetur",
        ConsentID: "recusandae",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoDeConsentimiento.GetConsentsConfirmationOfFunds(ctx, operations.GetConsentsConfirmationOfFundsRequest{
        Digest: "aspernatur",
        PSUAccept: tjsolerpsd2.String("minima"),
        PSUAcceptCharset: tjsolerpsd2.String("eaque"),
        PSUAcceptEncoding: tjsolerpsd2.String("a"),
        PSUAcceptLanguage: tjsolerpsd2.String("libero"),
        PSUDeviceID: tjsolerpsd2.String("aut"),
        PSUGeoLocation: tjsolerpsd2.String("aut"),
        PSUHTTPMethod: tjsolerpsd2.String("deleniti"),
        PSUIPAddress: tjsolerpsd2.String("impedit"),
        PSUIPPort: tjsolerpsd2.String("aliquam"),
        PSUUserAgent: tjsolerpsd2.String("fugit"),
        Signature: "accusamus",
        TPPNokRedirectURI: tjsolerpsd2.String("inventore"),
        TPPRedirectPreferred: tjsolerpsd2.String("non"),
        TPPRedirectURI: tjsolerpsd2.String("et"),
        TPPSignatureCertificate: "dolorum",
        XRequestID: "laborum",
        Aspsp: "placeat",
        ConsentID: "velit",
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

