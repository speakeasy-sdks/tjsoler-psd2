// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"net/http"
)

// GetEstadoSCAAutorizacionCancelacionPagoAspsp - Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
type GetEstadoSCAAutorizacionCancelacionPagoAspsp string

const (
	GetEstadoSCAAutorizacionCancelacionPagoAspspRedsys          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "redsys"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBbva            GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BBVA"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBbvapt          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BBVAPT"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBbvabe          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BBVABE"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBbvafr          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BBVAFR"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBbvaeuk         GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BBVAEUK"
	GetEstadoSCAAutorizacionCancelacionPagoAspspCaixabank       GetEstadoSCAAutorizacionCancelacionPagoAspsp = "caixabank"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBancSabadell    GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BancSabadell"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBancosantander  GetEstadoSCAAutorizacionCancelacionPagoAspsp = "bancosantander"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBancamarch      GetEstadoSCAAutorizacionCancelacionPagoAspsp = "bancamarch"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBankoa          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "bankoa"
	GetEstadoSCAAutorizacionCancelacionPagoAspspCajamar         GetEstadoSCAAutorizacionCancelacionPagoAspsp = "cajamar"
	GetEstadoSCAAutorizacionCancelacionPagoAspspArquia          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "arquia"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBff             GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BFF"
	GetEstadoSCAAutorizacionCancelacionPagoAspspColonya         GetEstadoSCAAutorizacionCancelacionPagoAspsp = "colonya"
	GetEstadoSCAAutorizacionCancelacionPagoAspspEurocajarural   GetEstadoSCAAutorizacionCancelacionPagoAspsp = "eurocajarural"
	GetEstadoSCAAutorizacionCancelacionPagoAspspEvobanco        GetEstadoSCAAutorizacionCancelacionPagoAspsp = "evobanco"
	GetEstadoSCAAutorizacionCancelacionPagoAspspFiarebancaetica GetEstadoSCAAutorizacionCancelacionPagoAspsp = "fiarebancaetica"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBancopichincha  GetEstadoSCAAutorizacionCancelacionPagoAspsp = "bancopichincha"
	GetEstadoSCAAutorizacionCancelacionPagoAspspUnicajabanco    GetEstadoSCAAutorizacionCancelacionPagoAspsp = "unicajabanco"
	GetEstadoSCAAutorizacionCancelacionPagoAspspCajasur         GetEstadoSCAAutorizacionCancelacionPagoAspsp = "cajasur"
	GetEstadoSCAAutorizacionCancelacionPagoAspspKutxabank       GetEstadoSCAAutorizacionCancelacionPagoAspsp = "kutxabank"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBankinter       GetEstadoSCAAutorizacionCancelacionPagoAspsp = "bankinter"
	GetEstadoSCAAutorizacionCancelacionPagoAspspRenta4          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "renta4"
	GetEstadoSCAAutorizacionCancelacionPagoAspspBce             GetEstadoSCAAutorizacionCancelacionPagoAspsp = "BCE"
	GetEstadoSCAAutorizacionCancelacionPagoAspspLaboralkutxa    GetEstadoSCAAutorizacionCancelacionPagoAspsp = "laboralkutxa"
	GetEstadoSCAAutorizacionCancelacionPagoAspspMediolanum      GetEstadoSCAAutorizacionCancelacionPagoAspsp = "mediolanum"
	GetEstadoSCAAutorizacionCancelacionPagoAspspOpenbank        GetEstadoSCAAutorizacionCancelacionPagoAspsp = "openbank"
	GetEstadoSCAAutorizacionCancelacionPagoAspspIbercaja        GetEstadoSCAAutorizacionCancelacionPagoAspsp = "ibercaja"
	GetEstadoSCAAutorizacionCancelacionPagoAspspSelfbank        GetEstadoSCAAutorizacionCancelacionPagoAspsp = "selfbank"
	GetEstadoSCAAutorizacionCancelacionPagoAspspInversis        GetEstadoSCAAutorizacionCancelacionPagoAspsp = "inversis"
	GetEstadoSCAAutorizacionCancelacionPagoAspspAndbank         GetEstadoSCAAutorizacionCancelacionPagoAspsp = "andbank"
	GetEstadoSCAAutorizacionCancelacionPagoAspspWizink          GetEstadoSCAAutorizacionCancelacionPagoAspsp = "wizink"
)

func (e GetEstadoSCAAutorizacionCancelacionPagoAspsp) ToPointer() *GetEstadoSCAAutorizacionCancelacionPagoAspsp {
	return &e
}

