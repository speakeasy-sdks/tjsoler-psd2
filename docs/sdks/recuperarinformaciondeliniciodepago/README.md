# RecuperarInformacionDelInicioDePago

### Available Operations

* [GetInfoPayment](#getinfopayment) - Recuperar información del Inicio de pago

## GetInfoPayment

Este mensaje es enviado por el TPP al HUB para solicitar información de un inicio de pago

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
    res, err := s.RecuperarInformacionDelInicioDePago.GetInfoPayment(ctx, operations.GetInfoPaymentRequest{
        Digest: "eveniet",
        PSUAccept: psd2cajarural.String("asperiores"),
        PSUAcceptCharset: psd2cajarural.String("facere"),
        PSUAcceptEncoding: psd2cajarural.String("veritatis"),
        PSUAcceptLanguage: psd2cajarural.String("consequuntur"),
        PSUDeviceID: psd2cajarural.String("quasi"),
        PSUGeoLocation: psd2cajarural.String("similique"),
        PSUHTTPMethod: psd2cajarural.String("culpa"),
        PSUIPAddress: "aliquid",
        PSUIPPort: psd2cajarural.String("tenetur"),
        PSUUserAgent: psd2cajarural.String("quae"),
        Signature: "earum",
        TPPSignatureCertificate: "vel",
        XRequestID: "in",
        Aspsp: "eius",
        PaymentID: "libero",
        PaymentProduct: operations.GetInfoPaymentPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.GetInfoPaymentPaymentServicePeriodicPayments,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetInfoPaymentRequest](../../models/operations/getinfopaymentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetInfoPaymentResponse](../../models/operations/getinfopaymentresponse.md), error**

