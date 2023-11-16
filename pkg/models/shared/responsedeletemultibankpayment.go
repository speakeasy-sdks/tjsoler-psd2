// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Links - Lista de hipervinculos
type Links struct {
}

// ChallengeData - NO SOPORTADO PARA ESTA VERSIÓN
type ChallengeData struct {
}

// ChosenScaMethod - NO SOPORTADO PARA ESTA VERSIÓN
type ChosenScaMethod struct {
}

// ScaMethods - Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación.
type ScaMethods struct {
}

// TppMessages - Mensaje para el TPP enviado a través del HUB.
type TppMessages struct {
}

type ResponseDeleteMultibankPayment struct {
	// Lista de hipervinculos
	Links *Links `json:"_links,omitempty"`
	// NO SOPORTADO PARA ESTA VERSIÓN
	ChallengeData *ChallengeData `json:"challengeData,omitempty"`
	// NO SOPORTADO PARA ESTA VERSIÓN
	ChosenScaMethod *ChosenScaMethod `json:"chosenScaMethod,omitempty"`
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string `json:"psuMessage,omitempty"`
	// Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación.
	ScaMethods *ScaMethods `json:"scaMethods,omitempty"`
	// Mensaje para el TPP enviado a través del HUB.
	TppMessages *TppMessages `json:"tppMessages,omitempty"`
	// Estado de la transacción
	TransactionStatus string `json:"transactionStatus"`
}

func (o *ResponseDeleteMultibankPayment) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *ResponseDeleteMultibankPayment) GetChallengeData() *ChallengeData {
	if o == nil {
		return nil
	}
	return o.ChallengeData
}

func (o *ResponseDeleteMultibankPayment) GetChosenScaMethod() *ChosenScaMethod {
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

func (o *ResponseDeleteMultibankPayment) GetScaMethods() *ScaMethods {
	if o == nil {
		return nil
	}
	return o.ScaMethods
}

func (o *ResponseDeleteMultibankPayment) GetTppMessages() *TppMessages {
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
