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
	"context"
	"log"
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiateRecurringPayment(ctx, operations.InitiateRecurringPaymentRequest{
        Digest: "OCR",
        PSUIPAddress: "hmph",
        RequestPeriodicPayment: shared.RequestPeriodicPayment{
            StartDate: "payment",
        },
        Signature: "vortals",
        TPPSignatureCertificate: "when",
        XRequestID: "national",
        Aspsp: "Berkshire",
        PaymentProduct: operations.InitiateRecurringPaymentPaymentProductCrossBorderCreditTransfers,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.InitiateRecurringPaymentRequest](../../models/operations/initiaterecurringpaymentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.InitiateRecurringPaymentResponse](../../models/operations/initiaterecurringpaymentresponse.md), error**


## InitiationBulkPayment

Este mensaje es enviado por el TPP hacia el HUB para la realización de inicio de pago bulk

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
    res, err := s.InicioDePago.InitiationBulkPayment(ctx, operations.InitiationBulkPaymentRequest{
        Digest: "rigidly",
        PSUIPAddress: "Demigender",
        RequestBulkPayment: shared.RequestBulkPayment{
            BatchBookingPreferred: tjsolerpsd2.Bool(true),
            DebtorAccount: shared.RequestBulkPaymentDebtorAccount{},
            Payments: []shared.RequestBulkPaymentPayments{
                shared.RequestBulkPaymentPayments{},
            },
        },
        Signature: "Infrastructure",
        TPPSignatureCertificate: "digital",
        XRequestID: "Buckinghamshire",
        Aspsp: "Checking",
        PaymentProduct: operations.InitiationBulkPaymentPaymentProductTarget2Payments,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.InitiationBulkPaymentRequest](../../models/operations/initiationbulkpaymentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.InitiationBulkPaymentResponse](../../models/operations/initiationbulkpaymentresponse.md), error**


## InitiationPayment

Mensaje enviado por el TPP al ASPSP a través del HUB para crear un inicio de pago Simple.

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
    res, err := s.InicioDePago.InitiationPayment(ctx, operations.InitiationPaymentRequest{
        Digest: "override",
        PSUIPAddress: "lavender",
        Signature: "Southwest",
        SinglePayment: shared.SinglePayment{
            ChargeBearer: tjsolerpsd2.String("SLEV"),
            CreditorAccount: &shared.SinglePaymentCreditorAccount{},
            CreditorAddress: &shared.SinglePaymentCreditorAddress{},
            CreditorAgent: tjsolerpsd2.String("XXXLESMMXXX"),
            CreditorName: tjsolerpsd2.String("Nombre"),
            DebtorAccount: &shared.SinglePaymentDebtorAccount{},
            ExchangeRateInformation: &shared.SinglePaymentExchangeRateInformation{},
            InstructedAmount: &shared.SinglePaymentInstructedAmount{},
            RemittanceInformationStructured: &shared.SinglePaymentRemittanceInformationStructured{},
            RemittanceInformationStructuredArray: &shared.SinglePaymentRemittanceInformationStructuredArray{},
            RemittanceInformationUnstructured: tjsolerpsd2.String("Informacion adicional"),
            RemittanceInformationUnstructuredArray: []string{
                "Sedan",
            },
        },
        TPPSignatureCertificate: "mobile",
        XRequestID: "Screen",
        Aspsp: "Concrete",
        PaymentProduct: operations.InitiationPaymentPaymentProductInstantSepaCreditTransfers,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.InitiationPaymentRequest](../../models/operations/initiationpaymentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.InitiationPaymentResponse](../../models/operations/initiationpaymentresponse.md), error**

