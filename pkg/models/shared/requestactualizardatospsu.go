// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RequestActualizarDatosPsu struct {
	// Identificador del método de autenticación.
	AuthenticationMethodID *string `json:"authenticationMethodId,omitempty"`
}

func (o *RequestActualizarDatosPsu) GetAuthenticationMethodID() *string {
	if o == nil {
		return nil
	}
	return o.AuthenticationMethodID
}