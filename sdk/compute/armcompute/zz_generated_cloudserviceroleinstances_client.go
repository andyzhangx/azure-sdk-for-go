// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CloudServiceRoleInstancesClient contains the methods for the CloudServiceRoleInstances group.
// Don't use this type directly, use NewCloudServiceRoleInstancesClient() instead.
type CloudServiceRoleInstancesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCloudServiceRoleInstancesClient creates a new instance of CloudServiceRoleInstancesClient with the specified values.
func NewCloudServiceRoleInstancesClient(con *armcore.Connection, subscriptionID string) *CloudServiceRoleInstancesClient {
	return &CloudServiceRoleInstancesClient{con: con, subscriptionID: subscriptionID}
}

// BeginDelete - Deletes a role instance from a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) BeginDelete(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CloudServiceRoleInstancesClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CloudServiceRoleInstancesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("CloudServiceRoleInstancesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a role instance from a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) deleteOperation(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CloudServiceRoleInstancesClient) deleteCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CloudServiceRoleInstancesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets a role instance from a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) Get(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetOptions) (RoleInstanceResponse, error) {
	req, err := client.getCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return RoleInstanceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleInstanceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleInstanceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudServiceRoleInstancesClient) getCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudServiceRoleInstancesClient) getHandleResponse(resp *azcore.Response) (RoleInstanceResponse, error) {
	var val *RoleInstance
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleInstanceResponse{}, err
	}
	return RoleInstanceResponse{RawResponse: resp.Response, RoleInstance: val}, nil
}

// getHandleError handles the Get error response.
func (client *CloudServiceRoleInstancesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetInstanceView - Retrieves information about the run-time state of a role instance in a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) GetInstanceView(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetInstanceViewOptions) (RoleInstanceViewResponse, error) {
	req, err := client.getInstanceViewCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return RoleInstanceViewResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleInstanceViewResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleInstanceViewResponse{}, client.getInstanceViewHandleError(resp)
	}
	return client.getInstanceViewHandleResponse(resp)
}

// getInstanceViewCreateRequest creates the GetInstanceView request.
func (client *CloudServiceRoleInstancesClient) getInstanceViewCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetInstanceViewOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/instanceView"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInstanceViewHandleResponse handles the GetInstanceView response.
func (client *CloudServiceRoleInstancesClient) getInstanceViewHandleResponse(resp *azcore.Response) (RoleInstanceViewResponse, error) {
	var val *RoleInstanceView
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleInstanceViewResponse{}, err
	}
	return RoleInstanceViewResponse{RawResponse: resp.Response, RoleInstanceView: val}, nil
}

// getInstanceViewHandleError handles the GetInstanceView error response.
func (client *CloudServiceRoleInstancesClient) getInstanceViewHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetRemoteDesktopFile - Gets a remote desktop file for a role instance in a cloud service.
// If the operation fails it returns a generic error.
func (client *CloudServiceRoleInstancesClient) GetRemoteDesktopFile(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetRemoteDesktopFileOptions) (*http.Response, error) {
	req, err := client.getRemoteDesktopFileCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getRemoteDesktopFileHandleError(resp)
	}
	return resp.Response, nil
}

// getRemoteDesktopFileCreateRequest creates the GetRemoteDesktopFile request.
func (client *CloudServiceRoleInstancesClient) getRemoteDesktopFileCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesGetRemoteDesktopFileOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/remoteDesktopFile"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.SkipBodyDownload()
	req.Header.Set("Accept", "application/x-rdp")
	return req, nil
}

// getRemoteDesktopFileHandleError handles the GetRemoteDesktopFile error response.
func (client *CloudServiceRoleInstancesClient) getRemoteDesktopFileHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets the list of all role instances in a cloud service. Use nextLink property in the response to get the next page of role instances. Do this
// till nextLink is null to fetch all the role instances.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) List(resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesListOptions) RoleInstanceListResultPager {
	return &roleInstanceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RoleInstanceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleInstanceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *CloudServiceRoleInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudServiceRoleInstancesClient) listHandleResponse(resp *azcore.Response) (RoleInstanceListResultResponse, error) {
	var val *RoleInstanceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleInstanceListResultResponse{}, err
	}
	return RoleInstanceListResultResponse{RawResponse: resp.Response, RoleInstanceListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *CloudServiceRoleInstancesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRebuild - The Rebuild Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles and initializes
// the storage resources that are used by them. If you do not
// want to initialize storage resources, you can use Reimage Role Instance.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) BeginRebuild(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRebuildOptions) (HTTPPollerResponse, error) {
	resp, err := client.rebuild(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CloudServiceRoleInstancesClient.Rebuild", "", resp, client.rebuildHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRebuild creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CloudServiceRoleInstancesClient) ResumeRebuild(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("CloudServiceRoleInstancesClient.Rebuild", token, client.rebuildHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Rebuild - The Rebuild Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles and initializes
// the storage resources that are used by them. If you do not
// want to initialize storage resources, you can use Reimage Role Instance.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) rebuild(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRebuildOptions) (*azcore.Response, error) {
	req, err := client.rebuildCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.rebuildHandleError(resp)
	}
	return resp, nil
}

// rebuildCreateRequest creates the Rebuild request.
func (client *CloudServiceRoleInstancesClient) rebuildCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRebuildOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/rebuild"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// rebuildHandleError handles the Rebuild error response.
func (client *CloudServiceRoleInstancesClient) rebuildHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginReimage - The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) BeginReimage(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginReimageOptions) (HTTPPollerResponse, error) {
	resp, err := client.reimage(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CloudServiceRoleInstancesClient.Reimage", "", resp, client.reimageHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeReimage creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CloudServiceRoleInstancesClient) ResumeReimage(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("CloudServiceRoleInstancesClient.Reimage", token, client.reimageHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Reimage - The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) reimage(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginReimageOptions) (*azcore.Response, error) {
	req, err := client.reimageCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.reimageHandleError(resp)
	}
	return resp, nil
}

// reimageCreateRequest creates the Reimage request.
func (client *CloudServiceRoleInstancesClient) reimageCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginReimageOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/reimage"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// reimageHandleError handles the Reimage error response.
func (client *CloudServiceRoleInstancesClient) reimageHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRestart - The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) BeginRestart(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRestartOptions) (HTTPPollerResponse, error) {
	resp, err := client.restart(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("CloudServiceRoleInstancesClient.Restart", "", resp, client.restartHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRestart creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *CloudServiceRoleInstancesClient) ResumeRestart(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("CloudServiceRoleInstancesClient.Restart", token, client.restartHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Restart - The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRoleInstancesClient) restart(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRestartOptions) (*azcore.Response, error) {
	req, err := client.restartCreateRequest(ctx, roleInstanceName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.restartHandleError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *CloudServiceRoleInstancesClient) restartCreateRequest(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRoleInstancesBeginRestartOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/restart"
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *CloudServiceRoleInstancesClient) restartHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
