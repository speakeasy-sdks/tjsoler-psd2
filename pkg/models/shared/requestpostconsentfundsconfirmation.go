// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Account - Cuenta sobre la que se realizará la consulta de fondos.
type Account struct {
}

type RequestPostConsentFundsConfirmation struct {
	// Cuenta sobre la que se realizará la consulta de fondos.
	Account Account
	// Fecha de caducidad de la tarjeta emitida por el PIISP. (Formato ISODate)
	CardExpiryDate *string
	// Explicación adicional del producto.
	CardInformation *string
	// Número de tarjeta de la tarjeta emitida por el PIISP. Debe ser enviada si está disponible.
	CardNumber *string
	// Información adicional acerca del proceso de registro para el PSU. Por ej. una referencia al contrato entre TPP/PSU
	RegistrationInformation *string
}

func (o *RequestPostConsentFundsConfirmation) GetAccount() Account {
	if o == nil {
		return Account{}
	}
	return o.Account
}

func (o *RequestPostConsentFundsConfirmation) GetCardExpiryDate() *string {
	if o == nil {
		return nil
	}
	return o.CardExpiryDate
}

func (o *RequestPostConsentFundsConfirmation) GetCardInformation() *string {
	if o == nil {
		return nil
	}
	return o.CardInformation
}

func (o *RequestPostConsentFundsConfirmation) GetCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.CardNumber
}

func (o *RequestPostConsentFundsConfirmation) GetRegistrationInformation() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationInformation
}
