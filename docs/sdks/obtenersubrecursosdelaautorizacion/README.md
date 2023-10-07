# ObtenerSubRecursosDeLaAutorizacion
(*ObtenerSubRecursosDeLaAutorizacion*)

### Available Operations

* [GetSubRecursosAutorizacionCancelacionPago](#getsubrecursosautorizacioncancelacionpago) - Endpoint en caso de obtener sub-recursos de la Autorización para Cancelación de Pago
* [GetSubRecursosAutorizacionConsentimientosAIS](#getsubrecursosautorizacionconsentimientosais) - Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos AIS
* [GetSubRecursosAutorizacionConsentimientosFCS](#getsubrecursosautorizacionconsentimientosfcs) - Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos FCS
* [GetSubRecursosAutorizacionInicioPago](#getsubrecursosautorizacioniniciopago) - Endpoint en caso de obtener sub-recursos de la Autorización para Inicio de Pago

## GetSubRecursosAutorizacionCancelacionPago

Endpoint en caso de obtener sub-recursos de la Autorización para Cancelación de Pago

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
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionCancelacionPago(ctx, operations.GetSubRecursosAutorizacionCancelacionPagoRequest{
        Digest: "Van emulation Mouse",
        PSUIPAddress: "Avon clove recontextualize",
        Signature: "Soul",
        TPPSignatureCertificate: "Bedfordshire",
        XRequestID: "Cab black male",
        Aspsp: operations.GetSubRecursosAutorizacionCancelacionPagoAspspCaixabank,
        PaymentID: "at monitoring Fermium",
        PaymentProduct: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentServiceBulkPayments,
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetSubRecursosAutorizacionCancelacionPagoRequest](../../models/operations/getsubrecursosautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetSubRecursosAutorizacionCancelacionPagoResponse](../../models/operations/getsubrecursosautorizacioncancelacionpagoresponse.md), error**


## GetSubRecursosAutorizacionConsentimientosAIS

Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos AIS

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
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionConsentimientosAIS(ctx, operations.GetSubRecursosAutorizacionConsentimientosAISRequest{
        Digest: "Harrisonburg",
        PSUIPAddress: "why haptic Virginia",
        Signature: "Accounts Hat Health",
        TPPSignatureCertificate: "wireless Loan Bahamian",
        XRequestID: "Northeast",
        Aspsp: "loyalty Electronic",
        ConsentIDPathParameter: "geez",
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetSubRecursosAutorizacionConsentimientosAISRequest](../../models/operations/getsubrecursosautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetSubRecursosAutorizacionConsentimientosAISResponse](../../models/operations/getsubrecursosautorizacionconsentimientosaisresponse.md), error**


## GetSubRecursosAutorizacionConsentimientosFCS

Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos FCS

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
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionConsentimientosFCS(ctx, operations.GetSubRecursosAutorizacionConsentimientosFCSRequest{
        Digest: "hybrid extend geez",
        PSUIPAddress: "Refined huzzah",
        Signature: "Azusa SCSI",
        TPPSignatureCertificate: "deposit",
        XRequestID: "HEX",
        Aspsp: "Loan Reactive Iraqi",
        ConsentIDPathParameter: "Central gray Soap",
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.GetSubRecursosAutorizacionConsentimientosFCSRequest](../../models/operations/getsubrecursosautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetSubRecursosAutorizacionConsentimientosFCSResponse](../../models/operations/getsubrecursosautorizacionconsentimientosfcsresponse.md), error**


## GetSubRecursosAutorizacionInicioPago

Endpoint en caso de obtener sub-recursos de la Autorización para Inicio de Pago

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
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionInicioPago(ctx, operations.GetSubRecursosAutorizacionInicioPagoRequest{
        Digest: "unsung Metal",
        PSUIPAddress: "Associate",
        Signature: "withdrawal Handmade",
        TPPSignatureCertificate: "teal strategic",
        XRequestID: "virtual",
        Aspsp: operations.GetSubRecursosAutorizacionInicioPagoAspspMediolanum,
        PaymentID: "Cis parse Music",
        PaymentProduct: operations.GetSubRecursosAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.GetSubRecursosAutorizacionInicioPagoPaymentServicePeriodicPayments,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetSubRecursosAutorizacionInicioPagoRequest](../../models/operations/getsubrecursosautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetSubRecursosAutorizacionInicioPagoResponse](../../models/operations/getsubrecursosautorizacioniniciopagoresponse.md), error**

