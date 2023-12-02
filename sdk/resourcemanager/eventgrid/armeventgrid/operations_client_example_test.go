//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationsListResult = armeventgrid.OperationsListResult{
		// 	Value: []*armeventgrid.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/register/action"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Registers the eventSubscription for the EventGrid resource provider and enables the creation of Event Grid subscriptions."),
		// 				Operation: to.Ptr("Registers the EventGrid Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("EventGrid Resource Provider"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/eventSubscriptions/write"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Create or update a eventSubscription"),
		// 				Operation: to.Ptr("Write EventSubscription"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("eventSubscriptions"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/eventSubscriptions/read"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Read a eventSubscription"),
		// 				Operation: to.Ptr("Read EventSubscription"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("eventSubscriptions"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/eventSubscriptions/delete"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Delete a eventSubscription"),
		// 				Operation: to.Ptr("Delete EventSubscription"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("eventSubscriptions"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/topics/write"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Create or update a topic"),
		// 				Operation: to.Ptr("Write Topic"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("topics"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/topics/read"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Read a topic"),
		// 				Operation: to.Ptr("Read Topic"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("topics"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventGrid/topics/delete"),
		// 			Display: &armeventgrid.OperationInfo{
		// 				Description: to.Ptr("Delete a topic"),
		// 				Operation: to.Ptr("Delete Topic"),
		// 				Provider: to.Ptr("Microsoft Event Grid"),
		// 				Resource: to.Ptr("topics"),
		// 			},
		// 			Origin: to.Ptr("UserAndSystem"),
		// 	}},
		// }
	}
}
