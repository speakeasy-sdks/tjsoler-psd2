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
        ConsentID: tjsolerpsd2.String("driver Metal"),
        Digest: "Recycled Northwest",
        PSUAccept: tjsolerpsd2.String("Connelly"),
        PSUAcceptCharset: tjsolerpsd2.String("Sleek"),
        PSUAcceptEncoding: tjsolerpsd2.String("grow sunt"),
        PSUAcceptLanguage: tjsolerpsd2.String("connecting Convertible"),
        PSUDeviceID: tjsolerpsd2.String("Directives outside blockchains"),
        PSUGeoLocation: tjsolerpsd2.String("jack transmit"),
        PSUHTTPMethod: tjsolerpsd2.String("SMS"),
        PSUIPAddress: "female",
        PSUIPPort: tjsolerpsd2.String("Gasoline Books Facilitator"),
        PSUUserAgent: tjsolerpsd2.String("Road"),
        Signature: "online",
        TPPNokRedirectURI: tjsolerpsd2.String("mole"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Views application"),
        TPPSignatureCertificate: "Rubber amongst",
        XRequestID: "Beauty siemens Litas",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspWizink,
        AuthorisationID: "untried deposit",
        PaymentID: "Coupe orchid",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePayments,
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