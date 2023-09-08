// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"net/http"
)

// GetSubRecursosAutorizacionCancelacionPagoAspsp - Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
type GetSubRecursosAutorizacionCancelacionPagoAspsp string

const (
	GetSubRecursosAutorizacionCancelacionPagoAspspRedsys          GetSubRecursosAutorizacionCancelacionPagoAspsp = "redsys"
	GetSubRecursosAutorizacionCancelacionPagoAspspBbva            GetSubRecursosAutorizacionCancelacionPagoAspsp = "BBVA"
	GetSubRecursosAutorizacionCancelacionPagoAspspBbvapt          GetSubRecursosAutorizacionCancelacionPagoAspsp = "BBVAPT"
	GetSubRecursosAutorizacionCancelacionPagoAspspBbvabe          GetSubRecursosAutorizacionCancelacionPagoAspsp = "BBVABE"
	GetSubRecursosAutorizacionCancelacionPagoAspspBbvafr          GetSubRecursosAutorizacionCancelacionPagoAspsp = "BBVAFR"
	GetSubRecursosAutorizacionCancelacionPagoAspspBbvaeuk         GetSubRecursosAutorizacionCancelacionPagoAspsp = "BBVAEUK"
	GetSubRecursosAutorizacionCancelacionPagoAspspCaixabank       GetSubRecursosAutorizacionCancelacionPagoAspsp = "caixabank"
	GetSubRecursosAutorizacionCancelacionPagoAspspBancSabadell    GetSubRecursosAutorizacionCancelacionPagoAspsp = "BancSabadell"
	GetSubRecursosAutorizacionCancelacionPagoAspspBancosantander  GetSubRecursosAutorizacionCancelacionPagoAspsp = "bancosantander"
	GetSubRecursosAutorizacionCancelacionPagoAspspBancamarch      GetSubRecursosAutorizacionCancelacionPagoAspsp = "bancamarch"
	GetSubRecursosAutorizacionCancelacionPagoAspspBankoa          GetSubRecursosAutorizacionCancelacionPagoAspsp = "bankoa"
	GetSubRecursosAutorizacionCancelacionPagoAspspCajamar         GetSubRecursosAutorizacionCancelacionPagoAspsp = "cajamar"
	GetSubRecursosAutorizacionCancelacionPagoAspspArquia          GetSubRecursosAutorizacionCancelacionPagoAspsp = "arquia"
	GetSubRecursosAutorizacionCancelacionPagoAspspBff             GetSubRecursosAutorizacionCancelacionPagoAspsp = "BFF"
	GetSubRecursosAutorizacionCancelacionPagoAspspColonya         GetSubRecursosAutorizacionCancelacionPagoAspsp = "colonya"
	GetSubRecursosAutorizacionCancelacionPagoAspspEurocajarural   GetSubRecursosAutorizacionCancelacionPagoAspsp = "eurocajarural"
	GetSubRecursosAutorizacionCancelacionPagoAspspEvobanco        GetSubRecursosAutorizacionCancelacionPagoAspsp = "evobanco"
	GetSubRecursosAutorizacionCancelacionPagoAspspFiarebancaetica GetSubRecursosAutorizacionCancelacionPagoAspsp = "fiarebancaetica"
	GetSubRecursosAutorizacionCancelacionPagoAspspBancopichincha  GetSubRecursosAutorizacionCancelacionPagoAspsp = "bancopichincha"
	GetSubRecursosAutorizacionCancelacionPagoAspspUnicajabanco    GetSubRecursosAutorizacionCancelacionPagoAspsp = "unicajabanco"
	GetSubRecursosAutorizacionCancelacionPagoAspspCajasur         GetSubRecursosAutorizacionCancelacionPagoAspsp = "cajasur"
	GetSubRecursosAutorizacionCancelacionPagoAspspKutxabank       GetSubRecursosAutorizacionCancelacionPagoAspsp = "kutxabank"
	GetSubRecursosAutorizacionCancelacionPagoAspspBankinter       GetSubRecursosAutorizacionCancelacionPagoAspsp = "bankinter"
	GetSubRecursosAutorizacionCancelacionPagoAspspRenta4          GetSubRecursosAutorizacionCancelacionPagoAspsp = "renta4"
	GetSubRecursosAutorizacionCancelacionPagoAspspBce             GetSubRecursosAutorizacionCancelacionPagoAspsp = "BCE"
	GetSubRecursosAutorizacionCancelacionPagoAspspLaboralkutxa    GetSubRecursosAutorizacionCancelacionPagoAspsp = "laboralkutxa"
	GetSubRecursosAutorizacionCancelacionPagoAspspMediolanum      GetSubRecursosAutorizacionCancelacionPagoAspsp = "mediolanum"
	GetSubRecursosAutorizacionCancelacionPagoAspspOpenbank        GetSubRecursosAutorizacionCancelacionPagoAspsp = "openbank"
	GetSubRecursosAutorizacionCancelacionPagoAspspIbercaja        GetSubRecursosAutorizacionCancelacionPagoAspsp = "ibercaja"
	GetSubRecursosAutorizacionCancelacionPagoAspspSelfbank        GetSubRecursosAutorizacionCancelacionPagoAspsp = "selfbank"
	GetSubRecursosAutorizacionCancelacionPagoAspspInversis        GetSubRecursosAutorizacionCancelacionPagoAspsp = "inversis"
	GetSubRecursosAutorizacionCancelacionPagoAspspAndbank         GetSubRecursosAutorizacionCancelacionPagoAspsp = "andbank"
	GetSubRecursosAutorizacionCancelacionPagoAspspWizink          GetSubRecursosAutorizacionCancelacionPagoAspsp = "wizink"
)

