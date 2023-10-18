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
        Digest: "Kiribati",
        PSUIPAddress: "Southwest",
        Signature: "Savings",
        TPPSignatureCertificate: "Kwanza",
        XRequestID: "female",
        AspName: "monitor",
        MultibancoPaymentType: "lime",
        PaymentID: "haptic",
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
        Digest: "Island",
        Signature: "sensor",
        TPPSignatureCertificate: "Senegal",
        XRequestID: "Bhutan",
        AspName: "cyan",
        MultibancoPaymentType: "West",
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
        Digest: "purple",
        PSUIPAddress: "World",
        Signature: "panel",
        TPPSignatureCertificate: "Xenogender",
        XRequestID: "calculate",
        AspName: "although",
        MultibancoPaymentType: "Bugatti",
        PaymentID: "Iraq",
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
        Digest: "Home",
        PSUIPAddress: "Rubber",
        Signature: "animi",
        TPPSignatureCertificate: "Hampshire",
        XRequestID: "Health",
        AspName: "Savings",
        MultibancoPaymentType: "earum",
        PaymentID: "primary",
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
        Digest: "Ameliorated",
        PSUIPAddress: "PNG",
        Signature: "female",
        TPPSignatureCertificate: "sip",
        XRequestID: "though",
        AspName: "Functionality",
        AuthorisationID: "Court",
        MultibancoPaymentType: "Hybrid",
        PaymentID: "transmitting",
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
        Digest: "Diesel",
        PSUIPAddress: "Frozen",
        Signature: "Wyoming",
        TPPSignatureCertificate: "Electric",
        XRequestID: "Digitized",
        AspName: "Marketing",
        MultibancoPaymentType: "withdrawal",
        PaymentID: "Hybrid",
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
        Digest: "Funk",
        PSUIPAddress: "newton",
        Signature: "Northeast",
        TPPSignatureCertificate: "teal",
        XRequestID: "JSON",
        AspName: "visionary",
        MultibancoPaymentType: "adventurously",
        PaymentID: "Virginia",
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
        Digest: "Taka",
        PSUIPAddress: "Bedfordshire",
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
        Signature: "International",
        TPPSignatureCertificate: "haptic",
        XRequestID: "mint",
        AspName: "Gasoline",
        MultibancoPaymentType: "Tuna",
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
        Digest: "Neon",
        PSUIPAddress: "transmitting",
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
        Signature: "liar",
        TPPSignatureCertificate: "Factors",
        XRequestID: "Shoes",
        AspName: "degree",
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
        Digest: "Claire",
        PSUIPAddress: "Teagan",
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: tjsolerpsd2.String("123"),
        },
        Signature: "partnerships",
        TPPSignatureCertificate: "Woman",
        XRequestID: "Buckinghamshire",
        AspName: "matrix",
        AuthorisationID: "female",
        MultibancoPaymentType: "Pickup",
        PaymentID: "West",
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

