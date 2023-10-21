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
        Digest: "string",
        PSUIPAddress: "string",
        Signature: "string",
        TPPSignatureCertificate: "string",
        XRequestID: "string",
        Aspsp: "string",
        PaymentID: "string",
        PaymentProduct: operations.StatusPaymentPaymentProductTarget2Payments,
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

