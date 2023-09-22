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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerListadoDeBeneficiariosDeConfianza.GetTrustedBeneficiaries(ctx, operations.GetTrustedBeneficiariesRequest{
        ConsentID: "esse",
        Digest: "esse",
        PSUAccept: tjsolerpsd2.String("rem"),
        PSUAcceptCharset: tjsolerpsd2.String("fuga"),
        PSUAcceptEncoding: tjsolerpsd2.String("reprehenderit"),
        PSUAcceptLanguage: tjsolerpsd2.String("quidem"),
        PSUDeviceID: tjsolerpsd2.String("fugiat"),
        PSUGeoLocation: tjsolerpsd2.String("ut"),
        PSUHTTPMethod: tjsolerpsd2.String("eum"),
        PSUIPAddress: tjsolerpsd2.String("suscipit"),
        PSUIPPort: tjsolerpsd2.String("assumenda"),
        PSUUserAgent: tjsolerpsd2.String("eos"),
        Signature: "praesentium",
        TPPSignatureCertificate: "quisquam",
        XRequestID: "veritatis",
        AccountID: tjsolerpsd2.String("ipsa"),
        Aspsp: "id",
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

