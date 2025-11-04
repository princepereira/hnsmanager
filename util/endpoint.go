package util

import (
	"fmt"
	"os"

	"github.com/Microsoft/hnslib/hcn"
)

func CreateEndpoint(networkId, epName, macAddress, ipAddress string) {
	// Delete network
	fmt.Println("Create endpoint initiated: ", ipAddress)
	epObj := hcn.HostComputeEndpoint{
		Name:               epName,
		HostComputeNetwork: networkId,
		MacAddress:         macAddress,
		IpConfigurations: []hcn.IpConfig{
			{
				IpAddress:    ipAddress,
				PrefixLength: 24,
			},
		},
	}
	newEpObj, err := epObj.Create()
	if err != nil {
		fmt.Println("Create endpoint failed: ", ipAddress, " Error : ", err, " Endpoint obj : ", epObj)
		return
	}
	fmt.Println("Create endpoint successful: ", ipAddress, " New Endpoint : ", newEpObj)
}

func DeleteEndpoint(endpointIdOrName string) {
	epObj, err := hcn.GetEndpointByID(endpointIdOrName)
	if err != nil {
		fmt.Println("Get endpoint failed: ", endpointIdOrName, " Error : ", err)
		return
	}
	// Delete network
	fmt.Println("Delete endpoint initiated: ", endpointIdOrName)

	err = epObj.Delete()
	if err != nil {
		fmt.Println("Delete endpoint failed: ", endpointIdOrName, " Error : ", err, " Endpoint obj : ", epObj)
		return
	}
	fmt.Println("Delete endpoint successful: ", endpointIdOrName)
}

func EndpointOperation(operation string) {
	switch operation {
	case Create:
		if len(os.Args) < 6 {
			fmt.Println("Endpoint create operation : hnsmanager.exe endpoint create <networkId> <endpointName> <ipAddress> <macAddress>")
			return
		}
		networkId := os.Args[3]
		epName := os.Args[4]
		ipAddress := os.Args[5]
		macAddress := os.Args[6]
		CreateEndpoint(networkId, epName, macAddress, ipAddress)
	case Delete:
		if len(os.Args) < 4 {
			fmt.Println("Endpoint delete operation : hnsmanager.exe endpoint delete <endpointIdOrName>")
			return
		}
		endpointIdOrName := os.Args[3]
		DeleteEndpoint(endpointIdOrName)
	}
}
