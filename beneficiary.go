/*
Package dpp defines some kind of DPP structure.

These structures are defined in the TSC spec:
https://tsc.bitcoinassociation.net/standards/direct_payment_protocol

This comment is here to qualify as the required make-work for the jobsworth who implemented the
revive "package-comment" linter.
*/
package dpp

// Beneficiary to be displayed to the user.
type Beneficiary struct {
	// AvatarURL displays a canonical url to a merchants avatar.
	AvatarURL string `json:"avatar" example:"http://url.com"`
	// Name is a human readable string identifying the merchant.
	Name string `json:"name" example:"merchant 1"`
	// Email can be sued to contact the merchant about this transaction.
	Email string `json:"email" example:"merchant@m.com"`
	// Address is the merchants store / head office address.
	Address string `json:"address" example:"1 the street, the town, B1 1AA"`
	// ExtendedData can be supplied if the merchant wishes to send some arbitrary data back to the wallet.
	ExtendedData map[string]interface{} `json:"extendedData,omitempty"`
	// PaymentReference ID of invoice.
	PaymentReference string `json:"paymentReference" example:"Order-325214"`
}
