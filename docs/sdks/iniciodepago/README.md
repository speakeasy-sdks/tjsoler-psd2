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
        ConsentID: tjsolerpsd2.String("VGA intrepid"),
        Digest: "vortals when",
        PSUAccept: tjsolerpsd2.String("Senior SUV"),
        PSUAcceptCharset: tjsolerpsd2.String("Sleek"),
        PSUAcceptEncoding: tjsolerpsd2.String("Consultant"),
        PSUAcceptLanguage: tjsolerpsd2.String("ASCII"),
        PSUCorporateID: tjsolerpsd2.String("indigo merrily"),
        PSUCorporateIDType: tjsolerpsd2.String("application steradian"),
        PSUDeviceID: tjsolerpsd2.String("models index panel"),
        PSUGeoLocation: tjsolerpsd2.String("international"),
        PSUHTTPMethod: tjsolerpsd2.String("architect mobile"),
        PsuID: tjsolerpsd2.String("whispered"),
        PSUIDType: tjsolerpsd2.String("chastise"),
        PSUIPAddress: "IP",
        PSUIPPort: tjsolerpsd2.String("hmph"),
        PSUUserAgent: tjsolerpsd2.String("Nitrogen productivity Recumbent"),
        RequestPeriodicPayment: shared.RequestPeriodicPayment{
            DayOfExecution: tjsolerpsd2.String("East Infrastructure since"),
            EndDate: tjsolerpsd2.String("boo Bicycle Administrator"),
            ExecutionRule: tjsolerpsd2.String("actuating Bedfordshire"),
            Frequency: tjsolerpsd2.String("Account"),
            StartDate: "card spirituality Customer",
        },
        Signature: "quia syndicate Elegant",
        TPPBrandLoggingInformation: tjsolerpsd2.String("neologism Pennsylvania Soap"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("Indiana Cheese"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("Territory North eek"),
        TPPNotificationURI: tjsolerpsd2.String("Wooden forenenst"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Northeast Applications"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("transmitter transgender SUV"),
        TPPSignatureCertificate: "Cisgender Electric",
        XRequestID: "analyzer female murder",
        Aspsp: "pixel state infrastructures",
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
        ConsentID: tjsolerpsd2.String("Orchestrator Generic online"),
        Digest: "Buckinghamshire",
        PSUAccept: tjsolerpsd2.String("North Assurance"),
        PSUAcceptCharset: tjsolerpsd2.String("fuchsia"),
        PSUAcceptEncoding: tjsolerpsd2.String("Well olive hence"),
        PSUAcceptLanguage: tjsolerpsd2.String("SCSI Oriental bold"),
        PSUCorporateID: tjsolerpsd2.String("Infrastructure solid Southeast"),
        PSUCorporateIDType: tjsolerpsd2.String("Handmade haptic scalable"),
        PSUDeviceID: tjsolerpsd2.String("Computer"),
        PSUGeoLocation: tjsolerpsd2.String("Wisconsin Austria"),
        PSUHTTPMethod: tjsolerpsd2.String("withdrawal"),
        PsuID: tjsolerpsd2.String("Southeast"),
        PSUIDType: tjsolerpsd2.String("RSS"),
        PSUIPAddress: "Applications Southwest",
        PSUIPPort: tjsolerpsd2.String("hack Northeast"),
        PSUUserAgent: tjsolerpsd2.String("Darmstadtium Handcrafted"),
        RequestBulkPayment: shared.RequestBulkPayment{
            BatchBookingPreferred: tjsolerpsd2.Bool(true),
            DebtorAccount: shared.RequestBulkPaymentDebtorAccount{},
            PaymentInformationID: tjsolerpsd2.String("Zimbabwe"),
            Payments: []shared.RequestBulkPaymentPayments{
                shared.RequestBulkPaymentPayments{},
            },
            RequestedExecutionDate: tjsolerpsd2.String("TLS mid"),
            RequestedExecutionTime: tjsolerpsd2.String("Northwest"),
        },
        Signature: "tesla female",
        TPPBrandLoggingInformation: tjsolerpsd2.String("Global Pines"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("shyly"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("orange input gold"),
        TPPNotificationURI: tjsolerpsd2.String("invoice"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Checking azure Cedi"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("Southwest"),
        TPPSignatureCertificate: "Mini",
        XRequestID: "Kenya aside",
        Aspsp: "ouch ad Berkshire",
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
        ConsentID: tjsolerpsd2.String("Jewelery bottom"),
        Digest: "Sedan mobile Screen",
        PSUAccept: tjsolerpsd2.String("Plastic"),
        PSUAcceptCharset: tjsolerpsd2.String("Southeast"),
        PSUAcceptEncoding: tjsolerpsd2.String("Mauritania North Designer"),
        PSUAcceptLanguage: tjsolerpsd2.String("pascal Female"),
        PSUCorporateID: tjsolerpsd2.String("Tricycle Maryland cruelly"),
        PSUCorporateIDType: tjsolerpsd2.String("Northeast anenst to"),
        PSUDeviceID: tjsolerpsd2.String("turquoise"),
        PSUGeoLocation: tjsolerpsd2.String("East radian"),
        PSUHTTPMethod: tjsolerpsd2.String("Grocery"),
        PsuID: tjsolerpsd2.String("Outdoors"),
        PSUIDType: tjsolerpsd2.String("worriedly circuit"),
        PSUIPAddress: "dynamic bypass Road",
        PSUIPPort: tjsolerpsd2.String("Forest"),
        PSUUserAgent: tjsolerpsd2.String("Technician engage"),
        Signature: "revolutionize Gasoline",
        SinglePayment: shared.SinglePayment{
            ChargeBearer: tjsolerpsd2.String("SLEV"),
            CreditorAccount: &shared.SinglePaymentCreditorAccount{},
            CreditorAddress: &shared.SinglePaymentCreditorAddress{},
            CreditorAgent: tjsolerpsd2.String("XXXLESMMXXX"),
            CreditorID: tjsolerpsd2.String("female array"),
            CreditorName: tjsolerpsd2.String("Nombre"),
            CreditorNameAndAddress: tjsolerpsd2.String("vice Balanced"),
            CurrencyOfTranfer: tjsolerpsd2.String("pascal Audi South"),
            DebtorAccount: &shared.SinglePaymentDebtorAccount{},
            DebtorID: tjsolerpsd2.String("Grenadines questioning"),
            DebtorName: tjsolerpsd2.String("olive"),
            EndToEndIdentification: tjsolerpsd2.String("Solutions"),
            ExchangeRateInformation: &shared.SinglePaymentExchangeRateInformation{},
            InstructedAmount: &shared.SinglePaymentInstructedAmount{},
            InstructionIdentification: tjsolerpsd2.String("Networked niches"),
            PurposeCode: tjsolerpsd2.String("Cis Cotton"),
            RemittanceInformationStructured: &shared.SinglePaymentRemittanceInformationStructured{},
            RemittanceInformationStructuredArray: &shared.SinglePaymentRemittanceInformationStructuredArray{},
            RemittanceInformationUnstructured: tjsolerpsd2.String("Informacion adicional"),
            RemittanceInformationUnstructuredArray: []string{
                "driver",
            },
            RequestedExecutionDate: tjsolerpsd2.String("curiously Hybrid Identity"),
            RequestedExecutionTime: tjsolerpsd2.String("definition clonk East"),
            ServiceLevel: tjsolerpsd2.String("cohesive"),
            UltimateDebtor: tjsolerpsd2.String("Jaguar"),
        },
        TPPBrandLoggingInformation: tjsolerpsd2.String("Architect briskly Curium"),
        TPPExplicitAuthorisationPreferred: tjsolerpsd2.Bool(false),
        TPPNokRedirectURI: tjsolerpsd2.String("yahoo sunbonnet"),
        TPPNotificationContentPreferred: tjsolerpsd2.String("monitor North"),
        TPPNotificationURI: tjsolerpsd2.String("but Aruba"),
        TPPRedirectPreferred: tjsolerpsd2.Bool(false),
        TPPRedirectURI: tjsolerpsd2.String("Philippines Keyboard parse"),
        TPPRejectionNoFundsPreferred: tjsolerpsd2.String("override"),
        TPPSignatureCertificate: "Manager",
        XRequestID: "District",
        Aspsp: "North Blues",
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

