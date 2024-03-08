<!-- Start SDK Example Usage [usage] -->
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
	res, err := s.ServiciosMultibanco.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
		Digest:                  "<value>",
		PSUIPAddress:            "<value>",
		Signature:               "<value>",
		TPPSignatureCertificate: "<value>",
		XRequestID:              "<value>",
		AspName:                 "<value>",
		MultibancoPaymentType:   "<value>",
		PaymentID:               "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseDeleteMultibankPayment != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->