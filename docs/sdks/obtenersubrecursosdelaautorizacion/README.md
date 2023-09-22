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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionCancelacionPago(ctx, operations.GetSubRecursosAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("quidem"),
        Digest: "neque",
        PSUAccept: tjsolerpsd2.String("quo"),
        PSUAcceptCharset: tjsolerpsd2.String("illum"),
        PSUAcceptEncoding: tjsolerpsd2.String("quo"),
        PSUAcceptLanguage: tjsolerpsd2.String("fuga"),
        PSUDeviceID: tjsolerpsd2.String("eius"),
        PSUGeoLocation: tjsolerpsd2.String("eos"),
        PSUHTTPMethod: tjsolerpsd2.String("voluptas"),
        PSUIPAddress: "ab",
        PSUIPPort: tjsolerpsd2.String("cupiditate"),
        PSUUserAgent: tjsolerpsd2.String("consequatur"),
        Signature: "tempora",
        TPPNokRedirectURI: tjsolerpsd2.String("debitis"),
        TPPRedirectPreferred: tjsolerpsd2.String("ipsam"),
        TPPRedirectURI: tjsolerpsd2.String("aspernatur"),
        TPPSignatureCertificate: "sequi",
        XRequestID: "quo",
        Aspsp: operations.GetSubRecursosAutorizacionCancelacionPagoAspspEurocajarural,
        PaymentID: "recusandae",
        PaymentProduct: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentProductSepaCreditTransfers,
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerSubRecursosDeLaAutorizacion.GetSubRecursosAutorizacionConsentimientosAIS(ctx, operations.GetSubRecursosAutorizacionConsentimientosAISRequest{
        ConsentID: tjsolerpsd2.String("quod"),
        Digest: "dignissimos",
        PSUAccept: tjsolerpsd2.String("inventore"),
        PSUAcceptCharset: tjsolerpsd2.String("nihil"),
        PSUAcceptEncoding: tjsolerpsd2.String("totam"),
        PSUAcceptLanguage: tjsolerpsd2.String("accusamus"),
        PSUDeviceID: tjsolerpsd2.String("aliquam"),
        PSUGeoLocation: tjsolerpsd2.String("odio"),
        PSUHTTPMethod: tjsolerpsd2.String("occaecati"),
        PSUIPAddress: "commodi",
        PSUIPPort: tjsolerpsd2.String("sapiente"),
        PSUUserAgent: tjsolerpsd2.String("dolores"),
        Signature: "deserunt",
        TPPNokRedirectURI: tjsolerpsd2.String("molestiae"),
        TPPRedirectPreferred: tjsolerpsd2.String("accusantium"),
        TPPRedirectURI: tjsolerpsd2.String("porro"),
        TPPSignatureCertificate: "eum",
        XRequestID: "quas",
        Aspsp: "praesentium",
        ConsentIDPathParameter: "consequuntur",
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
        ConsentID: tjsolerpsd2.String("deleniti"),
        Digest: "fugit",
        PSUAccept: tjsolerpsd2.String("fuga"),
        PSUAcceptCharset: tjsolerpsd2.String("mollitia"),
        PSUAcceptEncoding: tjsolerpsd2.String("incidunt"),
        PSUAcceptLanguage: tjsolerpsd2.String("atque"),
        PSUDeviceID: tjsolerpsd2.String("explicabo"),
        PSUGeoLocation: tjsolerpsd2.String("minima"),
        PSUHTTPMethod: tjsolerpsd2.String("nisi"),
        PSUIPAddress: "fugit",
        PSUIPPort: tjsolerpsd2.String("sapiente"),
        PSUUserAgent: tjsolerpsd2.String("consequuntur"),
        Signature: "ratione",
        TPPNokRedirectURI: tjsolerpsd2.String("explicabo"),
        TPPRedirectPreferred: tjsolerpsd2.String("saepe"),
        TPPRedirectURI: tjsolerpsd2.String("occaecati"),
        TPPSignatureCertificate: "atque",
        XRequestID: "et",
        Aspsp: "esse",
        ConsentIDPathParameter: "eveniet",
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
        ConsentID: tjsolerpsd2.String("accusamus"),
        Digest: "veritatis",
        PSUAccept: tjsolerpsd2.String("esse"),
        PSUAcceptCharset: tjsolerpsd2.String("quod"),
        PSUAcceptEncoding: tjsolerpsd2.String("nam"),
        PSUAcceptLanguage: tjsolerpsd2.String("vero"),
        PSUDeviceID: tjsolerpsd2.String("aliquid"),
        PSUGeoLocation: tjsolerpsd2.String("quasi"),
        PSUHTTPMethod: tjsolerpsd2.String("saepe"),
        PSUIPAddress: "vel",
        PSUIPPort: tjsolerpsd2.String("harum"),
        PSUUserAgent: tjsolerpsd2.String("molestiae"),
        Signature: "rerum",
        TPPNokRedirectURI: tjsolerpsd2.String("occaecati"),
        TPPRedirectPreferred: tjsolerpsd2.String("minima"),
        TPPRedirectURI: tjsolerpsd2.String("distinctio"),
        TPPSignatureCertificate: "eligendi",
        XRequestID: "sit",
        Aspsp: operations.GetSubRecursosAutorizacionInicioPagoAspspCajasur,
        PaymentID: "tempore",
        PaymentProduct: operations.GetSubRecursosAutorizacionInicioPagoPaymentProductInstantSepaCreditTransfers,
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

