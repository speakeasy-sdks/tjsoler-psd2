// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"net/http"
)

// InitiationBulkPaymentPaymentProduct
type InitiationBulkPaymentPaymentProduct string

const (
	InitiationBulkPaymentPaymentProductSepaCreditTransfers        InitiationBulkPaymentPaymentProduct = "sepa-credit-transfers"
	InitiationBulkPaymentPaymentProductInstantSepaCreditTransfers InitiationBulkPaymentPaymentProduct = "instant-sepa-credit-transfers"
	InitiationBulkPaymentPaymentProductTarget2Payments            InitiationBulkPaymentPaymentProduct = "target-2-payments"
	InitiationBulkPaymentPaymentProductCrossBorderCreditTransfers InitiationBulkPaymentPaymentProduct = "cross-border-credit-transfers"
)

func (e InitiationBulkPaymentPaymentProduct) ToPointer() *InitiationBulkPaymentPaymentProduct {
	return &e
}

func (e *InitiationBulkPaymentPaymentProduct) UnmarshalJSON(data []byte) error {
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
		*e = InitiationBulkPaymentPaymentProduct(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InitiationBulkPaymentPaymentProduct: %v", v)
	}
}

type InitiationBulkPaymentRequest struct {
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
	// Identificador de "empresa" en los Canales Online.
	PSUCorporateID *string `header:"style=simple,explode=false,name=PSU-Corporate-ID"`
	// Tipo del PSU-Corporate-ID necesario por el ASPSP para identificar su contenido.
	PSUCorporateIDType *string `header:"style=simple,explode=false,name=PSU-Corporate-ID-Type"`
	// UUID (Universally Unique Identifier) para un dispositivo. El UUID identifica al dispositivo o a una instalación de una aplicación en un dispositivo. Este ID no debe ser modificado hasta la desinstalación de la aplicación del dispositivo. Ej: PSU-Device-ID: 5b3ab8e8-0fd5-43d2-946e-d75958b172e7
	PSUDeviceID *string `header:"style=simple,explode=false,name=PSU-Device-ID"`
	// Localización correspondiente a la petición HTTP entre el PSU y el TPP. Ej: PSU-Geo-Location: GEO:90.023856;25.345963
	PSUGeoLocation *string `header:"style=simple,explode=false,name=PSU-Geo-Location"`
	// Método HTTP usado en la interfaz entre PSU y TPP. Valores permitidos: POST. Ej: PSU-Http-Method: POST
	PSUHTTPMethod *string `header:"style=simple,explode=false,name=PSU-Http-Method"`
	// Identificador que el PSU utiliza para identificarse en su ASPSP. Puede ser informado incluso si se está usando un token de OAuth y, en tal caso, el ASPSP podría comprobar que el PSU-ID y el token se corresponden.
	PsuID *string `header:"style=simple,explode=false,name=PSU-ID"`
	// Tipo del PSU-ID. Necesario en escenarios donde el PSU tiene varios PSU-IDs como posibilidades de acceso.
	PSUIDType *string `header:"style=simple,explode=false,name=PSU-ID-Type"`
	// Dirección IP de la petición HTPP entre el PSU y el TPP. Si no está disponible, el TPP debe usar la dirección IP usada por el TPP cuando envía esta petición. Ej: Ej: PSU-IP-Address: 192.168.16.5
	PSUIPAddress string `header:"style=simple,explode=false,name=PSU-IP-Address"`
	// Puerto IP de la petición HTTP entre el PSU y el TPP si está disponible. Ejemplo: PSU-IP-Port: 443
	PSUIPPort *string `header:"style=simple,explode=false,name=PSU-IP-Port"`
	// Navegador o sistema operativo de la petición HTTP entre el PSU y el TPP. Ejemplo: PSU-User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
	PSUUserAgent       *string                   `header:"style=simple,explode=false,name=PSU-User-Agent"`
	RequestBulkPayment shared.RequestBulkPayment `request:"mediaType=application/json"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Este campo podría ser usado por el TPP para informar al ASPSP acerca de la marca (Brand) usada por el TPP de cara al PSU. Esta información puede ser usada para mejorar la comunicación entre el ASPSP y el PSU o el ASPSP y el TPP.
	TPPBrandLoggingInformation *string `header:"style=simple,explode=false,name=TPP-Brand-Logging-Information"`
	// Si es igual a true, el TPP prefiere iniciar el proceso de autorización separadamente, por ej. debido a la necesidad de la autorización de un conjunto de operaciones simultáneamente.
	TPPExplicitAuthorisationPreferred *bool `header:"style=simple,explode=false,name=TPP-Explicit-Authorisation-Preferred"`
	// Si esta URI es contenida, el TPP está solicitando redirigir el flujo de la transacción a esta dirección en vez de al TPP-Redirect-URI en caso de un resultado negativo del método de SCA por redirección. "TPP-Nok-Redirect-URI":"https://www.tpp.com/cb/nok"
	TPPNokRedirectURI *string `header:"style=simple,explode=false,name=TPP-Nok-Redirect-URI"`
	// Nota: Este campo será ignorado en caso de venir informado por el TPP.
	TPPNotificationContentPreferred *string `header:"style=simple,explode=false,name=TPP-Notification-Content-Preferred"`
	// Nota: Este campo será ignorado en caso de venir informado por el TPP.
	TPPNotificationURI *string `header:"style=simple,explode=false,name=TPP-Notification-URI"`
	// Si es "true", el TPP ha comunicado al HUB que prefiere SCA por redirección. Si es "false", el TPP ha comunicado al HUB que prefiere no ser redireccionado para SCA y el procedimiento será por flujo desacoplado. Si el parámetro no es usado, el ASPSP elegirá el flujo SCA a aplicar dependiendo del método SCA elegido por el TPP/PSU.
	TPPRedirectPreferred *bool `header:"style=simple,explode=false,name=TPP-Redirect-Preferred"`
	// URI del TPP donde el flujo de la transacción debe ser redirigido después de alguna de las fases del SCA. Es recomendado usar siempre este campo de cabecera.En el futuro, este campo podría cambiar a obligatorio. Ej: TPP-Redirect-URI: https://www.tpp.com/cb
	TPPRedirectURI *string `header:"style=simple,explode=false,name=TPP-Redirect-URI"`
	// Nota: Este campo será ignorado en caso de venir informado por el TPP.
	TPPRejectionNoFundsPreferred *string `header:"style=simple,explode=false,name=TPP-Rejection-NoFunds-Preferred"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	// Nombre del ASPSP al que desea realizar la petición.
	Aspsp          string                              `pathParam:"style=simple,explode=false,name=aspsp"`
	PaymentProduct InitiationBulkPaymentPaymentProduct `pathParam:"style=simple,explode=false,name=payment-product"`
}

