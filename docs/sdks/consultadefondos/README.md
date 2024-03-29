# ConsultaDeFondos
(*ConsultaDeFondos*)

### Available Operations

* [FundsConfirmation](#fundsconfirmation) - Consulta de fondos

## FundsConfirmation

Este tipo de mensaje es utilizado en el servicio de consulta de fondos. El HUB consulta al ASPSP por la disponibilidad de fondos para una cantidad dada. El HUB se comunica con el ASPSP para preguntar si tiene fondos y, tras consultarlo, devuelve la respuesta al TPP.

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
    res, err := s.ConsultaDeFondos.FundsConfirmation(ctx, operations.FundsConfirmationRequest{
        ConsentID: "<value>",
        Digest: "<value>",
        RequestBody: []byte("0xFadC6da29A"),
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.FundsConfirmationRequest](../../pkg/models/operations/fundsconfirmationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.FundsConfirmationResponse](../../pkg/models/operations/fundsconfirmationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
