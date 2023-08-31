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
        ConsentID: psd2cajarural.String("esse"),
        Digest: "esse",
        PSUAccept: psd2cajarural.String("rem"),
        PSUAcceptCharset: psd2cajarural.String("fuga"),
        PSUAcceptEncoding: psd2cajarural.String("reprehenderit"),
        PSUAcceptLanguage: psd2cajarural.String("quidem"),
        PSUDeviceID: psd2cajarural.String("fugiat"),
        PSUGeoLocation: psd2cajarural.String("ut"),
        PSUHTTPMethod: psd2cajarural.String("eum"),
        PSUIPAddress: "suscipit",
        PSUIPPort: psd2cajarural.String("assumenda"),
        PSUUserAgent: psd2cajarural.String("eos"),
        Signature: "praesentium",
        TPPNokRedirectURI: psd2cajarural.String("quisquam"),
        TPPRedirectPreferred: psd2cajarural.String("veritatis"),
        TPPRedirectURI: psd2cajarural.String("ipsa"),
        TPPSignatureCertificate: "id",
        XRequestID: "quidem",
        Aspsp: operations.GetSubRecursosAutorizacionCancelacionPagoAspspCaixabank,
        PaymentID: "quo",
        PaymentProduct: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentServicePeriodicPayments,
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
        ConsentID: psd2cajarural.String("fuga"),
        Digest: "eius",
        PSUAccept: psd2cajarural.String("eos"),
        PSUAcceptCharset: psd2cajarural.String("voluptas"),
        PSUAcceptEncoding: psd2cajarural.String("ab"),
        PSUAcceptLanguage: psd2cajarural.String("cupiditate"),
        PSUDeviceID: psd2cajarural.String("consequatur"),
        PSUGeoLocation: psd2cajarural.String("tempora"),
        PSUHTTPMethod: psd2cajarural.String("debitis"),
        PSUIPAddress: "ipsam",
        PSUIPPort: psd2cajarural.String("aspernatur"),
        PSUUserAgent: psd2cajarural.String("sequi"),
        Signature: "quo",
        TPPNokRedirectURI: psd2cajarural.String("esse"),
        TPPRedirectPreferred: psd2cajarural.String("recusandae"),
        TPPRedirectURI: psd2cajarural.String("aperiam"),
        TPPSignatureCertificate: "distinctio",
        XRequestID: "quod",
        Aspsp: "dignissimos",
        ConsentIDPathParameter: "inventore",
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
        ConsentID: psd2cajarural.String("nihil"),
        Digest: "totam",
        PSUAccept: psd2cajarural.String("accusamus"),
        PSUAcceptCharset: psd2cajarural.String("aliquam"),
        PSUAcceptEncoding: psd2cajarural.String("odio"),
        PSUAcceptLanguage: psd2cajarural.String("occaecati"),
        PSUDeviceID: psd2cajarural.String("commodi"),
        PSUGeoLocation: psd2cajarural.String("sapiente"),
        PSUHTTPMethod: psd2cajarural.String("dolores"),
        PSUIPAddress: "deserunt",
        PSUIPPort: psd2cajarural.String("molestiae"),
        PSUUserAgent: psd2cajarural.String("accusantium"),
        Signature: "porro",
        TPPNokRedirectURI: psd2cajarural.String("eum"),
        TPPRedirectPreferred: psd2cajarural.String("quas"),
        TPPRedirectURI: psd2cajarural.String("praesentium"),
        TPPSignatureCertificate: "consequuntur",
        XRequestID: "deleniti",
        Aspsp: "fugit",
        ConsentIDPathParameter: "fuga",
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
        ConsentID: psd2cajarural.String("mollitia"),
        Digest: "incidunt",
        PSUAccept: psd2cajarural.String("atque"),
        PSUAcceptCharset: psd2cajarural.String("explicabo"),
        PSUAcceptEncoding: psd2cajarural.String("minima"),
        PSUAcceptLanguage: psd2cajarural.String("nisi"),
        PSUDeviceID: psd2cajarural.String("fugit"),
        PSUGeoLocation: psd2cajarural.String("sapiente"),
        PSUHTTPMethod: psd2cajarural.String("consequuntur"),
        PSUIPAddress: "ratione",
        PSUIPPort: psd2cajarural.String("explicabo"),
        PSUUserAgent: psd2cajarural.String("saepe"),
        Signature: "occaecati",
        TPPNokRedirectURI: psd2cajarural.String("atque"),
        TPPRedirectPreferred: psd2cajarural.String("et"),
        TPPRedirectURI: psd2cajarural.String("esse"),
        TPPSignatureCertificate: "eveniet",
        XRequestID: "accusamus",
        Aspsp: operations.GetSubRecursosAutorizacionInicioPagoAspspBbvapt,
        PaymentID: "esse",
        PaymentProduct: operations.GetSubRecursosAutorizacionInicioPagoPaymentProductDomesticFasterPaymentsUk,
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

