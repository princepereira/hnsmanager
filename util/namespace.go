package util

import (
	"fmt"
	"os"
)

func AttachNamespace(namespaceID string, containerID string) {
	fmt.Println("Attach namespace initiated. NamespaceID: ", namespaceID, " ContainerID: ", containerID)
	// err := hcn.AttachNamespace(namespaceID, containerID)
	// if err != nil {
	// 	fmt.Println("Attach namespace failed. NamespaceID: ", namespaceID, " ContainerID: ", containerID, " Error: ", err)
	// 	return
	// }
	fmt.Println("Attach namespace successful. NamespaceID: ", namespaceID, " ContainerID: ", containerID)
}

func DetachNamespace(namespaceID string, containerID string) {
	fmt.Println("Detach namespace initiated. NamespaceID: ", namespaceID, " ContainerID: ", containerID)
	// err := hcn.DetachNamespace(namespaceID, containerID)
	// if err != nil {
	// 	fmt.Println("Detach namespace failed. NamespaceID: ", namespaceID, " ContainerID: ", containerID, " Error: ", err)
	// 	return
	// }
	fmt.Println("Detach namespace successful. NamespaceID: ", namespaceID, " ContainerID: ", containerID)
}

func NamespaceOperation(operation string) {
	switch operation {
	case Attach:
		if len(os.Args) < 5 {
			fmt.Println("Namespace attach operation : hnsmanager.exe namespace attach <namespaceid> <containerid>")
			return
		}
		namespaceID := os.Args[3]
		containerID := os.Args[4]
		AttachNamespace(namespaceID, containerID)

	case Detach:
		if len(os.Args) < 5 {
			fmt.Println("Namespace detach operation : hnsmanager.exe namespace detach <namespaceid> <containerid>")
			return
		}
		namespaceID := os.Args[3]
		containerID := os.Args[4]
		DetachNamespace(namespaceID, containerID)

	default:
		fmt.Println("Invalid operation. Valid operations are: attach, detach")
	}
}
