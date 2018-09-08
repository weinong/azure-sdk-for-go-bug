# What is it

Repro for azure-sdk-for-go failing to return error response when delete resource group failed with non-conformant async behavior.
https://github.com/Azure/azure-sdk-for-go/issues/2645

## Steps

```sh
go build main.go
./main
```

## Output

```
2018/09/07 22:46:55 DELETE /subscriptions/subID/resourcegroups/resourcegroup?api-version=2018-05-01 HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /subscriptions/subID/resourcegroups/resourcegroup?api-version=2018-05-01 on loop 0
2018/09/07 22:46:55 HTTP/1.1 202 Accepted
Date: Sat, 08 Sep 2018 05:46:55 GMT
Location: http://127.0.0.1:20426/randomthing
Retry-After: 1
Content-Length: 0


2018/09/07 22:46:55 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 1
2018/09/07 22:46:55 HTTP/1.1 202 Accepted
Date: Sat, 08 Sep 2018 05:46:55 GMT
Location: http://127.0.0.1:20426/randomthing
Retry-After: 1
Content-Length: 0


2018/09/07 22:46:56 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 2
2018/09/07 22:46:56 HTTP/1.1 202 Accepted
Date: Sat, 08 Sep 2018 05:46:56 GMT
Location: http://127.0.0.1:20426/randomthing
Retry-After: 1
Content-Length: 0


2018/09/07 22:46:57 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 3
2018/09/07 22:46:57 HTTP/1.1 202 Accepted
Date: Sat, 08 Sep 2018 05:46:57 GMT
Location: http://127.0.0.1:20426/randomthing
Retry-After: 1
Content-Length: 0


2018/09/07 22:46:58 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 4
2018/09/07 22:46:58 HTTP/1.1 202 Accepted
Date: Sat, 08 Sep 2018 05:46:58 GMT
Location: http://127.0.0.1:20426/randomthing
Retry-After: 1
Content-Length: 0


2018/09/07 22:46:59 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 5
2018/09/07 22:46:59 HTTP/1.1 409 Conflict
Transfer-Encoding: chunked
Content-Type: application/json
Date: Sat, 08 Sep 2018 05:46:59 GMT

8bc

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
}

0


2018/09/07 22:46:59 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 6
2018/09/07 22:46:59 HTTP/1.1 409 Conflict
Transfer-Encoding: chunked
Content-Type: application/json
Date: Sat, 08 Sep 2018 05:46:59 GMT

8bc

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
}

0


2018/09/07 22:46:59 GET /randomthing HTTP/1.1
Host: 127.0.0.1:20426
User-Agent: Go/go1.10.4 (amd64-windows) go-autorest/v10.15.3 Azure-SDK-For-Go/v20.1.0 resources/2018-05-01
Accept-Encoding: gzip


received request /randomthing on loop 7
2018/09/07 22:46:59 HTTP/1.1 409 Conflict
Transfer-Encoding: chunked
Content-Type: application/json
Date: Sat, 08 Sep 2018 05:46:59 GMT

8bc

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
}

0


panic: pollingTrackerBase#updatePollingState: the response from the async operation has an invalid status code: StatusCode=0

goroutine 1 [running]:
main.main()
        E:/GoWork/src/github.com/weinong/bug/main.go:93 +0x445
```