func (e *GetEstadoSCAAutorizacionCancelacionPagoAspsp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redsys":
		fallthrough
	case "BBVA":
		fallthrough
	case "BBVAPT":
		fallthrough
	case "BBVABE":
		fallthrough
	case "BBVAFR":
		fallthrough
	case "BBVAEUK":
		fallthrough
	case "caixabank":
		fallthrough
	case "BancSabadell":
		fallthrough
	case "bancosantander":
		fallthrough
	case "bancamarch":
		fallthrough
	case "bankoa":
		fallthrough
	case "cajamar":
		fallthrough
	case "arquia":
		fallthrough
	case "BFF":
		fallthrough
	case "colonya":
		fallthrough
	case "eurocajarural":
		fallthrough
	case "evobanco":
		fallthrough
	case "fiarebancaetica":
		fallthrough
	case "bancopichincha":
		fallthrough
	case "unicajabanco":
		fallthrough
	case "cajasur":
		fallthrough
	case "kutxabank":
		fallthrough
	case "bankinter":
		fallthrough
	case "renta4":
		fallthrough
	case "BCE":
		fallthrough
	case "laboralkutxa":
		fallthrough
	case "mediolanum":
		fallthrough
	case "openbank":
		fallthrough
	case "ibercaja":
		fallthrough
	case "selfbank":
		fallthrough
	case "inversis":
		fallthrough
	case "andbank":
		fallthrough
	case "wizink":
		*e = GetEstadoSCAAutorizacionCancelacionPagoAspsp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEstadoSCAAutorizacionCancelacionPagoAspsp: %v", v)
	}
}

// GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct
type GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct string

const (
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductSepaCreditTransfers             GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "sepa-credit-transfers"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductInstantSepaCreditTransfers      GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "instant-sepa-credit-transfers"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductTarget2Payments                 GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "target-2-payments"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductCrossBorderCreditTransfers      GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "cross-border-credit-transfers"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticCrossCurrencyPaymentsUk GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "domestic-cross-currency-payments-uk"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticChapsPaymentsUk         GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "domestic-chaps-payments-uk"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk        GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "domestic-faster-payments-uk"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk          GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct = "domestic-bacs-payments-uk"
)

func (e GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct) ToPointer() *GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct {
	return &e
}

func (e *GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sepa-credit-transfers":
		fallthrough
	case "instant-sepa-credit-transfers":
		fallthrough
	case "target-2-payments":
		fallthrough
	case "cross-border-credit-transfers":
		fallthrough
	case "domestic-cross-currency-payments-uk":
		fallthrough
	case "domestic-chaps-payments-uk":
		fallthrough
	case "domestic-faster-payments-uk":
		fallthrough
	case "domestic-bacs-payments-uk":
		*e = GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct: %v", v)
	}
}

// GetEstadoSCAAutorizacionCancelacionPagoPaymentService
type GetEstadoSCAAutorizacionCancelacionPagoPaymentService string

const (
	GetEstadoSCAAutorizacionCancelacionPagoPaymentServicePayments         GetEstadoSCAAutorizacionCancelacionPagoPaymentService = "payments"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentServiceBulkPayments     GetEstadoSCAAutorizacionCancelacionPagoPaymentService = "bulk-payments"
	GetEstadoSCAAutorizacionCancelacionPagoPaymentServicePeriodicPayments GetEstadoSCAAutorizacionCancelacionPagoPaymentService = "periodic-payments"
)

func (e GetEstadoSCAAutorizacionCancelacionPagoPaymentService) ToPointer() *GetEstadoSCAAutorizacionCancelacionPagoPaymentService {
	return &e
}

func (e *GetEstadoSCAAutorizacionCancelacionPagoPaymentService) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payments":
		fallthrough
	case "bulk-payments":
		fallthrough
	case "periodic-payments":
		*e = GetEstadoSCAAutorizacionCancelacionPagoPaymentService(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEstadoSCAAutorizacionCancelacionPagoPaymentService: %v", v)
	}
}

