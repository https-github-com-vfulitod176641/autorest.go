package dategroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DateClient is the test Infrastructure for AutoRest
type DateClient struct {
	BaseClient
}

// NewDateClient creates an instance of the DateClient client.
func NewDateClient() DateClient {
	return NewDateClientWithBaseURI(DefaultBaseURI)
}

// NewDateClientWithBaseURI creates an instance of the DateClient client.
func NewDateClientWithBaseURI(baseURI string) DateClient {
	return DateClient{NewWithBaseURI(baseURI)}
}

// GetInvalidDate get invalid date value
func (client DateClient) GetInvalidDate(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetInvalidDate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetInvalidDatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetInvalidDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetInvalidDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetInvalidDate", resp, "Failure responding to request")
	}

	return
}

// GetInvalidDatePreparer prepares the GetInvalidDate request.
func (client DateClient) GetInvalidDatePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/invaliddate"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidDateSender sends the GetInvalidDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetInvalidDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidDateResponder handles the response to the GetInvalidDate request. The method always
// closes the http.Response Body.
func (client DateClient) GetInvalidDateResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMaxDate get max date value 9999-12-31
func (client DateClient) GetMaxDate(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetMaxDate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMaxDatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMaxDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMaxDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMaxDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetMaxDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMaxDate", resp, "Failure responding to request")
	}

	return
}

// GetMaxDatePreparer prepares the GetMaxDate request.
func (client DateClient) GetMaxDatePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/max"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMaxDateSender sends the GetMaxDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetMaxDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetMaxDateResponder handles the response to the GetMaxDate request. The method always
// closes the http.Response Body.
func (client DateClient) GetMaxDateResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMinDate get min date value 0000-01-01
func (client DateClient) GetMinDate(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetMinDate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMinDatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMinDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMinDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMinDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetMinDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetMinDate", resp, "Failure responding to request")
	}

	return
}

// GetMinDatePreparer prepares the GetMinDate request.
func (client DateClient) GetMinDatePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/min"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMinDateSender sends the GetMinDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetMinDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetMinDateResponder handles the response to the GetMinDate request. The method always
// closes the http.Response Body.
func (client DateClient) GetMinDateResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null date value
func (client DateClient) GetNull(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetNull")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client DateClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client DateClient) GetNullResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowDate get overflow date value
func (client DateClient) GetOverflowDate(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetOverflowDate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOverflowDatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetOverflowDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetOverflowDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetOverflowDate", resp, "Failure responding to request")
	}

	return
}

// GetOverflowDatePreparer prepares the GetOverflowDate request.
func (client DateClient) GetOverflowDatePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/overflowdate"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOverflowDateSender sends the GetOverflowDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetOverflowDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowDateResponder handles the response to the GetOverflowDate request. The method always
// closes the http.Response Body.
func (client DateClient) GetOverflowDateResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowDate get underflow date value
func (client DateClient) GetUnderflowDate(ctx context.Context) (result DateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.GetUnderflowDate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetUnderflowDatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetUnderflowDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetUnderflowDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "GetUnderflowDate", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowDatePreparer prepares the GetUnderflowDate request.
func (client DateClient) GetUnderflowDatePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/underflowdate"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUnderflowDateSender sends the GetUnderflowDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) GetUnderflowDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowDateResponder handles the response to the GetUnderflowDate request. The method always
// closes the http.Response Body.
func (client DateClient) GetUnderflowDateResponder(resp *http.Response) (result DateModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMaxDate put max date value 9999-12-31
func (client DateClient) PutMaxDate(ctx context.Context, dateBody date.Date) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.PutMaxDate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutMaxDatePreparer(ctx, dateBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMaxDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMaxDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMaxDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutMaxDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMaxDate", resp, "Failure responding to request")
	}

	return
}

// PutMaxDatePreparer prepares the PutMaxDate request.
func (client DateClient) PutMaxDatePreparer(ctx context.Context, dateBody date.Date) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/max"),
		autorest.WithJSON(dateBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMaxDateSender sends the PutMaxDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) PutMaxDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMaxDateResponder handles the response to the PutMaxDate request. The method always
// closes the http.Response Body.
func (client DateClient) PutMaxDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMinDate put min date value 0000-01-01
func (client DateClient) PutMinDate(ctx context.Context, dateBody date.Date) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DateClient.PutMinDate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutMinDatePreparer(ctx, dateBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMinDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMinDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMinDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutMinDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.DateClient", "PutMinDate", resp, "Failure responding to request")
	}

	return
}

// PutMinDatePreparer prepares the PutMinDate request.
func (client DateClient) PutMinDatePreparer(ctx context.Context, dateBody date.Date) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/min"),
		autorest.WithJSON(dateBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMinDateSender sends the PutMinDate request. The method will close the
// http.Response Body if it receives an error.
func (client DateClient) PutMinDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMinDateResponder handles the response to the PutMinDate request. The method always
// closes the http.Response Body.
func (client DateClient) PutMinDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
