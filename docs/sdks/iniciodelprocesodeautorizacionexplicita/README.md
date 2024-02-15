# InicioDelProcesoDeAutorizacionExplicita
(*InicioDelProcesoDeAutorizacionExplicita*)

### Available Operations

* [PostAutorizacionCancelacionPago](#postautorizacioncancelacionpago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Cancelación de Pago
* [PostAutorizacionConsentimientosAIS](#postautorizacionconsentimientosais) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS
* [PostAutorizacionConsentimientosFCS](#postautorizacionconsentimientosfcs) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS
* [PostAutorizacionInicioPago](#postautorizacioniniciopago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

## PostAutorizacionCancelacionPago

Endpoint en caso de Inicio del proceso de Autorización explícita para Cancelación de Pago

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionCancelacionPago(ctx, operations.PostAutorizacionCancelacionPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.PostAutorizacionCancelacionPagoPathParamAspspBbva,
        PaymentID: "<value>",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPathParamPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PostAutorizacionCancelacionPagoPathParamPaymentServiceBulkPayments,
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PostAutorizacionCancelacionPagoRequest](../../pkg/models/operations/postautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PostAutorizacionCancelacionPagoResponse](../../pkg/models/operations/postautorizacioncancelacionpagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAutorizacionConsentimientosAIS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosAIS(ctx, operations.PostAutorizacionConsentimientosAISRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PostAutorizacionConsentimientosAISRequest](../../pkg/models/operations/postautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PostAutorizacionConsentimientosAISResponse](../../pkg/models/operations/postautorizacionconsentimientosaisresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAutorizacionConsentimientosFCS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosFCS(ctx, operations.PostAutorizacionConsentimientosFCSRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PostAutorizacionConsentimientosFCSRequest](../../pkg/models/operations/postautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PostAutorizacionConsentimientosFCSResponse](../../pkg/models/operations/postautorizacionconsentimientosfcsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAutorizacionInicioPago

Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionInicioPago(ctx, operations.PostAutorizacionInicioPagoRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: operations.PostAutorizacionInicioPagoPathParamAspspCajasur,
        PaymentID: "<value>",
        PaymentProduct: operations.PostAutorizacionInicioPagoPathParamPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.PostAutorizacionInicioPagoPathParamPaymentServicePayments,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAutorizacionInicioPagoRequest](../../pkg/models/operations/postautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostAutorizacionInicioPagoResponse](../../pkg/models/operations/postautorizacioniniciopagoresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
