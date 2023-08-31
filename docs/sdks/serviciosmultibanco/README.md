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
        Digest: "accusantium",
        PSUAccept: psd2cajarural.String("aliquam"),
        PSUAcceptCharset: psd2cajarural.String("sapiente"),
        PSUAcceptEncoding: psd2cajarural.String("dicta"),
        PSUAcceptLanguage: psd2cajarural.String("ullam"),
        PSUDeviceID: psd2cajarural.String("reprehenderit"),
        PSUGeoLocation: psd2cajarural.String("ullam"),
        PSUHTTPMethod: psd2cajarural.String("nisi"),
        PSUIPAddress: "aut",
        PSUIPPort: psd2cajarural.String("voluptatum"),
        PSUUserAgent: psd2cajarural.String("qui"),
        Signature: "quibusdam",
        TPPSignatureCertificate: "ex",
        XRequestID: "deleniti",
        AspName: "itaque",
        MultibancoPaymentType: "dolorum",
        PaymentID: "architecto",
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
        Digest: "omnis",
        Signature: "tenetur",
        TPPSignatureCertificate: "quasi",
        XRequestID: "at",
        AspName: "et",
        InstructedAmount: psd2cajarural.String("voluptate"),
        MultibancoPaymentType: "ipsa",
        PaymentReference: psd2cajarural.String("minima"),
        RequestedExecutionDate: psd2cajarural.String("veritatis"),
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
        Digest: "consectetur",
        PSUAccept: psd2cajarural.String("adipisci"),
        PSUAcceptCharset: psd2cajarural.String("iste"),
        PSUAcceptEncoding: psd2cajarural.String("temporibus"),
        PSUAcceptLanguage: psd2cajarural.String("accusantium"),
        PSUDeviceID: psd2cajarural.String("rem"),
        PSUGeoLocation: psd2cajarural.String("aut"),
        PSUHTTPMethod: psd2cajarural.String("laudantium"),
        PSUIPAddress: "eum",
        PSUIPPort: psd2cajarural.String("mollitia"),
        PSUUserAgent: psd2cajarural.String("ab"),
        Signature: "corrupti",
        TPPSignatureCertificate: "non",
        XRequestID: "voluptatem",
        AspName: "dolor",
        MultibancoPaymentType: "occaecati",
        PaymentID: "numquam",
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
        Digest: "impedit",
        PSUAccept: psd2cajarural.String("explicabo"),
        PSUAcceptCharset: psd2cajarural.String("voluptas"),
        PSUAcceptEncoding: psd2cajarural.String("aut"),
        PSUAcceptLanguage: psd2cajarural.String("dignissimos"),
        PSUCorporateID: psd2cajarural.String("dicta"),
        PSUCorporateIDType: psd2cajarural.String("maiores"),
        PSUDeviceID: psd2cajarural.String("natus"),
        PSUGeoLocation: psd2cajarural.String("velit"),
        PSUHTTPMethod: psd2cajarural.String("voluptatibus"),
        PSUIPAddress: "voluptas",
        PSUIPPort: psd2cajarural.String("asperiores"),
        PSUUserAgent: psd2cajarural.String("aperiam"),
        Signature: "ea",
        TPPSignatureCertificate: "quaerat",
        XRequestID: "consequuntur",
        AspName: "repellendus",
        MultibancoPaymentType: "officia",
        PaymentID: "maxime",
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
        Digest: "dignissimos",
        PSUAccept: psd2cajarural.String("officia"),
        PSUAcceptCharset: psd2cajarural.String("asperiores"),
        PSUAcceptEncoding: psd2cajarural.String("nemo"),
        PSUAcceptLanguage: psd2cajarural.String("quae"),
        PSUCorporateID: psd2cajarural.String("quaerat"),
        PSUCorporateIDType: psd2cajarural.String("porro"),
        PSUDeviceID: psd2cajarural.String("quod"),
        PSUGeoLocation: psd2cajarural.String("labore"),
        PSUHTTPMethod: psd2cajarural.String("ab"),
        PSUIPAddress: "adipisci",
        PSUIPPort: psd2cajarural.String("fuga"),
        PSUUserAgent: psd2cajarural.String("id"),
        Signature: "suscipit",
        TPPSignatureCertificate: "velit",
        XRequestID: "culpa",
        AspName: "est",
        AuthorisationID: "recusandae",
        MultibancoPaymentType: "totam",
        PaymentID: "fugiat",
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
        Digest: "vel",
        PSUAccept: psd2cajarural.String("ducimus"),
        PSUAcceptCharset: psd2cajarural.String("quos"),
        PSUAcceptEncoding: psd2cajarural.String("vel"),
        PSUAcceptLanguage: psd2cajarural.String("labore"),
        PSUDeviceID: psd2cajarural.String("possimus"),
        PSUGeoLocation: psd2cajarural.String("facilis"),
        PSUHTTPMethod: psd2cajarural.String("cum"),
        PSUIPAddress: "commodi",
        PSUIPPort: psd2cajarural.String("in"),
        PSUUserAgent: psd2cajarural.String("corporis"),
        Signature: "reiciendis",
        TPPSignatureCertificate: "assumenda",
        XRequestID: "nemo",
        AspName: "recusandae",
        MultibancoPaymentType: "aliquid",
        PaymentID: "aperiam",
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
        Digest: "cum",
        PSUAccept: psd2cajarural.String("consectetur"),
        PSUAcceptCharset: psd2cajarural.String("in"),
        PSUAcceptEncoding: psd2cajarural.String("exercitationem"),
        PSUAcceptLanguage: psd2cajarural.String("earum"),
        PSUCorporateID: psd2cajarural.String("facere"),
        PSUCorporateIDType: psd2cajarural.String("numquam"),
        PSUDeviceID: psd2cajarural.String("doloribus"),
        PSUGeoLocation: psd2cajarural.String("suscipit"),
        PSUHTTPMethod: psd2cajarural.String("reiciendis"),
        PSUIPAddress: "quidem",
        PSUIPPort: psd2cajarural.String("saepe"),
        PSUUserAgent: psd2cajarural.String("necessitatibus"),
        Signature: "dolore",
        TPPSignatureCertificate: "sunt",
        XRequestID: "asperiores",
        AspName: "adipisci",
        MultibancoPaymentType: "non",
        PaymentID: "amet",
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
        Digest: "beatae",
        PSUAccept: psd2cajarural.String("dignissimos"),
        PSUAcceptCharset: psd2cajarural.String("a"),
        PSUAcceptEncoding: psd2cajarural.String("debitis"),
        PSUAcceptLanguage: psd2cajarural.String("consectetur"),
        PSUCorporateID: psd2cajarural.String("corporis"),
        PSUCorporateIDType: psd2cajarural.String("harum"),
        PSUDeviceID: psd2cajarural.String("laboriosam"),
        PSUGeoLocation: psd2cajarural.String("ipsa"),
        PSUHTTPMethod: psd2cajarural.String("voluptates"),
        PsuID: psd2cajarural.String("libero"),
        PSUIDType: psd2cajarural.String("vitae"),
        PSUIPAddress: "accusamus",
        PSUIPPort: psd2cajarural.String("similique"),
        PSUUserAgent: psd2cajarural.String("tempora"),
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
        Signature: "aspernatur",
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("voluptas"),
        TPPRedirectPreferred: psd2cajarural.String("voluptas"),
        TPPRedirectURI: psd2cajarural.String("voluptas"),
        TPPSignatureCertificate: "minima",
        XRequestID: "nobis",
        AspName: "dolorum",
        MultibancoPaymentType: "adipisci",
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
        Digest: "minus",
        PSUAccept: psd2cajarural.String("dolores"),
        PSUAcceptCharset: psd2cajarural.String("blanditiis"),
        PSUAcceptEncoding: psd2cajarural.String("in"),
        PSUAcceptLanguage: psd2cajarural.String("dolore"),
        PSUCorporateID: psd2cajarural.String("aliquam"),
        PSUCorporateIDType: psd2cajarural.String("officiis"),
        PSUDeviceID: psd2cajarural.String("temporibus"),
        PSUGeoLocation: psd2cajarural.String("ullam"),
        PSUHTTPMethod: psd2cajarural.String("adipisci"),
        PsuID: psd2cajarural.String("cum"),
        PSUIDType: psd2cajarural.String("blanditiis"),
        PSUIPAddress: "quas",
        PSUIPPort: psd2cajarural.String("hic"),
        PSUUserAgent: psd2cajarural.String("nesciunt"),
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
        Signature: "culpa",
        TPPSignatureCertificate: "corrupti",
        XRequestID: "pariatur",
        AspName: "totam",
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
        Digest: "hic",
        PSUAccept: psd2cajarural.String("exercitationem"),
        PSUAcceptCharset: psd2cajarural.String("nobis"),
        PSUAcceptEncoding: psd2cajarural.String("sit"),
        PSUAcceptLanguage: psd2cajarural.String("rerum"),
        PSUCorporateID: psd2cajarural.String("sed"),
        PSUCorporateIDType: psd2cajarural.String("reiciendis"),
        PSUDeviceID: psd2cajarural.String("explicabo"),
        PSUGeoLocation: psd2cajarural.String("asperiores"),
        PSUHTTPMethod: psd2cajarural.String("facilis"),
        PSUIPAddress: "voluptate",
        PSUIPPort: psd2cajarural.String("expedita"),
        PSUUserAgent: psd2cajarural.String("ab"),
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: psd2cajarural.String("123"),
        },
        Signature: "iste",
        TPPSignatureCertificate: "dolore",
        XRequestID: "laborum",
        AspName: "sed",
        AuthorisationID: "in",
        MultibancoPaymentType: "commodi",
        PaymentID: "quidem",
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

