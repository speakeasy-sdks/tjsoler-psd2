# ServiciosMultibanco

### Available Operations

* [DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID](#deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentid) - Esta petición permite iniciar la cancelación de un pago. Dependiendo del servicio de pago, el producto de pago y la implementación del ASPSP, esta petición podría ser suficiente para cancelar el pago o podría ser necesario una autorización. Si una autorización de la cancelación de pago es necesaria por el ASPSP, el link correspondiente será contenido en el mensaje de respuesta
* [GetMultibancoCatalogue](#getmultibancocatalogue) - Petición iniciada por el TPP para obtener el catálogo de pagos MULTIBANCO
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID](#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentid) - Este mensaje es enviado por el TPP hacia el ASPSP a través del HUB para la recuperación de información del inicio de pago MULTIBANCO.
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations](#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisations)
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID](#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationid)
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus](#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatus) - Mensaje enviado por el TPP al ASPSP a través del Hub para solicitar el estado en el que se encuentra una iniciación de pago.
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations](#postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisations)
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType](#postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttype) - Mensaje enviado por el TPP al ASPSP a través del Hub para crear un inicio de pago MULTIBANCO
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding](#postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholding) - Mensaje enviado por el TPP al ASPSP a través del Hub para conocer el valor del importe a pagar en la seguridad social
* [PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID](#putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationid)

## DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID

Esta petición permite iniciar la cancelación de un pago. Dependiendo del servicio de pago, el producto de pago y la implementación del ASPSP, esta petición podría ser suficiente para cancelar el pago o podría ser necesario una autorización. Si una autorización de la cancelación de pago es necesaria por el ASPSP, el link correspondiente será contenido en el mensaje de respuesta

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
    res, err := s.ServiciosMultibanco.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "in",
        PSUAccept: psd2cajarural.String("eius"),
        PSUAcceptCharset: psd2cajarural.String("libero"),
        PSUAcceptEncoding: psd2cajarural.String("illum"),
        PSUAcceptLanguage: psd2cajarural.String("soluta"),
        PSUDeviceID: psd2cajarural.String("accusantium"),
        PSUGeoLocation: psd2cajarural.String("aliquam"),
        PSUHTTPMethod: psd2cajarural.String("sapiente"),
        PSUIPAddress: "dicta",
        PSUIPPort: psd2cajarural.String("ullam"),
        PSUUserAgent: psd2cajarural.String("reprehenderit"),
        Signature: "ullam",
        TPPSignatureCertificate: "nisi",
        XRequestID: "aut",
        AspName: "voluptatum",
        MultibancoPaymentType: "qui",
        PaymentID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseDeleteMultibankPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                                            | [operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest](../../models/operations/deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                                           |


### Response

**[*operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDResponse](../../models/operations/deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidresponse.md), error**


## GetMultibancoCatalogue

Petición iniciada por el TPP para obtener el catálogo de pagos MULTIBANCO.

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
    res, err := s.ServiciosMultibanco.GetMultibancoCatalogue(ctx, operations.GetMultibancoCatalogueRequest{
        Digest: "ex",
        Signature: "deleniti",
        TPPSignatureCertificate: "itaque",
        XRequestID: "dolorum",
        AspName: "architecto",
        InstructedAmount: psd2cajarural.String("omnis"),
        MultibancoPaymentType: "tenetur",
        PaymentReference: psd2cajarural.String("quasi"),
        RequestedExecutionDate: psd2cajarural.String("at"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetMultibancoCatalogueRequest](../../models/operations/getmultibancocataloguerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetMultibancoCatalogueResponse](../../models/operations/getmultibancocatalogueresponse.md), error**


## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID

Este mensaje es enviado por el TPP hacia el ASPSP a través del HUB para la recuperación de información del inicio de pago MULTIBANCO.

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
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "et",
        PSUAccept: psd2cajarural.String("voluptate"),
        PSUAcceptCharset: psd2cajarural.String("ipsa"),
        PSUAcceptEncoding: psd2cajarural.String("minima"),
        PSUAcceptLanguage: psd2cajarural.String("veritatis"),
        PSUDeviceID: psd2cajarural.String("consectetur"),
        PSUGeoLocation: psd2cajarural.String("adipisci"),
        PSUHTTPMethod: psd2cajarural.String("iste"),
        PSUIPAddress: "temporibus",
        PSUIPPort: psd2cajarural.String("accusantium"),
        PSUUserAgent: psd2cajarural.String("rem"),
        Signature: "aut",
        TPPSignatureCertificate: "laudantium",
        XRequestID: "eum",
        AspName: "mollitia",
        MultibancoPaymentType: "ab",
        PaymentID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseGetMultibankPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                            |
| `request`                                                                                                                                                                                                                      | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                                     |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDResponse](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidresponse.md), error**


## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations

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
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "non",
        PSUAccept: psd2cajarural.String("voluptatem"),
        PSUAcceptCharset: psd2cajarural.String("dolor"),
        PSUAcceptEncoding: psd2cajarural.String("occaecati"),
        PSUAcceptLanguage: psd2cajarural.String("numquam"),
        PSUCorporateID: psd2cajarural.String("impedit"),
        PSUCorporateIDType: psd2cajarural.String("explicabo"),
        PSUDeviceID: psd2cajarural.String("voluptas"),
        PSUGeoLocation: psd2cajarural.String("aut"),
        PSUHTTPMethod: psd2cajarural.String("dignissimos"),
        PSUIPAddress: "dicta",
        PSUIPPort: psd2cajarural.String("maiores"),
        PSUUserAgent: psd2cajarural.String("natus"),
        Signature: "velit",
        TPPSignatureCertificate: "voluptatibus",
        XRequestID: "voluptas",
        AspName: "asperiores",
        MultibancoPaymentType: "aperiam",
        PaymentID: "ea",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseMultibankAuthorizationSubresources != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                                        |
| `request`                                                                                                                                                                                                                                                  | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                                                                 |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsResponse](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsresponse.md), error**


## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID

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
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "quaerat",
        PSUAccept: psd2cajarural.String("consequuntur"),
        PSUAcceptCharset: psd2cajarural.String("repellendus"),
        PSUAcceptEncoding: psd2cajarural.String("officia"),
        PSUAcceptLanguage: psd2cajarural.String("maxime"),
        PSUCorporateID: psd2cajarural.String("dignissimos"),
        PSUCorporateIDType: psd2cajarural.String("officia"),
        PSUDeviceID: psd2cajarural.String("asperiores"),
        PSUGeoLocation: psd2cajarural.String("nemo"),
        PSUHTTPMethod: psd2cajarural.String("quae"),
        PSUIPAddress: "quaerat",
        PSUIPPort: psd2cajarural.String("porro"),
        PSUUserAgent: psd2cajarural.String("quod"),
        Signature: "labore",
        TPPSignatureCertificate: "ab",
        XRequestID: "adipisci",
        AspName: "fuga",
        AuthorisationID: "id",
        MultibancoPaymentType: "suscipit",
        PaymentID: "velit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseMultibankGetSCAStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                                                                                | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                                                                                               |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidresponse.md), error**


## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus

Mensaje enviado por el TPP al ASPSP a través del Hub para solicitar el estado en el que se encuentra una iniciación de pago.

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
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusRequest{
        Digest: "culpa",
        PSUAccept: psd2cajarural.String("est"),
        PSUAcceptCharset: psd2cajarural.String("recusandae"),
        PSUAcceptEncoding: psd2cajarural.String("totam"),
        PSUAcceptLanguage: psd2cajarural.String("fugiat"),
        PSUDeviceID: psd2cajarural.String("vel"),
        PSUGeoLocation: psd2cajarural.String("ducimus"),
        PSUHTTPMethod: psd2cajarural.String("quos"),
        PSUIPAddress: "vel",
        PSUIPPort: psd2cajarural.String("labore"),
        PSUUserAgent: psd2cajarural.String("possimus"),
        Signature: "facilis",
        TPPSignatureCertificate: "cum",
        XRequestID: "commodi",
        AspName: "in",
        MultibancoPaymentType: "corporis",
        PaymentID: "reiciendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseGetStatusMultibankPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                        |
| `request`                                                                                                                                                                                                                                  | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusRequest](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatusrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                                                 |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusResponse](../../models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatusresponse.md), error**


## PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations

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
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "assumenda",
        PSUAccept: psd2cajarural.String("nemo"),
        PSUAcceptCharset: psd2cajarural.String("recusandae"),
        PSUAcceptEncoding: psd2cajarural.String("aliquid"),
        PSUAcceptLanguage: psd2cajarural.String("aperiam"),
        PSUCorporateID: psd2cajarural.String("cum"),
        PSUCorporateIDType: psd2cajarural.String("consectetur"),
        PSUDeviceID: psd2cajarural.String("in"),
        PSUGeoLocation: psd2cajarural.String("exercitationem"),
        PSUHTTPMethod: psd2cajarural.String("earum"),
        PSUIPAddress: "facere",
        PSUIPPort: psd2cajarural.String("numquam"),
        PSUUserAgent: psd2cajarural.String("doloribus"),
        Signature: "suscipit",
        TPPSignatureCertificate: "reiciendis",
        XRequestID: "quidem",
        AspName: "saepe",
        MultibancoPaymentType: "necessitatibus",
        PaymentID: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseMultibankAuthorizeConsentEstablishment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                                                    | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest](../../models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                                                   |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsResponse](../../models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsresponse.md), error**


## PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType

Mensaje enviado por el TPP al ASPSP a través del Hub para crear un inicio de pago MULTIBANCO

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeRequest{
        Digest: "sunt",
        PSUAccept: psd2cajarural.String("asperiores"),
        PSUAcceptCharset: psd2cajarural.String("adipisci"),
        PSUAcceptEncoding: psd2cajarural.String("non"),
        PSUAcceptLanguage: psd2cajarural.String("amet"),
        PSUCorporateID: psd2cajarural.String("beatae"),
        PSUCorporateIDType: psd2cajarural.String("dignissimos"),
        PSUDeviceID: psd2cajarural.String("a"),
        PSUGeoLocation: psd2cajarural.String("debitis"),
        PSUHTTPMethod: psd2cajarural.String("consectetur"),
        PsuID: psd2cajarural.String("corporis"),
        PSUIDType: psd2cajarural.String("harum"),
        PSUIPAddress: "laboriosam",
        PSUIPPort: psd2cajarural.String("ipsa"),
        PSUUserAgent: psd2cajarural.String("voluptates"),
        RequestStartMultibankPayment: shared.RequestStartMultibankPayment{
            DebtorAccount: &shared.RequestStartMultibankPaymentDebtorAccount{},
            EntityCode: psd2cajarural.Int64(10003),
            InstructedAmount: &shared.RequestStartMultibankPaymentInstructedAmount{},
            OperationReference: psd2cajarural.String("1501ab4e-6904-11ea-bc55-0242ac130003"),
            ParameterCode: psd2cajarural.Int64(3),
            PaymentPeriod: &shared.RequestStartMultibankPaymentPaymentPeriod{},
            PaymentReference: psd2cajarural.String("123456789"),
            PaymentType: psd2cajarural.Int64(1),
            PaymentTypeCode: psd2cajarural.String("05"),
            RequestedExecutionDate: psd2cajarural.String("2018-05-17"),
            TaxpayerIdentificationNumber: psd2cajarural.Int64(6244688226942976),
            TsuCenterCode: psd2cajarural.Int64(2698),
        },
        Signature: "libero",
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("vitae"),
        TPPRedirectPreferred: psd2cajarural.String("accusamus"),
        TPPRedirectURI: psd2cajarural.String("similique"),
        TPPSignatureCertificate: "tempora",
        XRequestID: "aspernatur",
        AspName: "voluptas",
        MultibancoPaymentType: "voluptas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseStartMultibankPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                      | Type                                                                                                                                                                                                           | Required                                                                                                                                                                                                       | Description                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                            |
| `request`                                                                                                                                                                                                      | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeRequest](../../models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttyperequest.md) | :heavy_check_mark:                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                     |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeResponse](../../models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttyperesponse.md), error**


## PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding

Mensaje enviado por el TPP al ASPSP a través del Hub para conocer el valor del importe a pagar en la seguridad social.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingRequest{
        Digest: "voluptas",
        PSUAccept: psd2cajarural.String("minima"),
        PSUAcceptCharset: psd2cajarural.String("nobis"),
        PSUAcceptEncoding: psd2cajarural.String("dolorum"),
        PSUAcceptLanguage: psd2cajarural.String("adipisci"),
        PSUCorporateID: psd2cajarural.String("minus"),
        PSUCorporateIDType: psd2cajarural.String("dolores"),
        PSUDeviceID: psd2cajarural.String("blanditiis"),
        PSUGeoLocation: psd2cajarural.String("in"),
        PSUHTTPMethod: psd2cajarural.String("dolore"),
        PsuID: psd2cajarural.String("aliquam"),
        PSUIDType: psd2cajarural.String("officiis"),
        PSUIPAddress: "temporibus",
        PSUIPPort: psd2cajarural.String("ullam"),
        PSUUserAgent: psd2cajarural.String("adipisci"),
        RequestSocialSecurityWithholding: shared.RequestSocialSecurityWithholding{
            DebtorAccount: shared.RequestSocialSecurityWithholdingDebtorAccount{},
            Niss: psd2cajarural.Int64(65136589331),
            Number: 10000000003,
            PaymentDate: "2020-03-04",
            PaymentNature: 4613978645,
            PaymentUnit: 1,
            RemunerationAmount: shared.RequestSocialSecurityWithholdingRemunerationAmount{},
            RemunerationCode: 1,
        },
        Signature: "cum",
        TPPSignatureCertificate: "blanditiis",
        XRequestID: "quas",
        AspName: "hic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseSocialSecurityWithholding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingRequest](../../models/operations/postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholdingrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingResponse](../../models/operations/postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholdingresponse.md), error**


## PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "nesciunt",
        PSUAccept: psd2cajarural.String("culpa"),
        PSUAcceptCharset: psd2cajarural.String("corrupti"),
        PSUAcceptEncoding: psd2cajarural.String("pariatur"),
        PSUAcceptLanguage: psd2cajarural.String("totam"),
        PSUCorporateID: psd2cajarural.String("hic"),
        PSUCorporateIDType: psd2cajarural.String("exercitationem"),
        PSUDeviceID: psd2cajarural.String("nobis"),
        PSUGeoLocation: psd2cajarural.String("sit"),
        PSUHTTPMethod: psd2cajarural.String("rerum"),
        PSUIPAddress: "sed",
        PSUIPPort: psd2cajarural.String("reiciendis"),
        PSUUserAgent: psd2cajarural.String("explicabo"),
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: psd2cajarural.String("123"),
        },
        Signature: "asperiores",
        TPPSignatureCertificate: "facilis",
        XRequestID: "voluptate",
        AspName: "expedita",
        AuthorisationID: "ab",
        MultibancoPaymentType: "iste",
        PaymentID: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseMultibankSelectSCAMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                                                                                | [operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest](../../models/operations/putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                                                                                               |


### Response

**[*operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse](../../models/operations/putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidresponse.md), error**

