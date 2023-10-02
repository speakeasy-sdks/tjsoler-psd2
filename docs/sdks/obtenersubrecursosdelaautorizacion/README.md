# ObtenerSubRecursosDeLaAutorizacion
(*ObtenerSubRecursosDeLaAutorizacion*)

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
        ConsentID: tjsolerpsd2.String("Van emulation Mouse"),
        Digest: "Avon clove recontextualize",
        PSUAccept: tjsolerpsd2.String("Soul"),
        PSUAcceptCharset: tjsolerpsd2.String("Bedfordshire"),
        PSUAcceptEncoding: tjsolerpsd2.String("Cab black male"),
        PSUAcceptLanguage: tjsolerpsd2.String("Tricycle"),
        PSUDeviceID: tjsolerpsd2.String("monitoring"),
        PSUGeoLocation: tjsolerpsd2.String("Wagon program blue"),
        PSUHTTPMethod: tjsolerpsd2.String("Bedfordshire BCEAO content"),
        PSUIPAddress: "from Associate",
        PSUIPPort: tjsolerpsd2.String("McGlynn Franc"),
        PSUUserAgent: tjsolerpsd2.String("Mountain"),
        Signature: "modular BMW Facilitator",
        TPPNokRedirectURI: tjsolerpsd2.String("mundane"),
        TPPRedirectPreferred: tjsolerpsd2.String("blue per"),
        TPPRedirectURI: tjsolerpsd2.String("Berkshire"),
        TPPSignatureCertificate: "Borders newton steradian",
        XRequestID: "Multigender white auxiliary",
        Aspsp: operations.GetSubRecursosAutorizacionCancelacionPagoAspspCaixabank,
        PaymentID: "bifurcated",
        PaymentProduct: operations.GetSubRecursosAutorizacionCancelacionPagoPaymentProductInstantSepaCreditTransfers,
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
        ConsentID: tjsolerpsd2.String("Harrisonburg"),
        Digest: "why haptic Virginia",
        PSUAccept: tjsolerpsd2.String("Accounts Hat Health"),
        PSUAcceptCharset: tjsolerpsd2.String("wireless Loan Bahamian"),
        PSUAcceptEncoding: tjsolerpsd2.String("Northeast"),
        PSUAcceptLanguage: tjsolerpsd2.String("loyalty Electronic"),
        PSUDeviceID: tjsolerpsd2.String("geez"),
        PSUGeoLocation: tjsolerpsd2.String("blanditiis"),
        PSUHTTPMethod: tjsolerpsd2.String("Principal inglenook gold"),
        PSUIPAddress: "success District",
        PSUIPPort: tjsolerpsd2.String("Egypt Diesel Unbranded"),
        PSUUserAgent: tjsolerpsd2.String("Synergistic sievert Canada"),
        Signature: "Touring",
        TPPNokRedirectURI: tjsolerpsd2.String("gold Mauritius"),
        TPPRedirectPreferred: tjsolerpsd2.String("Folk Legacy Practical"),
        TPPRedirectURI: tjsolerpsd2.String("holistic dolor Florida"),
        TPPSignatureCertificate: "second Tuna",
        XRequestID: "firewall Nakfa Northwest",
        Aspsp: "Malaysian",
        ConsentIDPathParameter: "capacitor azure Clothing",
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
        ConsentID: tjsolerpsd2.String("hybrid extend geez"),
        Digest: "Refined huzzah",
        PSUAccept: tjsolerpsd2.String("Azusa SCSI"),
        PSUAcceptCharset: tjsolerpsd2.String("deposit"),
        PSUAcceptEncoding: tjsolerpsd2.String("HEX"),
        PSUAcceptLanguage: tjsolerpsd2.String("Loan Reactive Iraqi"),
        PSUDeviceID: tjsolerpsd2.String("Central gray Soap"),
        PSUGeoLocation: tjsolerpsd2.String("from framework Tunisia"),
        PSUHTTPMethod: tjsolerpsd2.String("Bicycle district"),
        PSUIPAddress: "coagulate",
        PSUIPPort: tjsolerpsd2.String("West"),
        PSUUserAgent: tjsolerpsd2.String("Bentley"),
        Signature: "Euless",
        TPPNokRedirectURI: tjsolerpsd2.String("lumpy Account application"),
        TPPRedirectPreferred: tjsolerpsd2.String("Montana"),
        TPPRedirectURI: tjsolerpsd2.String("zowie Yuan Rolls"),
        TPPSignatureCertificate: "darling 24/7 Cab",
        XRequestID: "success",
        Aspsp: "rich Product",
        ConsentIDPathParameter: "Planner maroon Zambian",
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
        ConsentID: tjsolerpsd2.String("unsung Metal"),
        Digest: "Associate",
        PSUAccept: tjsolerpsd2.String("withdrawal Handmade"),
        PSUAcceptCharset: tjsolerpsd2.String("teal strategic"),
        PSUAcceptEncoding: tjsolerpsd2.String("virtual"),
        PSUAcceptLanguage: tjsolerpsd2.String("Orchestrator West parse"),
        PSUDeviceID: tjsolerpsd2.String("Southwest"),
        PSUGeoLocation: tjsolerpsd2.String("Metal Account West"),
        PSUHTTPMethod: tjsolerpsd2.String("Tin Corkery"),
        PSUIPAddress: "azure",
        PSUIPPort: tjsolerpsd2.String("Borders"),
        PSUUserAgent: tjsolerpsd2.String("blue Gorgeous Toyota"),
        Signature: "Llewellyn",
        TPPNokRedirectURI: tjsolerpsd2.String("hospitality index"),
        TPPRedirectPreferred: tjsolerpsd2.String("invoice Bedfordshire"),
        TPPRedirectURI: tjsolerpsd2.String("furiously ivory to"),
        TPPSignatureCertificate: "Security",
        XRequestID: "Granite Salad New",
        Aspsp: operations.GetSubRecursosAutorizacionInicioPagoAspspBbvapt,
        PaymentID: "Diesel",
        PaymentProduct: operations.GetSubRecursosAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.GetSubRecursosAutorizacionInicioPagoPaymentServicePayments,
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