func (e GetSubRecursosAutorizacionCancelacionPagoAspsp) ToPointer() *GetSubRecursosAutorizacionCancelacionPagoAspsp {
	return &e
}

func (e *GetSubRecursosAutorizacionCancelacionPagoAspsp) UnmarshalJSON(data []byte) error {
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
		*e = GetSubRecursosAutorizacionCancelacionPagoAspsp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSubRecursosAutorizacionCancelacionPagoAspsp: %v", v)
	}
}

type GetSubRecursosAutorizacionCancelacionPagoPaymentProduct string

const (
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductSepaCreditTransfers             GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "sepa-credit-transfers"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductInstantSepaCreditTransfers      GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "instant-sepa-credit-transfers"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductTarget2Payments                 GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "target-2-payments"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductCrossBorderCreditTransfers      GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "cross-border-credit-transfers"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticCrossCurrencyPaymentsUk GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "domestic-cross-currency-payments-uk"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticChapsPaymentsUk         GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "domestic-chaps-payments-uk"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticFasterPaymentsUk        GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "domestic-faster-payments-uk"
	GetSubRecursosAutorizacionCancelacionPagoPaymentProductDomesticBacsPaymentsUk          GetSubRecursosAutorizacionCancelacionPagoPaymentProduct = "domestic-bacs-payments-uk"
)

func (e GetSubRecursosAutorizacionCancelacionPagoPaymentProduct) ToPointer() *GetSubRecursosAutorizacionCancelacionPagoPaymentProduct {
	return &e
}

func (e *GetSubRecursosAutorizacionCancelacionPagoPaymentProduct) UnmarshalJSON(data []byte) error {
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
		*e = GetSubRecursosAutorizacionCancelacionPagoPaymentProduct(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSubRecursosAutorizacionCancelacionPagoPaymentProduct: %v", v)
	}
}

type GetSubRecursosAutorizacionCancelacionPagoPaymentService string

const (
	GetSubRecursosAutorizacionCancelacionPagoPaymentServicePayments         GetSubRecursosAutorizacionCancelacionPagoPaymentService = "payments"
	GetSubRecursosAutorizacionCancelacionPagoPaymentServiceBulkPayments     GetSubRecursosAutorizacionCancelacionPagoPaymentService = "bulk-payments"
	GetSubRecursosAutorizacionCancelacionPagoPaymentServicePeriodicPayments GetSubRecursosAutorizacionCancelacionPagoPaymentService = "periodic-payments"
)

func (e GetSubRecursosAutorizacionCancelacionPagoPaymentService) ToPointer() *GetSubRecursosAutorizacionCancelacionPagoPaymentService {
	return &e
}

func (e *GetSubRecursosAutorizacionCancelacionPagoPaymentService) UnmarshalJSON(data []byte) error {
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
		*e = GetSubRecursosAutorizacionCancelacionPagoPaymentService(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSubRecursosAutorizacionCancelacionPagoPaymentService: %v", v)
	}
}

type GetSubRecursosAutorizacionCancelacionPagoRequest struct {
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
	Aspsp GetSubRecursosAutorizacionCancelacionPagoAspsp `pathParam:"style=simple,explode=false,name=aspsp"`
	// Identificador del recurso que referencia a la iniciación de pago o consentimiento.
	PaymentID      string                                                  `pathParam:"style=simple,explode=false,name=payment-id"`
	PaymentProduct GetSubRecursosAutorizacionCancelacionPagoPaymentProduct `pathParam:"style=simple,explode=false,name=payment-product"`
	PaymentService GetSubRecursosAutorizacionCancelacionPagoPaymentService `pathParam:"style=simple,explode=false,name=payment-service"`
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetTPPRedirectPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetAspsp() GetSubRecursosAutorizacionCancelacionPagoAspsp {
	if o == nil {
		return GetSubRecursosAutorizacionCancelacionPagoAspsp("")
	}
	return o.Aspsp
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPaymentProduct() GetSubRecursosAutorizacionCancelacionPagoPaymentProduct {
	if o == nil {
		return GetSubRecursosAutorizacionCancelacionPagoPaymentProduct("")
	}
	return o.PaymentProduct
}

func (o *GetSubRecursosAutorizacionCancelacionPagoRequest) GetPaymentService() GetSubRecursosAutorizacionCancelacionPagoPaymentService {
	if o == nil {
		return GetSubRecursosAutorizacionCancelacionPagoPaymentService("")
	}
	return o.PaymentService
}

type GetSubRecursosAutorizacionCancelacionPagoResponse struct {
	ContentType string
	Headers     map[string][]string
	// HTTP/1.1 200 Ok
	ResponseGetSubResourcesAuthorization *shared.ResponseGetSubResourcesAuthorization
	StatusCode                           int
	RawResponse                          *http.Response
}

func (o *GetSubRecursosAutorizacionCancelacionPagoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSubRecursosAutorizacionCancelacionPagoResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetSubRecursosAutorizacionCancelacionPagoResponse) GetResponseGetSubResourcesAuthorization() *shared.ResponseGetSubResourcesAuthorization {
	if o == nil {
		return nil
	}
	return o.ResponseGetSubResourcesAuthorization
}

func (o *GetSubRecursosAutorizacionCancelacionPagoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSubRecursosAutorizacionCancelacionPagoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
