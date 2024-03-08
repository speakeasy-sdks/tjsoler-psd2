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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                      |
| `request`                                                                                                                                                                                                                                | [operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest](../../pkg/models/operations/deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                       | The request object to use for the request.                                                                                                                                                                                               |


### Response

**[*operations.DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDResponse](../../pkg/models/operations/deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMultibancoCatalogue

Petición iniciada por el TPP para obtener el catálogo de pagos MULTIBANCO.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetMultibancoCatalogue(ctx, operations.GetMultibancoCatalogueRequest{
        Digest: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetMultibancoCatalogueRequest](../../pkg/models/operations/getmultibancocataloguerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetMultibancoCatalogueResponse](../../pkg/models/operations/getmultibancocatalogueresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID

Este mensaje es enviado por el TPP hacia el ASPSP a través del HUB para la recuperación de información del inicio de pago MULTIBANCO.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                |
| `request`                                                                                                                                                                                                                          | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDRequest](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                                         |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDResponse](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                                                            |
| `request`                                                                                                                                                                                                                                                      | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                                                                     |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsResponse](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        AuthorisationID: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                                                                                    | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                                                                                   |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus

Mensaje enviado por el TPP al ASPSP a través del Hub para solicitar el estado en el que se encuentra una iniciación de pago.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus(ctx, operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                                            |
| `request`                                                                                                                                                                                                                                      | [operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusRequest](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatusrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                             | The request object to use for the request.                                                                                                                                                                                                     |


### Response

**[*operations.GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatusResponse](../../pkg/models/operations/getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                              |
| `request`                                                                                                                                                                                                                                                        | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsRequest](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                                                                       |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsResponse](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType

Mensaje enviado por el TPP al ASPSP a través del Hub para crear un inicio de pago MULTIBANCO

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestStartMultibankPayment: shared.RequestStartMultibankPayment{
            EntityCode: tjsolerpsd2.Int64(10003),
            OperationReference: tjsolerpsd2.String("1501ab4e-6904-11ea-bc55-0242ac130003"),
            ParameterCode: tjsolerpsd2.Int64(3),
            PaymentReference: tjsolerpsd2.String("123456789"),
            PaymentType: tjsolerpsd2.Int64(1),
            PaymentTypeCode: tjsolerpsd2.String("05"),
            RequestedExecutionDate: tjsolerpsd2.String("2018-05-17T00:00:00Z"),
            TaxpayerIdentificationNumber: tjsolerpsd2.Int64(6244688226942976),
            TsuCenterCode: tjsolerpsd2.Int64(2698),
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        MultibancoPaymentType: "<value>",
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

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |
| `request`                                                                                                                                                                                                          | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeRequest](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttyperequest.md) | :heavy_check_mark:                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                         |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypeResponse](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttyperesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding

Mensaje enviado por el TPP al ASPSP a través del Hub para conocer el valor del importe a pagar en la seguridad social.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding(ctx, operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestSocialSecurityWithholding: shared.RequestSocialSecurityWithholding{
            DebtorAccount: shared.DebtorAccount{},
            Niss: tjsolerpsd2.Int64(65136589331),
            Number: 10000000003,
            PaymentDate: "2020-03-04T00:00:00Z",
            PaymentNature: 4613978645,
            PaymentUnit: 1,
            RemunerationAmount: shared.RemunerationAmount{},
            RemunerationCode: 1,
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
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

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingRequest](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholdingrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |


### Response

**[*operations.PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholdingResponse](../../pkg/models/operations/postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholdingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.ServiciosMultibanco.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID(ctx, operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestActualizarDatosPsu: shared.RequestActualizarDatosPsu{
            AuthenticationMethodID: tjsolerpsd2.String("123"),
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        AspName: "<value>",
        AuthorisationID: "<value>",
        MultibancoPaymentType: "<value>",
        PaymentID: "<value>",
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

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `request`                                                                                                                                                                                                                                                                                    | [operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest](../../pkg/models/operations/putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidrequest.md) | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                                                                                                   |


### Response

**[*operations.PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse](../../pkg/models/operations/putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
