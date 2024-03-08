# ActualizarDatosDelPSUSeleccionarMetodoSCA
(*ActualizarDatosDelPSUSeleccionarMetodoSCA*)

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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.AspspBff,
        AuthorisationID: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.PaymentProductDomesticCrossCurrencyPaymentsUk,
        PaymentService: operations.PaymentServicePayments,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest](../../pkg/models/operations/putseleccionarscaautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.PutSeleccionarSCAAutorizacionCancelacionPagoResponse](../../pkg/models/operations/putseleccionarscaautorizacioncancelacionpagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutSeleccionarSCAAutorizacionConsentimientosAIS

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos AIS

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosAIS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosAISRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "<value>",
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.PathParamAspspOpenbank,
        AuthorisationID: "<value>",
        ConsentIDPathParameter: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.PutSeleccionarSCAAutorizacionConsentimientosAISRequest](../../pkg/models/operations/putseleccionarscaautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |


### Response

**[*operations.PutSeleccionarSCAAutorizacionConsentimientosAISResponse](../../pkg/models/operations/putseleccionarscaautorizacionconsentimientosaisresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutSeleccionarSCAAutorizacionConsentimientosFCS

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos FCS

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionConsentimientosFCS(ctx, operations.PutSeleccionarSCAAutorizacionConsentimientosFCSRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "<value>",
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosFCSPathParamAspspEurocajarural,
        AuthorisationID: "<value>",
        ConsentIDPathParameter: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.PutSeleccionarSCAAutorizacionConsentimientosFCSRequest](../../pkg/models/operations/putseleccionarscaautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |


### Response

**[*operations.PutSeleccionarSCAAutorizacionConsentimientosFCSResponse](../../pkg/models/operations/putseleccionarscaautorizacionconsentimientosfcsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutSeleccionarSCAAutorizacionInicioPago

Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Inicio de Pago

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionInicioPago(ctx, operations.PutSeleccionarSCAAutorizacionInicioPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "<value>",
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.PutSeleccionarSCAAutorizacionInicioPagoPathParamAspspArquia,
        AuthorisationID: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.PathParamPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PathParamPaymentServiceBulkPayments,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PutSeleccionarSCAAutorizacionInicioPagoRequest](../../pkg/models/operations/putseleccionarscaautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.PutSeleccionarSCAAutorizacionInicioPagoResponse](../../pkg/models/operations/putseleccionarscaautorizacioniniciopagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
