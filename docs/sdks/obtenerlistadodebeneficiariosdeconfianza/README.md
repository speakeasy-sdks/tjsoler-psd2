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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerListadoDeBeneficiariosDeConfianza.GetTrustedBeneficiaries(ctx, operations.GetTrustedBeneficiariesRequest{
        ConsentID: "string",
        Digest: "string",
        Signature: "string",
        TPPSignatureCertificate: "string",
        XRequestID: "string",
        Aspsp: "string",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetTrustedBeneficiariesRequest](../../pkg/models/operations/gettrustedbeneficiariesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetTrustedBeneficiariesResponse](../../pkg/models/operations/gettrustedbeneficiariesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
