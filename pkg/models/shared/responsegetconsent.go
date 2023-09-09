// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseGetConsentLinks - Lista de hipervínculos para ser reconocidos por el HUB. Tipos soportados en esta respuesta: scaRedirect: en caso de SCA por redirección. Link donde el navegador del PSU debe ser redireccionado por el Hub. startAuthorisation: en caso de que un inicio explícito de la autorización de la transacción sea necesario (no hay selección del método SCA) startAuthorisationWithAuthenticationMethodSelection: link al end-point de autorización donde el sub-recurso de autorización tiene que ser generado mientras se selecciona el método SCA. Este enlace es contenido bajo las mismas condiciones que el campo "scaMethods" self: link al recurso de inicio de pago creado por esta petición. status: link para recuperar el estado de la transacción del inicio de pago. scaStatus: link para consultar el estado SCA correspondiente al sub-recurso de autorización. Este link es solo contenido si un sub-recurso de autorización ha sido creado. (Si no hay autorización explícita este enlace debe venir obligatoriamente)
type ResponseGetConsentLinks struct {
}

// ResponseGetConsentChoosenScaMethod - NO SOPORTADO PARA ESTA VERSIÓN
type ResponseGetConsentChoosenScaMethod struct {
}

// ResponseGetConsentScaMethods - Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación. Si este dato es contenido también se informará el link "startAuthorisationWithAuthenticationMethodSelection”. Esto métodos deberán ser presentados al PSU. Nota: Solo si ASPSP soporta selección del método SCA
type ResponseGetConsentScaMethods struct {
}

// ResponseGetConsentTppMessage - Mensaje para el TPP enviado a través del HUB.
type ResponseGetConsentTppMessage struct {
}

type ResponseGetConsent struct {
	// Lista de hipervínculos para ser reconocidos por el HUB. Tipos soportados en esta respuesta: scaRedirect: en caso de SCA por redirección. Link donde el navegador del PSU debe ser redireccionado por el Hub. startAuthorisation: en caso de que un inicio explícito de la autorización de la transacción sea necesario (no hay selección del método SCA) startAuthorisationWithAuthenticationMethodSelection: link al end-point de autorización donde el sub-recurso de autorización tiene que ser generado mientras se selecciona el método SCA. Este enlace es contenido bajo las mismas condiciones que el campo "scaMethods" self: link al recurso de inicio de pago creado por esta petición. status: link para recuperar el estado de la transacción del inicio de pago. scaStatus: link para consultar el estado SCA correspondiente al sub-recurso de autorización. Este link es solo contenido si un sub-recurso de autorización ha sido creado. (Si no hay autorización explícita este enlace debe venir obligatoriamente)
	Links ResponseGetConsentLinks
	// NO SOPORTADO PARA ESTA VERSIÓN
	ChoosenScaMethod *ResponseGetConsentChoosenScaMethod
	// Identificador del recurso que referencia al consentimiento. Debe ser contenido si se generó un consentimiento.
	ConsentID string
	// Estado de autenticación del consentimiento.  Valores definidos en anexos.
	ConsentStatus string
	// Texto enviado al TPP a través del HUB para ser mostrado al PSU.
	PsuMessage *string
	// Este elemento es contenido si SCA es requerido y si el PSU puede elegir entre diferentes métodos de autenticación. Si este dato es contenido también se informará el link "startAuthorisationWithAuthenticationMethodSelection”. Esto métodos deberán ser presentados al PSU. Nota: Solo si ASPSP soporta selección del método SCA
	ScaMethods *ResponseGetConsentScaMethods
	// Mensaje para el TPP enviado a través del HUB.
	TppMessage *ResponseGetConsentTppMessage
}

func (o *ResponseGetConsent) GetLinks() ResponseGetConsentLinks {
	if o == nil {
		return ResponseGetConsentLinks{}
	}
	return o.Links
}

func (o *ResponseGetConsent) GetChoosenScaMethod() *ResponseGetConsentChoosenScaMethod {
	if o == nil {
		return nil
	}
	return o.ChoosenScaMethod
}

func (o *ResponseGetConsent) GetConsentID() string {
	if o == nil {
		return ""
	}
	return o.ConsentID
}

func (o *ResponseGetConsent) GetConsentStatus() string {
	if o == nil {
		return ""
	}
	return o.ConsentStatus
}

func (o *ResponseGetConsent) GetPsuMessage() *string {
	if o == nil {
		return nil
	}
	return o.PsuMessage
}

func (o *ResponseGetConsent) GetScaMethods() *ResponseGetConsentScaMethods {
	if o == nil {
		return nil
	}
	return o.ScaMethods
}

func (o *ResponseGetConsent) GetTppMessage() *ResponseGetConsentTppMessage {
	if o == nil {
		return nil
	}
	return o.TppMessage
}
