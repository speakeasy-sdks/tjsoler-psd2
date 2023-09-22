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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.CancelarInicioDePago.DeletePayment(ctx, operations.DeletePaymentRequest{
        Digest: "aut",
        PSUAccept: tjsolerpsd2.String("quasi"),
        PSUAcceptCharset: tjsolerpsd2.String("error"),
        PSUAcceptEncoding: tjsolerpsd2.String("temporibus"),
        PSUAcceptLanguage: tjsolerpsd2.String("laborum"),
        PSUDeviceID: tjsolerpsd2.String("quasi"),
        PSUGeoLocation: tjsolerpsd2.String("reiciendis"),
        PSUHTTPMethod: tjsolerpsd2.String("voluptatibus"),
        PSUIPAddress: tjsolerpsd2.String("vero"),
        PSUIPPort: tjsolerpsd2.String("nihil"),
        PSUUserAgent: tjsolerpsd2.String("praesentium"),
        Signature: "voluptatibus",
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("ipsa"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("omnis"),
        TPPSignatureCertificate: "voluptate",
        XRequestID: "cum",
        Aspsp: "perferendis",
        PaymentID: "doloremque",
        PaymentProduct: operations.DeletePaymentPaymentProductInstantSepaCreditTransfers,
        PaymentService: operations.DeletePaymentPaymentServicePayments,
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

