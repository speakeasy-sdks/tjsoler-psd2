# ServiciosMultibanco
(*ServiciosMultibanco*)

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
        Digest: "API",
        PSUAccept: tjsolerpsd2.String("Savings Kwanza female"),
        PSUAcceptCharset: tjsolerpsd2.String("male Praseodymium"),
        PSUAcceptEncoding: tjsolerpsd2.String("East Nissan"),
        PSUAcceptLanguage: tjsolerpsd2.String("meter Fantastic sievert"),
        PSUDeviceID: tjsolerpsd2.String("SMTP"),
        PSUGeoLocation: tjsolerpsd2.String("Associate circuit"),
        PSUHTTPMethod: tjsolerpsd2.String("Gambia"),
        PSUIPAddress: "Lead gosh",
        PSUIPPort: tjsolerpsd2.String("generate"),
        PSUUserAgent: tjsolerpsd2.String("exotic Concrete"),
        Signature: "Stage whistle Bacon",
        TPPSignatureCertificate: "program",
        XRequestID: "annotate male",
        AspName: "Florida Personal markets",
        MultibancoPaymentType: "Customer indigo teal",
        PaymentID: "buttonhole Gasoline SUV",
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
        Digest: "Trial",
        Signature: "Senegal Bhutan",
        TPPSignatureCertificate: "brown Gold CSS",
        XRequestID: "reprehenderit",
        AspName: "until West victoriously",
        InstructedAmount: tjsolerpsd2.String("Keyboard seamless Smart"),
        MultibancoPaymentType: "Solutions",
        PaymentReference: tjsolerpsd2.String("Rubber"),
        RequestedExecutionDate: tjsolerpsd2.String("wireless spectacle"),
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
        Digest: "Hagenes",
        PSUAccept: tjsolerpsd2.String("molder Interactions"),
        PSUAcceptCharset: tjsolerpsd2.String("Bugatti"),
        PSUAcceptEncoding: tjsolerpsd2.String("Account"),
        PSUAcceptLanguage: tjsolerpsd2.String("Outdoors"),
        PSUDeviceID: tjsolerpsd2.String("Extended Antimony Northwest"),
        PSUGeoLocation: tjsolerpsd2.String("capacitor"),
        PSUHTTPMethod: tjsolerpsd2.String("JSON"),
        PSUIPAddress: "SMTP",
        PSUIPPort: tjsolerpsd2.String("volt Health"),
        PSUUserAgent: tjsolerpsd2.String("Northeast"),
        Signature: "Classical",
        TPPSignatureCertificate: "Security Customer Credit",
        XRequestID: "calculating radiant man",
        AspName: "Hatchback",
        MultibancoPaymentType: "East asymmetric",
        PaymentID: "Dollar",
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
        Digest: "minima Rubber",
        PSUAccept: tjsolerpsd2.String("Direct earum"),
        PSUAcceptCharset: tjsolerpsd2.String("redundant"),
        PSUAcceptEncoding: tjsolerpsd2.String("base"),
        PSUAcceptLanguage: tjsolerpsd2.String("online Southwest"),
        PSUCorporateID: tjsolerpsd2.String("Toys Northeast static"),
        PSUCorporateIDType: tjsolerpsd2.String("Pop transmitting"),
        PSUDeviceID: tjsolerpsd2.String("candela"),
        PSUGeoLocation: tjsolerpsd2.String("Latin"),
        PSUHTTPMethod: tjsolerpsd2.String("Fitness Samarium"),
        PSUIPAddress: "architectures Refined Mouse",
        PSUIPPort: tjsolerpsd2.String("crisp actuating Liechtenstein"),
        PSUUserAgent: tjsolerpsd2.String("Unbranded against"),
        Signature: "Applications",
        TPPSignatureCertificate: "Diesel Fresh Fiat",
        XRequestID: "placeat sensor Rock",
        AspName: "TLS",
        MultibancoPaymentType: "anti",
        PaymentID: "Man Oklahoma",
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
        Digest: "Habra female",
        PSUAccept: tjsolerpsd2.String("Hybrid Investor Bronze"),
        PSUAcceptCharset: tjsolerpsd2.String("infrastructures"),
        PSUAcceptEncoding: tjsolerpsd2.String("Hybrid transmitting"),
        PSUAcceptLanguage: tjsolerpsd2.String("synthesizing quizzical"),
        PSUCorporateID: tjsolerpsd2.String("slight Auto VGA"),
        PSUCorporateIDType: tjsolerpsd2.String("afore cyan"),
        PSUDeviceID: tjsolerpsd2.String("Polygender navigating"),
        PSUGeoLocation: tjsolerpsd2.String("Baby Borders Northeast"),
        PSUHTTPMethod: tjsolerpsd2.String("delude Implementation Leannon"),
        PSUIPAddress: "Ridge",
        PSUIPPort: tjsolerpsd2.String("fuchsia Games"),
        PSUUserAgent: tjsolerpsd2.String("Avon internet"),
        Signature: "Adventure female",
        TPPSignatureCertificate: "East",
        XRequestID: "ability Lebanese Steel",
        AspName: "World",
        AuthorisationID: "hacking Cab East",
        MultibancoPaymentType: "World Buckinghamshire Account",
        PaymentID: "firewall",
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
        Digest: "Ball car skulk",
        PSUAccept: tjsolerpsd2.String("Digitized Marketing"),
        PSUAcceptCharset: tjsolerpsd2.String("networks and"),
        PSUAcceptEncoding: tjsolerpsd2.String("than female sans"),
        PSUAcceptLanguage: tjsolerpsd2.String("Wagon array navigate"),
        PSUDeviceID: tjsolerpsd2.String("hack"),
        PSUGeoLocation: tjsolerpsd2.String("voluptatum Robust Borders"),
        PSUHTTPMethod: tjsolerpsd2.String("opportunist silver"),
        PSUIPAddress: "hertz Frozen",
        PSUIPPort: tjsolerpsd2.String("Hawaii connecting"),
        PSUUserAgent: tjsolerpsd2.String("Checking"),
        Signature: "Missouri Cruz",
        TPPSignatureCertificate: "tongue tangible",
        XRequestID: "accidentally Account",
        AspName: "Interactions Liechtenstein navigate",
        MultibancoPaymentType: "Cambridgeshire but amusing",
        PaymentID: "Account Director disable",
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
        Digest: "SSD transform",
        PSUAccept: tjsolerpsd2.String("teal"),
        PSUAcceptCharset: tjsolerpsd2.String("FTP uncomfortable"),
        PSUAcceptEncoding: tjsolerpsd2.String("Virginia"),
        PSUAcceptLanguage: tjsolerpsd2.String("jovially unleash Lanka"),
        PSUCorporateID: tjsolerpsd2.String("smart Soft Electric"),
        PSUCorporateIDType: tjsolerpsd2.String("Krona transmitting"),
        PSUDeviceID: tjsolerpsd2.String("Metal Volvo ick"),
        PSUGeoLocation: tjsolerpsd2.String("Unbranded huddle"),
        PSUHTTPMethod: tjsolerpsd2.String("uniform Electric"),
        PSUIPAddress: "Skyway painfully",
        PSUIPPort: tjsolerpsd2.String("Mazda Southeast"),
        PSUUserAgent: tjsolerpsd2.String("Intranet Licensed"),
        Signature: "Shoes Balanced",
        TPPSignatureCertificate: "Hybrid",
        XRequestID: "redundant",
        AspName: "Computer",
        MultibancoPaymentType: "blanch Fluorine",
        PaymentID: "PNG",
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
        Digest: "Avon Tuna",
        PSUAccept: tjsolerpsd2.String("haptic mint Gasoline"),
        PSUAcceptCharset: tjsolerpsd2.String("Directives"),
        PSUAcceptEncoding: tjsolerpsd2.String("Transexual pink"),
        PSUAcceptLanguage: tjsolerpsd2.String("newton because Electric"),
        PSUCorporateID: tjsolerpsd2.String("navigate olive"),
        PSUCorporateIDType: tjsolerpsd2.String("Ohio elicit"),
        PSUDeviceID: tjsolerpsd2.String("Pants Executive"),
        PSUGeoLocation: tjsolerpsd2.String("customized ivory Rand"),
        PSUHTTPMethod: tjsolerpsd2.String("why"),
        PsuID: tjsolerpsd2.String("brr"),
        PSUIDType: tjsolerpsd2.String("3rd transgender"),
        PSUIPAddress: "product Jeep plum",
        PSUIPPort: tjsolerpsd2.String("Investment Table"),
        PSUUserAgent: tjsolerpsd2.String("Northwest altar Micronesia"),
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
        Signature: "SUV",
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("Iranian Coupe"),
        TPPRedirectPreferred: tjsolerpsd2.String("Male grey"),
        TPPRedirectURI: tjsolerpsd2.String("Hat"),
        TPPSignatureCertificate: "Southeast so Account",
        XRequestID: "Future leafy Chair",
        AspName: "Sports oof",
        MultibancoPaymentType: "purple",
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
        Digest: "Borders Hybrid compressing",
        PSUAccept: tjsolerpsd2.String("Factors Shoes degree"),
        PSUAcceptCharset: tjsolerpsd2.String("Personal"),
        PSUAcceptEncoding: tjsolerpsd2.String("Cab JBOD"),
        PSUAcceptLanguage: tjsolerpsd2.String("woman embrace Kazakhstan"),
        PSUCorporateID: tjsolerpsd2.String("Gasoline Aston Modern"),
        PSUCorporateIDType: tjsolerpsd2.String("capacitor gratefully"),
        PSUDeviceID: tjsolerpsd2.String("Soft"),
        PSUGeoLocation: tjsolerpsd2.String("Minivan yuck"),
        PSUHTTPMethod: tjsolerpsd2.String("Yttrium"),
        PsuID: tjsolerpsd2.String("Loan turquoise"),
        PSUIDType: tjsolerpsd2.String("Hatchback viciously"),
        PSUIPAddress: "Aluminium USB Genderqueer",
        PSUIPPort: tjsolerpsd2.String("Southeast lens Vista"),
        PSUUserAgent: tjsolerpsd2.String("Buckinghamshire regional withdrawal"),
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
        Signature: "violet",
        TPPSignatureCertificate: "system East",
        XRequestID: "Integration",
        AspName: "connecting",
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
        Digest: "Car",
        PSUAccept: tjsolerpsd2.String("Money Woman"),
        PSUAcceptCharset: tjsolerpsd2.String("Central"),
        PSUAcceptEncoding: tjsolerpsd2.String("networks West shy"),
        PSUAcceptLanguage: tjsolerpsd2.String("mobile reboot"),
        PSUCorporateID: tjsolerpsd2.String("Bronze"),
        PSUCorporateIDType: tjsolerpsd2.String("Jazz indeed Exclusive"),
        PSUDeviceID: tjsolerpsd2.String("hm Cis Forward"),
        PSUGeoLocation: tjsolerpsd2.String("Transexual Money"),
        PSUHTTPMethod: tjsolerpsd2.String("Diesel"),
        PSUIPAddress: "er District",
        PSUIPPort: tjsolerpsd2.String("Internal"),
        PSUUserAgent: tjsolerpsd2.String("haptic"),
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: tjsolerpsd2.String("123"),
        },
        Signature: "Northwest",
        TPPSignatureCertificate: "Fresh",
        XRequestID: "female commodi sievert",
        AspName: "Leu programming",
        AuthorisationID: "Regional Northwest",
        MultibancoPaymentType: "project",
        PaymentID: "Checking SUV eek",
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

