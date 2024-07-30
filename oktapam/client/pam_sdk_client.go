package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/atko-pam/pam-sdk-go/client/pam"
)

type SDKClientWrapper struct {
	SDKClient *pam.APIClient
	Team      string
}

func checkStatusCodeFromSDK(resp *http.Response, allowed ...int) (int, error) {
	received := resp.StatusCode
	for _, c := range allowed {
		if received == c {
			return received, nil
		}
	}

	return received, createErrorForInvalidCodeFromSDK(resp, allowed...)
}

func createErrorForInvalidCodeFromSDK(resp *http.Response, allowed ...int) error {
	received := resp.StatusCode
	if received == 403 {
		// a 403 will likely have UAMs in the body, which we don't want to show
		return fmt.Errorf("not authorized")
	}

	var body string
	if bodyBytes, err := io.ReadAll(resp.Body); err == nil {
		body = string(bodyBytes)
	} else {
		body = "could not read response body"
	}
	resp.Body.Close()

	if len(allowed) == 1 {
		return fmt.Errorf("call resulted in status of %d, expected a %d.\nResponse Body: %s", received, allowed[0], body)
	}

	return fmt.Errorf("call resulted in status of %d, expected one of %v.\nResponse Body: %s", received, allowed, body)
}
