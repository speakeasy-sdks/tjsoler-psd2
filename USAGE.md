<!-- Start SDK Example Usage -->


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
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        ConsentID: psd2cajarural.String("corrupti"),
        Digest: "provident",
        PSUAccept: psd2cajarural.String("distinctio"),
        PSUAcceptCharset: psd2cajarural.String("quibusdam"),
        PSUAcceptEncoding: psd2cajarural.String("unde"),
        PSUAcceptLanguage: psd2cajarural.String("nulla"),
        PSUDeviceID: psd2cajarural.String("corrupti"),
        PSUGeoLocation: psd2cajarural.String("illum"),
        PSUHTTPMethod: psd2cajarural.String("vel"),
        PSUIPAddress: "error",
        PSUIPPort: psd2cajarural.String("deserunt"),
        PSUUserAgent: psd2cajarural.String("suscipit"),
        Signature: "iure",
        TPPNokRedirectURI: psd2cajarural.String("magnam"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("debitis"),
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