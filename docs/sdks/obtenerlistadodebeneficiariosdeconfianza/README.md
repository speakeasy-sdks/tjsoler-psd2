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
        ConsentID: "autem",
        Digest: "nobis",
        PSUAccept: psd2cajarural.String("quas"),
        PSUAcceptCharset: psd2cajarural.String("assumenda"),
        PSUAcceptEncoding: psd2cajarural.String("nulla"),
        PSUAcceptLanguage: psd2cajarural.String("voluptas"),
        PSUDeviceID: psd2cajarural.String("libero"),
        PSUGeoLocation: psd2cajarural.String("quasi"),
        PSUHTTPMethod: psd2cajarural.String("tempora"),
        PSUIPAddress: psd2cajarural.String("numquam"),
        PSUIPPort: psd2cajarural.String("explicabo"),
        PSUUserAgent: psd2cajarural.String("provident"),
        Signature: "ipsa",
        TPPSignatureCertificate: "molestiae",
        XRequestID: "magnam",
        AccountID: psd2cajarural.String("odio"),
        Aspsp: "eius",
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

