# ConsentimientoDeInformacionSobreCuentasDePagoAIS
(*ConsentimientoDeInformacionSobreCuentasDePagoAIS*)

### Available Operations

* [PostConsents](#postconsents) - Solicitud de consentimiento AIS

## PostConsents

Con este servicio, un TPP a través del HUB puede solicitar un consentimiento para acceder a las cuentas del PSU. Esta solicitud puede ser sobre unas cuentas indicadas o no.

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
    res, err := s.ConsentimientoDeInformacionSobreCuentasDePagoAIS.PostConsents(ctx, operations.PostConsentsRequest{
        Digest: "string",
        PSUIPAddress: "string",
        RequestBody: []byte("0x99AC4b528c"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostConsentsRequest](../../pkg/models/operations/postconsentsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostConsentsResponse](../../pkg/models/operations/postconsentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
