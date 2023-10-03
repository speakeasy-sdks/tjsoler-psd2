# ObtenerEstadoDeConsentimiento
(*ObtenerEstadoDeConsentimiento*)

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
        Digest: "Universal policy",
        PSUAccept: tjsolerpsd2.String("swill applications Bedfordshire"),
        PSUAcceptCharset: tjsolerpsd2.String("appetizer Paradigm erosion"),
        PSUAcceptEncoding: tjsolerpsd2.String("usually pish Islands"),
        PSUAcceptLanguage: tjsolerpsd2.String("North"),
        PSUDeviceID: tjsolerpsd2.String("spicy"),
        PSUGeoLocation: tjsolerpsd2.String("Creative North Electric"),
        PSUHTTPMethod: tjsolerpsd2.String("input deploy Apopka"),
        PSUIPAddress: tjsolerpsd2.String("Metrics"),
        PSUIPPort: tjsolerpsd2.String("realise Bicycle"),
        PSUUserAgent: tjsolerpsd2.String("Cargo"),
        Signature: "architect core Southeast",
        TPPNokRedirectURI: tjsolerpsd2.String("yippee Northeast Robust"),
        TPPRedirectPreferred: tjsolerpsd2.String("Diesel Computer Bicycle"),
        TPPRedirectURI: tjsolerpsd2.String("pheasant Bedfordshire"),
        TPPSignatureCertificate: "York",
        XRequestID: "syndicate Northwest Gloves",
        Aspsp: "tube Tennessee Ferrari",
        ConsentID: "engage Account auxiliary",
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
        Digest: "Royce Cambridgeshire",
        PSUAccept: tjsolerpsd2.String("Arkansas"),
        PSUAcceptCharset: tjsolerpsd2.String("Northeast orange withdrawal"),
        PSUAcceptEncoding: tjsolerpsd2.String("partnerships East candela"),
        PSUAcceptLanguage: tjsolerpsd2.String("Investment Oval"),
        PSUDeviceID: tjsolerpsd2.String("copying male"),
        PSUGeoLocation: tjsolerpsd2.String("Practical within"),
        PSUHTTPMethod: tjsolerpsd2.String("Yuan technologies Milton"),
        PSUIPAddress: tjsolerpsd2.String("invoice Ball"),
        PSUIPPort: tjsolerpsd2.String("Creative copying"),
        PSUUserAgent: tjsolerpsd2.String("Gloves Oriental regional"),
        Signature: "Buckinghamshire",
        TPPNokRedirectURI: tjsolerpsd2.String("Northwest Northeast"),
        TPPRedirectPreferred: tjsolerpsd2.String("Southeast Pop"),
        TPPRedirectURI: tjsolerpsd2.String("ack"),
        TPPSignatureCertificate: "ideal Southwest",
        XRequestID: "transmitter Shoes Kazakhstan",
        Aspsp: "sexy",
        ConsentID: "Optional Clothing Promethium",
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

