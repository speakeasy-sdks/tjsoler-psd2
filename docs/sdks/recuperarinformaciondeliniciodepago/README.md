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
        Digest: "error",
        PSUAccept: psd2cajarural.String("sint"),
        PSUAcceptCharset: psd2cajarural.String("pariatur"),
        PSUAcceptEncoding: psd2cajarural.String("possimus"),
        PSUAcceptLanguage: psd2cajarural.String("quia"),
        PSUDeviceID: psd2cajarural.String("eveniet"),
        PSUGeoLocation: psd2cajarural.String("asperiores"),
        PSUHTTPMethod: psd2cajarural.String("facere"),
        PSUIPAddress: "veritatis",
        PSUIPPort: psd2cajarural.String("consequuntur"),
        PSUUserAgent: psd2cajarural.String("quasi"),
        Signature: "similique",
        TPPSignatureCertificate: "culpa",
        XRequestID: "aliquid",
        Aspsp: "tenetur",
        PaymentID: "quae",
        PaymentProduct: operations.GetInfoPaymentPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.GetInfoPaymentPaymentServiceBulkPayments,
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

