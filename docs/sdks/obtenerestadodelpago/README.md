# ObtenerEstadoDelPago
(*ObtenerEstadoDelPago*)

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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoDelPago.StatusPayment(ctx, operations.StatusPaymentRequest{
        Digest: "Electric oat Cruiser",
        PSUAccept: tjsolerpsd2.String("capacitor"),
        PSUAcceptCharset: tjsolerpsd2.String("Ball Intelligent"),
        PSUAcceptEncoding: tjsolerpsd2.String("iterate"),
        PSUAcceptLanguage: tjsolerpsd2.String("Car Movies carter"),
        PSUDeviceID: tjsolerpsd2.String("Falkland bifurcated"),
        PSUGeoLocation: tjsolerpsd2.String("manufacturer Steel"),
        PSUHTTPMethod: tjsolerpsd2.String("Virtual Optional"),
        PSUIPAddress: "Tandem Music distributed",
        PSUIPPort: tjsolerpsd2.String("Shoes lumen"),
        PSUUserAgent: tjsolerpsd2.String("male F2M lumen"),
        Signature: "AGP",
        TPPSignatureCertificate: "female Computer Rock",
        XRequestID: "monetize",
        Aspsp: "Bacon",
        PaymentID: "West ugh",
        PaymentProduct: operations.StatusPaymentPaymentProductSepaCreditTransfers,
        PaymentService: operations.StatusPaymentPaymentServicePeriodicPayments,
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

