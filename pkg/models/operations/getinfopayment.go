// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
)

// GetInfoPaymentPaymentProduct - Producto de pago a usar.
type GetInfoPaymentPaymentProduct string

const (
	GetInfoPaymentPaymentProductSepaCreditTransfers        GetInfoPaymentPaymentProduct = "sepa-credit-transfers"
	GetInfoPaymentPaymentProductInstantSepaCreditTransfers GetInfoPaymentPaymentProduct = "instant-sepa-credit-transfers"
	GetInfoPaymentPaymentProductTarget2Payments            GetInfoPaymentPaymentProduct = "target-2-payments"
	GetInfoPaymentPaymentProductCrossBorderCreditTransfers GetInfoPaymentPaymentProduct = "cross-border-credit-transfers"
)

func (e GetInfoPaymentPaymentProduct) ToPointer() *GetInfoPaymentPaymentProduct {
	return &e
}

func (e *GetInfoPaymentPaymentProduct) UnmarshalJSON(data []byte) error {
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
		*e = GetInfoPaymentPaymentProduct(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetInfoPaymentPaymentProduct: %v", v)
	}
}

// GetInfoPaymentPaymentService - Producto de pago a usar.
type GetInfoPaymentPaymentService string

const (
	GetInfoPaymentPaymentServicePayments         GetInfoPaymentPaymentService = "payments"
	GetInfoPaymentPaymentServiceBulkPayments     GetInfoPaymentPaymentService = "bulk-payments"
	GetInfoPaymentPaymentServicePeriodicPayments GetInfoPaymentPaymentService = "periodic-payments"
)

func (e GetInfoPaymentPaymentService) ToPointer() *GetInfoPaymentPaymentService {
	return &e
}

func (e *GetInfoPaymentPaymentService) UnmarshalJSON(data []byte) error {
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
		*e = GetInfoPaymentPaymentService(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetInfoPaymentPaymentService: %v", v)
	}
}

