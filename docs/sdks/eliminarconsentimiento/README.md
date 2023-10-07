# EliminarConsentimiento
(*EliminarConsentimiento*)

### Available Operations

* [DeleteConsentID](#deleteconsentid) - Eliminar consentimiento AIS
* [DeleteConsentsConfirmationOfFunds](#deleteconsentsconfirmationoffunds) - Eliminar consentimiento FCS

## DeleteConsentID

Este servicio permite al TPP eliminar una solicitud de consentimiento

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
    res, err := s.EliminarConsentimiento.DeleteConsentID(ctx, operations.DeleteConsentIDRequest{
        Digest: "Bicycle",
        Signature: "sans Cambridgeshire",
        TPPSignatureCertificate: "doubtfully",
        XRequestID: "Cambridgeshire",
        Aspsp: "Radon",
        ConsentID: "Bike conciliate",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteConsentIDRequest](../../models/operations/deleteconsentidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteConsentIDResponse](../../models/operations/deleteconsentidresponse.md), error**


## DeleteConsentsConfirmationOfFunds

Este servicio permite al TPP, a través del Hub, conocer el estado en el que se encuentra un recurso de consentimiento de confirmación de fondos en el ASPSP.

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
    res, err := s.EliminarConsentimiento.DeleteConsentsConfirmationOfFunds(ctx, operations.DeleteConsentsConfirmationOfFundsRequest{
        Digest: "deposit granular Money",
        Signature: "copying withdrawal farad",
        TPPSignatureCertificate: "Security Sausages Books",
        XRequestID: "coffee Concrete West",
        Aspsp: "calculating",
        ConsentID: "Gloves ratione",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteConsentsConfirmationOfFundsRequest](../../models/operations/deleteconsentsconfirmationoffundsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.DeleteConsentsConfirmationOfFundsResponse](../../models/operations/deleteconsentsconfirmationoffundsresponse.md), error**

