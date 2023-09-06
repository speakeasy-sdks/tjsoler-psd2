// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RequestGetConsentAccess struct {
}

type RequestGetConsent struct {
	Access RequestGetConsentAccess
	// Indica que un inicio de pago se realizará en la misma sesión. De momento el HUB no implementará servicios combinados por lo que este campo deberá tomar el valor de falso.
	CombinedServiceIndicator bool
	// Indica la frecuencia de acceso a la cuenta por día. 1 si es de un solo acceso.
	FrequencyPerDay int64
	// true: acceso recurrente a la cuenta. false: un solo acceso.
	RecurringIndicator bool
	// Fecha hasta la que el consentimiento solicita acceso. Para crear el consentimiento con el máximo tiempo de acceso posible se debe usar el valor: 9999-12-31 Cuando se recupere el consentimiento, la fecha máxima posible vendrá ajustada. La validez también incluye hasta la fecha solicitada.
	ValidUntil string
}

func (o *RequestGetConsent) GetAccess() RequestGetConsentAccess {
	if o == nil {
		return RequestGetConsentAccess{}
	}
	return o.Access
}

func (o *RequestGetConsent) GetCombinedServiceIndicator() bool {
	if o == nil {
		return false
	}
	return o.CombinedServiceIndicator
}

func (o *RequestGetConsent) GetFrequencyPerDay() int64 {
	if o == nil {
		return 0
	}
	return o.FrequencyPerDay
}

func (o *RequestGetConsent) GetRecurringIndicator() bool {
	if o == nil {
		return false
	}
	return o.RecurringIndicator
}

func (o *RequestGetConsent) GetValidUntil() string {
	if o == nil {
		return ""
	}
	return o.ValidUntil
}