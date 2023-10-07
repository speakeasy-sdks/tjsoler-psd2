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
        Digest: "driver Metal",
        PSUIPAddress: "Recycled Northwest",
        Signature: "Connelly",
        TPPSignatureCertificate: "Sleek",
        XRequestID: "grow sunt",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBff,
        AuthorisationID: "bus Dakota",
        PaymentID: "outside",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductTarget2Payments,
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