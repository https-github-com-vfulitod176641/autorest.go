// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getSubscriptionInMethodOperations(t *testing.T) azurespecialsgroup.SubscriptionInMethodOperations {
	client, err := azurespecialsgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.SubscriptionInMethodOperations()
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method.  pass in subscription id = null, client-side validation should prevent you from making this call
func TestPostMethodLocalNull(t *testing.T) {
	t.Skip("invalid test, missing x-nullable")
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostMethodLocalValid(t *testing.T) {
	client := getSubscriptionInMethodOperations(t)
	result, err := client.PostMethodLocalValid(context.Background(), "1234-5678-9012-3456")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostPathLocalValid(t *testing.T) {
	client := getSubscriptionInMethodOperations(t)
	result, err := client.PostPathLocalValid(context.Background(), "1234-5678-9012-3456")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostSwaggerLocalValid(t *testing.T) {
	client := getSubscriptionInMethodOperations(t)
	result, err := client.PostSwaggerLocalValid(context.Background(), "1234-5678-9012-3456")
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}