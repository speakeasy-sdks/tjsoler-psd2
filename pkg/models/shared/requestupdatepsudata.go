// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RequestUpdatePSUData struct {
	// Identificador del método de autenticación
	AuthenticationMethodID string `json:"authenticationMethodId"`
}

func (o *RequestUpdatePSUData) GetAuthenticationMethodID() string {
	if o == nil {
		return ""
	}
	return o.AuthenticationMethodID
}
