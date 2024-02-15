# InicioDePago
(*InicioDePago*)

### Available Operations

* [InitiateRecurringPayment](#initiaterecurringpayment) - Inicio de pago Recurrente/Periódico
* [InitiationBulkPayment](#initiationbulkpayment) - Inicio de pago Bulk
* [InitiationPayment](#initiationpayment) - Inicio de pago Simple o pago Futuro

## InitiateRecurringPayment

Mensaje enviado por el TPP al ASPSP a través del Hub para crear un inicio de pago recurrente/periódico. La funcionalidad de inicios de pagos recurrentes es cubierta por la especificación de Berlin Group como la iniciación de una orden específica permanente. Un TPP puede enviar un inicio de pago recurrente donde se proporciona la fecha de inicio, frecuencia y, condicionalmnete, fecha fin. Una vez autorizado por el PSU, el pago será ejecutado por el ASPSP, si es posible, siguiendo la “orden permanente” como fue enviada por el TPP. No se necesitan acciones adicionales por parte del TPP. En este contexto, este pago es considerado un pago periódico para diferenciar el pago de otros tipos de pagos recurrentes donde terceras partes están iniciando la misma cantidad de dinero. Nota: para las órdenes permanentes de inicios de pago, el ASPSP siempre pedirá SCA con Dynamic linking. No se permiten exenciones.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiateRecurringPayment(ctx, operations.InitiateRecurringPaymentRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestPeriodicPayment: shared.RequestPeriodicPayment{
            StartDate: "<value>",
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
        PaymentProduct: operations.InitiateRecurringPaymentPathParamPaymentProductInstantSepaCreditTransfers,
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
| `request`                                                                                                    | [operations.InitiateRecurringPaymentRequest](../../pkg/models/operations/initiaterecurringpaymentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.InitiateRecurringPaymentResponse](../../pkg/models/operations/initiaterecurringpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InitiationBulkPayment

Este mensaje es enviado por el TPP hacia el HUB para la realización de inicio de pago bulk

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiationBulkPayment(ctx, operations.InitiationBulkPaymentRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        RequestBulkPayment: shared.RequestBulkPayment{
            BatchBookingPreferred: tjsolerpsd2.Bool(true),
            DebtorAccount: shared.RequestBulkPaymentDebtorAccount{},
            Payments: []shared.Payments{
                shared.Payments{},
            },
        },
        Signature: "<value>",
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
        PaymentProduct: operations.InitiationBulkPaymentPathParamPaymentProductCrossBorderCreditTransfers,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.InitiationBulkPaymentRequest](../../pkg/models/operations/initiationbulkpaymentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.InitiationBulkPaymentResponse](../../pkg/models/operations/initiationbulkpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InitiationPayment

Mensaje enviado por el TPP al ASPSP a través del HUB para crear un inicio de pago Simple.

### Example Usage

```go
package main

import(
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"context"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiationPayment(ctx, operations.InitiationPaymentRequest{
        Digest: "<value>",
        PSUIPAddress: "<value>",
        Signature: "<value>",
        SinglePayment: shared.SinglePayment{
            ChargeBearer: tjsolerpsd2.String("SLEV"),
            CreditorAgent: tjsolerpsd2.String("XXXLESMMXXX"),
            CreditorName: tjsolerpsd2.String("Nombre"),
            RemittanceInformationUnstructured: tjsolerpsd2.String("Informacion adicional"),
        },
        TPPSignatureCertificate: "<value>",
        XRequestID: "<value>",
        Aspsp: "<value>",
        PaymentProduct: operations.InitiationPaymentPathParamPaymentProductTarget2Payments,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.InitiationPaymentRequest](../../pkg/models/operations/initiationpaymentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.InitiationPaymentResponse](../../pkg/models/operations/initiationpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
