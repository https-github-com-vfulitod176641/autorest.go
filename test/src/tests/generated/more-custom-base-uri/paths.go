package morecustombaseurigroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PathsClient is the test Infrastructure for AutoRest
type PathsClient struct {
	BaseClient
}

// NewPathsClient creates an instance of the PathsClient client.
func NewPathsClient(subscriptionID string) PathsClient {
	return PathsClient{New(subscriptionID)}
}

// GetEmpty get a 200 to test a valid base uri
// Parameters:
// vault - the vault name, e.g. https://myvault
// secret - secret value.
// keyName - the key name with value 'key1'.
// keyVersion - the key version. Default value 'v1'.
func (client PathsClient) GetEmpty(ctx context.Context, vault string, secret string, keyName string, keyVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PathsClient.GetEmpty")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEmptyPreparer(ctx, vault, secret, keyName, keyVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "morecustombaseurigroup.PathsClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "morecustombaseurigroup.PathsClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "morecustombaseurigroup.PathsClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client PathsClient) GetEmptyPreparer(ctx context.Context, vault string, secret string, keyName string, keyVersion string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"dnsSuffix": client.DNSSuffix,
		"secret":    secret,
		"vault":     vault,
	}

	pathParameters := map[string]interface{}{
		"keyName":        autorest.Encode("path", keyName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{}
	if len(keyVersion) > 0 {
		queryParameters["keyVersion"] = autorest.Encode("query", keyVersion)
	} else {
		queryParameters["keyVersion"] = autorest.Encode("query", "v1")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vault}{secret}{dnsSuffix}", urlParameters),
		autorest.WithPathParameters("/customuri/{subscriptionId}/{keyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client PathsClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client PathsClient) GetEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
