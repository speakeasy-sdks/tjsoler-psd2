// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
)

type GetAccountIdv11Request struct {
	// Identificador del consentimiento sobre el que iría la consulta de cuentas Ej: Consent-ID: 7890-asdf-4321
	ConsentID string `default:"7890-asdf-4321" header:"style=simple,explode=false,name=Consent-ID"`
	// Es contenido si viaja el campo Signature. Ej: Digest: SHA-256=NzdmZjA4YjY5M2M2NDYyMmVjOWFmMGNmYTZiNTU3MjVmNDI4NTRlMzJkYzE3ZmNmMDE3ZGFmMjhhNTc5OTU3OQ==
	Digest string `header:"style=simple,explode=false,name=Digest"`
	// Accept header de la petición HTTP entre PSU y el TPP. Ej: PSU-Accept: application/json
	PSUAccept *string `header:"style=simple,explode=false,name=PSU-Accept"`
	// Accept charset header de la petición HTTP entre PSU y el TPP. PSU-Accept-Charset: utf-8
	PSUAcceptCharset *string `header:"style=simple,explode=false,name=PSU-Accept-Charset"`
	// Accept language header de la petición HTTP entre PSU y el TPP. PSU-Accept-Language: gzip
	PSUAcceptEncoding *string `header:"style=simple,explode=false,name=PSU-Accept-Encoding"`
	// Accept language header de la petición HTTP entre PSU y el TPP. PSU-Accept-Language: es-ES
	PSUAcceptLanguage *string `header:"style=simple,explode=false,name=PSU-Accept-Language"`
	// UUID (Universally Unique Identifier) para un dispositivo. El UUID identifica al dispositivo o a una instalación de una aplicación en un dispositivo. Este ID no debe ser modificado hasta la desinstalación de la aplicación del dispositivo. Ej: PSU-Device-ID: 5b3ab8e8-0fd5-43d2-946e-d75958b172e7
	PSUDeviceID *string `header:"style=simple,explode=false,name=PSU-Device-ID"`
	// Localización correspondiente a la petición HTTP entre el PSU y el TPP. Ej: PSU-Geo-Location: GEO:90.023856;25.345963
	PSUGeoLocation *string `header:"style=simple,explode=false,name=PSU-Geo-Location"`
	// Método HTTP usado en la interfaz entre PSU y TPP. Valores permitidos: GET, POST, PUT, PATCH, DELETE. Ej: PSU-Http-Method: GET
	PSUHTTPMethod *string `default:"GET" header:"style=simple,explode=false,name=PSU-Http-Method"`
	// Dirección IP de la petición HTPP entre el PSU y el TPP. Si no está disponible, el TPP debe usar la dirección IP usada por el TPP cuando envía esta petición. Ej: Ej: PSU-IP-Address: 192.168.16.5
	PSUIPAddress *string `header:"style=simple,explode=false,name=PSU-IP-Address"`
	// Puerto IP de la petición HTTP entre el PSU y el TPP si está disponible. Ejemplo: PSU-IP-Port: 443
	PSUIPPort *string `header:"style=simple,explode=false,name=PSU-IP-Port"`
	// Navegador o sistema operativo de la petición HTTP entre el PSU y el TPP. Ejemplo: PSU-User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
	PSUUserAgent *string `header:"style=simple,explode=false,name=PSU-User-Agent"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `default:"TestTPPCertificate" header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	AccountID  string `pathParam:"style=simple,explode=false,name=account-id"`
	// Nombre del ASPSP al que desea realizar la petición.
	Aspsp string `pathParam:"style=simple,explode=false,name=aspsp"`
	// Si está incluido, esta función incluye los balances. Esta petición será rechazada si el acceso a balances no lo recoge el consentimiento o el ASPSP no soporta este parámetro.
	WithBalance *bool `default:"false" queryParam:"style=form,explode=true,name=withBalance"`
}

func (g GetAccountIdv11Request) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAccountIdv11Request) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAccountIdv11Request) GetConsentID() string {
	if o == nil {
		return ""
	}
	return o.ConsentID
}

func (o *GetAccountIdv11Request) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetAccountIdv11Request) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetAccountIdv11Request) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetAccountIdv11Request) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetAccountIdv11Request) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetAccountIdv11Request) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetAccountIdv11Request) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetAccountIdv11Request) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetAccountIdv11Request) GetPSUIPAddress() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPAddress
}

func (o *GetAccountIdv11Request) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetAccountIdv11Request) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetAccountIdv11Request) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetAccountIdv11Request) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetAccountIdv11Request) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetAccountIdv11Request) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetAccountIdv11Request) GetAspsp() string {
	if o == nil {
		return ""
	}
	return o.Aspsp
}

func (o *GetAccountIdv11Request) GetWithBalance() *bool {
	if o == nil {
		return nil
	}
	return o.WithBalance
}

type GetAccountIdv11Response struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountIdv11Response) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetAccountIdv11Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountIdv11Response) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetAccountIdv11Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountIdv11Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
