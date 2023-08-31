# ObtenerEstadoDelPago

### Available Operations

* [StatusPayment](#statuspayment) - Obtener información del Estado de pago

## StatusPayment

Este mensaje es enviado por el TPP al HUB para solicitar información del estado en el que se encuentra la iniciación de pago que solicitó el TPP

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
    res, err := s.ObtenerEstadoDelPago.StatusPayment(ctx, operations.StatusPaymentRequest{
        Digest: "eaque",
        PSUAccept: psd2cajarural.String("a"),
        PSUAcceptCharset: psd2cajarural.String("libero"),
        PSUAcceptEncoding: psd2cajarural.String("aut"),
        PSUAcceptLanguage: psd2cajarural.String("aut"),
        PSUDeviceID: psd2cajarural.String("deleniti"),
        PSUGeoLocation: psd2cajarural.String("impedit"),
        PSUHTTPMethod: psd2cajarural.String("aliquam"),
        PSUIPAddress: "fugit",
        PSUIPPort: psd2cajarural.String("accusamus"),
        PSUUserAgent: psd2cajarural.String("inventore"),
        Signature: "non",
        TPPSignatureCertificate: "et",
        XRequestID: "dolorum",
        Aspsp: "laborum",
        PaymentID: "placeat",
        PaymentProduct: operations.StatusPaymentPaymentProductSepaCreditTransfers,
        PaymentService: operations.StatusPaymentPaymentServiceBulkPayments,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.StatusPaymentRequest](../../models/operations/statuspaymentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.StatusPaymentResponse](../../models/operations/statuspaymentresponse.md), error**

