# RecuperarInformacionDelInicioDePago
(*RecuperarInformacionDelInicioDePago*)

### Available Operations

* [GetInfoPayment](#getinfopayment) - Recuperar información del Inicio de pago

## GetInfoPayment

Este mensaje es enviado por el TPP al HUB para solicitar información de un inicio de pago

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
    res, err := s.RecuperarInformacionDelInicioDePago.GetInfoPayment(ctx, operations.GetInfoPaymentRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.GetInfoPaymentPathParamPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.GetInfoPaymentPathParamPaymentServiceBulkPayments,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetInfoPaymentRequest](../../pkg/models/operations/getinfopaymentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetInfoPaymentResponse](../../pkg/models/operations/getinfopaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
