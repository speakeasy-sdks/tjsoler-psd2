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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "omnis",
        PSUAccept: tjsolerpsd2.String("tenetur"),
        PSUAcceptCharset: tjsolerpsd2.String("quasi"),
        PSUAcceptEncoding: tjsolerpsd2.String("at"),
        PSUAcceptLanguage: tjsolerpsd2.String("et"),
        PSUDeviceID: tjsolerpsd2.String("voluptate"),
        PSUGeoLocation: tjsolerpsd2.String("ipsa"),
        PSUHTTPMethod: tjsolerpsd2.String("minima"),
        PSUIPAddress: "veritatis",
        PSUIPPort: tjsolerpsd2.String("consectetur"),
        PSUUserAgent: tjsolerpsd2.String("adipisci"),
        Signature: "iste",
        TPPSignatureCertificate: "temporibus",
        XRequestID: "accusantium",
        AspName: "rem",
        MultibancoPaymentType: "aut",
        PaymentID: "laudantium",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetMultibancoCatalogue(ctx, operations.GetMultibancoCatalogueRequest{
        Digest: "eum",
        Signature: "mollitia",
        TPPSignatureCertificate: "ab",
        XRequestID: "corrupti",
        AspName: "non",
        InstructedAmount: tjsolerpsd2.String("voluptatem"),
        MultibancoPaymentType: "dolor",
        PaymentReference: tjsolerpsd2.String("occaecati"),
        RequestedExecutionDate: tjsolerpsd2.String("numquam"),
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "impedit",
        PSUAccept: tjsolerpsd2.String("explicabo"),
        PSUAcceptCharset: tjsolerpsd2.String("voluptas"),
        PSUAcceptEncoding: tjsolerpsd2.String("aut"),
        PSUAcceptLanguage: tjsolerpsd2.String("dignissimos"),
        PSUDeviceID: tjsolerpsd2.String("dicta"),
        PSUGeoLocation: tjsolerpsd2.String("maiores"),
        PSUHTTPMethod: tjsolerpsd2.String("natus"),
        PSUIPAddress: "velit",
        PSUIPPort: tjsolerpsd2.String("voluptatibus"),
        PSUUserAgent: tjsolerpsd2.String("voluptas"),
        Signature: "asperiores",
        TPPSignatureCertificate: "aperiam",
        XRequestID: "ea",
        AspName: "quaerat",
        MultibancoPaymentType: "consequuntur",
        PaymentID: "repellendus",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "officia",
        PSUAccept: tjsolerpsd2.String("maxime"),
        PSUAcceptCharset: tjsolerpsd2.String("dignissimos"),
        PSUAcceptEncoding: tjsolerpsd2.String("officia"),
        PSUAcceptLanguage: tjsolerpsd2.String("asperiores"),
        PSUCorporateID: tjsolerpsd2.String("nemo"),
        PSUCorporateIDType: tjsolerpsd2.String("quae"),
        PSUDeviceID: tjsolerpsd2.String("quaerat"),
        PSUGeoLocation: tjsolerpsd2.String("porro"),
        PSUHTTPMethod: tjsolerpsd2.String("quod"),
        PSUIPAddress: "labore",
        PSUIPPort: tjsolerpsd2.String("ab"),
        PSUUserAgent: tjsolerpsd2.String("adipisci"),
        Signature: "fuga",
        TPPSignatureCertificate: "id",
        XRequestID: "suscipit",
        AspName: "velit",
        MultibancoPaymentType: "culpa",
        PaymentID: "est",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "recusandae",
        PSUAccept: tjsolerpsd2.String("totam"),
        PSUAcceptCharset: tjsolerpsd2.String("fugiat"),
        PSUAcceptEncoding: tjsolerpsd2.String("vel"),
        PSUAcceptLanguage: tjsolerpsd2.String("ducimus"),
        PSUCorporateID: tjsolerpsd2.String("quos"),
        PSUCorporateIDType: tjsolerpsd2.String("vel"),
        PSUDeviceID: tjsolerpsd2.String("labore"),
        PSUGeoLocation: tjsolerpsd2.String("possimus"),
        PSUHTTPMethod: tjsolerpsd2.String("facilis"),
        PSUIPAddress: "cum",
        PSUIPPort: tjsolerpsd2.String("commodi"),
        PSUUserAgent: tjsolerpsd2.String("in"),
        Signature: "corporis",
        TPPSignatureCertificate: "reiciendis",
        XRequestID: "assumenda",
        AspName: "nemo",
        AuthorisationID: "recusandae",
        MultibancoPaymentType: "aliquid",
        PaymentID: "aperiam",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusRequest{
        Digest: "cum",
        PSUAccept: tjsolerpsd2.String("consectetur"),
        PSUAcceptCharset: tjsolerpsd2.String("in"),
        PSUAcceptEncoding: tjsolerpsd2.String("exercitationem"),
        PSUAcceptLanguage: tjsolerpsd2.String("earum"),
        PSUDeviceID: tjsolerpsd2.String("facere"),
        PSUGeoLocation: tjsolerpsd2.String("numquam"),
        PSUHTTPMethod: tjsolerpsd2.String("doloribus"),
        PSUIPAddress: "suscipit",
        PSUIPPort: tjsolerpsd2.String("reiciendis"),
        PSUUserAgent: tjsolerpsd2.String("quidem"),
        Signature: "saepe",
        TPPSignatureCertificate: "necessitatibus",
        XRequestID: "dolore",
        AspName: "sunt",
        MultibancoPaymentType: "asperiores",
        PaymentID: "adipisci",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "non",
        PSUAccept: tjsolerpsd2.String("amet"),
        PSUAcceptCharset: tjsolerpsd2.String("beatae"),
        PSUAcceptEncoding: tjsolerpsd2.String("dignissimos"),
        PSUAcceptLanguage: tjsolerpsd2.String("a"),
        PSUCorporateID: tjsolerpsd2.String("debitis"),
        PSUCorporateIDType: tjsolerpsd2.String("consectetur"),
        PSUDeviceID: tjsolerpsd2.String("corporis"),
        PSUGeoLocation: tjsolerpsd2.String("harum"),
        PSUHTTPMethod: tjsolerpsd2.String("laboriosam"),
        PSUIPAddress: "ipsa",
        PSUIPPort: tjsolerpsd2.String("voluptates"),
        PSUUserAgent: tjsolerpsd2.String("libero"),
        Signature: "vitae",
        TPPSignatureCertificate: "accusamus",
        XRequestID: "similique",
        AspName: "tempora",
        MultibancoPaymentType: "aspernatur",
        PaymentID: "voluptas",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeRequest{
        Digest: "voluptas",
        PSUAccept: tjsolerpsd2.String("voluptas"),
        PSUAcceptCharset: tjsolerpsd2.String("minima"),
        PSUAcceptEncoding: tjsolerpsd2.String("nobis"),
        PSUAcceptLanguage: tjsolerpsd2.String("dolorum"),
        PSUCorporateID: tjsolerpsd2.String("adipisci"),
        PSUCorporateIDType: tjsolerpsd2.String("minus"),
        PSUDeviceID: tjsolerpsd2.String("dolores"),
        PSUGeoLocation: tjsolerpsd2.String("blanditiis"),
        PSUHTTPMethod: tjsolerpsd2.String("in"),
        PsuID: tjsolerpsd2.String("dolore"),
        PSUIDType: tjsolerpsd2.String("aliquam"),
        PSUIPAddress: "officiis",
        PSUIPPort: tjsolerpsd2.String("temporibus"),
        PSUUserAgent: tjsolerpsd2.String("ullam"),
        RequestStartMultibankPayment: shared.RequestStartMultibankPayment{
            DebtorAccount: &shared.RequestStartMultibankPaymentDebtorAccount{},
            EntityCode: tjsolerpsd2.Int64(10003),
            InstructedAmount: &shared.RequestStartMultibankPaymentInstructedAmount{},
            OperationReference: tjsolerpsd2.String("1501ab4e-6904-11ea-bc55-0242ac130003"),
            ParameterCode: tjsolerpsd2.Int64(3),
            PaymentPeriod: &shared.RequestStartMultibankPaymentPaymentPeriod{},
            PaymentReference: tjsolerpsd2.String("123456789"),
            PaymentType: tjsolerpsd2.Int64(1),
            PaymentTypeCode: tjsolerpsd2.String("05"),
            RequestedExecutionDate: tjsolerpsd2.String("2018-05-17"),
            TaxpayerIdentificationNumber: tjsolerpsd2.Int64(6244688226942976),
            TsuCenterCode: tjsolerpsd2.Int64(2698),
        },
        Signature: "adipisci",
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("cum"),
        TPPRedirectPreferred: tjsolerpsd2.String("blanditiis"),
        TPPRedirectURI: tjsolerpsd2.String("quas"),
        TPPSignatureCertificate: "hic",
        XRequestID: "nesciunt",
        AspName: "culpa",
        MultibancoPaymentType: "corrupti",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingRequest{
        Digest: "pariatur",
        PSUAccept: tjsolerpsd2.String("totam"),
        PSUAcceptCharset: tjsolerpsd2.String("hic"),
        PSUAcceptEncoding: tjsolerpsd2.String("exercitationem"),
        PSUAcceptLanguage: tjsolerpsd2.String("nobis"),
        PSUCorporateID: tjsolerpsd2.String("sit"),
        PSUCorporateIDType: tjsolerpsd2.String("rerum"),
        PSUDeviceID: tjsolerpsd2.String("sed"),
        PSUGeoLocation: tjsolerpsd2.String("reiciendis"),
        PSUHTTPMethod: tjsolerpsd2.String("explicabo"),
        PsuID: tjsolerpsd2.String("asperiores"),
        PSUIDType: tjsolerpsd2.String("facilis"),
        PSUIPAddress: "voluptate",
        PSUIPPort: tjsolerpsd2.String("expedita"),
        PSUUserAgent: tjsolerpsd2.String("ab"),
        RequestSocialSecurityWithholding: shared.RequestSocialSecurityWithholding{
            DebtorAccount: shared.RequestSocialSecurityWithholdingDebtorAccount{},
            Niss: tjsolerpsd2.Int64(65136589331),
            Number: 10000000003,
            PaymentDate: "2020-03-04",
            PaymentNature: 4613978645,
            PaymentUnit: 1,
            RemunerationAmount: shared.RequestSocialSecurityWithholdingRemunerationAmount{},
            RemunerationCode: 1,
        },
        Signature: "iste",
        TPPSignatureCertificate: "dolore",
        XRequestID: "laborum",
        AspName: "sed",
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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "in",
        PSUAccept: tjsolerpsd2.String("commodi"),
        PSUAcceptCharset: tjsolerpsd2.String("quidem"),
        PSUAcceptEncoding: tjsolerpsd2.String("explicabo"),
        PSUAcceptLanguage: tjsolerpsd2.String("voluptas"),
        PSUCorporateID: tjsolerpsd2.String("unde"),
        PSUCorporateIDType: tjsolerpsd2.String("architecto"),
        PSUDeviceID: tjsolerpsd2.String("suscipit"),
        PSUGeoLocation: tjsolerpsd2.String("sapiente"),
        PSUHTTPMethod: tjsolerpsd2.String("debitis"),
        PSUIPAddress: "illo",
        PSUIPPort: tjsolerpsd2.String("reiciendis"),
        PSUUserAgent: tjsolerpsd2.String("perferendis"),
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: tjsolerpsd2.String("123"),
        },
        Signature: "corrupti",
        TPPSignatureCertificate: "maiores",
        XRequestID: "incidunt",
        AspName: "sed",
        AuthorisationID: "provident",
        MultibancoPaymentType: "eius",
        PaymentID: "necessitatibus",
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

