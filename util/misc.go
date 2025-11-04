package util

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: hnsmanager <entity> <operation> <id / name>")
	fmt.Println("Supported entities: ", SupportedEntities)
	fmt.Println("Supported operations: ", SupportedOperations)
}

func IsEntitySupported(entity string) bool {
	for _, supportedEntity := range SupportedEntities {
		if entity == supportedEntity {
			fmt.Println("Entity supported: ", entity)
			return true
		}
	}
	fmt.Println("Entity not supported: ", entity)
	return false
}

func IsOperationSupported(operation string) bool {
	for _, supportedOperation := range SupportedOperations {
		if operation == supportedOperation {
			fmt.Println("Operation supported: ", operation)
			return true
		}
	}
	fmt.Println("Operation not supported: ", operation)
	return false
}

func IsValidParams(entity string, operation string) bool {
	if IsEntitySupported(entity) && IsOperationSupported(operation) {
		return true
	}
	return false
}
