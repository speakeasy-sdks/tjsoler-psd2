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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        Digest: "Tugrik",
        PSUIPAddress: "South",
        Signature: "Northeast",
        TPPSignatureCertificate: "Recycled",
        XRequestID: "Northwest",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBbvafr,
        AuthorisationID: "Connelly",
        PaymentID: "B2B",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticCrossCurrencyPaymentsUk,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServiceBulkPayments,
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
        Digest: "Hybrid",
        PSUIPAddress: "lightly",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "Kirkland",
        },
        Signature: "turquoise",
        TPPSignatureCertificate: "back",
        XRequestID: "male",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosAISAspspBankoa,
        AuthorisationID: "IB",
        ConsentIDPathParameter: "Cisgender",
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
        Digest: "toolset",
        PSUIPAddress: "mobile",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "Cotton",
        },
        Signature: "Ameliorated",
        TPPSignatureCertificate: "Shoes",
        XRequestID: "Designer",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBce,
        AuthorisationID: "Volvo",
        ConsentIDPathParameter: "Crew",
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
        Digest: "model",
        PSUIPAddress: "pixel",
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "Extended",
        },
        Signature: "Extended",
        TPPSignatureCertificate: "Southwest",
        XRequestID: "attenuation",
        Aspsp: operations.PutSeleccionarSCAAutorizacionInicioPagoAspspMediolanum,
        AuthorisationID: "Nihonium",
        PaymentID: "Palladium",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
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

