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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoDelPago.StatusPayment(ctx, operations.StatusPaymentRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.StatusPaymentPathParamPaymentProductTarget2Payments,
        PaymentService: operations.StatusPaymentPathParamPaymentServicePeriodicPayments,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.StatusPaymentRequest](../../pkg/models/operations/statuspaymentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.StatusPaymentResponse](../../pkg/models/operations/statuspaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
