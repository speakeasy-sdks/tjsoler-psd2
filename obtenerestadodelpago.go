// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package tjsolerpsd2

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/tjsoler-psd2/internal/hooks"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/operations"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/tjsoler-psd2/pkg/utils"
	"io"
	"net/http"
)

type ObtenerEstadoDelPago struct {
	sdkConfiguration sdkConfiguration
}

func newObtenerEstadoDelPago(sdkConfig sdkConfiguration) *ObtenerEstadoDelPago {
	return &ObtenerEstadoDelPago{
		sdkConfiguration: sdkConfig,
	}
}

// StatusPayment - Obtener información del Estado de pago
// Este mensaje es enviado por el TPP al HUB para solicitar información del estado en el que se encuentra la iniciación de pago que solicitó el TPP
func (s *ObtenerEstadoDelPago) StatusPayment(ctx context.Context, request operations.StatusPaymentRequest) (*operations.StatusPaymentResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "statusPayment"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api-entrada-xs2a/services/{aspsp}/v1.1/{payment-service}/{payment-product}/{payment-id}/status", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.StatusPaymentResponse{
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
		case utils.MatchContentType(contentType, `*/*`):
			res.Body = rawBody
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
