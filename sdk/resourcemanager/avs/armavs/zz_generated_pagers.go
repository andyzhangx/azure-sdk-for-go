//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AddonsClientListPager provides operations for iterating over paged responses.
type AddonsClientListPager struct {
	client    *AddonsClient
	current   AddonsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AddonsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AddonsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AddonsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AddonList.NextLink == nil || len(*p.current.AddonList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AddonsClientListResponse page.
func (p *AddonsClientListPager) PageResponse() AddonsClientListResponse {
	return p.current
}

// AuthorizationsClientListPager provides operations for iterating over paged responses.
type AuthorizationsClientListPager struct {
	client    *AuthorizationsClient
	current   AuthorizationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AuthorizationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AuthorizationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AuthorizationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExpressRouteAuthorizationList.NextLink == nil || len(*p.current.ExpressRouteAuthorizationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AuthorizationsClientListResponse page.
func (p *AuthorizationsClientListPager) PageResponse() AuthorizationsClientListResponse {
	return p.current
}

// CloudLinksClientListPager provides operations for iterating over paged responses.
type CloudLinksClientListPager struct {
	client    *CloudLinksClient
	current   CloudLinksClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CloudLinksClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CloudLinksClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CloudLinksClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CloudLinkList.NextLink == nil || len(*p.current.CloudLinkList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CloudLinksClientListResponse page.
func (p *CloudLinksClientListPager) PageResponse() CloudLinksClientListResponse {
	return p.current
}

// ClustersClientListPager provides operations for iterating over paged responses.
type ClustersClientListPager struct {
	client    *ClustersClient
	current   ClustersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterList.NextLink == nil || len(*p.current.ClusterList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListResponse page.
func (p *ClustersClientListPager) PageResponse() ClustersClientListResponse {
	return p.current
}

// DatastoresClientListPager provides operations for iterating over paged responses.
type DatastoresClientListPager struct {
	client    *DatastoresClient
	current   DatastoresClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatastoresClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DatastoresClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DatastoresClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatastoreList.NextLink == nil || len(*p.current.DatastoreList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DatastoresClientListResponse page.
func (p *DatastoresClientListPager) PageResponse() DatastoresClientListResponse {
	return p.current
}

// GlobalReachConnectionsClientListPager provides operations for iterating over paged responses.
type GlobalReachConnectionsClientListPager struct {
	client    *GlobalReachConnectionsClient
	current   GlobalReachConnectionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, GlobalReachConnectionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *GlobalReachConnectionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *GlobalReachConnectionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.GlobalReachConnectionList.NextLink == nil || len(*p.current.GlobalReachConnectionList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current GlobalReachConnectionsClientListResponse page.
func (p *GlobalReachConnectionsClientListPager) PageResponse() GlobalReachConnectionsClientListResponse {
	return p.current
}

// HcxEnterpriseSitesClientListPager provides operations for iterating over paged responses.
type HcxEnterpriseSitesClientListPager struct {
	client    *HcxEnterpriseSitesClient
	current   HcxEnterpriseSitesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, HcxEnterpriseSitesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *HcxEnterpriseSitesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *HcxEnterpriseSitesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HcxEnterpriseSiteList.NextLink == nil || len(*p.current.HcxEnterpriseSiteList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current HcxEnterpriseSitesClientListResponse page.
func (p *HcxEnterpriseSitesClientListPager) PageResponse() HcxEnterpriseSitesClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// PlacementPoliciesClientListPager provides operations for iterating over paged responses.
type PlacementPoliciesClientListPager struct {
	client    *PlacementPoliciesClient
	current   PlacementPoliciesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PlacementPoliciesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PlacementPoliciesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PlacementPoliciesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PlacementPoliciesList.NextLink == nil || len(*p.current.PlacementPoliciesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PlacementPoliciesClientListResponse page.
func (p *PlacementPoliciesClientListPager) PageResponse() PlacementPoliciesClientListResponse {
	return p.current
}

// PrivateCloudsClientListInSubscriptionPager provides operations for iterating over paged responses.
type PrivateCloudsClientListInSubscriptionPager struct {
	client    *PrivateCloudsClient
	current   PrivateCloudsClientListInSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateCloudsClientListInSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateCloudsClientListInSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateCloudsClientListInSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateCloudList.NextLink == nil || len(*p.current.PrivateCloudList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listInSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateCloudsClientListInSubscriptionResponse page.
func (p *PrivateCloudsClientListInSubscriptionPager) PageResponse() PrivateCloudsClientListInSubscriptionResponse {
	return p.current
}

// PrivateCloudsClientListPager provides operations for iterating over paged responses.
type PrivateCloudsClientListPager struct {
	client    *PrivateCloudsClient
	current   PrivateCloudsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateCloudsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateCloudsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateCloudsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateCloudList.NextLink == nil || len(*p.current.PrivateCloudList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateCloudsClientListResponse page.
func (p *PrivateCloudsClientListPager) PageResponse() PrivateCloudsClientListResponse {
	return p.current
}

// ScriptCmdletsClientListPager provides operations for iterating over paged responses.
type ScriptCmdletsClientListPager struct {
	client    *ScriptCmdletsClient
	current   ScriptCmdletsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptCmdletsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptCmdletsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptCmdletsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptCmdletsList.NextLink == nil || len(*p.current.ScriptCmdletsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ScriptCmdletsClientListResponse page.
func (p *ScriptCmdletsClientListPager) PageResponse() ScriptCmdletsClientListResponse {
	return p.current
}

// ScriptExecutionsClientListPager provides operations for iterating over paged responses.
type ScriptExecutionsClientListPager struct {
	client    *ScriptExecutionsClient
	current   ScriptExecutionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptExecutionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptExecutionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptExecutionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptExecutionsList.NextLink == nil || len(*p.current.ScriptExecutionsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ScriptExecutionsClientListResponse page.
func (p *ScriptExecutionsClientListPager) PageResponse() ScriptExecutionsClientListResponse {
	return p.current
}

// ScriptPackagesClientListPager provides operations for iterating over paged responses.
type ScriptPackagesClientListPager struct {
	client    *ScriptPackagesClient
	current   ScriptPackagesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ScriptPackagesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ScriptPackagesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ScriptPackagesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ScriptPackagesList.NextLink == nil || len(*p.current.ScriptPackagesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ScriptPackagesClientListResponse page.
func (p *ScriptPackagesClientListPager) PageResponse() ScriptPackagesClientListResponse {
	return p.current
}

// VirtualMachinesClientListPager provides operations for iterating over paged responses.
type VirtualMachinesClientListPager struct {
	client    *VirtualMachinesClient
	current   VirtualMachinesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualMachinesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualMachinesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualMachinesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualMachinesList.NextLink == nil || len(*p.current.VirtualMachinesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VirtualMachinesClientListResponse page.
func (p *VirtualMachinesClientListPager) PageResponse() VirtualMachinesClientListResponse {
	return p.current
}

// WorkloadNetworksClientListDNSServicesPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListDNSServicesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListDNSServicesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListDNSServicesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListDNSServicesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListDNSServicesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDNSServicesList.NextLink == nil || len(*p.current.WorkloadNetworkDNSServicesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listDNSServicesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListDNSServicesResponse page.
func (p *WorkloadNetworksClientListDNSServicesPager) PageResponse() WorkloadNetworksClientListDNSServicesResponse {
	return p.current
}

// WorkloadNetworksClientListDNSZonesPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListDNSZonesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListDNSZonesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListDNSZonesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListDNSZonesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListDNSZonesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDNSZonesList.NextLink == nil || len(*p.current.WorkloadNetworkDNSZonesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listDNSZonesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListDNSZonesResponse page.
func (p *WorkloadNetworksClientListDNSZonesPager) PageResponse() WorkloadNetworksClientListDNSZonesResponse {
	return p.current
}

// WorkloadNetworksClientListDhcpPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListDhcpPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListDhcpResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListDhcpResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListDhcpPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListDhcpPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkDhcpList.NextLink == nil || len(*p.current.WorkloadNetworkDhcpList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listDhcpHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListDhcpResponse page.
func (p *WorkloadNetworksClientListDhcpPager) PageResponse() WorkloadNetworksClientListDhcpResponse {
	return p.current
}

// WorkloadNetworksClientListGatewaysPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListGatewaysPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListGatewaysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListGatewaysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListGatewaysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListGatewaysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkGatewayList.NextLink == nil || len(*p.current.WorkloadNetworkGatewayList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listGatewaysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListGatewaysResponse page.
func (p *WorkloadNetworksClientListGatewaysPager) PageResponse() WorkloadNetworksClientListGatewaysResponse {
	return p.current
}

// WorkloadNetworksClientListPortMirroringPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListPortMirroringPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListPortMirroringResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListPortMirroringResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListPortMirroringPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListPortMirroringPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkPortMirroringList.NextLink == nil || len(*p.current.WorkloadNetworkPortMirroringList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listPortMirroringHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListPortMirroringResponse page.
func (p *WorkloadNetworksClientListPortMirroringPager) PageResponse() WorkloadNetworksClientListPortMirroringResponse {
	return p.current
}

// WorkloadNetworksClientListPublicIPsPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListPublicIPsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListPublicIPsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListPublicIPsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListPublicIPsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListPublicIPsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkPublicIPsList.NextLink == nil || len(*p.current.WorkloadNetworkPublicIPsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listPublicIPsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListPublicIPsResponse page.
func (p *WorkloadNetworksClientListPublicIPsPager) PageResponse() WorkloadNetworksClientListPublicIPsResponse {
	return p.current
}

// WorkloadNetworksClientListSegmentsPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListSegmentsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListSegmentsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListSegmentsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListSegmentsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListSegmentsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkSegmentsList.NextLink == nil || len(*p.current.WorkloadNetworkSegmentsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listSegmentsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListSegmentsResponse page.
func (p *WorkloadNetworksClientListSegmentsPager) PageResponse() WorkloadNetworksClientListSegmentsResponse {
	return p.current
}

// WorkloadNetworksClientListVMGroupsPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListVMGroupsPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListVMGroupsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListVMGroupsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListVMGroupsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListVMGroupsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkVMGroupsList.NextLink == nil || len(*p.current.WorkloadNetworkVMGroupsList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listVMGroupsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListVMGroupsResponse page.
func (p *WorkloadNetworksClientListVMGroupsPager) PageResponse() WorkloadNetworksClientListVMGroupsResponse {
	return p.current
}

// WorkloadNetworksClientListVirtualMachinesPager provides operations for iterating over paged responses.
type WorkloadNetworksClientListVirtualMachinesPager struct {
	client    *WorkloadNetworksClient
	current   WorkloadNetworksClientListVirtualMachinesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkloadNetworksClientListVirtualMachinesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkloadNetworksClientListVirtualMachinesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkloadNetworksClientListVirtualMachinesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadNetworkVirtualMachinesList.NextLink == nil || len(*p.current.WorkloadNetworkVirtualMachinesList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listVirtualMachinesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkloadNetworksClientListVirtualMachinesResponse page.
func (p *WorkloadNetworksClientListVirtualMachinesPager) PageResponse() WorkloadNetworksClientListVirtualMachinesResponse {
	return p.current
}