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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionCancelacionPago(ctx, operations.GetEstadoSCAAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("soluta"),
        Digest: "nobis",
        PSUAccept: tjsolerpsd2.String("et"),
        PSUAcceptCharset: tjsolerpsd2.String("saepe"),
        PSUAcceptEncoding: tjsolerpsd2.String("ipsum"),
        PSUAcceptLanguage: tjsolerpsd2.String("veritatis"),
        PSUDeviceID: tjsolerpsd2.String("nobis"),
        PSUGeoLocation: tjsolerpsd2.String("quos"),
        PSUHTTPMethod: tjsolerpsd2.String("tempore"),
        PSUIPAddress: "cupiditate",
        PSUIPPort: tjsolerpsd2.String("aperiam"),
        PSUUserAgent: tjsolerpsd2.String("delectus"),
        Signature: "dolorem",
        TPPNokRedirectURI: tjsolerpsd2.String("dolore"),
        TPPRedirectPreferred: tjsolerpsd2.String("labore"),
        TPPRedirectURI: tjsolerpsd2.String("adipisci"),
        TPPSignatureCertificate: "dolorum",
        XRequestID: "architecto",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoAspspBbvapt,
        AuthorisationID: "aut",
        PaymentID: "quas",
        PaymentProduct: operations.GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk,
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosAIS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosAISRequest{
        ConsentID: tjsolerpsd2.String("est"),
        Digest: "repellendus",
        PSUAccept: tjsolerpsd2.String("porro"),
        PSUAcceptCharset: tjsolerpsd2.String("doloribus"),
        PSUAcceptEncoding: tjsolerpsd2.String("ut"),
        PSUAcceptLanguage: tjsolerpsd2.String("facilis"),
        PSUDeviceID: tjsolerpsd2.String("cupiditate"),
        PSUGeoLocation: tjsolerpsd2.String("qui"),
        PSUHTTPMethod: tjsolerpsd2.String("quae"),
        PSUIPAddress: "laudantium",
        PSUIPPort: tjsolerpsd2.String("odio"),
        PSUUserAgent: tjsolerpsd2.String("occaecati"),
        Signature: "voluptatibus",
        TPPNokRedirectURI: tjsolerpsd2.String("quisquam"),
        TPPRedirectPreferred: tjsolerpsd2.String("vero"),
        TPPRedirectURI: tjsolerpsd2.String("omnis"),
        TPPSignatureCertificate: "quis",
        XRequestID: "ipsum",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISAspspAndbank,
        AuthorisationID: "voluptate",
        ConsentIDPathParameter: "consectetur",
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
        ConsentID: tjsolerpsd2.String("vero"),
        Digest: "tenetur",
        PSUAccept: tjsolerpsd2.String("dignissimos"),
        PSUAcceptCharset: tjsolerpsd2.String("hic"),
        PSUAcceptEncoding: tjsolerpsd2.String("distinctio"),
        PSUAcceptLanguage: tjsolerpsd2.String("quod"),
        PSUDeviceID: tjsolerpsd2.String("odio"),
        PSUGeoLocation: tjsolerpsd2.String("similique"),
        PSUHTTPMethod: tjsolerpsd2.String("facilis"),
        PSUIPAddress: "vero",
        PSUIPPort: tjsolerpsd2.String("ducimus"),
        PSUUserAgent: tjsolerpsd2.String("dolore"),
        Signature: "quibusdam",
        TPPNokRedirectURI: tjsolerpsd2.String("illum"),
        TPPRedirectPreferred: tjsolerpsd2.String("sequi"),
        TPPRedirectURI: tjsolerpsd2.String("natus"),
        TPPSignatureCertificate: "impedit",
        XRequestID: "aut",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSAspspWizink,
        AuthorisationID: "exercitationem",
        ConsentIDPathParameter: "nulla",
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
        ConsentID: tjsolerpsd2.String("fugit"),
        Digest: "porro",
        PSUAccept: tjsolerpsd2.String("maiores"),
        PSUAcceptCharset: tjsolerpsd2.String("doloribus"),
        PSUAcceptEncoding: tjsolerpsd2.String("iusto"),
        PSUAcceptLanguage: tjsolerpsd2.String("eligendi"),
        PSUDeviceID: tjsolerpsd2.String("ducimus"),
        PSUGeoLocation: tjsolerpsd2.String("alias"),
        PSUHTTPMethod: tjsolerpsd2.String("officia"),
        PSUIPAddress: "tempora",
        PSUIPPort: tjsolerpsd2.String("ipsam"),
        PSUUserAgent: tjsolerpsd2.String("ea"),
        Signature: "aspernatur",
        TPPNokRedirectURI: tjsolerpsd2.String("vel"),
        TPPRedirectPreferred: tjsolerpsd2.String("possimus"),
        TPPRedirectURI: tjsolerpsd2.String("magnam"),
        TPPSignatureCertificate: "ratione",
        XRequestID: "ex",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoAspspEvobanco,
        AuthorisationID: "dicta",
        PaymentID: "dolor",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
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

