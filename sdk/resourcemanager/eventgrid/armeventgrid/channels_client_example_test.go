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

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_Get.json
func ExampleChannelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().Get(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Channel = armeventgrid.Channel{
	// 	Name: to.Ptr("exampleChannelName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerNamespaces/channels"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerNamespaces/examplePartnerNamespaceName1/changes/exampleChannelName1"),
	// 	Properties: &armeventgrid.ChannelProperties{
	// 		ChannelType: to.Ptr(armeventgrid.ChannelTypePartnerTopic),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Example message to approver"),
	// 		PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
	// 			Name: to.Ptr("examplePartnerTopic1"),
	// 			AzureSubscriptionID: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40"),
	// 			ResourceGroupName: to.Ptr("examplerg2"),
	// 			Source: to.Ptr("ContosoCorp.Accounts.User1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armeventgrid.ChannelProvisioningStateSucceeded),
	// 		ReadinessState: to.Ptr(armeventgrid.ReadinessStateNeverActivated),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_CreateOrUpdate.json
func ExampleChannelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().CreateOrUpdate(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", armeventgrid.Channel{
		Properties: &armeventgrid.ChannelProperties{
			ChannelType:                     to.Ptr(armeventgrid.ChannelTypePartnerTopic),
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t }()),
			MessageForActivation:            to.Ptr("Example message to approver"),
			PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
				Name:                to.Ptr("examplePartnerTopic1"),
				AzureSubscriptionID: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40"),
				ResourceGroupName:   to.Ptr("examplerg2"),
				Source:              to.Ptr("ContosoCorp.Accounts.User1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Channel = armeventgrid.Channel{
	// 	Name: to.Ptr("exampleChannelName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerNamespaces/channels"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerNamespaces/examplePartnerNamespaceName1/changes/exampleChannelName1"),
	// 	Properties: &armeventgrid.ChannelProperties{
	// 		ChannelType: to.Ptr(armeventgrid.ChannelTypePartnerTopic),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Example message to approver"),
	// 		PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
	// 			Name: to.Ptr("examplePartnerTopic1"),
	// 			AzureSubscriptionID: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40"),
	// 			ResourceGroupName: to.Ptr("examplerg2"),
	// 			Source: to.Ptr("ContosoCorp.Accounts.User1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_Delete.json
func ExampleChannelsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChannelsClient().BeginDelete(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleEventChannelName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_Update.json
func ExampleChannelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewChannelsClient().Update(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", armeventgrid.ChannelUpdateParameters{
		Properties: &armeventgrid.ChannelUpdateParametersProperties{
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T23:06:11.785Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_ListByPartnerNamespace.json
func ExampleChannelsClient_NewListByPartnerNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChannelsClient().NewListByPartnerNamespacePager("examplerg", "examplePartnerNamespaceName1", &armeventgrid.ChannelsClientListByPartnerNamespaceOptions{Filter: nil,
		Top: nil,
	})
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
		// page.ChannelsListResult = armeventgrid.ChannelsListResult{
		// 	Value: []*armeventgrid.Channel{
		// 		{
		// 			Name: to.Ptr("exampleChannelName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerNamespaces/channels"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerNamespaces/examplePartnerNamespaceName1/changes/exampleChannelName1"),
		// 			Properties: &armeventgrid.ChannelProperties{
		// 				ChannelType: to.Ptr(armeventgrid.ChannelTypePartnerTopic),
		// 				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
		// 				MessageForActivation: to.Ptr("Example message to approver"),
		// 				PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
		// 					Name: to.Ptr("examplePartnerTopic1"),
		// 					AzureSubscriptionID: to.Ptr("8f6b6269-84f2-4d09-9e31-1127efcd1e40"),
		// 					ResourceGroupName: to.Ptr("examplerg2"),
		// 					Source: to.Ptr("ContosoCorp.Accounts.User1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.ChannelProvisioningStateSucceeded),
		// 				ReadinessState: to.Ptr(armeventgrid.ReadinessStateNeverActivated),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_GetFullUrl.json
func ExampleChannelsClient_GetFullURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().GetFullURL(ctx, "examplerg", "examplenamespace", "examplechannel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscriptionFullURL = armeventgrid.EventSubscriptionFullURL{
	// 	EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
	// }
}
