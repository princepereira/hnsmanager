package util

// Entities
const (
	Network      = "network"
	Namespace    = "namespace"
	Endpoint     = "endpoint"
	LoadBalancer = "loadbalancer"
	Helper       = "help"
)

// Operations
const (
	Create = "create"
	Delete = "delete"
	List   = "list"
	Attach = "attach"
	Detach = "detach"
)

var SupportedEntities = []string{Helper, Network, Endpoint, Namespace}
var SupportedOperations = []string{Create, Delete, Attach, Detach}
