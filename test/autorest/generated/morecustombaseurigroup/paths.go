// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package morecustombaseurigroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PathsOperations contains the methods for the Paths group.
type PathsOperations interface {
	// GetEmpty - Get a 200 to test a valid base uri
	GetEmpty(ctx context.Context, vault string, secret string, keyName string, pathsGetEmptyOptions *PathsGetEmptyOptions) (*http.Response, error)
}

// pathsOperations implements the PathsOperations interface.
type pathsOperations struct {
	*Client
	dnsSuffix      string
	subscriptionID string
}

// GetEmpty - Get a 200 to test a valid base uri
func (client *pathsOperations) GetEmpty(ctx context.Context, vault string, secret string, keyName string, pathsGetEmptyOptions *PathsGetEmptyOptions) (*http.Response, error) {
	req, err := client.getEmptyCreateRequest(vault, secret, keyName, pathsGetEmptyOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *pathsOperations) getEmptyCreateRequest(vault string, secret string, keyName string, pathsGetEmptyOptions *PathsGetEmptyOptions) (*azcore.Request, error) {
	host := "{vault}{secret}{dnsSuffix}"
	host = strings.ReplaceAll(host, "{vault}", vault)
	host = strings.ReplaceAll(host, "{secret}", secret)
	host = strings.ReplaceAll(host, "{dnsSuffix}", client.dnsSuffix)
	urlPath := "/customuri/{subscriptionId}/{keyName}"
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(host + urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if pathsGetEmptyOptions != nil && pathsGetEmptyOptions.KeyVersion != nil {
		query.Set("keyVersion", *pathsGetEmptyOptions.KeyVersion)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *pathsOperations) getEmptyHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
