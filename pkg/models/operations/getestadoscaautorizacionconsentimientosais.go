// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
)

// GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp - Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
type GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp string

const (
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspRedsys          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "redsys"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbva            GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BBVA"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbvapt          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BBVAPT"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbvabe          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BBVABE"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbvafr          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BBVAFR"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBbvaeuk         GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BBVAEUK"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspCaixabank       GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "caixabank"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBancSabadell    GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BancSabadell"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBancosantander  GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "bancosantander"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBancamarch      GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "bancamarch"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBankoa          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "bankoa"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspCajamar         GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "cajamar"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspArquia          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "arquia"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBff             GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BFF"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspColonya         GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "colonya"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspEurocajarural   GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "eurocajarural"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspEvobanco        GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "evobanco"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspFiarebancaetica GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "fiarebancaetica"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBancopichincha  GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "bancopichincha"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspUnicajabanco    GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "unicajabanco"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspCajasur         GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "cajasur"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspKutxabank       GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "kutxabank"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBankinter       GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "bankinter"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspRenta4          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "renta4"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspBce             GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "BCE"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspLaboralkutxa    GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "laboralkutxa"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspMediolanum      GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "mediolanum"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspOpenbank        GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "openbank"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspIbercaja        GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "ibercaja"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspSelfbank        GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "selfbank"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspInversis        GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "inversis"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspAndbank         GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "andbank"
	GetEstadoSCAAutorizacionConsentimientosAISPathParamAspspWizink          GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp = "wizink"
)

func (e GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp) ToPointer() *GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp {
	return &e
}

func (e *GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp) UnmarshalJSON(data []byte) error {
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
		*e = GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp: %v", v)
	}
}

type GetEstadoSCAAutorizacionConsentimientosAISRequest struct {
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
	PSUHTTPMethod *string `default:"POST" header:"style=simple,explode=false,name=PSU-Http-Method"`
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
	Aspsp GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp `pathParam:"style=simple,explode=false,name=aspsp"`
	// Identificador del subrecurso asociado al ID de autorización
	AuthorisationID string `pathParam:"style=simple,explode=false,name=authorisation-id"`
	// Identificador asociado al ID de consentimiento, creado previamente
	ConsentIDPathParameter string `pathParam:"style=simple,explode=false,name=consent-id"`
}

func (g GetEstadoSCAAutorizacionConsentimientosAISRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetEstadoSCAAutorizacionConsentimientosAISRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetTPPRedirectPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetAspsp() GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp {
	if o == nil {
		return GetEstadoSCAAutorizacionConsentimientosAISPathParamAspsp("")
	}
	return o.Aspsp
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetAuthorisationID() string {
	if o == nil {
		return ""
	}
	return o.AuthorisationID
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISRequest) GetConsentIDPathParameter() string {
	if o == nil {
		return ""
	}
	return o.ConsentIDPathParameter
}

type GetEstadoSCAAutorizacionConsentimientosAISResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP/1.1 200 Ok
	ResponseGetSCAStatus *shared.ResponseGetSCAStatus
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISResponse) GetResponseGetSCAStatus() *shared.ResponseGetSCAStatus {
	if o == nil {
		return nil
	}
	return o.ResponseGetSCAStatus
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEstadoSCAAutorizacionConsentimientosAISResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
