<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
	s := tjsolerpsd2.New()

	ctx := context.Background()
	res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
		Digest:                  "string",
		PSUIPAddress:            "string",
		Signature:               "string",
		TPPSignatureCertificate: "string",
		XRequestID:              "string",
		Aspsp:                   operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBff,
		AuthorisationID:         "string",
		PaymentID:               "string",
		PaymentProduct:          operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticCrossCurrencyPaymentsUk,
		PaymentService:          operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePayments,
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