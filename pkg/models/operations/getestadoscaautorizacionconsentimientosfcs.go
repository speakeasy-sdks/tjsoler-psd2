// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
)

// GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp - Nombre del ASPSP al que desea realizar la petición.(BBVA, Banco Santander ...)
type GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp string

const (
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspRedsys          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "redsys"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBbva            GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BBVA"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBbvapt          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BBVAPT"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBbvabe          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BBVABE"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBbvafr          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BBVAFR"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBbvaeuk         GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BBVAEUK"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspCaixabank       GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "caixabank"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBancSabadell    GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BancSabadell"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBancosantander  GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "bancosantander"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBancamarch      GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "bancamarch"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBankoa          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "bankoa"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspCajamar         GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "cajamar"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspArquia          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "arquia"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBff             GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BFF"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspColonya         GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "colonya"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspEurocajarural   GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "eurocajarural"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspEvobanco        GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "evobanco"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspFiarebancaetica GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "fiarebancaetica"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBancopichincha  GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "bancopichincha"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspUnicajabanco    GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "unicajabanco"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspCajasur         GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "cajasur"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspKutxabank       GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "kutxabank"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBankinter       GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "bankinter"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspRenta4          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "renta4"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspBce             GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "BCE"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspLaboralkutxa    GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "laboralkutxa"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspMediolanum      GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "mediolanum"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspOpenbank        GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "openbank"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspIbercaja        GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "ibercaja"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspSelfbank        GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "selfbank"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspInversis        GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "inversis"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspAndbank         GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "andbank"
	GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspspWizink          GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp = "wizink"
)

func (e GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp) ToPointer() *GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp {
	return &e
}

func (e *GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp) UnmarshalJSON(data []byte) error {
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
		*e = GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp: %v", v)
	}
}

type GetEstadoSCAAutorizacionConsentimientosFCSRequest struct {
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
	Aspsp GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp `pathParam:"style=simple,explode=false,name=aspsp"`
	// Identificador del subrecurso asociado al ID de autorización
	AuthorisationID string `pathParam:"style=simple,explode=false,name=authorisation-id"`
	// Identificador asociado al ID de consentimiento, creado previamente
	ConsentIDPathParameter string `pathParam:"style=simple,explode=false,name=consent-id"`
}

func (g GetEstadoSCAAutorizacionConsentimientosFCSRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetEstadoSCAAutorizacionConsentimientosFCSRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetConsentID() *string {
	if o == nil {
		return nil
	}
	return o.ConsentID
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetTPPNokRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPNokRedirectURI
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetTPPRedirectPreferred() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectPreferred
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetTPPRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.TPPRedirectURI
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetAspsp() GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp {
	if o == nil {
		return GetEstadoSCAAutorizacionConsentimientosFCSPathParamAspsp("")
	}
	return o.Aspsp
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetAuthorisationID() string {
	if o == nil {
		return ""
	}
	return o.AuthorisationID
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSRequest) GetConsentIDPathParameter() string {
	if o == nil {
		return ""
	}
	return o.ConsentIDPathParameter
}

type GetEstadoSCAAutorizacionConsentimientosFCSResponse struct {
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

func (o *GetEstadoSCAAutorizacionConsentimientosFCSResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSResponse) GetResponseGetSCAStatus() *shared.ResponseGetSCAStatus {
	if o == nil {
		return nil
	}
	return o.ResponseGetSCAStatus
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEstadoSCAAutorizacionConsentimientosFCSResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
