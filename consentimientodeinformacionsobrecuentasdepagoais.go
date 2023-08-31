// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package psd2cajarural

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"io"
	"net/http"
)

type consentimientoDeInformacionSobreCuentasDePagoAIS struct {
	sdkConfiguration sdkConfiguration
}

func newConsentimientoDeInformacionSobreCuentasDePagoAIS(sdkConfig sdkConfiguration) *consentimientoDeInformacionSobreCuentasDePagoAIS {
	return &consentimientoDeInformacionSobreCuentasDePagoAIS{
		sdkConfiguration: sdkConfig,
	}
}

// PostConsents - Solicitud de consentimiento AIS
// Con este servicio, un TPP a través del HUB puede solicitar un consentimiento para acceder a las cuentas del PSU. Esta solicitud puede ser sobre unas cuentas indicadas o no.
func (s *consentimientoDeInformacionSobreCuentasDePagoAIS) PostConsents(ctx context.Context, request operations.PostConsentsRequest) (*operations.PostConsentsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v1.1/consents", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "raw")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostConsentsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `*/*`):
			res.Body = rawBody
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
