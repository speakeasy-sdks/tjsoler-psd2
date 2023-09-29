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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionCancelacionPago(ctx, operations.GetEstadoSCAAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("ADP"),
        Digest: "Fish vortals",
        PSUAccept: tjsolerpsd2.String("female shudder"),
        PSUAcceptCharset: tjsolerpsd2.String("local index streamline"),
        PSUAcceptEncoding: tjsolerpsd2.String("Land volt er"),
        PSUAcceptLanguage: tjsolerpsd2.String("Hybrid"),
        PSUDeviceID: tjsolerpsd2.String("Division"),
        PSUGeoLocation: tjsolerpsd2.String("Lebanon Montana West"),
        PSUHTTPMethod: tjsolerpsd2.String("Intersex Bicycle HTTP"),
        PSUIPAddress: "Practical Convertible Shirt",
        PSUIPPort: tjsolerpsd2.String("Executive"),
        PSUUserAgent: tjsolerpsd2.String("quasi ivory Gloves"),
        Signature: "Steel 24/365 non",
        TPPNokRedirectURI: tjsolerpsd2.String("Northwest Operations"),
        TPPRedirectPreferred: tjsolerpsd2.String("Ergonomic"),
        TPPRedirectURI: tjsolerpsd2.String("copy"),
        TPPSignatureCertificate: "Rap",
        XRequestID: "South newton grey",
        Aspsp: operations.GetEstadoSCAAutorizacionCancelacionPagoAspspKutxabank,
        AuthorisationID: "back",
        PaymentID: "SUV",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ObtenerEstadoSCA.GetEstadoSCAAutorizacionConsentimientosAIS(ctx, operations.GetEstadoSCAAutorizacionConsentimientosAISRequest{
        ConsentID: tjsolerpsd2.String("Account"),
        Digest: "Carolina reinvent beatae",
        PSUAccept: tjsolerpsd2.String("Malawi algorithm Quality"),
        PSUAcceptCharset: tjsolerpsd2.String("cart East"),
        PSUAcceptEncoding: tjsolerpsd2.String("silently primary tangible"),
        PSUAcceptLanguage: tjsolerpsd2.String("Buckinghamshire"),
        PSUDeviceID: tjsolerpsd2.String("Architect"),
        PSUGeoLocation: tjsolerpsd2.String("lest South"),
        PSUHTTPMethod: tjsolerpsd2.String("primary Sedan damage"),
        PSUIPAddress: "red",
        PSUIPPort: tjsolerpsd2.String("Waipahu Bentley"),
        PSUUserAgent: tjsolerpsd2.String("ugh tomorrow"),
        Signature: "autem",
        TPPNokRedirectURI: tjsolerpsd2.String("Bicycle Hybrid haptic"),
        TPPRedirectPreferred: tjsolerpsd2.String("Incredible"),
        TPPRedirectURI: tjsolerpsd2.String("Southwest Dollar COM"),
        TPPSignatureCertificate: "disunite orchid",
        XRequestID: "up bus Shilling",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosAISAspspBankoa,
        AuthorisationID: "Zambia vicious Bohrium",
        ConsentIDPathParameter: "holistic evolve",
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
        ConsentID: tjsolerpsd2.String("Bespoke connecting Cotton"),
        Digest: "Outdoors Savings",
        PSUAccept: tjsolerpsd2.String("peaceful"),
        PSUAcceptCharset: tjsolerpsd2.String("Central Barthelemy Ezra"),
        PSUAcceptEncoding: tjsolerpsd2.String("North"),
        PSUAcceptLanguage: tjsolerpsd2.String("Hat Northwest Southwest"),
        PSUDeviceID: tjsolerpsd2.String("Architect New"),
        PSUGeoLocation: tjsolerpsd2.String("Zirconium sensitivity"),
        PSUHTTPMethod: tjsolerpsd2.String("Flerovium Gasoline Colon"),
        PSUIPAddress: "phew",
        PSUIPPort: tjsolerpsd2.String("Bespoke"),
        PSUUserAgent: tjsolerpsd2.String("Fish hacking Vallejo"),
        Signature: "Liechtenstein Agent",
        TPPNokRedirectURI: tjsolerpsd2.String("Rath platforms"),
        TPPRedirectPreferred: tjsolerpsd2.String("beast Strategist mint"),
        TPPRedirectURI: tjsolerpsd2.String("South Hybrid optical"),
        TPPSignatureCertificate: "Decentralized",
        XRequestID: "North black",
        Aspsp: operations.GetEstadoSCAAutorizacionConsentimientosFCSAspspBancosantander,
        AuthorisationID: "Corporate Gasoline",
        ConsentIDPathParameter: "quizzically male visualize",
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
        ConsentID: tjsolerpsd2.String("withdrawal Southeast"),
        Digest: "testy Avon hacking",
        PSUAccept: tjsolerpsd2.String("West Connecticut gummy"),
        PSUAcceptCharset: tjsolerpsd2.String("Westminster"),
        PSUAcceptEncoding: tjsolerpsd2.String("hic disintermediate"),
        PSUAcceptLanguage: tjsolerpsd2.String("International Response"),
        PSUDeviceID: tjsolerpsd2.String("parse"),
        PSUGeoLocation: tjsolerpsd2.String("Forward revolutionize"),
        PSUHTTPMethod: tjsolerpsd2.String("Taiwan Rustic"),
        PSUIPAddress: "exuding SSL Hybrid",
        PSUIPPort: tjsolerpsd2.String("indigo Keyboard"),
        PSUUserAgent: tjsolerpsd2.String("Rubber"),
        Signature: "gig",
        TPPNokRedirectURI: tjsolerpsd2.String("CSS"),
        TPPRedirectPreferred: tjsolerpsd2.String("Peso Diesel Denesik"),
        TPPRedirectURI: tjsolerpsd2.String("grey"),
        TPPSignatureCertificate: "Aileen Architect",
        XRequestID: "Europium",
        Aspsp: operations.GetEstadoSCAAutorizacionInicioPagoAspspCajamar,
        AuthorisationID: "purple Hatchback spectate",
        PaymentID: "Buckinghamshire",
        PaymentProduct: operations.GetEstadoSCAAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.GetEstadoSCAAutorizacionInicioPagoPaymentServiceBulkPayments,
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

