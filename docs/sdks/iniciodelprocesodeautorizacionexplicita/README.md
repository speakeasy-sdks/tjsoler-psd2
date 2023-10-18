# InicioDelProcesoDeAutorizacionExplicita
(*InicioDelProcesoDeAutorizacionExplicita*)

### Available Operations

* [PostAutorizacionCancelacionPago](#postautorizacioncancelacionpago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Cancelación de Pago
* [PostAutorizacionConsentimientosAIS](#postautorizacionconsentimientosais) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS
* [PostAutorizacionConsentimientosFCS](#postautorizacionconsentimientosfcs) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS
* [PostAutorizacionInicioPago](#postautorizacioniniciopago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

## PostAutorizacionCancelacionPago

Endpoint en caso de Inicio del proceso de Autorización explícita para Cancelación de Pago

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionCancelacionPago(ctx, operations.PostAutorizacionCancelacionPagoRequest{
        Digest: "Island",
        PSUIPAddress: "Cheese",
        Signature: "win",
        TPPSignatureCertificate: "in",
        XRequestID: "expedite",
        Aspsp: operations.PostAutorizacionCancelacionPagoAspspBbvapt,
        PaymentID: "mmm",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.PostAutorizacionCancelacionPagoPaymentServicePayments,
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PostAutorizacionCancelacionPagoRequest](../../models/operations/postautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PostAutorizacionCancelacionPagoResponse](../../models/operations/postautorizacioncancelacionpagoresponse.md), error**


## PostAutorizacionConsentimientosAIS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosAIS(ctx, operations.PostAutorizacionConsentimientosAISRequest{
        Digest: "indeed",
        PSUIPAddress: "card",
        Signature: "Ball",
        TPPSignatureCertificate: "New",
        XRequestID: "Direct",
        Aspsp: "deposit",
        ConsentIDPathParameter: "Practical",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PostAutorizacionConsentimientosAISRequest](../../models/operations/postautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PostAutorizacionConsentimientosAISResponse](../../models/operations/postautorizacionconsentimientosaisresponse.md), error**


## PostAutorizacionConsentimientosFCS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosFCS(ctx, operations.PostAutorizacionConsentimientosFCSRequest{
        Digest: "Books",
        PSUIPAddress: "Lead",
        Signature: "vice",
        TPPSignatureCertificate: "niches",
        XRequestID: "female",
        Aspsp: "Southeast",
        ConsentIDPathParameter: "West",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PostAutorizacionConsentimientosFCSRequest](../../models/operations/postautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PostAutorizacionConsentimientosFCSResponse](../../models/operations/postautorizacionconsentimientosfcsresponse.md), error**


## PostAutorizacionInicioPago

Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionInicioPago(ctx, operations.PostAutorizacionInicioPagoRequest{
        Digest: "aggregate",
        PSUIPAddress: "Trans",
        Signature: "seamless",
        TPPSignatureCertificate: "API",
        XRequestID: "contingency",
        Aspsp: operations.PostAutorizacionInicioPagoAspspLaboralkutxa,
        PaymentID: "Demigender",
        PaymentProduct: operations.PostAutorizacionInicioPagoPaymentProductInstantSepaCreditTransfers,
        PaymentService: operations.PostAutorizacionInicioPagoPaymentServicePayments,
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostAutorizacionInicioPagoRequest](../../models/operations/postautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostAutorizacionInicioPagoResponse](../../models/operations/postautorizacioniniciopagoresponse.md), error**

