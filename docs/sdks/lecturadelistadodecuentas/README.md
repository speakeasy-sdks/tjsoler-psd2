# LecturaDeListadoDeCuentas
(*LecturaDeListadoDeCuentas*)

### Available Operations

* [GetAccountListv11](#getaccountlistv11) - Lectura de listado de cuentas

## GetAccountListv11

Este servicio permite obtener un listado de cuentas del PSU, incluyendo los balances de las cuentas si ha sido requerido. Como requisito, se asume que el PSU ha dado su consentimiento para este acceso y ha sido almacenado por el ASPSP. El funcionamiento del servicio según el tipo de acceso indicado en el consentimiento.

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
    res, err := s.LecturaDeListadoDeCuentas.GetAccountListv11(ctx, operations.GetAccountListv11Request{
        ConsentID: "multimedia Fresh coulomb",
        Digest: "red Ferrari female",
        Signature: "female methodologies Assimilated",
        TPPSignatureCertificate: "empower",
        XRequestID: "Future Kwacha female",
        Aspsp: "to middleware to",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAccountListv11Request](../../models/operations/getaccountlistv11request.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAccountListv11Response](../../models/operations/getaccountlistv11response.md), error**

