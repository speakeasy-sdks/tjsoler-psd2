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
	tjsolerpsd2 "github.com/speakeasy-sdks/tjsoler-psd2"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
)

func main() {
    s := tjsolerpsd2.New()

    ctx := context.Background()
    res, err := s.InicioDePago.InitiateRecurringPayment(ctx, operations.InitiateRecurringPaymentRequest{
        ConsentID: tjsolerpsd2.String("sapiente"),
        Digest: "amet",
        PSUAccept: tjsolerpsd2.String("deserunt"),
        PSUAcceptCharset: tjsolerpsd2.String("nisi"),
        PSUAcceptEncoding: tjsolerpsd2.String("vel"),
        PSUAcceptLanguage: tjsolerpsd2.String("natus"),
        PSUCorporateID: tjsolerpsd2.String("omnis"),
        PSUCorporateIDType: tjsolerpsd2.String("molestiae"),
        PSUDeviceID: tjsolerpsd2.String("perferendis"),
        PSUGeoLocation: tjsolerpsd2.String("nihil"),
        PSUHTTPMethod: tjsolerpsd2.String("magnam"),
        PsuID: tjsolerpsd2.String("distinctio"),
        PSUIDType: tjsolerpsd2.String("id"),
        PSUIPAddress: "labore",
        PSUIPPort: tjsolerpsd2.String("labore"),
        PSUUserAgent: tjsolerpsd2.String("suscipit"),
        RequestPeriodicPayment: shared.RequestPeriodicPayment{
            DayOfExecution: tjsolerpsd2.String("natus"),
            EndDate: tjsolerpsd2.String("nobis"),
            ExecutionRule: tjsolerpsd2.String("eum"),
            Frequency: tjsolerpsd2.String("vero"),
            StartDate: "aspernatur",
        },
        Signature: "architecto",
        TPPBrandLoggingInformation: tjsolerpsd2.String("magnam"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("et"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("excepturi"),
        TPPNotificationURI: tjsolerpsd2.String("ullam"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("provident"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("quos"),
        TPPSignatureCertificate: "sint",
        XRequestID: "accusantium",
        Aspsp: "mollitia",
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
        ConsentID: tjsolerpsd2.String("mollitia"),
        Digest: "ad",
        PSUAccept: tjsolerpsd2.String("eum"),
        PSUAcceptCharset: tjsolerpsd2.String("dolor"),
        PSUAcceptEncoding: tjsolerpsd2.String("necessitatibus"),
        PSUAcceptLanguage: tjsolerpsd2.String("odit"),
        PSUCorporateID: tjsolerpsd2.String("nemo"),
        PSUCorporateIDType: tjsolerpsd2.String("quasi"),
        PSUDeviceID: tjsolerpsd2.String("iure"),
        PSUGeoLocation: tjsolerpsd2.String("doloribus"),
        PSUHTTPMethod: tjsolerpsd2.String("debitis"),
        PsuID: tjsolerpsd2.String("eius"),
        PSUIDType: tjsolerpsd2.String("maxime"),
        PSUIPAddress: "deleniti",
        PSUIPPort: tjsolerpsd2.String("facilis"),
        PSUUserAgent: tjsolerpsd2.String("in"),
        RequestBulkPayment: shared.RequestBulkPayment{
            BatchBookingPreferred: tjsolerpsd2.Bool(true),
            DebtorAccount: shared.RequestBulkPaymentDebtorAccount{},
            PaymentInformationID: tjsolerpsd2.String("architecto"),
            Payments: []shared.RequestBulkPaymentPayments{
                shared.RequestBulkPaymentPayments{},
            },
            RequestedExecutionDate: tjsolerpsd2.String("architecto"),
            RequestedExecutionTime: tjsolerpsd2.String("repudiandae"),
        },
        Signature: "ullam",
        TPPBrandLoggingInformation: tjsolerpsd2.String("expedita"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("nihil"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("repellat"),
        TPPNotificationURI: tjsolerpsd2.String("quibusdam"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("sed"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("saepe"),
        TPPSignatureCertificate: "pariatur",
        XRequestID: "accusantium",
        Aspsp: "consequuntur",
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
        ConsentID: tjsolerpsd2.String("natus"),
        Digest: "magni",
        PSUAccept: tjsolerpsd2.String("sunt"),
        PSUAcceptCharset: tjsolerpsd2.String("quo"),
        PSUAcceptEncoding: tjsolerpsd2.String("illum"),
        PSUAcceptLanguage: tjsolerpsd2.String("pariatur"),
        PSUCorporateID: tjsolerpsd2.String("maxime"),
        PSUCorporateIDType: tjsolerpsd2.String("ea"),
        PSUDeviceID: tjsolerpsd2.String("excepturi"),
        PSUGeoLocation: tjsolerpsd2.String("odit"),
        PSUHTTPMethod: tjsolerpsd2.String("ea"),
        PsuID: tjsolerpsd2.String("accusantium"),
        PSUIDType: tjsolerpsd2.String("ab"),
        PSUIPAddress: "maiores",
        PSUIPPort: tjsolerpsd2.String("quidem"),
        PSUUserAgent: tjsolerpsd2.String("ipsam"),
        Signature: "voluptate",
        SinglePayment: shared.SinglePayment{
            ChargeBearer: tjsolerpsd2.String("SLEV"),
            CreditorAccount: &shared.SinglePaymentCreditorAccount{},
            CreditorAddress: &shared.SinglePaymentCreditorAddress{},
            CreditorAgent: tjsolerpsd2.String("XXXLESMMXXX"),
            CreditorID: tjsolerpsd2.String("autem"),
            CreditorName: tjsolerpsd2.String("Nombre"),
            CreditorNameAndAddress: tjsolerpsd2.String("nam"),
            CurrencyOfTranfer: tjsolerpsd2.String("eaque"),
            DebtorAccount: &shared.SinglePaymentDebtorAccount{},
            DebtorID: tjsolerpsd2.String("pariatur"),
            DebtorName: tjsolerpsd2.String("nemo"),
            EndToEndIdentification: tjsolerpsd2.String("voluptatibus"),
            ExchangeRateInformation: &shared.SinglePaymentExchangeRateInformation{},
            InstructedAmount: &shared.SinglePaymentInstructedAmount{},
            InstructionIdentification: tjsolerpsd2.String("perferendis"),
            PurposeCode: tjsolerpsd2.String("fugiat"),
            RemittanceInformationStructured: &shared.SinglePaymentRemittanceInformationStructured{},
            RemittanceInformationStructuredArray: &shared.SinglePaymentRemittanceInformationStructuredArray{},
            RemittanceInformationUnstructured: tjsolerpsd2.String("Informacion adicional"),
            RemittanceInformationUnstructuredArray: []string{
                "amet",
            },
            RequestedExecutionDate: tjsolerpsd2.String("aut"),
            RequestedExecutionTime: tjsolerpsd2.String("cumque"),
            ServiceLevel: tjsolerpsd2.String("corporis"),
            UltimateDebtor: tjsolerpsd2.String("hic"),
        },
        TPPBrandLoggingInformation: tjsolerpsd2.String("libero"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("nobis"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("dolores"),
        TPPNotificationURI: tjsolerpsd2.String("quis"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("totam"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("dignissimos"),
        TPPSignatureCertificate: "eaque",
        XRequestID: "quis",
        Aspsp: "nesciunt",
        PaymentProduct: operations.InitiationPaymentPaymentProductSepaCreditTransfers,
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

