// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseDeleteMultibankPaymentLinks - Lista de hipervinculos
type ResponseDeleteMultibankPaymentLinks struct {
}

// ResponseDeleteMultibankPaymentChallengeData - NO SOPORTADO PARA ESTA VERSIÓN
type ResponseDeleteMultibankPaymentChallengeData struct {
}

// ResponseDeleteMultibankPaymentChosenScaMethod - NO SOPORTADO PARA ESTA VERSIÓN
type ResponseDeleteMultibankPaymentChosenScaMethod struct {
}

// ResponseDeleteMultibankPaymentScaMethods - Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación.
type ResponseDeleteMultibankPaymentScaMethods struct {
}

// ResponseDeleteMultibankPaymentTppMessages - Mensaje para el TPP enviado a través del HUB.
type ResponseDeleteMultibankPaymentTppMessages struct {
}

// ResponseDeleteMultibankPayment - Deleted
type ResponseDeleteMultibankPayment struct {
	// Lista de hipervinculos
	Links *ResponseDeleteMultibankPaymentLinks `json:"_links,omitempty"`
	// NO SOPORTADO PARA ESTA VERSIÓN
	ChallengeData *ResponseDeleteMultibankPaymentChallengeData `json:"challengeData,omitempty"`
	// NO SOPORTADO PARA ESTA VERSIÓN
	ChosenScaMethod *ResponseDeleteMultibankPaymentChosenScaMethod `json:"chosenScaMethod,omitempty"`
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string `json:"psuMessage,omitempty"`
	// Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación.
	ScaMethods *ResponseDeleteMultibankPaymentScaMethods `json:"scaMethods,omitempty"`
	// Mensaje para el TPP enviado a través del HUB.
	TppMessages *ResponseDeleteMultibankPaymentTppMessages `json:"tppMessages,omitempty"`
	// Estado de la transacción
	TransactionStatus string `json:"transactionStatus"`
}

func (o *ResponseDeleteMultibankPayment) GetLinks() *ResponseDeleteMultibankPaymentLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *ResponseDeleteMultibankPayment) GetChallengeData() *ResponseDeleteMultibankPaymentChallengeData {
	if o == nil {
		return nil
	}
	return o.ChallengeData
}

func (o *ResponseDeleteMultibankPayment) GetChosenScaMethod() *ResponseDeleteMultibankPaymentChosenScaMethod {
	if o == nil {
		return nil
	}
	return o.ChosenScaMethod
}

func (o *ResponseDeleteMultibankPayment) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseDeleteMultibankPayment) GetScaMethods() *ResponseDeleteMultibankPaymentScaMethods {
	if o == nil {
		return nil
	}
	return o.ScaMethods
}

func (o *ResponseDeleteMultibankPayment) GetTppMessages() *ResponseDeleteMultibankPaymentTppMessages {
	if o == nil {
		return nil
	}
	return o.TppMessages
}

func (o *ResponseDeleteMultibankPayment) GetTransactionStatus() string {
	if o == nil {
		return ""
	}
	return o.TransactionStatus
}
