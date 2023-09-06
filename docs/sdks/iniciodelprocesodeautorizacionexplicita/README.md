# InicioDelProcesoDeAutorizacionExplicita

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionCancelacionPago(ctx, operations.PostAutorizacionCancelacionPagoRequest{
        ConsentID: psd2cajarural.String("nam"),
        Digest: "eaque",
        PSUAccept: psd2cajarural.String("pariatur"),
        PSUAcceptCharset: psd2cajarural.String("nemo"),
        PSUAcceptEncoding: psd2cajarural.String("voluptatibus"),
        PSUAcceptLanguage: psd2cajarural.String("perferendis"),
        PSUDeviceID: psd2cajarural.String("fugiat"),
        PSUGeoLocation: psd2cajarural.String("amet"),
        PSUHTTPMethod: psd2cajarural.String("aut"),
        PSUIPAddress: "cumque",
        PSUIPPort: psd2cajarural.String("corporis"),
        PSUUserAgent: psd2cajarural.String("hic"),
        Signature: "libero",
        TPPNokRedirectURI: psd2cajarural.String("nobis"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("dolores"),
        TPPSignatureCertificate: "quis",
        XRequestID: "totam",
        Aspsp: operations.PostAutorizacionCancelacionPagoAspspEvobanco,
        PaymentID: "eaque",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPaymentProductTarget2Payments,
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosAIS(ctx, operations.PostAutorizacionConsentimientosAISRequest{
        ConsentID: psd2cajarural.String("eos"),
        Digest: "perferendis",
        PSUAccept: psd2cajarural.String("dolores"),
        PSUAcceptCharset: psd2cajarural.String("minus"),
        PSUAcceptEncoding: psd2cajarural.String("quam"),
        PSUAcceptLanguage: psd2cajarural.String("dolor"),
        PSUCorporateID: psd2cajarural.String("vero"),
        PSUCorporateIDType: psd2cajarural.String("nostrum"),
        PSUDeviceID: psd2cajarural.String("hic"),
        PSUGeoLocation: psd2cajarural.String("recusandae"),
        PSUHTTPMethod: psd2cajarural.String("omnis"),
        PsuID: psd2cajarural.String("facilis"),
        PSUIDType: psd2cajarural.String("perspiciatis"),
        PSUIPAddress: "voluptatem",
        PSUIPPort: psd2cajarural.String("porro"),
        PSUUserAgent: psd2cajarural.String("consequuntur"),
        Signature: "blanditiis",
        TPPNokRedirectURI: psd2cajarural.String("error"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("eaque"),
        TPPSignatureCertificate: "occaecati",
        XRequestID: "rerum",
        Aspsp: "adipisci",
        ConsentIDPathParameter: "asperiores",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionConsentimientosFCS(ctx, operations.PostAutorizacionConsentimientosFCSRequest{
        ConsentID: psd2cajarural.String("earum"),
        Digest: "modi",
        PSUAccept: psd2cajarural.String("iste"),
        PSUAcceptCharset: psd2cajarural.String("dolorum"),
        PSUAcceptEncoding: psd2cajarural.String("deleniti"),
        PSUAcceptLanguage: psd2cajarural.String("pariatur"),
        PSUCorporateID: psd2cajarural.String("provident"),
        PSUCorporateIDType: psd2cajarural.String("nobis"),
        PSUDeviceID: psd2cajarural.String("libero"),
        PSUGeoLocation: psd2cajarural.String("delectus"),
        PSUHTTPMethod: psd2cajarural.String("quaerat"),
        PsuID: psd2cajarural.String("quos"),
        PSUIDType: psd2cajarural.String("aliquid"),
        PSUIPAddress: "dolorem",
        PSUIPPort: psd2cajarural.String("dolorem"),
        PSUUserAgent: psd2cajarural.String("dolor"),
        Signature: "qui",
        TPPNokRedirectURI: psd2cajarural.String("ipsum"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("hic"),
        TPPSignatureCertificate: "excepturi",
        XRequestID: "cum",
        Aspsp: "voluptate",
        ConsentIDPathParameter: "dignissimos",
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionInicioPago(ctx, operations.PostAutorizacionInicioPagoRequest{
        ConsentID: psd2cajarural.String("reiciendis"),
        Digest: "amet",
        PSUAccept: psd2cajarural.String("dolorum"),
        PSUAcceptCharset: psd2cajarural.String("numquam"),
        PSUAcceptEncoding: psd2cajarural.String("veritatis"),
        PSUAcceptLanguage: psd2cajarural.String("ipsa"),
        PSUDeviceID: psd2cajarural.String("ipsa"),
        PSUGeoLocation: psd2cajarural.String("iure"),
        PSUHTTPMethod: psd2cajarural.String("odio"),
        PSUIPAddress: "quaerat",
        PSUIPPort: psd2cajarural.String("accusamus"),
        PSUUserAgent: psd2cajarural.String("quidem"),
        Signature: "voluptatibus",
        TPPNokRedirectURI: psd2cajarural.String("voluptas"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("natus"),
        TPPSignatureCertificate: "eos",
        XRequestID: "atque",
        Aspsp: operations.PostAutorizacionInicioPagoAspspRedsys,
        PaymentID: "fugiat",
        PaymentProduct: operations.PostAutorizacionInicioPagoPaymentProductSepaCreditTransfers,
        PaymentService: operations.PostAutorizacionInicioPagoPaymentServicePeriodicPayments,
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

