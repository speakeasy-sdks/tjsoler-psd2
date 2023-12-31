// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseMultibankAuthorizationSubresourcesTppMessages - Mensaje para el TPP enviado a través del HUB.
type ResponseMultibankAuthorizationSubresourcesTppMessages struct {
}

type ResponseMultibankAuthorizationSubresources struct {
	// Array de authorisationIds.
	AuthorisationIds []string `json:"authorisationIds,omitempty"`
	// Array de cancellationIds conectados al recurso de pago. Nota: obligatorio si se trata de una cancelación
	CancellationIds []string `json:"cancellationIds,omitempty"`
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string `json:"psuMessage,omitempty"`
	// Mensaje para el TPP enviado a través del HUB.
	TppMessages *ResponseMultibankAuthorizationSubresourcesTppMessages `json:"tppMessages,omitempty"`
}

func (o *ResponseMultibankAuthorizationSubresources) GetAuthorisationIds() []string {
	if o == nil {
		return nil
	}
	return o.AuthorisationIds
}

func (o *ResponseMultibankAuthorizationSubresources) GetCancellationIds() []string {
	if o == nil {
		return nil
	}
	return o.CancellationIds
}

func (o *ResponseMultibankAuthorizationSubresources) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseMultibankAuthorizationSubresources) GetTppMessages() *ResponseMultibankAuthorizationSubresourcesTppMessages {
	if o == nil {
		return nil
	}
	return o.TppMessages
}
