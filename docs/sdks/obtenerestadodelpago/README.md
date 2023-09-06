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
        Digest: "quo",
        PSUAccept: psd2cajarural.String("consectetur"),
        PSUAcceptCharset: psd2cajarural.String("recusandae"),
        PSUAcceptEncoding: psd2cajarural.String("aspernatur"),
        PSUAcceptLanguage: psd2cajarural.String("minima"),
        PSUDeviceID: psd2cajarural.String("eaque"),
        PSUGeoLocation: psd2cajarural.String("a"),
        PSUHTTPMethod: psd2cajarural.String("libero"),
        PSUIPAddress: "aut",
        PSUIPPort: psd2cajarural.String("aut"),
        PSUUserAgent: psd2cajarural.String("deleniti"),
        Signature: "impedit",
        TPPSignatureCertificate: "aliquam",
        XRequestID: "fugit",
        Aspsp: "accusamus",
        PaymentID: "inventore",
        PaymentProduct: operations.StatusPaymentPaymentProductInstantSepaCreditTransfers,
        PaymentService: operations.StatusPaymentPaymentServicePayments,
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

