/*
 * Kubtest API
 *
 * Kubtest provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: kubtest@kubshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kubtest

// scripts execution request body
type ExecutionRequest struct {
	// script execution custom name
	Name string `json:"name,omitempty"`
	// script kubernetes namespace (\"default\" when not set)
	Namespace string `json:"namespace,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
}
