package storagecache

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// StorageTargetClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from
// either NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage
// Caches.
type StorageTargetClient struct {
	BaseClient
}

// NewStorageTargetClient creates an instance of the StorageTargetClient client.
func NewStorageTargetClient(subscriptionID string) StorageTargetClient {
	return NewStorageTargetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageTargetClientWithBaseURI creates an instance of the StorageTargetClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewStorageTargetClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetClient {
	return StorageTargetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Flush tells the cache to write all dirty data to the Storage Target's backend storage. Client requests to this
// storage target's namespace will return errors until the flush operation completes.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetClient) Flush(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetFlushFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetClient.Flush")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetClient", "Flush", err.Error())
	}

	req, err := client.FlushPreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Flush", nil, "Failure preparing request")
		return
	}

	result, err = client.FlushSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Flush", result.Response(), "Failure sending request")
		return
	}

	return
}

// FlushPreparer prepares the Flush request.
func (client StorageTargetClient) FlushPreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/flush", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FlushSender sends the Flush request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetClient) FlushSender(req *http.Request) (future StorageTargetFlushFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// FlushResponder handles the response to the Flush request. The method always
// closes the http.Response Body.
func (client StorageTargetClient) FlushResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Invalidate invalidate all cached data for a storage target. Cached files are discarded and fetched from the back end
// on the next request.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetClient) Invalidate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetInvalidateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetClient.Invalidate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetClient", "Invalidate", err.Error())
	}

	req, err := client.InvalidatePreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Invalidate", nil, "Failure preparing request")
		return
	}

	result, err = client.InvalidateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Invalidate", result.Response(), "Failure sending request")
		return
	}

	return
}

// InvalidatePreparer prepares the Invalidate request.
func (client StorageTargetClient) InvalidatePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/invalidate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// InvalidateSender sends the Invalidate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetClient) InvalidateSender(req *http.Request) (future StorageTargetInvalidateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// InvalidateResponder handles the response to the Invalidate request. The method always
// closes the http.Response Body.
func (client StorageTargetClient) InvalidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Resume resumes client access to a previously suspended storage target.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetClient) Resume(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetResumeFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetClient.Resume")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetClient", "Resume", err.Error())
	}

	req, err := client.ResumePreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Resume", nil, "Failure preparing request")
		return
	}

	result, err = client.ResumeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Resume", result.Response(), "Failure sending request")
		return
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client StorageTargetClient) ResumePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetClient) ResumeSender(req *http.Request) (future StorageTargetResumeFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client StorageTargetClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Suspend suspends client access to a storage target.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetClient) Suspend(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetSuspendFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetClient.Suspend")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetClient", "Suspend", err.Error())
	}

	req, err := client.SuspendPreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Suspend", nil, "Failure preparing request")
		return
	}

	result, err = client.SuspendSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetClient", "Suspend", result.Response(), "Failure sending request")
		return
	}

	return
}

// SuspendPreparer prepares the Suspend request.
func (client StorageTargetClient) SuspendPreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/suspend", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SuspendSender sends the Suspend request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetClient) SuspendSender(req *http.Request) (future StorageTargetSuspendFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// SuspendResponder handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (client StorageTargetClient) SuspendResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}