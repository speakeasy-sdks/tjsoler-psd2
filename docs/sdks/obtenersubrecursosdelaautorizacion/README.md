# ObtenerSubRecursosDeLaAutorizacion

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionCancelacionPago(ctx, operations.GetSubRecursosAutorizacionCancelacionPagoRequest{
        ConsentID: psd2cajarural.String("ipsa"),
        Digest: "molestiae",
        PSUAccept: psd2cajarural.String("magnam"),
        PSUAcceptCharset: psd2cajarural.String("odio"),
        PSUAcceptEncoding: psd2cajarural.String("eius"),
        PSUAcceptLanguage: psd2cajarural.String("esse"),
        PSUDeviceID: psd2cajarural.String("esse"),
        PSUGeoLocation: psd2cajarural.String("rem"),
        PSUHTTPMethod: psd2cajarural.String("fuga"),
        PSUIPAddress: "reprehenderit",
        PSUIPPort: psd2cajarural.String("quidem"),
        PSUUserAgent: psd2cajarural.String("fugiat"),
        Signature: "ut",
        TPPNokRedirectURI: psd2cajarural.String("eum"),
        TPPRedirectPreferred: psd2cajarural.String("suscipit"),
        TPPRedirectURI: psd2cajarural.String("assumenda"),
        TPPSignatureCertificate: "eos",
        XRequestID: "praesentium",
        Aspsp: operations.GetSubRecursosAutorizacionCancelacionPagoAspspMediolanum,
        PaymentID: "veritatis",
        PaymentProduct: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentProductSepaCreditTransfers,
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionConsentimientosAIS(ctx, operations.GetSubRecursosAutorizacionConsentimientosAISRequest{
        ConsentID: psd2cajarural.String("quidem"),
        Digest: "neque",
        PSUAccept: psd2cajarural.String("quo"),
        PSUAcceptCharset: psd2cajarural.String("illum"),
        PSUAcceptEncoding: psd2cajarural.String("quo"),
        PSUAcceptLanguage: psd2cajarural.String("fuga"),
        PSUDeviceID: psd2cajarural.String("eius"),
        PSUGeoLocation: psd2cajarural.String("eos"),
        PSUHTTPMethod: psd2cajarural.String("voluptas"),
        PSUIPAddress: "ab",
        PSUIPPort: psd2cajarural.String("cupiditate"),
        PSUUserAgent: psd2cajarural.String("consequatur"),
        Signature: "tempora",
        TPPNokRedirectURI: psd2cajarural.String("debitis"),
        TPPRedirectPreferred: psd2cajarural.String("ipsam"),
        TPPRedirectURI: psd2cajarural.String("aspernatur"),
        TPPSignatureCertificate: "sequi",
        XRequestID: "quo",
        Aspsp: "esse",
        ConsentIDPathParameter: "recusandae",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionConsentimientosFCS(ctx, operations.GetSubRecursosAutorizacionConsentimientosFCSRequest{
        ConsentID: psd2cajarural.String("aperiam"),
        Digest: "distinctio",
        PSUAccept: psd2cajarural.String("quod"),
        PSUAcceptCharset: psd2cajarural.String("dignissimos"),
        PSUAcceptEncoding: psd2cajarural.String("inventore"),
        PSUAcceptLanguage: psd2cajarural.String("nihil"),
        PSUDeviceID: psd2cajarural.String("totam"),
        PSUGeoLocation: psd2cajarural.String("accusamus"),
        PSUHTTPMethod: psd2cajarural.String("aliquam"),
        PSUIPAddress: "odio",
        PSUIPPort: psd2cajarural.String("occaecati"),
        PSUUserAgent: psd2cajarural.String("commodi"),
        Signature: "sapiente",
        TPPNokRedirectURI: psd2cajarural.String("dolores"),
        TPPRedirectPreferred: psd2cajarural.String("deserunt"),
        TPPRedirectURI: psd2cajarural.String("molestiae"),
        TPPSignatureCertificate: "accusantium",
        XRequestID: "porro",
        Aspsp: "eum",
        ConsentIDPathParameter: "quas",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionInicioPago(ctx, operations.GetSubRecursosAutorizacionInicioPagoRequest{
        ConsentID: psd2cajarural.String("praesentium"),
        Digest: "consequuntur",
        PSUAccept: psd2cajarural.String("deleniti"),
        PSUAcceptCharset: psd2cajarural.String("fugit"),
        PSUAcceptEncoding: psd2cajarural.String("fuga"),
        PSUAcceptLanguage: psd2cajarural.String("mollitia"),
        PSUDeviceID: psd2cajarural.String("incidunt"),
        PSUGeoLocation: psd2cajarural.String("atque"),
        PSUHTTPMethod: psd2cajarural.String("explicabo"),
        PSUIPAddress: "minima",
        PSUIPPort: psd2cajarural.String("nisi"),
        PSUUserAgent: psd2cajarural.String("fugit"),
        Signature: "sapiente",
        TPPNokRedirectURI: psd2cajarural.String("consequuntur"),
        TPPRedirectPreferred: psd2cajarural.String("ratione"),
        TPPRedirectURI: psd2cajarural.String("explicabo"),
        TPPSignatureCertificate: "saepe",
        XRequestID: "occaecati",
        Aspsp: operations.GetSubRecursosAutorizacionInicioPagoAspspFiarebancaetica,
        PaymentID: "et",
        PaymentProduct: operations.GetSubRecursosAutorizacionInicioPagoPaymentProductCrossBorderCreditTransfers,
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

