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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.CancelarInicioDePago.DeletePayment(ctx, operations.DeletePaymentRequest{
        Digest: "Australian",
        PSUAccept: tjsolerpsd2.String("Bronze but"),
        PSUAcceptCharset: tjsolerpsd2.String("Brand Sports Brand"),
        PSUAcceptEncoding: tjsolerpsd2.String("disinherit local Rutherfordium"),
        PSUAcceptLanguage: tjsolerpsd2.String("repurpose Chevrolet Unbranded"),
        PSUDeviceID: tjsolerpsd2.String("Hybrid Electric Account"),
        PSUGeoLocation: tjsolerpsd2.String("South"),
        PSUHTTPMethod: tjsolerpsd2.String("olive Won National"),
        PSUIPAddress: tjsolerpsd2.String("bluetooth Bronze Technician"),
        PSUIPPort: tjsolerpsd2.String("immunise Technician"),
        PSUUserAgent: tjsolerpsd2.String("Utah payment"),
        Signature: "kelvin",
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("Intelligent"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("hacking Corporate groupware"),
        TPPSignatureCertificate: "dicta",
        XRequestID: "Cambridgeshire worth Neon",
        Aspsp: "Wagon",
        PaymentID: "invoice mmm",
        PaymentProduct: operations.DeletePaymentPaymentProductInstantSepaCreditTransfers,
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

