package util

import (
	"fmt"
	"os"

	"github.com/Microsoft/hnslib/hcn"
)

// func CreateNetwork() {
// 	// Delete network
// 	fmt.Println("Delete network initiated: ", networkIdOrName)
// 	networkObj, err := hcn.GetNetworkByID(networkIdOrName)
// 	if err != nil {
// 		networkObj, err = hcn.GetNetworkByName(networkIdOrName)
// 	}
// 	if err != nil {
// 		fmt.Println("Delete network failed: ", networkIdOrName, " Unable to find network. Error : ", err)
// 		return
// 	}
// 	err = networkObj.Delete()
// 	if err != nil {
// 		fmt.Println("Delete network failed: ", networkIdOrName, " with Error : ", err)
// 		return
// 	}
// 	fmt.Println("Delete network successful: ", networkIdOrName)
// }

func DeleteNetwork(networkIdOrName string) {
	// Delete network
	fmt.Println("Delete network initiated: ", networkIdOrName)
	networkObj, err := hcn.GetNetworkByID(networkIdOrName)
	if err != nil {
		networkObj, err = hcn.GetNetworkByName(networkIdOrName)
	}
	if err != nil {
		fmt.Println("Delete network failed: ", networkIdOrName, " Unable to find network. Error : ", err)
		return
	}
	err = networkObj.Delete()
	if err != nil {
		fmt.Println("Delete network failed: ", networkIdOrName, " with Error : ", err)
		return
	}
	fmt.Println("Delete network successful: ", networkIdOrName)
}

func NetworkOperation(operation string) {
	switch operation {
	case Delete:
		if len(os.Args) < 4 {
			fmt.Println("Network delete operation requires network id or name")
			return
		}
		networkIdOrName := os.Args[3]
		DeleteNetwork(networkIdOrName)
	}
}
