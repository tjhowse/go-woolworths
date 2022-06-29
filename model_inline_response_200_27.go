/*
 * flows_filtered Mitmproxy2Swagger
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package woolworths

type InlineResponse20027 struct {
	PartialSuccess bool `json:"PartialSuccess,omitempty"`
	TransactionReceipt *interface{} `json:"TransactionReceipt,omitempty"`
	PaymentResponses *interface{} `json:"PaymentResponses,omitempty"`
	ErrorCode string `json:"ErrorCode,omitempty"`
	ErrorDetail *interface{} `json:"ErrorDetail,omitempty"`
	Result bool `json:"Result,omitempty"`
	ErrorMessage *interface{} `json:"ErrorMessage,omitempty"`
	PlacedOrderId *interface{} `json:"PlacedOrderId,omitempty"`
	OrderReference *interface{} `json:"OrderReference,omitempty"`
	OrderPlacementErrors []interface{} `json:"OrderPlacementErrors,omitempty"`
	Model *interface{} `json:"Model,omitempty"`
}
