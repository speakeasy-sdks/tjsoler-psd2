# CancelarInicioDePago

### Available Operations

* [DeletePayment](#deletepayment) - Cancelar Inicio de pago

## DeletePayment

Esta petición es enviada por el TPP al ASPSP a través del Hub y permite iniciar la cancelación de un pago. Dependiendo del servicio de pago, el producto de pago y la implementación del ASPSP, esta petición podríar ser suficiente para cancelar el pago o podría ser necesario una autorización.

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
    res, err := s.CancelarInicioDePago.DeletePayment(ctx, operations.DeletePaymentRequest{
        Digest: "repellat",
        PSUAccept: psd2cajarural.String("mollitia"),
        PSUAcceptCharset: psd2cajarural.String("occaecati"),
        PSUAcceptEncoding: psd2cajarural.String("numquam"),
        PSUAcceptLanguage: psd2cajarural.String("commodi"),
        PSUDeviceID: psd2cajarural.String("quam"),
        PSUGeoLocation: psd2cajarural.String("molestiae"),
        PSUHTTPMethod: psd2cajarural.String("velit"),
        PSUIPAddress: psd2cajarural.String("error"),
        PSUIPPort: psd2cajarural.String("quia"),
        PSUUserAgent: psd2cajarural.String("quis"),
        Signature: "vitae",
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("laborum"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("animi"),
        TPPSignatureCertificate: "enim",
        XRequestID: "odit",
        Aspsp: "quo",
        PaymentID: "sequi",
        PaymentProduct: operations.DeletePaymentPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.DeletePaymentPaymentServiceBulkPayments,
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
| `request`                                                                          | [operations.DeletePaymentRequest](../../models/operations/deletepaymentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeletePaymentResponse](../../models/operations/deletepaymentresponse.md), error**