type GetEstadoSCAAutorizacionCancelacionPagoRequest struct {
	// Este dato es contenido si la transacción de inicio de pago forma parte de una sesión (combinación de AIS/PIS). Contendrá el consentId del consentimiento AIS que se realizó previo al inicio de pago. . Ej: Consent-ID: 123-qwer-456
	ConsentID *string `header:"style=simple,explode=false,name=Consent-ID"`
	// Es contenido si viaja el campo Signature. Ej: Digest: SHA-256=NzdmZjA4YjY5M2M2NDYyMmVjOWFmMGNmYTZiNTU3MjVmNDI4NTRlMzJkYzE3ZmNmMDE3ZGFmMjhhNTc5OTU3OQ==
	Digest string `header:"style=simple,explode=false,name=Digest"`
	// Accept header de la petición HTTP entre PSU y el TPP. Ej: PSU-Accept: application/json
	PSUAccept *string `header:"style=simple,explode=false,name=PSU-Accept"`
	// Accept charset header de la petición HTTP entre PSU y el TPP. PSU-Accept-Charset: utf-8
	PSUAcceptCharset *string `header:"style=simple,explode=false,name=PSU-Accept-Charset"`
	// Accept encoding header de la petición HTTP entre PSU y el TPP. PSU-Accept-Language: gzip
	PSUAcceptEncoding *string `header:"style=simple,explode=false,name=PSU-Accept-Encoding"`
	// Accept language header de la petición HTTP entre PSU y el TPP. PSU-Accept-Language: es-ES
	PSUAcceptLanguage *string `header:"style=simple,explode=false,name=PSU-Accept-Language"`
	// UUID (Universally Unique Identifier) para un dispositivo. El UUID identifica al dispositivo o a una instalación de una aplicación en un dispositivo. Este ID no debe ser modificado hasta la desinstalación de la aplicación del dispositivo. Ej: PSU-Device-ID: 5b3ab8e8-0fd5-43d2-946e-d75958b172e7
	PSUDeviceID *string `header:"style=simple,explode=false,name=PSU-Device-ID"`
	// Localización correspondiente a la petición HTTP entre el PSU y el TPP. Ej: PSU-Geo-Location: GEO:90.023856;25.345963
	PSUGeoLocation *string `header:"style=simple,explode=false,name=PSU-Geo-Location"`
	// Método HTTP usado en la interfaz entre PSU y TPP. Valores permitidos: POST. Ej: PSU-Http-Method: POST
	PSUHTTPMethod *string `header:"style=simple,explode=false,name=PSU-Http-Method"`
	// Dirección IP de la petición HTPP entre el PSU y el TPP. Si no está disponible, el TPP debe usar la dirección IP usada por el TPP cuando envía esta petición. Ej: Ej: PSU-IP-Address: 192.168.16.5
	PSUIPAddress string `header:"style=simple,explode=false,name=PSU-IP-Address"`
	// Puerto IP de la petición HTTP entre el PSU y el TPP si está disponible. Ejemplo: PSU-IP-Port: 443
	PSUIPPort *string `header:"style=simple,explode=false,name=PSU-IP-Port"`
	// Navegador o sistema operativo de la petición HTTP entre el PSU y el TPP. Ejemplo: PSU-User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
	PSUUserAgent *string `header:"style=simple,explode=false,name=PSU-User-Agent"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Si esta URI es contenida, el TPP está solicitando redirigir el flujo de la transacción a esta dirección en vez de al TPP-Redirect-URI en caso de un resultado negativo del método de SCA por redirección. "TPP-Nok-Redirect-URI":"https://www.tpp.com/cb/nok"
	TPPNokRedirectURI *string `header:"style=simple,explode=false,name=TPP-Nok-Redirect-URI"`
	// Si es "true", el TPP ha comunicado al HUB que prefiere SCA por redirección. Si es "false", el TPP ha comunicado al HUB que prefiere no ser redireccionado para SCA y el procedimiento será por flujo desacoplado. Si el parámetro no es usado, el ASPSP elegirá el flujo SCA a aplicar dependiendo del método SCA elegido por el TPP/PSU.
	TPPRedirectPreferred *string `header:"style=simple,explode=false,name=TPP-Redirect-Preferred"`
	// URI del TPP donde el flujo de la transacción debe ser redirigido después de alguna de las fases del SCA. Es recomendado usar siempre este campo de cabecera.En el futuro, este campo podría cambiar a obligatorio. Ej: TPP-Redirect-URI: https://www.tpp.com/cb
	TPPRedirectURI *string `header:"style=simple,explode=false,name=TPP-Redirect-URI"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	// Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
	Aspsp           GetEstadoSCAAutorizacionCancelacionPagoAspsp `pathParam:"style=simple,explode=false,name=aspsp"`
	AuthorisationID string                                       `pathParam:"style=simple,explode=false,name=authorisation-id"`
	// Identificador del recurso que referencia a la iniciación de pago o consentimiento.
	PaymentID      string                                                `pathParam:"style=simple,explode=false,name=payment-id"`
	PaymentProduct GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct `pathParam:"style=simple,explode=false,name=payment-product"`
	PaymentService GetEstadoSCAAutorizacionCancelacionPagoPaymentService `pathParam:"style=simple,explode=false,name=payment-service"`
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetTPPRedirectPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetAspsp() GetEstadoSCAAutorizacionCancelacionPagoAspsp {
	if o == nil {
		return GetEstadoSCAAutorizacionCancelacionPagoAspsp("")
	}
	return o.Aspsp
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetAuthorisationID() string {
	if o == nil {
		return ""
	}
	return o.AuthorisationID
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPaymentProduct() GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct {
	if o == nil {
		return GetEstadoSCAAutorizacionCancelacionPagoPaymentProduct("")
	}
	return o.PaymentProduct
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoRequest) GetPaymentService() GetEstadoSCAAutorizacionCancelacionPagoPaymentService {
	if o == nil {
		return GetEstadoSCAAutorizacionCancelacionPagoPaymentService("")
	}
	return o.PaymentService
}

type GetEstadoSCAAutorizacionCancelacionPagoResponse struct {
	ContentType string
	Headers     map[string][]string
	// HTTP/1.1 200 Ok
	ResponseGetSCAStatus *shared.ResponseGetSCAStatus
	StatusCode           int
	RawResponse          *http.Response
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoResponse) GetResponseGetSCAStatus() *shared.ResponseGetSCAStatus {
	if o == nil {
		return nil
	}
	return o.ResponseGetSCAStatus
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEstadoSCAAutorizacionCancelacionPagoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
