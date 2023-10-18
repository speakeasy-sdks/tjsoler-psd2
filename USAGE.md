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
		Digest:                  "Tugrik",
		PSUIPAddress:            "South",
		Signature:               "Northeast",
		TPPSignatureCertificate: "Recycled",
		XRequestID:              "Northwest",
		Aspsp:                   operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBbvafr,
		AuthorisationID:         "Connelly",
		PaymentID:               "B2B",
		PaymentProduct:          operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticCrossCurrencyPaymentsUk,
		PaymentService:          operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServiceBulkPayments,
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