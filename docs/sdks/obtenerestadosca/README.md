# ObtenerEstadoSCA

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionCancelacionPago(ctx, operations.GetEstadoSCAAutorizacionCancelacionPagoRequest{
        ConsentID: psd2cajarural.String("laborum"),
        Digest: "totam",
        PSUAccept: psd2cajarural.String("incidunt"),
        PSUAcceptCharset: psd2cajarural.String("aspernatur"),
        PSUAcceptEncoding: psd2cajarural.String("dolores"),
        PSUAcceptLanguage: psd2cajarural.String("distinctio"),
        PSUDeviceID: psd2cajarural.String("facilis"),
        PSUGeoLocation: psd2cajarural.String("aliquid"),
        PSUHTTPMethod: psd2cajarural.String("quam"),
        PSUIPAddress: "molestias",
        PSUIPPort: psd2cajarural.String("temporibus"),
        PSUUserAgent: psd2cajarural.String("qui"),
        Signature: "neque",
        TPPNokRedirectURI: psd2cajarural.String("fugit"),
        TPPRedirectPreferred: psd2cajarural.String("magni"),
        TPPRedirectURI: psd2cajarural.String("odio"),
        TPPSignatureCertificate: "sunt",
        XRequestID: "ullam",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoAspspRenta4,
        AuthorisationID: "hic",
        PaymentID: "voluptatem",
        PaymentProduct: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentServicePeriodicPayments,
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosAIS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosAISRequest{
        ConsentID: psd2cajarural.String("nobis"),
        Digest: "et",
        PSUAccept: psd2cajarural.String("saepe"),
        PSUAcceptCharset: psd2cajarural.String("ipsum"),
        PSUAcceptEncoding: psd2cajarural.String("veritatis"),
        PSUAcceptLanguage: psd2cajarural.String("nobis"),
        PSUDeviceID: psd2cajarural.String("quos"),
        PSUGeoLocation: psd2cajarural.String("tempore"),
        PSUHTTPMethod: psd2cajarural.String("cupiditate"),
        PSUIPAddress: "aperiam",
        PSUIPPort: psd2cajarural.String("delectus"),
        PSUUserAgent: psd2cajarural.String("dolorem"),
        Signature: "dolore",
        TPPNokRedirectURI: psd2cajarural.String("labore"),
        TPPRedirectPreferred: psd2cajarural.String("adipisci"),
        TPPRedirectURI: psd2cajarural.String("dolorum"),
        TPPSignatureCertificate: "architecto",
        XRequestID: "quae",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISAspspRedsys,
        AuthorisationID: "quas",
        ConsentIDPathParameter: "itaque",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosFCS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest{
        ConsentID: psd2cajarural.String("consequatur"),
        Digest: "est",
        PSUAccept: psd2cajarural.String("repellendus"),
        PSUAcceptCharset: psd2cajarural.String("porro"),
        PSUAcceptEncoding: psd2cajarural.String("doloribus"),
        PSUAcceptLanguage: psd2cajarural.String("ut"),
        PSUDeviceID: psd2cajarural.String("facilis"),
        PSUGeoLocation: psd2cajarural.String("cupiditate"),
        PSUHTTPMethod: psd2cajarural.String("qui"),
        PSUIPAddress: "quae",
        PSUIPPort: psd2cajarural.String("laudantium"),
        PSUUserAgent: psd2cajarural.String("odio"),
        Signature: "occaecati",
        TPPNokRedirectURI: psd2cajarural.String("voluptatibus"),
        TPPRedirectPreferred: psd2cajarural.String("quisquam"),
        TPPRedirectURI: psd2cajarural.String("vero"),
        TPPSignatureCertificate: "omnis",
        XRequestID: "quis",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSAspspBancSabadell,
        AuthorisationID: "delectus",
        ConsentIDPathParameter: "voluptate",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionInicioPago(ctx, operations.GetEstadoSCAAutorizacionInicioPagoRequest{
        ConsentID: psd2cajarural.String("consectetur"),
        Digest: "vero",
        PSUAccept: psd2cajarural.String("tenetur"),
        PSUAcceptCharset: psd2cajarural.String("dignissimos"),
        PSUAcceptEncoding: psd2cajarural.String("hic"),
        PSUAcceptLanguage: psd2cajarural.String("distinctio"),
        PSUDeviceID: psd2cajarural.String("quod"),
        PSUGeoLocation: psd2cajarural.String("odio"),
        PSUHTTPMethod: psd2cajarural.String("similique"),
        PSUIPAddress: "facilis",
        PSUIPPort: psd2cajarural.String("vero"),
        PSUUserAgent: psd2cajarural.String("ducimus"),
        Signature: "dolore",
        TPPNokRedirectURI: psd2cajarural.String("quibusdam"),
        TPPRedirectPreferred: psd2cajarural.String("illum"),
        TPPRedirectURI: psd2cajarural.String("sequi"),
        TPPSignatureCertificate: "natus",
        XRequestID: "impedit",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoAspspRedsys,
        AuthorisationID: "voluptatibus",
        PaymentID: "exercitationem",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.GetEstadoSCAAutorizacionInicioPagoPaymentServicePayments,
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