func (o *InitiationBulkPaymentRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *InitiationBulkPaymentRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *InitiationBulkPaymentRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *InitiationBulkPaymentRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *InitiationBulkPaymentRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *InitiationBulkPaymentRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *InitiationBulkPaymentRequest) GetPSUCorporateID() *string {
	if o == nil {
		return nil
	}
	return o.PSUCorporateID
}

func (o *InitiationBulkPaymentRequest) GetPSUCorporateIDType() *string {
	if o == nil {
		return nil
	}
	return o.PSUCorporateIDType
}

func (o *InitiationBulkPaymentRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *InitiationBulkPaymentRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *InitiationBulkPaymentRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *InitiationBulkPaymentRequest) GetPsuID() *string {
	if o == nil {
		return nil
	}
	return o.PsuID
}

func (o *InitiationBulkPaymentRequest) GetPSUIDType() *string {
	if o == nil {
		return nil
	}
	return o.PSUIDType
}

func (o *InitiationBulkPaymentRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *InitiationBulkPaymentRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *InitiationBulkPaymentRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *InitiationBulkPaymentRequest) GetRequestBulkPayment() shared.RequestBulkPayment {
	if o == nil {
		return shared.RequestBulkPayment{}
	}
	return o.RequestBulkPayment
}

func (o *InitiationBulkPaymentRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *InitiationBulkPaymentRequest) GetTPPBrandLoggingInformation() *string {
	if o == nil {
		return nil
	}
	return o.TPPBrandLoggingInformation
}

func (o *InitiationBulkPaymentRequest) GetTPPExplicitAuthorisationPreferred() *bool {
	if o == nil {
		return nil
	}
	return o.TPPExplicitAuthorisationPreferred
}

func (o *InitiationBulkPaymentRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *InitiationBulkPaymentRequest) GetTPPNotificationContentPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPNotificationContentPreferred
}

func (o *InitiationBulkPaymentRequest) GetTPPNotificationURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNotificationURI
}

func (o *InitiationBulkPaymentRequest) GetTPPRedirectPreferred() *bool {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *InitiationBulkPaymentRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *InitiationBulkPaymentRequest) GetTPPRejectionNoFundsPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRejectionNoFundsPreferred
}

func (o *InitiationBulkPaymentRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *InitiationBulkPaymentRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *InitiationBulkPaymentRequest) GetAspsp() string {
	if o == nil {
		return ""
	}
	return o.Aspsp
}

func (o *InitiationBulkPaymentRequest) GetPaymentProduct() InitiationBulkPaymentPaymentProduct {
	if o == nil {
		return InitiationBulkPaymentPaymentProduct("")
	}
	return o.PaymentProduct
}

type InitiationBulkPaymentResponse struct {
	ContentType string
	Headers     map[string][]string
	// HTTP/1.1 201 Created
	ResponsePaymentInitiation *shared.ResponsePaymentInitiation
	StatusCode                int
	RawResponse               *http.Response
}

func (o *InitiationBulkPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InitiationBulkPaymentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *InitiationBulkPaymentResponse) GetResponsePaymentInitiation() *shared.ResponsePaymentInitiation {
	if o == nil {
		return nil
	}
	return o.ResponsePaymentInitiation
}

func (o *InitiationBulkPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InitiationBulkPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
