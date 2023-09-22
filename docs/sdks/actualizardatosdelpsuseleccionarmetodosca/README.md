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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("esse"),
        Digest: "totam",
        PSUAccept: tjsolerpsd2.String("porro"),
        PSUAcceptCharset: tjsolerpsd2.String("dolorum"),
        PSUAcceptEncoding: tjsolerpsd2.String("dicta"),
        PSUAcceptLanguage: tjsolerpsd2.String("nam"),
        PSUDeviceID: tjsolerpsd2.String("officia"),
        PSUGeoLocation: tjsolerpsd2.String("occaecati"),
        PSUHTTPMethod: tjsolerpsd2.String("fugit"),
        PSUIPAddress: "deleniti",
        PSUIPPort: tjsolerpsd2.String("hic"),
        PSUUserAgent: tjsolerpsd2.String("optio"),
        Signature: "totam",
        TPPNokRedirectURI: tjsolerpsd2.String("beatae"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("commodi"),
        TPPSignatureCertificate: "molestiae",
        XRequestID: "modi",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspCaixabank,
        AuthorisationID: "impedit",
        PaymentID: "cum",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePayments,
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosAIS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosAISRequest{
        ConsentID: tjsolerpsd2.String("excepturi"),
        Digest: "aspernatur",
        PSUAccept: tjsolerpsd2.String("perferendis"),
        PSUAcceptCharset: tjsolerpsd2.String("ad"),
        PSUAcceptEncoding: tjsolerpsd2.String("natus"),
        PSUAcceptLanguage: tjsolerpsd2.String("sed"),
        PSUDeviceID: tjsolerpsd2.String("iste"),
        PSUGeoLocation: tjsolerpsd2.String("dolor"),
        PSUHTTPMethod: tjsolerpsd2.String("natus"),
        PSUIPAddress: "laboriosam",
        PSUIPPort: tjsolerpsd2.String("hic"),
        PSUUserAgent: tjsolerpsd2.String("saepe"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "fuga",
        },
        Signature: "in",
        TPPNokRedirectURI: tjsolerpsd2.String("corporis"),
        TPPRedirectPreferred: tjsolerpsd2.String("iste"),
        TPPRedirectURI: tjsolerpsd2.String("iure"),
        TPPSignatureCertificate: "saepe",
        XRequestID: "quidem",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosAISAspspBbvabe,
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosFCS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosFCSRequest{
        ConsentID: tjsolerpsd2.String("est"),
        Digest: "mollitia",
        PSUAccept: tjsolerpsd2.String("laborum"),
        PSUAcceptCharset: tjsolerpsd2.String("dolores"),
        PSUAcceptEncoding: tjsolerpsd2.String("dolorem"),
        PSUAcceptLanguage: tjsolerpsd2.String("corporis"),
        PSUDeviceID: tjsolerpsd2.String("explicabo"),
        PSUGeoLocation: tjsolerpsd2.String("nobis"),
        PSUHTTPMethod: tjsolerpsd2.String("enim"),
        PSUIPAddress: "omnis",
        PSUIPPort: tjsolerpsd2.String("nemo"),
        PSUUserAgent: tjsolerpsd2.String("minima"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "excepturi",
        },
        Signature: "accusantium",
        TPPNokRedirectURI: tjsolerpsd2.String("iure"),
        TPPRedirectPreferred: tjsolerpsd2.String("culpa"),
        TPPRedirectURI: tjsolerpsd2.String("doloribus"),
        TPPSignatureCertificate: "sapiente",
        XRequestID: "architecto",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosFCSAspspKutxabank,
        AuthorisationID: "dolorem",
        ConsentIDPathParameter: "culpa",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionInicioPago(ctx, operations.PutSeleccionarSCAAutorizacionInicioPagoRequest{
        ConsentID: tjsolerpsd2.String("consequuntur"),
        Digest: "repellat",
        PSUAccept: tjsolerpsd2.String("mollitia"),
        PSUAcceptCharset: tjsolerpsd2.String("occaecati"),
        PSUAcceptEncoding: tjsolerpsd2.String("numquam"),
        PSUAcceptLanguage: tjsolerpsd2.String("commodi"),
        PSUDeviceID: tjsolerpsd2.String("quam"),
        PSUGeoLocation: tjsolerpsd2.String("molestiae"),
        PSUHTTPMethod: tjsolerpsd2.String("velit"),
        PSUIPAddress: "error",
        PSUIPPort: tjsolerpsd2.String("quia"),
        PSUUserAgent: tjsolerpsd2.String("quis"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "vitae",
        },
        Signature: "laborum",
        TPPNokRedirectURI: tjsolerpsd2.String("animi"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("enim"),
        TPPSignatureCertificate: "odit",
        XRequestID: "quo",
        Aspsp: operations.PutSeleccionarSCAAutorizacionInicioPagoAspspCaixabank,
        AuthorisationID: "tenetur",
        PaymentID: "ipsam",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentServicePeriodicPayments,
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

