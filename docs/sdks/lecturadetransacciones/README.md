# LecturaDeTransacciones

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.LecturaDeTransacciones.AccountsTrasactions(ctx, operations.AccountsTrasactionsRequest{
        ConsentID: "expedita",
        Digest: "neque",
        PSUAccept: psd2cajarural.String("sed"),
        PSUAcceptCharset: psd2cajarural.String("vel"),
        PSUAcceptEncoding: psd2cajarural.String("libero"),
        PSUAcceptLanguage: psd2cajarural.String("voluptas"),
        PSUDeviceID: psd2cajarural.String("deserunt"),
        PSUGeoLocation: psd2cajarural.String("quam"),
        PSUHTTPMethod: psd2cajarural.String("ipsum"),
        PSUIPAddress: psd2cajarural.String("incidunt"),
        PSUIPPort: psd2cajarural.String("qui"),
        PSUUserAgent: psd2cajarural.String("cupiditate"),
        Signature: "maxime",
        TPPSignatureCertificate: "pariatur",
        XRequestID: "soluta",
        AccountID: "dicta",
        Aspsp: "laborum",
        BookingStatus: "totam",
        DateFrom: psd2cajarural.String("incidunt"),
        DateTo: psd2cajarural.String("aspernatur"),
        DeltaList: psd2cajarural.Bool(false),
        EntryReferenceFrom: psd2cajarural.String("dolores"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountsTrasactionsRequest](../../models/operations/accountstrasactionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.AccountsTrasactionsResponse](../../models/operations/accountstrasactionsresponse.md), error**

