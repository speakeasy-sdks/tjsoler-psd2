# ActualizarDatosDelPSUSeleccionarMetodoSCA

### Available Operations

* [PutSeleccionarSCAAutorizacionCancelacionPago](#putseleccionarscaautorizacioncancelacionpago) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Cancelación de Pago
* [PutSeleccionarSCAAutorizacionConsentimientosAIS](#putseleccionarscaautorizacionconsentimientosais) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos AIS
* [PutSeleccionarSCAAutorizacionConsentimientosFCS](#putseleccionarscaautorizacionconsentimientosfcs) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos FCS
* [PutSeleccionarSCAAutorizacionInicioPago](#putseleccionarscaautorizacioniniciopago) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Inicio de Pago

## PutSeleccionarSCAAutorizacionCancelacionPago

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Cancelación de Pago

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
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        ConsentID: psd2cajarural.String("voluptatum"),
        Digest: "iusto",
        PSUAccept: psd2cajarural.String("excepturi"),
        PSUAcceptCharset: psd2cajarural.String("nisi"),
        PSUAcceptEncoding: psd2cajarural.String("recusandae"),
        PSUAcceptLanguage: psd2cajarural.String("temporibus"),
        PSUDeviceID: psd2cajarural.String("ab"),
        PSUGeoLocation: psd2cajarural.String("quis"),
        PSUHTTPMethod: psd2cajarural.String("veritatis"),
        PSUIPAddress: "deserunt",
        PSUIPPort: psd2cajarural.String("perferendis"),
        PSUUserAgent: psd2cajarural.String("ipsam"),
        Signature: "repellendus",
        TPPNokRedirectURI: psd2cajarural.String("sapiente"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("quo"),
        TPPSignatureCertificate: "odit",
        XRequestID: "at",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspIbercaja,
        AuthorisationID: "maiores",
        PaymentID: "molestiae",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePeriodicPayments,
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
| `request`                                                                                                                                        | [operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest](../../models/operations/putseleccionarscaautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.PutSeleccionarSCAAutorizacionCancelacionPagoResponse](../../models/operations/putseleccionarscaautorizacioncancelacionpagoresponse.md), error**


## PutSeleccionarSCAAutorizacionConsentimientosAIS

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos AIS

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosAIS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosAISRequest{
        ConsentID: psd2cajarural.String("esse"),
        Digest: "totam",
        PSUAccept: psd2cajarural.String("porro"),
        PSUAcceptCharset: psd2cajarural.String("dolorum"),
        PSUAcceptEncoding: psd2cajarural.String("dicta"),
        PSUAcceptLanguage: psd2cajarural.String("nam"),
        PSUDeviceID: psd2cajarural.String("officia"),
        PSUGeoLocation: psd2cajarural.String("occaecati"),
        PSUHTTPMethod: psd2cajarural.String("fugit"),
        PSUIPAddress: "deleniti",
        PSUIPPort: psd2cajarural.String("hic"),
        PSUUserAgent: psd2cajarural.String("optio"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "totam",
        },
        Signature: "beatae",
        TPPNokRedirectURI: psd2cajarural.String("commodi"),
        TPPRedirectPreferred: psd2cajarural.String("molestiae"),
        TPPRedirectURI: psd2cajarural.String("modi"),
        TPPSignatureCertificate: "qui",
        XRequestID: "impedit",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosAISAspspBce,
        AuthorisationID: "esse",
        ConsentIDPathParameter: "ipsum",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.PutSeleccionarSCAAutorizacionConsentimientosAISRequest](../../models/operations/putseleccionarscaautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PutSeleccionarSCAAutorizacionConsentimientosAISResponse](../../models/operations/putseleccionarscaautorizacionconsentimientosaisresponse.md), error**


## PutSeleccionarSCAAutorizacionConsentimientosFCS

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos FCS

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosFCS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosFCSRequest{
        ConsentID: psd2cajarural.String("excepturi"),
        Digest: "aspernatur",
        PSUAccept: psd2cajarural.String("perferendis"),
        PSUAcceptCharset: psd2cajarural.String("ad"),
        PSUAcceptEncoding: psd2cajarural.String("natus"),
        PSUAcceptLanguage: psd2cajarural.String("sed"),
        PSUDeviceID: psd2cajarural.String("iste"),
        PSUGeoLocation: psd2cajarural.String("dolor"),
        PSUHTTPMethod: psd2cajarural.String("natus"),
        PSUIPAddress: "laboriosam",
        PSUIPPort: psd2cajarural.String("hic"),
        PSUUserAgent: psd2cajarural.String("saepe"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "fuga",
        },
        Signature: "in",
        TPPNokRedirectURI: psd2cajarural.String("corporis"),
        TPPRedirectPreferred: psd2cajarural.String("iste"),
        TPPRedirectURI: psd2cajarural.String("iure"),
        TPPSignatureCertificate: "saepe",
        XRequestID: "quidem",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbvabe,
        AuthorisationID: "ipsa",
        ConsentIDPathParameter: "reiciendis",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.PutSeleccionarSCAAutorizacionConsentimientosFCSRequest](../../models/operations/putseleccionarscaautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PutSeleccionarSCAAutorizacionConsentimientosFCSResponse](../../models/operations/putseleccionarscaautorizacionconsentimientosfcsresponse.md), error**


## PutSeleccionarSCAAutorizacionInicioPago

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Inicio de Pago

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionInicioPago(ctx, operations.PutSeleccionarSCAAutorizacionInicioPagoRequest{
        ConsentID: psd2cajarural.String("est"),
        Digest: "mollitia",
        PSUAccept: psd2cajarural.String("laborum"),
        PSUAcceptCharset: psd2cajarural.String("dolores"),
        PSUAcceptEncoding: psd2cajarural.String("dolorem"),
        PSUAcceptLanguage: psd2cajarural.String("corporis"),
        PSUDeviceID: psd2cajarural.String("explicabo"),
        PSUGeoLocation: psd2cajarural.String("nobis"),
        PSUHTTPMethod: psd2cajarural.String("enim"),
        PSUIPAddress: "omnis",
        PSUIPPort: psd2cajarural.String("nemo"),
        PSUUserAgent: psd2cajarural.String("minima"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "excepturi",
        },
        Signature: "accusantium",
        TPPNokRedirectURI: psd2cajarural.String("iure"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("culpa"),
        TPPSignatureCertificate: "doloribus",
        XRequestID: "sapiente",
        Aspsp: operations.PutSeleccionarSCAAutorizacionInicioPagoAspspBbvabe,
        AuthorisationID: "mollitia",
        PaymentID: "dolorem",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentServicePayments,
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
| `request`                                                                                                                              | [operations.PutSeleccionarSCAAutorizacionInicioPagoRequest](../../models/operations/putseleccionarscaautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.PutSeleccionarSCAAutorizacionInicioPagoResponse](../../models/operations/putseleccionarscaautorizacioniniciopagoresponse.md), error**

