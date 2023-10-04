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
        ConsentID: tjsolerpsd2.String("driver Metal"),
        Digest: "Recycled Northwest",
        PSUAccept: tjsolerpsd2.String("Connelly"),
        PSUAcceptCharset: tjsolerpsd2.String("Sleek"),
        PSUAcceptEncoding: tjsolerpsd2.String("grow sunt"),
        PSUAcceptLanguage: tjsolerpsd2.String("connecting Convertible"),
        PSUDeviceID: tjsolerpsd2.String("Directives outside blockchains"),
        PSUGeoLocation: tjsolerpsd2.String("jack transmit"),
        PSUHTTPMethod: tjsolerpsd2.String("SMS"),
        PSUIPAddress: "female",
        PSUIPPort: tjsolerpsd2.String("Gasoline Books Facilitator"),
        PSUUserAgent: tjsolerpsd2.String("Road"),
        Signature: "online",
        TPPNokRedirectURI: tjsolerpsd2.String("mole"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Views application"),
        TPPSignatureCertificate: "Rubber amongst",
        XRequestID: "Beauty siemens Litas",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspWizink,
        AuthorisationID: "untried deposit",
        PaymentID: "Coupe orchid",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductDomesticChapsPaymentsUk,
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
        ConsentID: tjsolerpsd2.String("Trial Kirkland turquoise"),
        Digest: "Dunwoody male",
        PSUAccept: tjsolerpsd2.String("IB"),
        PSUAcceptCharset: tjsolerpsd2.String("Arizona Belgium"),
        PSUAcceptEncoding: tjsolerpsd2.String("Pop repurpose"),
        PSUAcceptLanguage: tjsolerpsd2.String("engineer BMW Executive"),
        PSUDeviceID: tjsolerpsd2.String("solid"),
        PSUGeoLocation: tjsolerpsd2.String("Bedfordshire"),
        PSUHTTPMethod: tjsolerpsd2.String("recusandae Diesel gold"),
        PSUIPAddress: "invoice",
        PSUIPPort: tjsolerpsd2.String("Executive male Wells"),
        PSUUserAgent: tjsolerpsd2.String("East man"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "yet",
        },
        Signature: "Cotton Lodi",
        TPPNokRedirectURI: tjsolerpsd2.String("announce"),
        TPPRedirectPreferred: tjsolerpsd2.String("olive"),
        TPPRedirectURI: tjsolerpsd2.String("digital Road"),
        TPPSignatureCertificate: "Internal Generic",
        XRequestID: "West",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosAISAspspEurocajarural,
        AuthorisationID: "Hybrid Fresh AI",
        ConsentIDPathParameter: "West",
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
        ConsentID: tjsolerpsd2.String("Loaf Cotton"),
        Digest: "Cicero gray",
        PSUAccept: tjsolerpsd2.String("siemens to"),
        PSUAcceptCharset: tjsolerpsd2.String("Fish Frozen"),
        PSUAcceptEncoding: tjsolerpsd2.String("Quality Garrison Garden"),
        PSUAcceptLanguage: tjsolerpsd2.String("West Books"),
        PSUDeviceID: tjsolerpsd2.String("innovate"),
        PSUGeoLocation: tjsolerpsd2.String("Interface farad"),
        PSUHTTPMethod: tjsolerpsd2.String("yellow IP"),
        PSUIPAddress: "orchid",
        PSUIPPort: tjsolerpsd2.String("and"),
        PSUUserAgent: tjsolerpsd2.String("online"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "Place Southeast East",
        },
        Signature: "IP initiatives Baby",
        TPPNokRedirectURI: tjsolerpsd2.String("Electronic Convertible"),
        TPPRedirectPreferred: tjsolerpsd2.String("API"),
        TPPRedirectURI: tjsolerpsd2.String("West female"),
        TPPSignatureCertificate: "firewall generation Bridgeport",
        XRequestID: "Specialist Agent",
        Aspsp: operations.PutSeleccionarSCAAutorizacionConsentimientosFCSAspspArquia,
        AuthorisationID: "Granite",
        ConsentIDPathParameter: "quasi whereas",
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
        ConsentID: tjsolerpsd2.String("Operations Configuration"),
        Digest: "Southeast bifurcated",
        PSUAccept: tjsolerpsd2.String("Wagon"),
        PSUAcceptCharset: tjsolerpsd2.String("Gasoline"),
        PSUAcceptEncoding: tjsolerpsd2.String("Palladium commonly Principal"),
        PSUAcceptLanguage: tjsolerpsd2.String("Avon"),
        PSUDeviceID: tjsolerpsd2.String("Accounts"),
        PSUGeoLocation: tjsolerpsd2.String("Northeast Southeast"),
        PSUHTTPMethod: tjsolerpsd2.String("Games"),
        PSUIPAddress: "Leu deposit Recumbent",
        PSUIPPort: tjsolerpsd2.String("Southeast forecast coulomb"),
        PSUUserAgent: tjsolerpsd2.String("Diesel Principal"),
        RequestUpdatePSUData: shared.RequestUpdatePSUData{
            AuthenticationMethodID: "Modern North",
        },
        Signature: "Shoes Soft Northwest",
        TPPNokRedirectURI: tjsolerpsd2.String("Porsche"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("sticky Handcrafted sky"),
        TPPSignatureCertificate: "withdrawal Southwest",
        XRequestID: "fuel East",
        Aspsp: operations.PutSeleccionarSCAAutorizacionInicioPagoAspspEvobanco,
        AuthorisationID: "eum",
        PaymentID: "Buckridge Electronic circuit",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionInicioPagoPaymentProductSepaCreditTransfers,
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

