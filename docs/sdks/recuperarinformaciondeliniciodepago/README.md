# RecuperarInformacionDelInicioDePago

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
        Digest: "soluta",
        PSUAccept: tjsolerpsd2.String("accusantium"),
        PSUAcceptCharset: tjsolerpsd2.String("aliquam"),
        PSUAcceptEncoding: tjsolerpsd2.String("sapiente"),
        PSUAcceptLanguage: tjsolerpsd2.String("dicta"),
        PSUDeviceID: tjsolerpsd2.String("ullam"),
        PSUGeoLocation: tjsolerpsd2.String("reprehenderit"),
        PSUHTTPMethod: tjsolerpsd2.String("ullam"),
        PSUIPAddress: "nisi",
        PSUIPPort: tjsolerpsd2.String("aut"),
        PSUUserAgent: tjsolerpsd2.String("voluptatum"),
        Signature: "qui",
        TPPSignatureCertificate: "quibusdam",
        XRequestID: "ex",
        Aspsp: "deleniti",
        PaymentID: "itaque",
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

