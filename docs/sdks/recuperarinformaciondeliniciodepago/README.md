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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.RecuperarInformacionDelInicioDePago.GetInfoPayment(ctx, operations.GetInfoPaymentRequest{
        Digest: "moderator Non Lead",
        PSUAccept: tjsolerpsd2.String("payment Neon wireless"),
        PSUAcceptCharset: tjsolerpsd2.String("intuitive Zambia orange"),
        PSUAcceptEncoding: tjsolerpsd2.String("Bedfordshire Fresh"),
        PSUAcceptLanguage: tjsolerpsd2.String("Martin"),
        PSUDeviceID: tjsolerpsd2.String("Incredible Shoes"),
        PSUGeoLocation: tjsolerpsd2.String("incentivize"),
        PSUHTTPMethod: tjsolerpsd2.String("local"),
        PSUIPAddress: "complexity Frozen Grocery",
        PSUIPPort: tjsolerpsd2.String("Buckinghamshire"),
        PSUUserAgent: tjsolerpsd2.String("ADP Home"),
        Signature: "invoice punctual",
        TPPSignatureCertificate: "International",
        XRequestID: "Hybrid Minnesota",
        Aspsp: "Egypt leading",
        PaymentID: "compelling Electronics Implemented",
        PaymentProduct: operations.GetInfoPaymentPaymentProductTarget2Payments,
        PaymentService: operations.GetInfoPaymentPaymentServicePayments,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetInfoPaymentRequest](../../models/operations/getinfopaymentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetInfoPaymentResponse](../../models/operations/getinfopaymentresponse.md), error**

