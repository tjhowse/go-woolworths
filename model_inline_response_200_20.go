/*
 * flows_filtered Mitmproxy2Swagger
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package woolworths

type InlineResponse20020 struct {
	CreditCard *interface{} `json:"CreditCard,omitempty"`
	GiftCard *interface{} `json:"GiftCard,omitempty"`
	Paypal *interface{} `json:"Paypal,omitempty"`
	Account *interface{} `json:"Account,omitempty"`
	Openpay *interface{} `json:"Openpay,omitempty"`
	Success bool `json:"Success,omitempty"`
	ErrorMessage *interface{} `json:"ErrorMessage,omitempty"`
	LastUsedDate string `json:"LastUsedDate,omitempty"`
	IsImpersonated bool `json:"IsImpersonated,omitempty"`
}