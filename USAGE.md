<!-- Start SDK Example Usage -->


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
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("corrupti"),
        Digest: "provident",
        PSUAccept: tjsolerpsd2.String("distinctio"),
        PSUAcceptCharset: tjsolerpsd2.String("quibusdam"),
        PSUAcceptEncoding: tjsolerpsd2.String("unde"),
        PSUAcceptLanguage: tjsolerpsd2.String("nulla"),
        PSUDeviceID: tjsolerpsd2.String("corrupti"),
        PSUGeoLocation: tjsolerpsd2.String("illum"),
        PSUHTTPMethod: tjsolerpsd2.String("vel"),
        PSUIPAddress: "error",
        PSUIPPort: tjsolerpsd2.String("deserunt"),
        PSUUserAgent: tjsolerpsd2.String("suscipit"),
        Signature: "iure",
        TPPNokRedirectURI: tjsolerpsd2.String("magnam"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("debitis"),
        TPPSignatureCertificate: "ipsa",
        XRequestID: "delectus",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBancosantander,
        AuthorisationID: "suscipit",
        PaymentID: "molestiae",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePeriodicPayments,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->