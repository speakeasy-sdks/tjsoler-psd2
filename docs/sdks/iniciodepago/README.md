# InicioDePago

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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiateRecurringPayment(ctx, operations.InitiateRecurringPaymentRequest{
        ConsentID: psd2cajarural.String("illum"),
        Digest: "maiores",
        PSUAccept: psd2cajarural.String("rerum"),
        PSUAcceptCharset: psd2cajarural.String("dicta"),
        PSUAcceptEncoding: psd2cajarural.String("magnam"),
        PSUAcceptLanguage: psd2cajarural.String("cumque"),
        PSUCorporateID: psd2cajarural.String("facere"),
        PSUCorporateIDType: psd2cajarural.String("ea"),
        PSUDeviceID: psd2cajarural.String("aliquid"),
        PSUGeoLocation: psd2cajarural.String("laborum"),
        PSUHTTPMethod: psd2cajarural.String("accusamus"),
        PsuID: psd2cajarural.String("non"),
        PSUIDType: psd2cajarural.String("occaecati"),
        PSUIPAddress: "enim",
        PSUIPPort: psd2cajarural.String("accusamus"),
        PSUUserAgent: psd2cajarural.String("delectus"),
        RequestPeriodicPayment: shared.RequestPeriodicPayment{
            DayOfExecution: psd2cajarural.String("quidem"),
            EndDate: psd2cajarural.String("provident"),
            ExecutionRule: psd2cajarural.String("nam"),
            Frequency: psd2cajarural.String("id"),
            StartDate: "blanditiis",
        },
        Signature: "deleniti",
        TPPBrandLoggingInformation: psd2cajarural.String("sapiente"),
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("amet"),
        TPPNotificationContentPreferred: psd2cajarural.String("deserunt"),
        TPPNotificationURI: psd2cajarural.String("nisi"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("vel"),
        TPPRejectionNoFundsPreferred: psd2cajarural.String("natus"),
        TPPSignatureCertificate: "omnis",
        XRequestID: "molestiae",
        Aspsp: "perferendis",
        PaymentProduct: operations.InitiateRecurringPaymentPaymentProductInstantSepaCreditTransfers,
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiationBulkPayment(ctx, operations.InitiationBulkPaymentRequest{
        ConsentID: psd2cajarural.String("magnam"),
        Digest: "distinctio",
        PSUAccept: psd2cajarural.String("id"),
        PSUAcceptCharset: psd2cajarural.String("labore"),
        PSUAcceptEncoding: psd2cajarural.String("labore"),
        PSUAcceptLanguage: psd2cajarural.String("suscipit"),
        PSUCorporateID: psd2cajarural.String("natus"),
        PSUCorporateIDType: psd2cajarural.String("nobis"),
        PSUDeviceID: psd2cajarural.String("eum"),
        PSUGeoLocation: psd2cajarural.String("vero"),
        PSUHTTPMethod: psd2cajarural.String("aspernatur"),
        PsuID: psd2cajarural.String("architecto"),
        PSUIDType: psd2cajarural.String("magnam"),
        PSUIPAddress: "et",
        PSUIPPort: psd2cajarural.String("excepturi"),
        PSUUserAgent: psd2cajarural.String("ullam"),
        RequestBulkPayment: shared.RequestBulkPayment{
            BatchBookingPreferred: psd2cajarural.Bool(true),
            DebtorAccount: shared.RequestBulkPaymentDebtorAccount{},
            PaymentInformationID: psd2cajarural.String("provident"),
            Payments: []shared.RequestBulkPaymentPayments{
                shared.RequestBulkPaymentPayments{},
                shared.RequestBulkPaymentPayments{},
                shared.RequestBulkPaymentPayments{},
            },
            RequestedExecutionDate: psd2cajarural.String("sint"),
            RequestedExecutionTime: psd2cajarural.String("accusantium"),
        },
        Signature: "mollitia",
        TPPBrandLoggingInformation: psd2cajarural.String("reiciendis"),
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("mollitia"),
        TPPNotificationContentPreferred: psd2cajarural.String("ad"),
        TPPNotificationURI: psd2cajarural.String("eum"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("dolor"),
        TPPRejectionNoFundsPreferred: psd2cajarural.String("necessitatibus"),
        TPPSignatureCertificate: "odit",
        XRequestID: "nemo",
        Aspsp: "quasi",
        PaymentProduct: operations.InitiationBulkPaymentPaymentProductInstantSepaCreditTransfers,
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
	"github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := psd2cajarural.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiationPayment(ctx, operations.InitiationPaymentRequest{
        ConsentID: psd2cajarural.String("doloribus"),
        Digest: "debitis",
        PSUAccept: psd2cajarural.String("eius"),
        PSUAcceptCharset: psd2cajarural.String("maxime"),
        PSUAcceptEncoding: psd2cajarural.String("deleniti"),
        PSUAcceptLanguage: psd2cajarural.String("facilis"),
        PSUCorporateID: psd2cajarural.String("in"),
        PSUCorporateIDType: psd2cajarural.String("architecto"),
        PSUDeviceID: psd2cajarural.String("architecto"),
        PSUGeoLocation: psd2cajarural.String("repudiandae"),
        PSUHTTPMethod: psd2cajarural.String("ullam"),
        PsuID: psd2cajarural.String("expedita"),
        PSUIDType: psd2cajarural.String("nihil"),
        PSUIPAddress: "repellat",
        PSUIPPort: psd2cajarural.String("quibusdam"),
        PSUUserAgent: psd2cajarural.String("sed"),
        Signature: "saepe",
        SinglePayment: shared.SinglePayment{
            ChargeBearer: psd2cajarural.String("SLEV"),
            CreditorAccount: &shared.SinglePaymentCreditorAccount{},
            CreditorAddress: &shared.SinglePaymentCreditorAddress{},
            CreditorAgent: psd2cajarural.String("XXXLESMMXXX"),
            CreditorID: psd2cajarural.String("pariatur"),
            CreditorName: psd2cajarural.String("Nombre"),
            CreditorNameAndAddress: psd2cajarural.String("accusantium"),
            CurrencyOfTranfer: psd2cajarural.String("consequuntur"),
            DebtorAccount: &shared.SinglePaymentDebtorAccount{},
            DebtorID: psd2cajarural.String("praesentium"),
            DebtorName: psd2cajarural.String("natus"),
            EndToEndIdentification: psd2cajarural.String("magni"),
            ExchangeRateInformation: &shared.SinglePaymentExchangeRateInformation{},
            InstructedAmount: &shared.SinglePaymentInstructedAmount{},
            InstructionIdentification: psd2cajarural.String("sunt"),
            PurposeCode: psd2cajarural.String("quo"),
            RemittanceInformationStructured: &shared.SinglePaymentRemittanceInformationStructured{},
            RemittanceInformationStructuredArray: &shared.SinglePaymentRemittanceInformationStructuredArray{},
            RemittanceInformationUnstructured: psd2cajarural.String("Informacion adicional"),
            RemittanceInformationUnstructuredArray: []string{
                "pariatur",
                "maxime",
                "ea",
                "excepturi",
            },
            RequestedExecutionDate: psd2cajarural.String("odit"),
            RequestedExecutionTime: psd2cajarural.String("ea"),
            ServiceLevel: psd2cajarural.String("accusantium"),
            UltimateDebtor: psd2cajarural.String("ab"),
        },
        TPPBrandLoggingInformation: psd2cajarural.String("maiores"),
        TPPExplicitAuthorisationPreferred: psd2cajarural.Bool(false),
        TPPNokRedirectURI: psd2cajarural.String("quidem"),
        TPPNotificationContentPreferred: psd2cajarural.String("ipsam"),
        TPPNotificationURI: psd2cajarural.String("voluptate"),
        TPPRedirectPreferred: psd2cajarural.Bool(false),
        TPPRedirectURI: psd2cajarural.String("autem"),
        TPPRejectionNoFundsPreferred: psd2cajarural.String("nam"),
        TPPSignatureCertificate: "eaque",
        XRequestID: "pariatur",
        Aspsp: "nemo",
        PaymentProduct: operations.InitiationPaymentPaymentProductCrossBorderCreditTransfers,
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

