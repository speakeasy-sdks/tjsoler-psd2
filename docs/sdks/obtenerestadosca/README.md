# ObtenerEstadoSCA
(*ObtenerEstadoSCA*)

### Available Operations

* [GetEstadoSCAAutorizacionCancelacionPago](#getestadoscaautorizacioncancelacionpago) - Endpoint en caso de obtener el Estado SCA para Autorización de Cancelación de Pago
* [GetEstadoSCAAutorizacionConsentimientosAIS](#getestadoscaautorizacionconsentimientosais) - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS
* [GetEstadoSCAAutorizacionConsentimientosFCS](#getestadoscaautorizacionconsentimientosfcs) - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS
* [GetEstadoSCAAutorizacionInicioPago](#getestadoscaautorizacioniniciopago) - Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago

## GetEstadoSCAAutorizacionCancelacionPago

Endpoint en caso de obtener el Estado SCA para Autorización de Cancelación de Pago

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
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionCancelacionPago(ctx, operations.GetEstadoSCAAutorizacionCancelacionPagoRequest{
        Digest: "ADP",
        PSUIPAddress: "Fish vortals",
        Signature: "female shudder",
        TPPSignatureCertificate: "local index streamline",
        XRequestID: "Land volt er",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoAspspBancSabadell,
        AuthorisationID: "gleefully Division Diesel",
        PaymentID: "Montana West",
        PaymentProduct: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentServiceBulkPayments,
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetEstadoSCAAutorizacionCancelacionPagoRequest](../../models/operations/getestadoscaautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetEstadoSCAAutorizacionCancelacionPagoResponse](../../models/operations/getestadoscaautorizacioncancelacionpagoresponse.md), error**


## GetEstadoSCAAutorizacionConsentimientosAIS

Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS

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
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosAIS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosAISRequest{
        Digest: "Account",
        PSUIPAddress: "Carolina reinvent beatae",
        Signature: "Malawi algorithm Quality",
        TPPSignatureCertificate: "cart East",
        XRequestID: "silently primary tangible",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISAspspBbva,
        AuthorisationID: "Architect",
        ConsentIDPathParameter: "lest South",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetEstadoSCAAutorizacionConsentimientosAISRequest](../../models/operations/getestadoscaautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetEstadoSCAAutorizacionConsentimientosAISResponse](../../models/operations/getestadoscaautorizacionconsentimientosaisresponse.md), error**


## GetEstadoSCAAutorizacionConsentimientosFCS

Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS

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
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosFCS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest{
        Digest: "Bespoke connecting Cotton",
        PSUIPAddress: "Outdoors Savings",
        Signature: "peaceful",
        TPPSignatureCertificate: "Central Barthelemy Ezra",
        XRequestID: "North",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSAspspWizink,
        AuthorisationID: "bandwidth",
        ConsentIDPathParameter: "Southwest",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest](../../models/operations/getestadoscaautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.GetEstadoSCAAutorizacionConsentimientosFCSResponse](../../models/operations/getestadoscaautorizacionconsentimientosfcsresponse.md), error**


## GetEstadoSCAAutorizacionInicioPago

Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago

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
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionInicioPago(ctx, operations.GetEstadoSCAAutorizacionInicioPagoRequest{
        Digest: "withdrawal Southeast",
        PSUIPAddress: "testy Avon hacking",
        Signature: "West Connecticut gummy",
        TPPSignatureCertificate: "Westminster",
        XRequestID: "hic disintermediate",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoAspspUnicajabanco,
        AuthorisationID: "Facilitator East Response",
        PaymentID: "parse",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.GetEstadoSCAAutorizacionInicioPagoPaymentServicePeriodicPayments,
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
| `request`                                                                                                                    | [operations.GetEstadoSCAAutorizacionInicioPagoRequest](../../models/operations/getestadoscaautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetEstadoSCAAutorizacionInicioPagoResponse](../../models/operations/getestadoscaautorizacioniniciopagoresponse.md), error**

