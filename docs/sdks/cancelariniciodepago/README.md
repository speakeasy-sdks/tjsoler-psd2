# CancelarInicioDePago
(*CancelarInicioDePago*)

### Available Operations

* [DeletePayment](#deletepayment) - Cancelar Inicio de pago

## DeletePayment

Esta petición es enviada por el TPP al ASPSP a través del Hub y permite iniciar la cancelación de un pago. Dependiendo del servicio de pago, el producto de pago y la implementación del ASPSP, esta petición podríar ser suficiente para cancelar el pago o podría ser necesario una autorización.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.CancelarInicioDePago.DeletePayment(ctx, operations.DeletePaymentRequest{
        Digest: "string",
        Signature: "string",
        TPPSignatureCertificate: "string",
        XRequestID: "string",
        Aspsp: "string",
        PaymentID: "string",
        PaymentProduct: operations.DeletePaymentPathParamPaymentProductSepaCreditTransfers,
        PaymentService: operations.DeletePaymentPathParamPaymentServiceBulkPayments,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeletePaymentRequest](../../pkg/models/operations/deletepaymentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeletePaymentResponse](../../pkg/models/operations/deletepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
