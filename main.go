package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

const (
	errorResponse = `
{
    "Error": {
        "Code": "ResourceGroupDeletionBlocked",
        "Target": null,
        "Message": "Deletion of resource group '<redacted>' failed as resources with identifiers 'Microsoft.Network/routeTables/<redacted>-routetable,Microsoft.Network/virtualNetworks/<redacted>,Microsoft.Network/networkSecurityGroups/<redacted>-nsg' could not be deleted. The provisioning state of the resource group will be rolled back. The tracking Id is '<redacted>'. Please check audit logs for more details.",
        "Details": [{
            "Code": null,
            "Target": "/subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/routeTables/<redacted>-routetable",
            "Message": "{\"error\":{\"code\":\"InUseRouteTableCannotBeDeleted\",\"message\":\"Route table <redacted>-routetable is in use and cannot be deleted.\",\"details\":[]}}",
            "Details": null,
            "AdditionalInfo": null
        }, {
            "Code": null,
            "Target": "/subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/virtualNetworks/<redacted>",
            "Message": "{\"error\":{\"code\":\"InUseSubnetCannotBeDeleted\",\"message\":\"Subnet <redacted> is in use by /subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/networkInterfaces/<redacted>-nic-4/ipConfigurations/ipconfig1 and cannot be deleted.\",\"details\":[]}}",
            "Details": null,
            "AdditionalInfo": null
        }, {
            "Code": null,
            "Target": "/subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/networkSecurityGroups/<redacted>-nsg",
            "Message": "{\"error\":{\"code\":\"InUseNetworkSecurityGroupCannotBeDeleted\",\"message\":\"Network security group /subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/networkSecurityGroups/<redacted>-nsg cannot be deleted because it is in use by the following resources: /subscriptions/<redacted>/resourceGroups/<redacted>/providers/Microsoft.Network/virtualNetworks/<redacted>/subnets/<redacted>.\",\"details\":[]}}",
            "Details": null,
            "AdditionalInfo": null
        }],
        "AdditionalInfo": null
    }
}`
)

func main() {
	var (
		subscriptionID    = "subID"
		resourceGroupName = "resourcegroup"
		loop              = 0
		acceptedHandler   func(w http.ResponseWriter, r *http.Request)
		terminalHandler   func(w http.ResponseWriter, r *http.Request)
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("received request %s on loop %d\n", r.URL, loop)
		if loop < 5 {
			acceptedHandler(w, r)
		} else {
			terminalHandler(w, r)
		}
		loop++
	}))
	defer ts.Close()

	acceptedHandler = func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Location", ts.URL+"/randomthing")
		header.Set("Retry-After", "1")
		w.WriteHeader(http.StatusAccepted)
	}

	terminalHandler = func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, errorResponse)
	}

	client := resources.NewGroupsClientWithBaseURI(ts.URL, subscriptionID)
	client.RequestInspector = logRequest()
	client.ResponseInspector = logResponse()
	client.RetryAttempts = 2
	client.RetryDuration = time.Duration(2)
	ctx := context.Background()
	future, err := client.Delete(ctx, resourceGroupName)
	if err != nil {
		panic(err)
	}

	if err = future.WaitForCompletion(ctx, client.Client); err != nil {
		panic(err)
	}

	fmt.Println("completed. but never really reach here..")
}

func logRequest() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				log.Println(err)
			}
			dump, _ := httputil.DumpRequestOut(r, true)
			log.Println(string(dump))
			return r, err
		})
	}
}

func logResponse() autorest.RespondDecorator {
	return func(p autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(r *http.Response) error {
			err := p.Respond(r)
			if err != nil {
				log.Println(err)
			}
			dump, _ := httputil.DumpResponse(r, true)
			log.Println(string(dump))
			return err
		})
	}
}
