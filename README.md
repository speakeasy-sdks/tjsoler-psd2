# github.com/speakeasy-sdks/tjsoler-psd2

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    
    <a href="https://https://github.com/speakeasy-sdks/tjsoler-psd2.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tjsoler-psd2
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
    res, err := s.ActualizarDatosDelPSUSeleccionarMetodoSCA.PutSeleccionarSCAAutorizacionCancelacionPago(ctx, operations.PutSeleccionarSCAAutorizacionCancelacionPagoRequest{
        Digest: "driver Metal",
        PSUIPAddress: "Recycled Northwest",
        Signature: "Connelly",
        TPPSignatureCertificate: "Sleek",
        XRequestID: "grow sunt",
        Aspsp: operations.PutSeleccionarSCAAutorizacionCancelacionPagoAspspBff,
        AuthorisationID: "bus Dakota",
        PaymentID: "outside",
        PaymentProduct: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentProductTarget2Payments,
        PaymentService: operations.PutSeleccionarSCAAutorizacionCancelacionPagoPaymentServicePeriodicPayments,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [ActualizarDatosDelPSUSeleccionarMetodoSCA](docs/sdks/actualizardatosdelpsuseleccionarmetodosca/README.md)

* [PutSeleccionarSCAAutorizacionCancelacionPago](docs/sdks/actualizardatosdelpsuseleccionarmetodosca/README.md#putseleccionarscaautorizacioncancelacionpago) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Cancelación de Pago
* [PutSeleccionarSCAAutorizacionConsentimientosAIS](docs/sdks/actualizardatosdelpsuseleccionarmetodosca/README.md#putseleccionarscaautorizacionconsentimientosais) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos AIS
* [PutSeleccionarSCAAutorizacionConsentimientosFCS](docs/sdks/actualizardatosdelpsuseleccionarmetodosca/README.md#putseleccionarscaautorizacionconsentimientosfcs) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Consentimientos FCS
* [PutSeleccionarSCAAutorizacionInicioPago](docs/sdks/actualizardatosdelpsuseleccionarmetodosca/README.md#putseleccionarscaautorizacioniniciopago) - Endpoint en caso de actualizacion de datos PSU (seleccionar método SCA) - Autorización de Inicio de Pago

### [CancelarInicioDePago](docs/sdks/cancelariniciodepago/README.md)

* [DeletePayment](docs/sdks/cancelariniciodepago/README.md#deletepayment) - Cancelar Inicio de pago

### [ConsentimientoDeConfirmacionDeFondosFCS](docs/sdks/consentimientodeconfirmaciondefondosfcs/README.md)

* [PostConsentsConfirmationOfFunds](docs/sdks/consentimientodeconfirmaciondefondosfcs/README.md#postconsentsconfirmationoffunds) - Solicitud de consentimiento FCS

### [ConsentimientoDeInformacionSobreCuentasDePagoAIS](docs/sdks/consentimientodeinformacionsobrecuentasdepagoais/README.md)

* [PostConsents](docs/sdks/consentimientodeinformacionsobrecuentasdepagoais/README.md#postconsents) - Solicitud de consentimiento AIS

### [ConsultaDeFondos](docs/sdks/consultadefondos/README.md)

* [FundsConfirmation](docs/sdks/consultadefondos/README.md#fundsconfirmation) - Consulta de fondos

### [EliminarConsentimiento](docs/sdks/eliminarconsentimiento/README.md)

* [DeleteConsentID](docs/sdks/eliminarconsentimiento/README.md#deleteconsentid) - Eliminar consentimiento AIS
* [DeleteConsentsConfirmationOfFunds](docs/sdks/eliminarconsentimiento/README.md#deleteconsentsconfirmationoffunds) - Eliminar consentimiento FCS

### [InicioDePago](docs/sdks/iniciodepago/README.md)

* [InitiateRecurringPayment](docs/sdks/iniciodepago/README.md#initiaterecurringpayment) - Inicio de pago Recurrente/Periódico
* [InitiationBulkPayment](docs/sdks/iniciodepago/README.md#initiationbulkpayment) - Inicio de pago Bulk
* [InitiationPayment](docs/sdks/iniciodepago/README.md#initiationpayment) - Inicio de pago Simple o pago Futuro

### [InicioDelProcesoDeAutorizacionExplicita](docs/sdks/iniciodelprocesodeautorizacionexplicita/README.md)

* [PostAutorizacionCancelacionPago](docs/sdks/iniciodelprocesodeautorizacionexplicita/README.md#postautorizacioncancelacionpago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Cancelación de Pago
* [PostAutorizacionConsentimientosAIS](docs/sdks/iniciodelprocesodeautorizacionexplicita/README.md#postautorizacionconsentimientosais) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos AIS
* [PostAutorizacionConsentimientosFCS](docs/sdks/iniciodelprocesodeautorizacionexplicita/README.md#postautorizacionconsentimientosfcs) - Endpoint en caso de Inicio del proceso de Autorización explícita para Consentimientos FCS
* [PostAutorizacionInicioPago](docs/sdks/iniciodelprocesodeautorizacionexplicita/README.md#postautorizacioniniciopago) - Endpoint en caso de Inicio del proceso de Autorización explícita para Inicio de Pago

### [LecturaDeBalances](docs/sdks/lecturadebalances/README.md)

* [GetAccountBalances](docs/sdks/lecturadebalances/README.md#getaccountbalances) - Lectura de balances de una cuenta

### [LecturaDeDetallesDeCuenta](docs/sdks/lecturadedetallesdecuenta/README.md)

* [GetAccountIdv11](docs/sdks/lecturadedetallesdecuenta/README.md#getaccountidv11) - Lectura de detalles de una cuenta

### [LecturaDeListadoDeCuentas](docs/sdks/lecturadelistadodecuentas/README.md)

* [GetAccountListv11](docs/sdks/lecturadelistadodecuentas/README.md#getaccountlistv11) - Lectura de listado de cuentas

### [LecturaDeTransacciones](docs/sdks/lecturadetransacciones/README.md)

* [AccountsTrasactions](docs/sdks/lecturadetransacciones/README.md#accountstrasactions) - Lectura de transacciones de una cuenta

### [ObtenerEstadoSCA](docs/sdks/obtenerestadosca/README.md)

* [GetEstadoSCAAutorizacionCancelacionPago](docs/sdks/obtenerestadosca/README.md#getestadoscaautorizacioncancelacionpago) - Endpoint en caso de obtener el Estado SCA para Autorización de Cancelación de Pago
* [GetEstadoSCAAutorizacionConsentimientosAIS](docs/sdks/obtenerestadosca/README.md#getestadoscaautorizacionconsentimientosais) - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS
* [GetEstadoSCAAutorizacionConsentimientosFCS](docs/sdks/obtenerestadosca/README.md#getestadoscaautorizacionconsentimientosfcs) - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS
* [GetEstadoSCAAutorizacionInicioPago](docs/sdks/obtenerestadosca/README.md#getestadoscaautorizacioniniciopago) - Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago

### [ObtenerEstadoDeConsentimiento](docs/sdks/obtenerestadodeconsentimiento/README.md)

* [GetConsentStatus](docs/sdks/obtenerestadodeconsentimiento/README.md#getconsentstatus) - Estado de consentimiento AIS
* [GetConsentsConfirmationOfFunds](docs/sdks/obtenerestadodeconsentimiento/README.md#getconsentsconfirmationoffunds) - Estado de consentimiento FCS

### [ObtenerEstadoDelPago](docs/sdks/obtenerestadodelpago/README.md)

* [StatusPayment](docs/sdks/obtenerestadodelpago/README.md#statuspayment) - Obtener información del Estado de pago

### [ObtenerListadoDeBeneficiariosDeConfianza](docs/sdks/obtenerlistadodebeneficiariosdeconfianza/README.md)

* [GetTrustedBeneficiaries](docs/sdks/obtenerlistadodebeneficiariosdeconfianza/README.md#gettrustedbeneficiaries) - Obtener listado de beneficiarios de confianza

### [ObtenerSubRecursosDeLaAutorizacion](docs/sdks/obtenersubrecursosdelaautorizacion/README.md)

* [GetSubRecursosAutorizacionCancelacionPago](docs/sdks/obtenersubrecursosdelaautorizacion/README.md#getsubrecursosautorizacioncancelacionpago) - Endpoint en caso de obtener sub-recursos de la Autorización para Cancelación de Pago
* [GetSubRecursosAutorizacionConsentimientosAIS](docs/sdks/obtenersubrecursosdelaautorizacion/README.md#getsubrecursosautorizacionconsentimientosais) - Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos AIS
* [GetSubRecursosAutorizacionConsentimientosFCS](docs/sdks/obtenersubrecursosdelaautorizacion/README.md#getsubrecursosautorizacionconsentimientosfcs) - Endpoint en caso de obtener sub-recursos de la Autorización para Consentimientos FCS
* [GetSubRecursosAutorizacionInicioPago](docs/sdks/obtenersubrecursosdelaautorizacion/README.md#getsubrecursosautorizacioniniciopago) - Endpoint en caso de obtener sub-recursos de la Autorización para Inicio de Pago

### [RecuperarInformacionDeConsentimiento](docs/sdks/recuperarinformaciondeconsentimiento/README.md)

* [GetConsentIDDetails](docs/sdks/recuperarinformaciondeconsentimiento/README.md#getconsentiddetails) - Recuperar información de consentimiento AIS
* [GetConsentsConfirmationOfFundsInfo](docs/sdks/recuperarinformaciondeconsentimiento/README.md#getconsentsconfirmationoffundsinfo) - Recuperar información de consentimiento FCS

### [RecuperarInformacionDelInicioDePago](docs/sdks/recuperarinformaciondeliniciodepago/README.md)

* [GetInfoPayment](docs/sdks/recuperarinformaciondeliniciodepago/README.md#getinfopayment) - Recuperar información del Inicio de pago

### [ServiciosMultibanco](docs/sdks/serviciosmultibanco/README.md)

* [DeleteAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID](docs/sdks/serviciosmultibanco/README.md#deleteapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentid) - Esta petición permite iniciar la cancelación de un pago. Dependiendo del servicio de pago, el producto de pago y la implementación del ASPSP, esta petición podría ser suficiente para cancelar el pago o podría ser necesario una autorización. Si una autorización de la cancelación de pago es necesaria por el ASPSP, el link correspondiente será contenido en el mensaje de respuesta
* [GetMultibancoCatalogue](docs/sdks/serviciosmultibanco/README.md#getmultibancocatalogue) - Petición iniciada por el TPP para obtener el catálogo de pagos MULTIBANCO
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentID](docs/sdks/serviciosmultibanco/README.md#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentid) - Este mensaje es enviado por el TPP hacia el ASPSP a través del HUB para la recuperación de información del inicio de pago MULTIBANCO.
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations](docs/sdks/serviciosmultibanco/README.md#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisations)
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID](docs/sdks/serviciosmultibanco/README.md#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationid)
* [GetAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDStatus](docs/sdks/serviciosmultibanco/README.md#getapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidstatus) - Mensaje enviado por el TPP al ASPSP a través del Hub para solicitar el estado en el que se encuentra una iniciación de pago.
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisations](docs/sdks/serviciosmultibanco/README.md#postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisations)
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentType](docs/sdks/serviciosmultibanco/README.md#postapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttype) - Mensaje enviado por el TPP al ASPSP a través del Hub para crear un inicio de pago MULTIBANCO
* [PostAPIEntradaXs2aServicesAspNameV11MultibancoSocialSecurityWithholding](docs/sdks/serviciosmultibanco/README.md#postapientradaxs2aservicesaspnamev11multibancosocialsecuritywithholding) - Mensaje enviado por el TPP al ASPSP a través del Hub para conocer el valor del importe a pagar en la seguridad social
* [PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationID](docs/sdks/serviciosmultibanco/README.md#putapientradaxs2aservicesaspnamev11multibancopaymentsmultibancopaymenttypepaymentidauthorisationsauthorisationid)
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
