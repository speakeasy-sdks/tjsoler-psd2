# LecturaDeDetallesDeCuenta
(*LecturaDeDetallesDeCuenta*)

### Available Operations

* [GetAccountIdv11](#getaccountidv11) - Lectura de detalles de una cuenta

## GetAccountIdv11

Este servicio permite leer los detalles de una cuenta con los balances si son requeridos.Como requisito, se asume que el PSU ha dado su consentimiento para este acceso y ha sido almacenado por el ASPSP. El funcionamiento del servicio según el tipo de acceso indicado en el consentimiento.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.LecturaDeDetallesDeCuenta.GetAccountIdv11(ctx, operations.GetAccountIdv11Request{
        ConsentID: "7890-asdf-4321",
        Digest: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "TestTPPCertificate",
        XRequestID: "<value>",
        AccountID: "<value>",
        Aspsp: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAccountIdv11Request](../../pkg/models/operations/getaccountidv11request.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAccountIdv11Response](../../pkg/models/operations/getaccountidv11response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
