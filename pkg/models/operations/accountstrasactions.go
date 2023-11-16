// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
)

type AccountsTrasactionsRequest struct {
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
	// Id de la cuenta
	AccountID string `pathParam:"style=simple,explode=false,name=account-id"`
	// Nombre del ASPSP al que desea realizar la petición.
	Aspsp string `pathParam:"style=simple,explode=false,name=aspsp"`
	// Estados de las transacciones devueltas. Los codigos de estado permitidos son “booked”, “pending”, "information" y “both”. Los obligatorios para los ASPSPs son “booked”.
	BookingStatus string `queryParam:"style=form,explode=true,name=bookingStatus"`
	// Fecha de inicio de consulta. Es incluido si no es requerido un acceso delta.
	DateFrom *string `queryParam:"style=form,explode=true,name=dateFrom"`
	// Fecha de fin de consulta. Su valor por defecto es la facha actual si no es dado. Ej: 2017-11-05
	DateTo *string `queryParam:"style=form,explode=true,name=dateTo"`
	// Indica que el AISP está a favor de obtener todas las transacciones después del último acceso de informe para esta PSU y cuenta. Este indicador podría ser rechazado por el ASPSP si esta función no es compatible.
	DeltaList *bool `queryParam:"style=form,explode=true,name=deltaList"`
	// Al ser indicado, nos daría los resultados desde la llamada con transactionId anterior al dado. Alternativo a dateFrom y dateTo
	EntryReferenceFrom *string `queryParam:"style=form,explode=true,name=entryReferenceFrom"`
	// Si está incluido, esta función incluye los balances. Esta petición será rechazada si el acceso a balances no lo recoge el consentimiento o el ASPSP no soporta este parámetro.
	WithBalance *bool `default:"false" queryParam:"style=form,explode=true,name=withBalance"`
}

func (a AccountsTrasactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountsTrasactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountsTrasactionsRequest) GetConsentID() string {
	if o == nil {
		return ""
	}
	return o.ConsentID
}

func (o *AccountsTrasactionsRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *AccountsTrasactionsRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *AccountsTrasactionsRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *AccountsTrasactionsRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *AccountsTrasactionsRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *AccountsTrasactionsRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *AccountsTrasactionsRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *AccountsTrasactionsRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *AccountsTrasactionsRequest) GetPSUIPAddress() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPAddress
}

func (o *AccountsTrasactionsRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *AccountsTrasactionsRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *AccountsTrasactionsRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *AccountsTrasactionsRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *AccountsTrasactionsRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *AccountsTrasactionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsTrasactionsRequest) GetAspsp() string {
	if o == nil {
		return ""
	}
	return o.Aspsp
}

func (o *AccountsTrasactionsRequest) GetBookingStatus() string {
	if o == nil {
		return ""
	}
	return o.BookingStatus
}

func (o *AccountsTrasactionsRequest) GetDateFrom() *string {
	if o == nil {
		return nil
	}
	return o.DateFrom
}

func (o *AccountsTrasactionsRequest) GetDateTo() *string {
	if o == nil {
		return nil
	}
	return o.DateTo
}

func (o *AccountsTrasactionsRequest) GetDeltaList() *bool {
	if o == nil {
		return nil
	}
	return o.DeltaList
}

func (o *AccountsTrasactionsRequest) GetEntryReferenceFrom() *string {
	if o == nil {
		return nil
	}
	return o.EntryReferenceFrom
}

func (o *AccountsTrasactionsRequest) GetWithBalance() *bool {
	if o == nil {
		return nil
	}
	return o.WithBalance
}

type AccountsTrasactionsResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AccountsTrasactionsResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *AccountsTrasactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountsTrasactionsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *AccountsTrasactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountsTrasactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
