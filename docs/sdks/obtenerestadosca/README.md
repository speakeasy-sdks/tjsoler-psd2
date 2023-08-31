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
        ConsentID: psd2cajarural.String("distinctio"),
        Digest: "facilis",
        PSUAccept: psd2cajarural.String("aliquid"),
        PSUAcceptCharset: psd2cajarural.String("quam"),
        PSUAcceptEncoding: psd2cajarural.String("molestias"),
        PSUAcceptLanguage: psd2cajarural.String("temporibus"),
        PSUDeviceID: psd2cajarural.String("qui"),
        PSUGeoLocation: psd2cajarural.String("neque"),
        PSUHTTPMethod: psd2cajarural.String("fugit"),
        PSUIPAddress: "magni",
        PSUIPPort: psd2cajarural.String("odio"),
        PSUUserAgent: psd2cajarural.String("sunt"),
        Signature: "ullam",
        TPPNokRedirectURI: psd2cajarural.String("nam"),
        TPPRedirectPreferred: psd2cajarural.String("hic"),
        TPPRedirectURI: psd2cajarural.String("voluptatem"),
        TPPSignatureCertificate: "cumque",
        XRequestID: "soluta",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoAspspBce,
        AuthorisationID: "et",
        PaymentID: "saepe",
        PaymentProduct: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentProductInstantSepaCreditTransfers,
        PaymentService: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentServicePayments,
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
        Digest: "quos",
        PSUAccept: psd2cajarural.String("tempore"),
        PSUAcceptCharset: psd2cajarural.String("cupiditate"),
        PSUAcceptEncoding: psd2cajarural.String("aperiam"),
        PSUAcceptLanguage: psd2cajarural.String("delectus"),
        PSUDeviceID: psd2cajarural.String("dolorem"),
        PSUGeoLocation: psd2cajarural.String("dolore"),
        PSUHTTPMethod: psd2cajarural.String("labore"),
        PSUIPAddress: "adipisci",
        PSUIPPort: psd2cajarural.String("dolorum"),
        PSUUserAgent: psd2cajarural.String("architecto"),
        Signature: "quae",
        TPPNokRedirectURI: psd2cajarural.String("aut"),
        TPPRedirectPreferred: psd2cajarural.String("quas"),
        TPPRedirectURI: psd2cajarural.String("itaque"),
        TPPSignatureCertificate: "consequatur",
        XRequestID: "est",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISAspspOpenbank,
        AuthorisationID: "porro",
        ConsentIDPathParameter: "doloribus",
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
        ConsentID: psd2cajarural.String("ut"),
        Digest: "facilis",
        PSUAccept: psd2cajarural.String("cupiditate"),
        PSUAcceptCharset: psd2cajarural.String("qui"),
        PSUAcceptEncoding: psd2cajarural.String("quae"),
        PSUAcceptLanguage: psd2cajarural.String("laudantium"),
        PSUDeviceID: psd2cajarural.String("odio"),
        PSUGeoLocation: psd2cajarural.String("occaecati"),
        PSUHTTPMethod: psd2cajarural.String("voluptatibus"),
        PSUIPAddress: "quisquam",
        PSUIPPort: psd2cajarural.String("vero"),
        PSUUserAgent: psd2cajarural.String("omnis"),
        Signature: "quis",
        TPPNokRedirectURI: psd2cajarural.String("ipsum"),
        TPPRedirectPreferred: psd2cajarural.String("delectus"),
        TPPRedirectURI: psd2cajarural.String("voluptate"),
        TPPSignatureCertificate: "consectetur",
        XRequestID: "vero",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSAspspAndbank,
        AuthorisationID: "dignissimos",
        ConsentIDPathParameter: "hic",
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
        ConsentID: psd2cajarural.String("distinctio"),
        Digest: "quod",
        PSUAccept: psd2cajarural.String("odio"),
        PSUAcceptCharset: psd2cajarural.String("similique"),
        PSUAcceptEncoding: psd2cajarural.String("facilis"),
        PSUAcceptLanguage: psd2cajarural.String("vero"),
        PSUDeviceID: psd2cajarural.String("ducimus"),
        PSUGeoLocation: psd2cajarural.String("dolore"),
        PSUHTTPMethod: psd2cajarural.String("quibusdam"),
        PSUIPAddress: "illum",
        PSUIPPort: psd2cajarural.String("sequi"),
        PSUUserAgent: psd2cajarural.String("natus"),
        Signature: "impedit",
        TPPNokRedirectURI: psd2cajarural.String("aut"),
        TPPRedirectPreferred: psd2cajarural.String("voluptatibus"),
        TPPRedirectURI: psd2cajarural.String("exercitationem"),
        TPPSignatureCertificate: "nulla",
        XRequestID: "fugit",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoAspspLaboralkutxa,
        AuthorisationID: "maiores",
        PaymentID: "doloribus",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPaymentProductCrossBorderCreditTransfers,
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

