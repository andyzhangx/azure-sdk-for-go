// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("AZURE_STORAGE_ACCOUNT")
	if url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "AZURE_STORAGE_ACCOUNT is not set")
		return
	}

	var client *azblob.Client
	cred, err := azidentity.NewManagedIdentityCredential(nil)
	if err == nil {
		client, err = azblob.NewClient(url, cred, nil)
		if err == nil {
			_, err = client.NewListContainersPager(nil).NextPage(context.Background())
		}
	}

	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
}

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Printf("listening at http://127.0.0.1:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
