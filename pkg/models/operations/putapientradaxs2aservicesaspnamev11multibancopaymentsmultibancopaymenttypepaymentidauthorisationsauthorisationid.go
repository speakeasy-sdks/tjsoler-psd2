// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"net/http"
)

type PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest struct {
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
	// Dirección IP de la petición HTPP entre el PSU y el TPP. Si no está disponible, el TPP debe usar la dirección IP usada por el TPP cuando envía esta petición. Ej: Ej: PSU-IP-Address: 192.168.16.5
	PSUIPAddress string `header:"style=simple,explode=false,name=PSU-IP-Address"`
	// Puerto IP de la petición HTTP entre el PSU y el TPP si está disponible. Ejemplo: PSU-IP-Port: 443
	PSUIPPort *string `header:"style=simple,explode=false,name=PSU-IP-Port"`
	// Navegador o sistema operativo de la petición HTTP entre el PSU y el TPP. Ejemplo: PSU-User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
	PSUUserAgent              *string                          `header:"style=simple,explode=false,name=PSU-User-Agent"`
	RequestActualizarDatosPsu shared.RequestActualizarDatosPsu `request:"mediaType=application/json"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	// Nombre del ASPSP al que desea realizar la petición.
	AspName string `pathParam:"style=simple,explode=false,name=asp-name"`
	// Identificador del sub-recurso asociado al inicio de pago Multibanco
	AuthorisationID string `pathParam:"style=simple,explode=false,name=authorisationId"`
	// Tipo de pago multibanco
	MultibancoPaymentType string `pathParam:"style=simple,explode=false,name=multibanco-payment-type"`
	// Identificador del recurso que referencia a la iniciación de pago
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUCorporateID() *string {
	if o == nil {
		return nil
	}
	return o.PSUCorporateID
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUCorporateIDType() *string {
	if o == nil {
		return nil
	}
	return o.PSUCorporateIDType
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetRequestActualizarDatosPsu() shared.RequestActualizarDatosPsu {
	if o == nil {
		return shared.RequestActualizarDatosPsu{}
	}
	return o.RequestActualizarDatosPsu
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetAspName() string {
	if o == nil {
		return ""
	}
	return o.AspName
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetAuthorisationID() string {
	if o == nil {
		return ""
	}
	return o.AuthorisationID
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetMultibancoPaymentType() string {
	if o == nil {
		return ""
	}
	return o.MultibancoPaymentType
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

type PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse struct {
	ContentType                      string
	ResponseMultibankSelectSCAMethod *shared.ResponseMultibankSelectSCAMethod
	StatusCode                       int
	RawResponse                      *http.Response
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse) GetResponseMultibankSelectSCAMethod() *shared.ResponseMultibankSelectSCAMethod {
	if o == nil {
		return nil
	}
	return o.ResponseMultibankSelectSCAMethod
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIEntradaXs2aServicesAspNameV11MultibancoPaymentsMultibancoPaymentTypePaymentIDAuthorisationsAuthorisationIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
