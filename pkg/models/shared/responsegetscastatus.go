// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseGetSCAStatusTppMessages - Mensaje para el TPP enviado a través del HUB.
type ResponseGetSCAStatusTppMessages struct {
}

// ResponseGetSCAStatus - HTTP/1.1 200 Ok
type ResponseGetSCAStatus struct {
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string `json:"psuMessage,omitempty"`
	// Estado SCA
	ScaStatus string `json:"scaStatus"`
	// Mensaje para el TPP enviado a través del HUB.
	TppMessages *ResponseGetSCAStatusTppMessages `json:"tppMessages,omitempty"`
	// Con este flag el ASPSP opcionalmente podría comunicar al TPP que el creditor formaba parte del listado de beneficiarios de confianza. Este atributo solo se contiene en caso de un estado final del scaStatus.
	TrustedBeneficiaryFlag *bool `json:"trustedBeneficiaryFlag,omitempty"`
}

func (o *ResponseGetSCAStatus) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseGetSCAStatus) GetScaStatus() string {
	if o == nil {
		return ""
	}
	return o.ScaStatus
}

func (o *ResponseGetSCAStatus) GetTppMessages() *ResponseGetSCAStatusTppMessages {
	if o == nil {
		return nil
	}
	return o.TppMessages
}

func (o *ResponseGetSCAStatus) GetTrustedBeneficiaryFlag() *bool {
	if o == nil {
		return nil
	}
	return o.TrustedBeneficiaryFlag
}
