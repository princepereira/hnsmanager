package main

import (
	"os"
	"strings"

	"github.com/hnsmanager/util"
)

func main() {

	if len(os.Args) == 1 {
		util.Help()
		return
	}

	if len(os.Args) < 2 && strings.ToLower(os.Args[1]) == util.Helper {
		util.Help()
		return
	}

	if len(os.Args) < 3 {
		util.Help()
		return
	}

	entity := strings.ToLower(os.Args[1])
	operation := strings.ToLower(os.Args[2])

	if !util.IsValidParams(entity, operation) {
		util.Help()
		return
	}

	switch entity {
	case util.Network:
		util.NetworkOperation(operation) // Commenting for size constraints to do kubectl cp
	case util.Endpoint:
		util.EndpointOperation(operation)
		// case util.Namespace:
		// 	util.NamespaceOperation(operation)
	}
}