type GetInfoPaymentRequest struct {
	// Es contenido si viaja el campo Signature. Ej: Digest: SHA-256=NzdmZjA4YjY5M2M2NDYyMmVjOWFmMGNmYTZiNTU3MjVmNDI4NTRlMzJkYzE3ZmNmMDE3ZGFmMjhhNTc5OTU3OQ==
	Digest string `header:"style=simple,explode=false,name=Digest"`
	// Accept header de la petición HTTP entre PSU y el TPP. Ej: PSU-Accept: application/json
	PSUAccept *string `header:"style=simple,explode=false,name=PSU-Accept"`
	// Accept charset header de la petición HTTP entre PSU y el TPP. PSU-Accept-Charset: utf-8
	PSUAcceptCharset *string `header:"style=simple,explode=false,name=PSU-Accept-Charset"`
	// Accept encoding header de la petición HTTP entre PSU y el TPP. PSU-Accept-Encoding: gzip
	PSUAcceptEncoding *string `header:"style=simple,explode=false,name=PSU-Accept-Encoding"`
	// Accept language header de la petición HTTP entre PSU y el TPP. PSU-Accept-Language: es-ES
	PSUAcceptLanguage *string `header:"style=simple,explode=false,name=PSU-Accept-Language"`
	// UUID (Universally Unique Identifier) para un dispositivo. El UUID identifica al dispositivo o a una instalación de una aplicación en un dispositivo. Este ID no debe ser modificado hasta la desinstalación de la aplicación del dispositivo. Ej: PSU-Device-ID: 5b3ab8e8-0fd5-43d2-946e-d75958b172e7
	PSUDeviceID *string `header:"style=simple,explode=false,name=PSU-Device-ID"`
	// Localización correspondiente a la petición HTTP entre el PSU y el TPP. Ej: PSU-Geo-Location: GEO:90.023856;25.345963
	PSUGeoLocation *string `header:"style=simple,explode=false,name=PSU-Geo-Location"`
	// Método HTTP usado en la interfaz entre PSU y TPP. Valores permitidos: GET. Ej: PSU-Http-Method: GET
	PSUHTTPMethod *string `default:"GET" header:"style=simple,explode=false,name=PSU-Http-Method"`
	// Dirección IP de la petición HTPP entre el PSU y el TPP. Si no está disponible, el TPP debe usar la dirección IP usada por el TPP cuando envía esta petición. Ej: Ej: PSU-IP-Address: 192.168.16.5
	PSUIPAddress string `header:"style=simple,explode=false,name=PSU-IP-Address"`
	// Puerto IP de la petición HTTP entre el PSU y el TPP si está disponible. Ejemplo: PSU-IP-Port: 443
	PSUIPPort *string `header:"style=simple,explode=false,name=PSU-IP-Port"`
	// Navegador o sistema operativo de la petición HTTP entre el PSU y el TPP. Ejemplo: PSU-User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
	PSUUserAgent *string `header:"style=simple,explode=false,name=PSU-User-Agent"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	// Nombre del ASPSP al que desea realizar la petición.
	Aspsp string `pathParam:"style=simple,explode=false,name=aspsp"`
	// Identificador del recurso que referencia a la iniciación de pago. Enviado previamente como respuesta a un mensaje de iniciación de pago del TPP al HUB.
	PaymentID string `pathParam:"style=simple,explode=false,name=payment-id"`
	// Producto de pago a usar.
	PaymentProduct GetInfoPaymentPaymentProduct `pathParam:"style=simple,explode=false,name=payment-product"`
	// Producto de pago a usar.
	PaymentService GetInfoPaymentPaymentService `pathParam:"style=simple,explode=false,name=payment-service"`
}

func (g GetInfoPaymentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetInfoPaymentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetInfoPaymentRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetInfoPaymentRequest) GetPSUAccept() *string {
	if o == nil {
		return nil
	}
	return o.PSUAccept
}

func (o *GetInfoPaymentRequest) GetPSUAcceptCharset() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptCharset
}

func (o *GetInfoPaymentRequest) GetPSUAcceptEncoding() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptEncoding
}

func (o *GetInfoPaymentRequest) GetPSUAcceptLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PSUAcceptLanguage
}

func (o *GetInfoPaymentRequest) GetPSUDeviceID() *string {
	if o == nil {
		return nil
	}
	return o.PSUDeviceID
}

func (o *GetInfoPaymentRequest) GetPSUGeoLocation() *string {
	if o == nil {
		return nil
	}
	return o.PSUGeoLocation
}

func (o *GetInfoPaymentRequest) GetPSUHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.PSUHTTPMethod
}

func (o *GetInfoPaymentRequest) GetPSUIPAddress() string {
	if o == nil {
		return ""
	}
	return o.PSUIPAddress
}

func (o *GetInfoPaymentRequest) GetPSUIPPort() *string {
	if o == nil {
		return nil
	}
	return o.PSUIPPort
}

func (o *GetInfoPaymentRequest) GetPSUUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.PSUUserAgent
}

func (o *GetInfoPaymentRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetInfoPaymentRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetInfoPaymentRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetInfoPaymentRequest) GetAspsp() string {
	if o == nil {
		return ""
	}
	return o.Aspsp
}

func (o *GetInfoPaymentRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

func (o *GetInfoPaymentRequest) GetPaymentProduct() GetInfoPaymentPaymentProduct {
	if o == nil {
		return GetInfoPaymentPaymentProduct("")
	}
	return o.PaymentProduct
}

func (o *GetInfoPaymentRequest) GetPaymentService() GetInfoPaymentPaymentService {
	if o == nil {
		return GetInfoPaymentPaymentService("")
	}
	return o.PaymentService
}

type GetInfoPaymentResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetInfoPaymentResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetInfoPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInfoPaymentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetInfoPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInfoPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
