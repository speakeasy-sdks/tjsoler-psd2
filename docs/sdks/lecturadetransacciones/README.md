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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.LecturaDeTransacciones.AccountsTrasactions(ctx, operations.AccountsTrasactionsRequest{
        ConsentID: "totam",
        Digest: "incidunt",
        PSUAccept: tjsolerpsd2.String("aspernatur"),
        PSUAcceptCharset: tjsolerpsd2.String("dolores"),
        PSUAcceptEncoding: tjsolerpsd2.String("distinctio"),
        PSUAcceptLanguage: tjsolerpsd2.String("facilis"),
        PSUDeviceID: tjsolerpsd2.String("aliquid"),
        PSUGeoLocation: tjsolerpsd2.String("quam"),
        PSUHTTPMethod: tjsolerpsd2.String("molestias"),
        PSUIPAddress: tjsolerpsd2.String("temporibus"),
        PSUIPPort: tjsolerpsd2.String("qui"),
        PSUUserAgent: tjsolerpsd2.String("neque"),
        Signature: "fugit",
        TPPSignatureCertificate: "magni",
        XRequestID: "odio",
        AccountID: "sunt",
        Aspsp: "ullam",
        BookingStatus: "nam",
        DateFrom: tjsolerpsd2.String("hic"),
        DateTo: tjsolerpsd2.String("voluptatem"),
        DeltaList: tjsolerpsd2.Bool(false),
        EntryReferenceFrom: tjsolerpsd2.String("cumque"),
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

