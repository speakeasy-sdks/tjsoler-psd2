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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDelProcesoDeAutorizacionExplicita.PostAutorizacionCancelacionPago(ctx, operations.PostAutorizacionCancelacionPagoRequest{
        ConsentID: tjsolerpsd2.String("perferendis"),
        Digest: "dolores",
        PSUAccept: tjsolerpsd2.String("minus"),
        PSUAcceptCharset: tjsolerpsd2.String("quam"),
        PSUAcceptEncoding: tjsolerpsd2.String("dolor"),
        PSUAcceptLanguage: tjsolerpsd2.String("vero"),
        PSUDeviceID: tjsolerpsd2.String("nostrum"),
        PSUGeoLocation: tjsolerpsd2.String("hic"),
        PSUHTTPMethod: tjsolerpsd2.String("recusandae"),
        PSUIPAddress: "omnis",
        PSUIPPort: tjsolerpsd2.String("facilis"),
        PSUUserAgent: tjsolerpsd2.String("perspiciatis"),
        Signature: "voluptatem",
        TPPNokRedirectURI: tjsolerpsd2.String("porro"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("consequuntur"),
        TPPSignatureCertificate: "blanditiis",
        XRequestID: "error",
        Aspsp: operations.PostAutorizacionCancelacionPagoAspspBbva,
        PaymentID: "occaecati",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPaymentProductDomesticChapsPaymentsUk,
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
        ConsentID: tjsolerpsd2.String("asperiores"),
        Digest: "earum",
        PSUAccept: tjsolerpsd2.String("modi"),
        PSUAcceptCharset: tjsolerpsd2.String("iste"),
        PSUAcceptEncoding: tjsolerpsd2.String("dolorum"),
        PSUAcceptLanguage: tjsolerpsd2.String("deleniti"),
        PSUCorporateID: tjsolerpsd2.String("pariatur"),
        PSUCorporateIDType: tjsolerpsd2.String("provident"),
        PSUDeviceID: tjsolerpsd2.String("nobis"),
        PSUGeoLocation: tjsolerpsd2.String("libero"),
        PSUHTTPMethod: tjsolerpsd2.String("delectus"),
        PsuID: tjsolerpsd2.String("quaerat"),
        PSUIDType: tjsolerpsd2.String("quos"),
        PSUIPAddress: "aliquid",
        PSUIPPort: tjsolerpsd2.String("dolorem"),
        PSUUserAgent: tjsolerpsd2.String("dolorem"),
        Signature: "dolor",
        TPPNokRedirectURI: tjsolerpsd2.String("qui"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("ipsum"),
        TPPSignatureCertificate: "hic",
        XRequestID: "excepturi",
        Aspsp: "cum",
        ConsentIDPathParameter: "voluptate",
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
        ConsentID: tjsolerpsd2.String("dignissimos"),
        Digest: "reiciendis",
        PSUAccept: tjsolerpsd2.String("amet"),
        PSUAcceptCharset: tjsolerpsd2.String("dolorum"),
        PSUAcceptEncoding: tjsolerpsd2.String("numquam"),
        PSUAcceptLanguage: tjsolerpsd2.String("veritatis"),
        PSUCorporateID: tjsolerpsd2.String("ipsa"),
        PSUCorporateIDType: tjsolerpsd2.String("ipsa"),
        PSUDeviceID: tjsolerpsd2.String("iure"),
        PSUGeoLocation: tjsolerpsd2.String("odio"),
        PSUHTTPMethod: tjsolerpsd2.String("quaerat"),
        PsuID: tjsolerpsd2.String("accusamus"),
        PSUIDType: tjsolerpsd2.String("quidem"),
        PSUIPAddress: "voluptatibus",
        PSUIPPort: tjsolerpsd2.String("voluptas"),
        PSUUserAgent: tjsolerpsd2.String("natus"),
        Signature: "eos",
        TPPNokRedirectURI: tjsolerpsd2.String("atque"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("sit"),
        TPPSignatureCertificate: "fugiat",
        XRequestID: "ab",
        Aspsp: "soluta",
        ConsentIDPathParameter: "dolorum",
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
        ConsentID: tjsolerpsd2.String("iusto"),
        Digest: "voluptate",
        PSUAccept: tjsolerpsd2.String("dolorum"),
        PSUAcceptCharset: tjsolerpsd2.String("deleniti"),
        PSUAcceptEncoding: tjsolerpsd2.String("omnis"),
        PSUAcceptLanguage: tjsolerpsd2.String("necessitatibus"),
        PSUDeviceID: tjsolerpsd2.String("distinctio"),
        PSUGeoLocation: tjsolerpsd2.String("asperiores"),
        PSUHTTPMethod: tjsolerpsd2.String("nihil"),
        PSUIPAddress: "ipsum",
        PSUIPPort: tjsolerpsd2.String("voluptate"),
        PSUUserAgent: tjsolerpsd2.String("id"),
        Signature: "saepe",
        TPPNokRedirectURI: tjsolerpsd2.String("eius"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("aspernatur"),
        TPPSignatureCertificate: "perferendis",
        XRequestID: "amet",
        Aspsp: operations.PostAutorizacionInicioPagoAspspLaboralkutxa,
        PaymentID: "accusamus",
        PaymentProduct: operations.PostAutorizacionInicioPagoPaymentProductTarget2Payments,
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

