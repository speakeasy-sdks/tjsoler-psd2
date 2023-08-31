// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseGetSubResourcesAuthorizationTppMessages - Mensaje para el TPP enviado a través del HUB.
type ResponseGetSubResourcesAuthorizationTppMessages struct {
}

// ResponseGetSubResourcesAuthorization - HTTP/1.1 200 Ok
type ResponseGetSubResourcesAuthorization struct {
	// Array de authorisationIds.
	AuthorisationIds []string `json:"authorisationIds,omitempty"`
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string `json:"psuMessage,omitempty"`
	// Mensaje para el TPP enviado a través del HUB.
	TppMessages *ResponseGetSubResourcesAuthorizationTppMessages `json:"tppMessages,omitempty"`
}

func (o *ResponseGetSubResourcesAuthorization) GetAuthorisationIds() []string {
	if o == nil {
		return nil
	}
	return o.AuthorisationIds
}

func (o *ResponseGetSubResourcesAuthorization) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseGetSubResourcesAuthorization) GetTppMessages() *ResponseGetSubResourcesAuthorizationTppMessages {
	if o == nil {
		return nil
	}
	return o.TppMessages
}
