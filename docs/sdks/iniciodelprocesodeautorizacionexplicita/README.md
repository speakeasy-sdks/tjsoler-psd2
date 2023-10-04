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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionCancelacionPago(ctx, operations.PostAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("Global"),
        Digest: "Tricycle",
        PSUAccept: tjsolerpsd2.String("in expedite Cambridgeshire"),
        PSUAcceptCharset: tjsolerpsd2.String("conspirator yum SSD"),
        PSUAcceptEncoding: tjsolerpsd2.String("to male"),
        PSUAcceptLanguage: tjsolerpsd2.String("Cambridgeshire female"),
        PSUDeviceID: tjsolerpsd2.String("Saint"),
        PSUGeoLocation: tjsolerpsd2.String("Passenger before"),
        PSUHTTPMethod: tjsolerpsd2.String("Usability"),
        PSUIPAddress: "schemas Division Plastic",
        PSUIPPort: tjsolerpsd2.String("Usability"),
        PSUUserAgent: tjsolerpsd2.String("plum ivory"),
        Signature: "Electronic",
        TPPNokRedirectURI: tjsolerpsd2.String("from Southeast"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("vortals"),
        TPPSignatureCertificate: "gold Hatchback Rican",
        XRequestID: "protocol Gloves Division",
        Aspsp: operations.PostAutorizacionCancelacionPagoAspspRedsys,
        PaymentID: "Northwest Coordinator",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.PostAutorizacionCancelacionPagoPaymentServicePayments,
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PostAutorizacionCancelacionPagoRequest](../../models/operations/postautorizacioncancelacionpagorequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PostAutorizacionCancelacionPagoResponse](../../models/operations/postautorizacioncancelacionpagoresponse.md), error**


## PostAutorizacionConsentimientosAIS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosAIS(ctx, operations.PostAutorizacionConsentimientosAISRequest{
        ConsentID: tjsolerpsd2.String("interactive aggregate Shirt"),
        Digest: "deposit",
        PSUAccept: tjsolerpsd2.String("deposit"),
        PSUAcceptCharset: tjsolerpsd2.String("Credit"),
        PSUAcceptEncoding: tjsolerpsd2.String("yowza Hatchback cap"),
        PSUAcceptLanguage: tjsolerpsd2.String("drat"),
        PSUCorporateID: tjsolerpsd2.String("Account bypass yippee"),
        PSUCorporateIDType: tjsolerpsd2.String("Gasoline"),
        PSUDeviceID: tjsolerpsd2.String("Investment Division neural"),
        PSUGeoLocation: tjsolerpsd2.String("Elegant"),
        PSUHTTPMethod: tjsolerpsd2.String("Awesome"),
        PsuID: tjsolerpsd2.String("eveniet"),
        PSUIDType: tjsolerpsd2.String("International turquoise"),
        PSUIPAddress: "Northeast traverse",
        PSUIPPort: tjsolerpsd2.String("black Gorgeous"),
        PSUUserAgent: tjsolerpsd2.String("Hat Bahamas South"),
        Signature: "orchestrate up",
        TPPNokRedirectURI: tjsolerpsd2.String("whereas"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Greenland Cisgender South"),
        TPPSignatureCertificate: "networks Toyota West",
        XRequestID: "coffee yahoo gah",
        Aspsp: "North lux Electric",
        ConsentIDPathParameter: "Infrastructure Wagon Berkshire",
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
| `request`                                                                                                                    | [operations.PostAutorizacionConsentimientosAISRequest](../../models/operations/postautorizacionconsentimientosaisrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PostAutorizacionConsentimientosAISResponse](../../models/operations/postautorizacionconsentimientosaisresponse.md), error**


## PostAutorizacionConsentimientosFCS

Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosFCS(ctx, operations.PostAutorizacionConsentimientosFCSRequest{
        ConsentID: tjsolerpsd2.String("South"),
        Digest: "vice",
        PSUAccept: tjsolerpsd2.String("orchestration"),
        PSUAcceptCharset: tjsolerpsd2.String("Southeast"),
        PSUAcceptEncoding: tjsolerpsd2.String("over"),
        PSUAcceptLanguage: tjsolerpsd2.String("Reduced Southwest"),
        PSUCorporateID: tjsolerpsd2.String("Visionary"),
        PSUCorporateIDType: tjsolerpsd2.String("Account Principal"),
        PSUDeviceID: tjsolerpsd2.String("Regional matrix Angelina"),
        PSUGeoLocation: tjsolerpsd2.String("Minivan Fish"),
        PSUHTTPMethod: tjsolerpsd2.String("SDD greatly"),
        PsuID: tjsolerpsd2.String("SMS Borders Potassium"),
        PSUIDType: tjsolerpsd2.String("Rap"),
        PSUIPAddress: "transition Bedfordshire",
        PSUIPPort: tjsolerpsd2.String("Usability Rock aw"),
        PSUUserAgent: tjsolerpsd2.String("Optimized Automotive"),
        Signature: "Panama",
        TPPNokRedirectURI: tjsolerpsd2.String("stain Bicycle South"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Southeast lamb scale"),
        TPPSignatureCertificate: "Volkswagen",
        XRequestID: "New dolor magenta",
        Aspsp: "Northwest Car",
        ConsentIDPathParameter: "program Buckinghamshire hack",
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
| `request`                                                                                                                    | [operations.PostAutorizacionConsentimientosFCSRequest](../../models/operations/postautorizacionconsentimientosfcsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PostAutorizacionConsentimientosFCSResponse](../../models/operations/postautorizacionconsentimientosfcsresponse.md), error**


## PostAutorizacionInicioPago

Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

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
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionInicioPago(ctx, operations.PostAutorizacionInicioPagoRequest{
        ConsentID: tjsolerpsd2.String("yum Soap"),
        Digest: "Bedfordshire Intersex contingency",
        PSUAccept: tjsolerpsd2.String("Demigender Delaware Carolina"),
        PSUAcceptCharset: tjsolerpsd2.String("kilogram Metrics Sedan"),
        PSUAcceptEncoding: tjsolerpsd2.String("er Industrial"),
        PSUAcceptLanguage: tjsolerpsd2.String("Chromium Dinar"),
        PSUDeviceID: tjsolerpsd2.String("occaecati Sharable blanditiis"),
        PSUGeoLocation: tjsolerpsd2.String("Applications"),
        PSUHTTPMethod: tjsolerpsd2.String("Electronic"),
        PSUIPAddress: "inasmuch Cruiser",
        PSUIPPort: tjsolerpsd2.String("Senior"),
        PSUUserAgent: tjsolerpsd2.String("Account toolset protocol"),
        Signature: "Polestar",
        TPPNokRedirectURI: tjsolerpsd2.String("Tools female discrete"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Electric initiatives Chips"),
        TPPSignatureCertificate: "SMS Steel Screen",
        XRequestID: "Supervisor hear Electric",
        Aspsp: operations.PostAutorizacionInicioPagoAspspFiarebancaetica,
        PaymentID: "West Liaison calculating",
        PaymentProduct: operations.PostAutorizacionInicioPagoPaymentProductDomesticBacsPaymentsUk,
        PaymentService: operations.PostAutorizacionInicioPagoPaymentServicePayments,
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostAutorizacionInicioPagoRequest](../../models/operations/postautorizacioniniciopagorequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostAutorizacionInicioPagoResponse](../../models/operations/postautorizacioniniciopagoresponse.md), error**

