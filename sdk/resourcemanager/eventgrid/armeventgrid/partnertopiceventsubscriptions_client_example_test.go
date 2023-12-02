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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_Get.json
func ExamplePartnerTopicEventSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().Get(ctx, "examplerg", "examplePartnerTopic1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("examplesubscription1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerTopics/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1/eventSubscriptions/examplesubscription1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
	// 			Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
	// 				QueueName: to.Ptr("que"),
	// 				ResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IncludedEventTypes: []*string{
	// 				to.Ptr("Microsoft.Storage.BlobCreated"),
	// 				to.Ptr("Microsoft.Storage.BlobDeleted")},
	// 				IsSubjectCaseSensitive: to.Ptr(false),
	// 				SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 				SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 			},
	// 			Labels: []*string{
	// 				to.Ptr("label1"),
	// 				to.Ptr("label2")},
	// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 				RetryPolicy: &armeventgrid.RetryPolicy{
	// 					EventTimeToLiveInMinutes: to.Ptr[int32](1440),
	// 					MaxDeliveryAttempts: to.Ptr[int32](30),
	// 				},
	// 				Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_CreateOrUpdate.json
func ExamplePartnerTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "examplerg", "examplePartnerTopic1", "exampleEventSubscriptionName1", armeventgrid.EventSubscription{
		Properties: &armeventgrid.EventSubscriptionProperties{
			Destination: &armeventgrid.WebHookEventSubscriptionDestination{
				EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
				Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
					EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
				},
			},
			Filter: &armeventgrid.EventSubscriptionFilter{
				IsSubjectCaseSensitive: to.Ptr(false),
				SubjectBeginsWith:      to.Ptr("ExamplePrefix"),
				SubjectEndsWith:        to.Ptr("ExampleSuffix"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("exampleEventSubscriptionName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerTopics/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1/eventSubscriptions/exampleEventSubscriptionName1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.WebHookEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
	// 			Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
	// 				EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IsSubjectCaseSensitive: to.Ptr(false),
	// 			SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 			SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 		},
	// 		ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 		RetryPolicy: &armeventgrid.RetryPolicy{
	// 			EventTimeToLiveInMinutes: to.Ptr[int32](1440),
	// 			MaxDeliveryAttempts: to.Ptr[int32](30),
	// 		},
	// 		Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_Delete.json
func ExamplePartnerTopicEventSubscriptionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().BeginDelete(ctx, "examplerg", "examplePartnerTopic1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_Update.json
func ExamplePartnerTopicEventSubscriptionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().BeginUpdate(ctx, "examplerg", "examplePartnerTopic1", "exampleEventSubscriptionName1", armeventgrid.EventSubscriptionUpdateParameters{
		Destination: &armeventgrid.WebHookEventSubscriptionDestination{
			EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
			Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
				EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
			},
		},
		Filter: &armeventgrid.EventSubscriptionFilter{
			IsSubjectCaseSensitive: to.Ptr(true),
			SubjectBeginsWith:      to.Ptr("existingPrefix"),
			SubjectEndsWith:        to.Ptr("newSuffix"),
		},
		Labels: []*string{
			to.Ptr("label1"),
			to.Ptr("label2")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_GetFullUrl.json
func ExamplePartnerTopicEventSubscriptionsClient_GetFullURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().GetFullURL(ctx, "examplerg", "examplePartnerTopic1", "examplesubscription1", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_ListByPartnerTopic.json
func ExamplePartnerTopicEventSubscriptionsClient_NewListByPartnerTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPartnerTopicEventSubscriptionsClient().NewListByPartnerTopicPager("examplerg", "examplePartnerTopic1", &armeventgrid.PartnerTopicEventSubscriptionsClientListByPartnerTopicOptions{Filter: nil,
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
		// page.EventSubscriptionsListResult = armeventgrid.EventSubscriptionsListResult{
		// 	Value: []*armeventgrid.EventSubscription{
		// 		{
		// 			Name: to.Ptr("examplesubscription1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerTopics/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1/eventSubscriptions/examplesubscription1"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
		// 					Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
		// 						QueueName: to.Ptr("que"),
		// 						ResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					SubjectBeginsWith: to.Ptr(""),
		// 					SubjectEndsWith: to.Ptr(""),
		// 				},
		// 				Labels: []*string{
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 				RetryPolicy: &armeventgrid.RetryPolicy{
		// 					EventTimeToLiveInMinutes: to.Ptr[int32](1440),
		// 					MaxDeliveryAttempts: to.Ptr[int32](10),
		// 				},
		// 				Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplesubscription2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerTopics/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1/eventSubscriptions/examplesubscription2"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
		// 					Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
		// 						QueueName: to.Ptr("que"),
		// 						ResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IncludedEventTypes: []*string{
		// 						to.Ptr("Microsoft.Storage.BlobCreated"),
		// 						to.Ptr("Microsoft.Storage.BlobDeleted")},
		// 						IsSubjectCaseSensitive: to.Ptr(false),
		// 						SubjectBeginsWith: to.Ptr("ExamplePrefix"),
		// 						SubjectEndsWith: to.Ptr("ExampleSuffix"),
		// 					},
		// 					Labels: []*string{
		// 						to.Ptr("label1"),
		// 						to.Ptr("label2")},
		// 						ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 						RetryPolicy: &armeventgrid.RetryPolicy{
		// 							EventTimeToLiveInMinutes: to.Ptr[int32](1440),
		// 							MaxDeliveryAttempts: to.Ptr[int32](30),
		// 						},
		// 						Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopicEventSubscriptions_GetDeliveryAttributes.json
func ExamplePartnerTopicEventSubscriptionsClient_GetDeliveryAttributes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().GetDeliveryAttributes(ctx, "examplerg", "examplePartnerTopic1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeliveryAttributeListResult = armeventgrid.DeliveryAttributeListResult{
	// 	Value: []armeventgrid.DeliveryAttributeMappingClassification{
	// 		&armeventgrid.StaticDeliveryAttributeMapping{
	// 			Name: to.Ptr("header1"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeStatic),
	// 			Properties: &armeventgrid.StaticDeliveryAttributeMappingProperties{
	// 				IsSecret: to.Ptr(false),
	// 				Value: to.Ptr("NormalValue"),
	// 			},
	// 		},
	// 		&armeventgrid.DynamicDeliveryAttributeMapping{
	// 			Name: to.Ptr("header2"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeDynamic),
	// 			Properties: &armeventgrid.DynamicDeliveryAttributeMappingProperties{
	// 				SourceField: to.Ptr("data.foo"),
	// 			},
	// 		},
	// 		&armeventgrid.StaticDeliveryAttributeMapping{
	// 			Name: to.Ptr("header3"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeStatic),
	// 			Properties: &armeventgrid.StaticDeliveryAttributeMappingProperties{
	// 				IsSecret: to.Ptr(true),
	// 				Value: to.Ptr("mySecretValue"),
	// 			},
	// 	}},
	// }
}
