# LecturaDeDetallesDeCuenta

### Available Operations

* [GetAccountIdv11](#getaccountidv11) - Lectura de detalles de una cuenta

## GetAccountIdv11

Este servicio permite leer los detalles de una cuenta con los balances si son requeridos.Como requisito, se asume que el PSU ha dado su consentimiento para este acceso y ha sido almacenado por el ASPSP. El funcionamiento del servicio según el tipo de acceso indicado en el consentimiento.

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
    res, err := s.LecturaDeDetallesDeCuenta.GetAccountIdv11(ctx, operations.GetAccountIdv11Request{
        ConsentID: "amet",
        Digest: "optio",
        PSUAccept: psd2cajarural.String("accusamus"),
        PSUAcceptCharset: psd2cajarural.String("ad"),
        PSUAcceptEncoding: psd2cajarural.String("saepe"),
        PSUAcceptLanguage: psd2cajarural.String("suscipit"),
        PSUDeviceID: psd2cajarural.String("deserunt"),
        PSUGeoLocation: psd2cajarural.String("provident"),
        PSUHTTPMethod: psd2cajarural.String("minima"),
        PSUIPAddress: psd2cajarural.String("repellendus"),
        PSUIPPort: psd2cajarural.String("totam"),
        PSUUserAgent: psd2cajarural.String("similique"),
        Signature: "alias",
        TPPSignatureCertificate: "at",
        XRequestID: "quaerat",
        AccountID: "tempora",
        Aspsp: "vel",
        WithBalance: psd2cajarural.Bool(false),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAccountIdv11Request](../../models/operations/getaccountidv11request.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetAccountIdv11Response](../../models/operations/getaccountidv11response.md), error**
