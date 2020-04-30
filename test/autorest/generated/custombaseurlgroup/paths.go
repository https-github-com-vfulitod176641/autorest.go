// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package custombaseurlgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// PathsOperations contains the methods for the Paths group.
type PathsOperations interface {
	// GetEmpty - Get a 200 to test a valid base uri
	GetEmpty(ctx context.Context, accountName string) (*http.Response, error)
}

// pathsOperations implements the PathsOperations interface.
type pathsOperations struct {
	*Client
}

// GetEmpty - Get a 200 to test a valid base uri
func (client *pathsOperations) GetEmpty(ctx context.Context, accountName string) (*http.Response, error) {
	req, err := client.getEmptyCreateRequest(accountName)
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
func (client *pathsOperations) getEmptyCreateRequest(accountName string) (*azcore.Request, error) {
	host := "http://{accountName}{host}"
	host = strings.ReplaceAll(host, "{accountName}", accountName)
	host = strings.ReplaceAll(host, "{host}", client.u.Host)
	urlPath := "/customuri"
	u, err := client.u.Parse(host + urlPath)
	if err != nil {
		return nil, err
	}
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
