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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ConsentimientoDeInformacionSobreCuentasDePagoAIS.PostConsents(ctx, operations.PostConsentsRequest{
        Digest: "Account Sleek",
        PSUIPAddress: "Maryland",
        RequestBody: []byte("T9ccp7rf*J"),
        Signature: "Electronic Omnigender",
        TPPSignatureCertificate: "Latin",
        XRequestID: "bandwidth",
        Aspsp: "Shoes",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PostConsentsRequest](../../models/operations/postconsentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PostConsentsResponse](../../models/operations/postconsentsresponse.md), error**

