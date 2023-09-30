# LecturaDeTransacciones
(*LecturaDeTransacciones*)

### Available Operations

* [AccountsTrasactions](#accountstrasactions) - Lectura de transacciones de una cuenta

## AccountsTrasactions

Este servicio permite obtener las transacciones de una cuenta determinada por su identificador. Este servicio permite obtener los balances de una cuenta determinada por su identificador. Como requisito, se asume que el PSU ha dado su consentimiento para este acceso y ha sido almacenado por el ASPSP.

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
    res, err := s.LecturaDeTransacciones.AccountsTrasactions(ctx, operations.AccountsTrasactionsRequest{
        ConsentID: "Convertible during Northwest",
        Digest: "granular Folk microchip",
        PSUAccept: tjsolerpsd2.String("minus"),
        PSUAcceptCharset: tjsolerpsd2.String("Smart"),
        PSUAcceptEncoding: tjsolerpsd2.String("tangible atop joule"),
        PSUAcceptLanguage: tjsolerpsd2.String("Hat City"),
        PSUDeviceID: tjsolerpsd2.String("female Tesla"),
        PSUGeoLocation: tjsolerpsd2.String("Newton monitor"),
        PSUHTTPMethod: tjsolerpsd2.String("quantifying Northwest Advanced"),
        PSUIPAddress: tjsolerpsd2.String("Hybrid Sandy Buckinghamshire"),
        PSUIPPort: tjsolerpsd2.String("orange"),
        PSUUserAgent: tjsolerpsd2.String("portals"),
        Signature: "calculating XSS",
        TPPSignatureCertificate: "candela Gasoline",
        XRequestID: "streamline",
        AccountID: "Portugal South",
        Aspsp: "quaerat Iowa",
        BookingStatus: "incremental",
        DateFrom: tjsolerpsd2.String("Factors"),
        DateTo: tjsolerpsd2.String("female"),
        DeltaList: tjsolerpsd2.Bool(false),
        EntryReferenceFrom: tjsolerpsd2.String("connect SUV"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountsTrasactionsRequest](../../models/operations/accountstrasactionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.AccountsTrasactionsResponse](../../models/operations/accountstrasactionsresponse.md), error**

