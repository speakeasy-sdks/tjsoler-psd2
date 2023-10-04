# ObtenerListadoDeBeneficiariosDeConfianza
(*ObtenerListadoDeBeneficiariosDeConfianza*)

### Available Operations

* [GetTrustedBeneficiaries](#gettrustedbeneficiaries) - Obtener listado de beneficiarios de confianza

## GetTrustedBeneficiaries

Obtiene el listado de los beneficiarios de confianza del PSU, el cual ha dado un consentimiento explícito.

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
    res, err := s.ObtenerListadoDeBeneficiariosDeConfianza.GetTrustedBeneficiaries(ctx, operations.GetTrustedBeneficiariesRequest{
        ConsentID: "Virginia Regional",
        Digest: "Cambridgeshire",
        PSUAccept: tjsolerpsd2.String("anenst"),
        PSUAcceptCharset: tjsolerpsd2.String("ability"),
        PSUAcceptEncoding: tjsolerpsd2.String("alarm"),
        PSUAcceptLanguage: tjsolerpsd2.String("Blues"),
        PSUDeviceID: tjsolerpsd2.String("ultimately"),
        PSUGeoLocation: tjsolerpsd2.String("Michigan"),
        PSUHTTPMethod: tjsolerpsd2.String("Configurable"),
        PSUIPAddress: tjsolerpsd2.String("Group City"),
        PSUIPPort: tjsolerpsd2.String("mobile Bromine"),
        PSUUserAgent: tjsolerpsd2.String("Solomon microchip Account"),
        Signature: "Iron Jazz",
        TPPSignatureCertificate: "streamline",
        XRequestID: "Minivan quirky",
        AccountID: tjsolerpsd2.String("SMS"),
        Aspsp: "Account",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetTrustedBeneficiariesRequest](../../models/operations/gettrustedbeneficiariesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetTrustedBeneficiariesResponse](../../models/operations/gettrustedbeneficiariesresponse.md), error**

