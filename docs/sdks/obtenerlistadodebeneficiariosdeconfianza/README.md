# ObtenerListadoDeBeneficiariosDeConfianza

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerListadoDeBeneficiariosDeConfianza.GetTrustedBeneficiaries(ctx, operations.GetTrustedBeneficiariesRequest{
        ConsentID: "dolorum",
        Digest: "laborum",
        PSUAccept: psd2cajarural.String("placeat"),
        PSUAcceptCharset: psd2cajarural.String("velit"),
        PSUAcceptEncoding: psd2cajarural.String("eum"),
        PSUAcceptLanguage: psd2cajarural.String("autem"),
        PSUDeviceID: psd2cajarural.String("nobis"),
        PSUGeoLocation: psd2cajarural.String("quas"),
        PSUHTTPMethod: psd2cajarural.String("assumenda"),
        PSUIPAddress: psd2cajarural.String("nulla"),
        PSUIPPort: psd2cajarural.String("voluptas"),
        PSUUserAgent: psd2cajarural.String("libero"),
        Signature: "quasi",
        TPPSignatureCertificate: "tempora",
        XRequestID: "numquam",
        AccountID: psd2cajarural.String("explicabo"),
        Aspsp: "provident",
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

