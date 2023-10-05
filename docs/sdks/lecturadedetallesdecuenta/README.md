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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.LecturaDeDetallesDeCuenta.GetAccountIdv11(ctx, operations.GetAccountIdv11Request{
        ConsentID: "BMW",
        Digest: "Metal South Incredible",
        PSUAccept: tjsolerpsd2.String("Brand Paucek"),
        PSUAcceptCharset: tjsolerpsd2.String("Burbank"),
        PSUAcceptEncoding: tjsolerpsd2.String("Bicycle male Directives"),
        PSUAcceptLanguage: tjsolerpsd2.String("Ghana Highway Wyoming"),
        PSUDeviceID: tjsolerpsd2.String("male"),
        PSUGeoLocation: tjsolerpsd2.String("pink"),
        PSUHTTPMethod: tjsolerpsd2.String("Human Rubber laborum"),
        PSUIPAddress: tjsolerpsd2.String("Consultant"),
        PSUIPPort: tjsolerpsd2.String("merrily bluetooth"),
        PSUUserAgent: tjsolerpsd2.String("Dinar Global Bicycle"),
        Signature: "Gloves",
        TPPSignatureCertificate: "Health Upgradable",
        XRequestID: "Gambia Waterloo",
        AccountID: "payment product Credit",
        Aspsp: "function",
        WithBalance: tjsolerpsd2.Bool(false),
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

