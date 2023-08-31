// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"net/http"
)

// PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp - Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
type PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp string

const (
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspRedsys          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "redsys"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbva            PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BBVA"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbvapt          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BBVAPT"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbvabe          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BBVABE"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbvafr          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BBVAFR"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBbvaeuk         PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BBVAEUK"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspCaixabank       PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "caixabank"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBancSabadell    PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BancSabadell"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBancosantander  PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "bancosantander"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBancamarch      PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "bancamarch"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBankoa          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "bankoa"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspCajamar         PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "cajamar"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspArquia          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "arquia"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBff             PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BFF"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspColonya         PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "colonya"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspEurocajarural   PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "eurocajarural"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspEvobanco        PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "evobanco"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspFiarebancaetica PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "fiarebancaetica"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBancopichincha  PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "bancopichincha"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspUnicajabanco    PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "unicajabanco"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspCajasur         PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "cajasur"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspKutxabank       PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "kutxabank"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBankinter       PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "bankinter"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspRenta4          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "renta4"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspBce             PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "BCE"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspLaboralkutxa    PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "laboralkutxa"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspMediolanum      PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "mediolanum"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspOpenbank        PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "openbank"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspIbercaja        PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "ibercaja"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspSelfbank        PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "selfbank"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspInversis        PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "inversis"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspAndbank         PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "andbank"
	PutSeleccionarSCAAutorizacionConsentimientosFCSAspspWizink          PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp = "wizink"
)

func (e PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp) ToPointer() *PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp {
	return &e
}

func (e *PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp) UnmarshalJSON(data []byte) error {
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
		*e = PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp: %v", v)
	}
}

type PutSeleccionarSCAAutorizacionConsentimientosFCSRequest struct {
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
	PSUUserAgent         *string                     `header:"style=simple,explode=false,name=PSU-User-Agent"`
	RequestUpdatePSUData shared.RequestUpdatePSUData `request:"mediaType=application/json"`
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
	Aspsp PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp `pathParam:"style=simple,explode=false,name=aspsp"`
	// Identificador del subrecurso asociado al ID de autorización
	AuthorisationID string `pathParam:"style=simple,explode=false,name=authorisation-id"`
	// Identificador asociado al ID de consentimiento, creado previamente
	ConsentIDPathParameter string `pathParam:"style=simple,explode=false,name=consent-id"`
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetRequestUpdatePSUData() shared.RequestUpdatePSUData {
	if o == nil {
		return shared.RequestUpdatePSUData{}
	}
	return o.RequestUpdatePSUData
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetTPPRedirectPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetAspsp() PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp {
	if o == nil {
		return PutSeleccionarSCAAutorizacionConsentimientosFCSAspsp("")
	}
	return o.Aspsp
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetAuthorisationID() string {
	if o == nil {
		return ""
	}
	return o.AuthorisationID
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSRequest) GetConsentIDPathParameter() string {
	if o == nil {
		return ""
	}
	return o.ConsentIDPathParameter
}

type PutSeleccionarSCAAutorizacionConsentimientosFCSResponse struct {
	ContentType string
	Headers     map[string][]string
	// HTTP/1.1 200 Ok
	ResponseUpdatePSUData *shared.ResponseUpdatePSUData
	StatusCode            int
	RawResponse           *http.Response
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSResponse) GetResponseUpdatePSUData() *shared.ResponseUpdatePSUData {
	if o == nil {
		return nil
	}
	return o.ResponseUpdatePSUData
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSeleccionarSCAAutorizacionConsentimientosFCSResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
