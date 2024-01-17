// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package tjsolerpsd2

import (
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https:///api-entrada-xs2a",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Psd2CajaRural - API para PSD2: API de Redsys para PSD2
type Psd2CajaRural struct {
	ServiciosMultibanco                              *ServiciosMultibanco
	LecturaDeListadoDeCuentas                        *LecturaDeListadoDeCuentas
	LecturaDeDetallesDeCuenta                        *LecturaDeDetallesDeCuenta
	LecturaDeBalances                                *LecturaDeBalances
	LecturaDeTransacciones                           *LecturaDeTransacciones
	InicioDePago                                     *InicioDePago
	ConsentimientoDeInformacionSobreCuentasDePagoAIS *ConsentimientoDeInformacionSobreCuentasDePagoAIS
	EliminarConsentimiento                           *EliminarConsentimiento
	RecuperarInformacionDeConsentimiento             *RecuperarInformacionDeConsentimiento
	ObtenerSubRecursosDeLaAutorizacion               *ObtenerSubRecursosDeLaAutorizacion
	InicioDelProcesoDeAutorizacionExplicita          *InicioDelProcesoDeAutorizacionExplicita
	ObtenerEstadoSCA                                 *ObtenerEstadoSCA
	ActualizarDatosDelPSUSeleccionarMetodoSCA        *ActualizarDatosDelPSUSeleccionarMetodoSCA
	ObtenerEstadoDeConsentimiento                    *ObtenerEstadoDeConsentimiento
	ConsultaDeFondos                                 *ConsultaDeFondos
	ObtenerListadoDeBeneficiariosDeConfianza         *ObtenerListadoDeBeneficiariosDeConfianza
	CancelarInicioDePago                             *CancelarInicioDePago
	RecuperarInformacionDelInicioDePago              *RecuperarInformacionDelInicioDePago
	ObtenerEstadoDelPago                             *ObtenerEstadoDelPago
	ConsentimientoDeConfirmacionDeFondosFCS          *ConsentimientoDeConfirmacionDeFondosFCS

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Psd2CajaRural)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Psd2CajaRural) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Psd2CajaRural) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Psd2CajaRural) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Psd2CajaRural) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Psd2CajaRural) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Psd2CajaRural {
	sdk := &Psd2CajaRural{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.1",
			SDKVersion:        "0.6.2",
			GenVersion:        "2.234.3",
			UserAgent:         "speakeasy-sdk/go 0.6.2 2.234.3 1.1 github.com/speakeasy-sdks/tjsoler-psd2",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.ServiciosMultibanco = newServiciosMultibanco(sdk.sdkConfiguration)

	sdk.LecturaDeListadoDeCuentas = newLecturaDeListadoDeCuentas(sdk.sdkConfiguration)

	sdk.LecturaDeDetallesDeCuenta = newLecturaDeDetallesDeCuenta(sdk.sdkConfiguration)

	sdk.LecturaDeBalances = newLecturaDeBalances(sdk.sdkConfiguration)

	sdk.LecturaDeTransacciones = newLecturaDeTransacciones(sdk.sdkConfiguration)

	sdk.InicioDePago = newInicioDePago(sdk.sdkConfiguration)

	sdk.ConsentimientoDeInformacionSobreCuentasDePagoAIS = newConsentimientoDeInformacionSobreCuentasDePagoAIS(sdk.sdkConfiguration)

	sdk.EliminarConsentimiento = newEliminarConsentimiento(sdk.sdkConfiguration)

	sdk.RecuperarInformacionDeConsentimiento = newRecuperarInformacionDeConsentimiento(sdk.sdkConfiguration)

	sdk.ObtenerSubRecursosDeLaAutorizacion = newObtenerSubRecursosDeLaAutorizacion(sdk.sdkConfiguration)

	sdk.InicioDelProcesoDeAutorizacionExplicita = newInicioDelProcesoDeAutorizacionExplicita(sdk.sdkConfiguration)

	sdk.ObtenerEstadoSCA = newObtenerEstadoSCA(sdk.sdkConfiguration)

	sdk.ActualizarDatosDelPSUSeleccionarMetodoSCA = newActualizarDatosDelPSUSeleccionarMetodoSCA(sdk.sdkConfiguration)

	sdk.ObtenerEstadoDeConsentimiento = newObtenerEstadoDeConsentimiento(sdk.sdkConfiguration)

	sdk.ConsultaDeFondos = newConsultaDeFondos(sdk.sdkConfiguration)

	sdk.ObtenerListadoDeBeneficiariosDeConfianza = newObtenerListadoDeBeneficiariosDeConfianza(sdk.sdkConfiguration)

	sdk.CancelarInicioDePago = newCancelarInicioDePago(sdk.sdkConfiguration)

	sdk.RecuperarInformacionDelInicioDePago = newRecuperarInformacionDelInicioDePago(sdk.sdkConfiguration)

	sdk.ObtenerEstadoDelPago = newObtenerEstadoDelPago(sdk.sdkConfiguration)

	sdk.ConsentimientoDeConfirmacionDeFondosFCS = newConsentimientoDeConfirmacionDeFondosFCS(sdk.sdkConfiguration)

	return sdk
}
