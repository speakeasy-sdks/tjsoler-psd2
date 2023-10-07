// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package tjsolerpsd2

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/shared"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"io"
	"net/http"
)

type obtenerEstadoSCA struct {
	sdkConfiguration sdkConfiguration
}

func newObtenerEstadoSCA(sdkConfig sdkConfiguration) *obtenerEstadoSCA {
	return &obtenerEstadoSCA{
		sdkConfiguration: sdkConfig,
	}
}

// GetEstadoSCAAutorizacionCancelacionPago - Endpoint en caso de obtener el Estado SCA para Autorización de Cancelación de Pago
// Endpoint en caso de obtener el Estado SCA para Autorización de Cancelación de Pago
func (s *obtenerEstadoSCA) GetEstadoSCAAutorizacionCancelacionPago(ctx context.Context, request operations.GetEstadoSCAAutorizacionCancelacionPagoRequest) (*operations.GetEstadoSCAAutorizacionCancelacionPagoResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v1.1/{payment-service}/{payment-product}/{payment-id}/cancellation-authorisations/{authorisation-id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEstadoSCAAutorizacionCancelacionPagoResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ResponseGetSCAStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ResponseGetSCAStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetEstadoSCAAutorizacionConsentimientosAIS - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS
// Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos AIS
func (s *obtenerEstadoSCA) GetEstadoSCAAutorizacionConsentimientosAIS(ctx context.Context, request operations.GetEstadoSCAAutorizacionConsentimientosAISRequest) (*operations.GetEstadoSCAAutorizacionConsentimientosAISResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v1.1/consents/{consent-id}/authorisations/{authorisation-id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEstadoSCAAutorizacionConsentimientosAISResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ResponseGetSCAStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ResponseGetSCAStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetEstadoSCAAutorizacionConsentimientosFCS - Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS
// Endpoint en caso de obtener el Estado SCA para Autorización de Consentimientos FCS
func (s *obtenerEstadoSCA) GetEstadoSCAAutorizacionConsentimientosFCS(ctx context.Context, request operations.GetEstadoSCAAutorizacionConsentimientosFCSRequest) (*operations.GetEstadoSCAAutorizacionConsentimientosFCSResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v2.1/consents/confirmation-of-funds/{consent-id}/authorisations/{authorisation-id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEstadoSCAAutorizacionConsentimientosFCSResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ResponseGetSCAStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ResponseGetSCAStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetEstadoSCAAutorizacionInicioPago - Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago
// Endpoint en caso de obtener el Estado SCA para Autorización de Inicio de Pago
func (s *obtenerEstadoSCA) GetEstadoSCAAutorizacionInicioPago(ctx context.Context, request operations.GetEstadoSCAAutorizacionInicioPagoRequest) (*operations.GetEstadoSCAAutorizacionInicioPagoResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v1.1/{payment-service}/{payment-product}/{payment-id}/authorisations/{authorisation-id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEstadoSCAAutorizacionInicioPagoResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ResponseGetSCAStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ResponseGetSCAStatus = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
