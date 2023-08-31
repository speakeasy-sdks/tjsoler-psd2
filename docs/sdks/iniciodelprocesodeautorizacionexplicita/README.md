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
        ConsentID: psd2cajarural.String("perferendis"),
        Digest: "fugiat",
        PSUAccept: psd2cajarural.String("amet"),
        PSUAcceptCharset: psd2cajarural.String("aut"),
        PSUAcceptEncoding: psd2cajarural.String("cumque"),
        PSUAcceptLanguage: psd2cajarural.String("corporis"),
        PSUDeviceID: psd2cajarural.String("hic"),
        PSUGeoLocation: psd2cajarural.String("libero"),
        PSUHTTPMethod: psd2cajarural.String("nobis"),
        PSUIPAddress: "dolores",
        PSUIPPort: psd2cajarural.String("quis"),
        PSUUserAgent: psd2cajarural.String("totam"),
        Signature: "dignissimos",
        TPPNokRedirectURI: psd2cajarural.String("eaque"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("quis"),
        TPPSignatureCertificate: "nesciunt",
        XRequestID: "eos",
        Aspsp: operations.PostAutorizacionCancelacionPagoAspspRedsys,
        PaymentID: "dolores",
        PaymentProduct: operations.PostAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk,
        PaymentService: operations.PostAutorizacionCancelacionPagoPaymentServiceBulkPayments,
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
        ConsentID: psd2cajarural.String("dolor"),
        Digest: "vero",
        PSUAccept: psd2cajarural.String("nostrum"),
        PSUAcceptCharset: psd2cajarural.String("hic"),
        PSUAcceptEncoding: psd2cajarural.String("recusandae"),
        PSUAcceptLanguage: psd2cajarural.String("omnis"),
        PSUCorporateID: psd2cajarural.String("facilis"),
        PSUCorporateIDType: psd2cajarural.String("perspiciatis"),
        PSUDeviceID: psd2cajarural.String("voluptatem"),
        PSUGeoLocation: psd2cajarural.String("porro"),
        PSUHTTPMethod: psd2cajarural.String("consequuntur"),
        PsuID: psd2cajarural.String("blanditiis"),
        PSUIDType: psd2cajarural.String("error"),
        PSUIPAddress: "eaque",
        PSUIPPort: psd2cajarural.String("occaecati"),
        PSUUserAgent: psd2cajarural.String("rerum"),
        Signature: "adipisci",
        TPPNokRedirectURI: psd2cajarural.String("asperiores"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("earum"),
        TPPSignatureCertificate: "modi",
        XRequestID: "iste",
        Aspsp: "dolorum",
        ConsentIDPathParameter: "deleniti",
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
        ConsentID: psd2cajarural.String("pariatur"),
        Digest: "provident",
        PSUAccept: psd2cajarural.String("nobis"),
        PSUAcceptCharset: psd2cajarural.String("libero"),
        PSUAcceptEncoding: psd2cajarural.String("delectus"),
        PSUAcceptLanguage: psd2cajarural.String("quaerat"),
        PSUCorporateID: psd2cajarural.String("quos"),
        PSUCorporateIDType: psd2cajarural.String("aliquid"),
        PSUDeviceID: psd2cajarural.String("dolorem"),
        PSUGeoLocation: psd2cajarural.String("dolorem"),
        PSUHTTPMethod: psd2cajarural.String("dolor"),
        PsuID: psd2cajarural.String("qui"),
        PSUIDType: psd2cajarural.String("ipsum"),
        PSUIPAddress: "hic",
        PSUIPPort: psd2cajarural.String("excepturi"),
        PSUUserAgent: psd2cajarural.String("cum"),
        Signature: "voluptate",
        TPPNokRedirectURI: psd2cajarural.String("dignissimos"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("reiciendis"),
        TPPSignatureCertificate: "amet",
        XRequestID: "dolorum",
        Aspsp: "numquam",
        ConsentIDPathParameter: "veritatis",
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
        ConsentID: psd2cajarural.String("ipsa"),
        Digest: "ipsa",
        PSUAccept: psd2cajarural.String("iure"),
        PSUAcceptCharset: psd2cajarural.String("odio"),
        PSUAcceptEncoding: psd2cajarural.String("quaerat"),
        PSUAcceptLanguage: psd2cajarural.String("accusamus"),
        PSUDeviceID: psd2cajarural.String("quidem"),
        PSUGeoLocation: psd2cajarural.String("voluptatibus"),
        PSUHTTPMethod: psd2cajarural.String("voluptas"),
        PSUIPAddress: "natus",
        PSUIPPort: psd2cajarural.String("eos"),
        PSUUserAgent: psd2cajarural.String("atque"),
        Signature: "sit",
        TPPNokRedirectURI: psd2cajarural.String("fugiat"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("ab"),
        TPPSignatureCertificate: "soluta",
        XRequestID: "dolorum",
        Aspsp: operations.PostAutorizacionInicioPagoAspspEurocajarural,
        PaymentID: "voluptate",
        PaymentProduct: operations.PostAutorizacionInicioPagoPaymentProductDomesticChapsPaymentsUk,
        PaymentService: operations.PostAutorizacionInicioPagoPaymentServiceBulkPayments,
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

