// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

var (
	c      = flag.String("c", "", "optional client ID of a user assigned identity. Mutually exclusive with r.")
	r      = flag.String("r", "", "optional resource ID of a user assigned identity. Mutually exclusive with c.")
	listen = flag.Bool("listen", false, "listen for HTTP requests")
	scope  = flag.String("s", "https://storage.azure.com/.default", "scope for token request")
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "test passed")
}

func main() {
	flag.Parse()
	if *listen {
		listenAddr := ":8080"
		if port, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
			listenAddr = ":" + port
		}
		http.HandleFunc("/", handler)
		log.Println("Listening on", listenAddr)
		log.Fatal(http.ListenAndServe(listenAddr, nil))
	} else {
		opts := &azidentity.ManagedIdentityCredentialOptions{}
		if *c != "" {
			if *r != "" {
				log.Fatal(`"c" and "r" are mutually exclusive`)
			}
			opts.ID = azidentity.ClientID(*c)
		} else if *r != "" {
			opts.ID = azidentity.ResourceID(*r)
		}
		cred, err := azidentity.NewManagedIdentityCredential(opts)
		if err != nil {
			log.Fatal(err)
		}
		_, err = azblob.NewClient("localhost", cred, nil)
		if err != nil {
			log.Fatal(err)
		}
		_, err = cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{*scope}})
		if err != nil {
			log.Fatal(err)
		}
	}
}
