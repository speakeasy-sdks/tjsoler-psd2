# ObtenerEstadoSCA
(*ObtenerEstadoSCA*)

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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionCancelacionPago(ctx, operations.GetEstadoSCAAutorizacionCancelacionPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoPathParamAspspBbvafr,
        AuthorisationID: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.GetEstadoSCAAutorizacionCancelacionPagoPathParamPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.GetEstadoSCAAutorizacionCancelacionPagoPathParamPaymentServicePayments,
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
| `request`                                                                                                                                  | [operations.GetEstadoSCAAutorizacionCancelacionPagoRequest](../../pkg/models/operations/getestadoscaautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetEstadoSCAAutorizacionCancelacionPagoResponse](../../pkg/models/operations/getestadoscaautorizacioncancelacionpagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetEstadoSCAAutorizacionConsentimientosAIS

Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosAIS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosAISRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbvaeuk,
        AuthorisationID: "<value>",
        ConsentIDPathParameter: "<value>",
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
| `request`                                                                                                                                        | [operations.GetEstadoSCAAutorizacionConsentimientosAISRequest](../../pkg/models/operations/getestadoscaautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetEstadoSCAAutorizacionConsentimientosAISResponse](../../pkg/models/operations/getestadoscaautorizacionconsentimientosaisresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetEstadoSCAAutorizacionConsentimientosFCS

Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosFCS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspWizink,
        AuthorisationID: "<value>",
        ConsentIDPathParameter: "<value>",
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
| `request`                                                                                                                                        | [operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest](../../pkg/models/operations/getestadoscaautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |


### Response

**[*operations.GetEstadoSCAAutorizacionConsentimientosFCSResponse](../../pkg/models/operations/getestadoscaautorizacionconsentimientosfcsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetEstadoSCAAutorizacionInicioPago

Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionInicioPago(ctx, operations.GetEstadoSCAAutorizacionInicioPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoPathParamAspspBff,
        AuthorisationID: "<value>",
        PaymentID: "<value>",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPathParamPaymentProductCrossBorderCreditTransfers,
        PaymentService: operations.GetEstadoSCAAutorizacionInicioPagoPathParamPaymentServicePayments,
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
| `request`                                                                                                                        | [operations.GetEstadoSCAAutorizacionInicioPagoRequest](../../pkg/models/operations/getestadoscaautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetEstadoSCAAutorizacionInicioPagoResponse](../../pkg/models/operations/getestadoscaautorizacioniniciopagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
