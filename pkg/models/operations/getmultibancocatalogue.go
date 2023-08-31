// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetMultibancoCatalogueRequest struct {
	// Es contenido si viaja el campo Signature. Ej: Digest: SHA-256=NzdmZjA4YjY5M2M2NDYyMmVjOWFmMGNmYTZiNTU3MjVmNDI4NTRlMzJkYzE3ZmNmMDE3ZGFmMjhhNTc5OTU3OQ==
	Digest string `header:"style=simple,explode=false,name=Digest"`
	// Firma de la petición por el TPP.
	Signature string `header:"style=simple,explode=false,name=Signature"`
	// Certificado del TPP usado para firmar la petición, en base64, sin cabecera, pie ni saltos de linea. Ej: TPP-Signature-Certificate: MIIHgzCCBmugAwIBAgIIZzZvBQlt0UcwDQYJ………….KoZIhvcNAQELBQAwSTELMAkGA1UEBhMCVVMxEzARBgNVBA
	TPPSignatureCertificate string `header:"style=simple,explode=false,name=TPP-Signature-Certificate"`
	// Identificador único de la transacción asignado por el TPP. Ej: X-Request-ID: 1b3ab8e8-0fd5-43d2-946e-d75958b172e7
	XRequestID string `header:"style=simple,explode=false,name=X-Request-ID"`
	// Nombre del ASPSP al que desea realizar la petición.
	AspName string `pathParam:"style=simple,explode=false,name=asp-name"`
	// Importe de pago
	InstructedAmount *string `queryParam:"style=form,explode=true,name=instructedAmount"`
	// Tipo de pago multibanco
	MultibancoPaymentType string `pathParam:"style=simple,explode=false,name=multibanco-payment-type"`
	// Referencia de la operación
	PaymentReference *string `queryParam:"style=form,explode=true,name=paymentReference"`
	// Fecha de la ejecución del pago
	RequestedExecutionDate *string `queryParam:"style=form,explode=true,name=requestedExecutionDate"`
}

func (o *GetMultibancoCatalogueRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetMultibancoCatalogueRequest) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *GetMultibancoCatalogueRequest) GetTPPSignatureCertificate() string {
	if o == nil {
		return ""
	}
	return o.TPPSignatureCertificate
}

func (o *GetMultibancoCatalogueRequest) GetXRequestID() string {
	if o == nil {
		return ""
	}
	return o.XRequestID
}

func (o *GetMultibancoCatalogueRequest) GetAspName() string {
	if o == nil {
		return ""
	}
	return o.AspName
}

func (o *GetMultibancoCatalogueRequest) GetInstructedAmount() *string {
	if o == nil {
		return nil
	}
	return o.InstructedAmount
}

func (o *GetMultibancoCatalogueRequest) GetMultibancoPaymentType() string {
	if o == nil {
		return ""
	}
	return o.MultibancoPaymentType
}

func (o *GetMultibancoCatalogueRequest) GetPaymentReference() *string {
	if o == nil {
		return nil
	}
	return o.PaymentReference
}

func (o *GetMultibancoCatalogueRequest) GetRequestedExecutionDate() *string {
	if o == nil {
		return nil
	}
	return o.RequestedExecutionDate
}

type GetMultibancoCatalogueResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetMultibancoCatalogueResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetMultibancoCatalogueResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMultibancoCatalogueResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMultibancoCatalogueResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
