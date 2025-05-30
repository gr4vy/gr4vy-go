// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TransactionStatus string

const (
	TransactionStatusProcessing               TransactionStatus = "processing"
	TransactionStatusAuthorizationSucceeded   TransactionStatus = "authorization_succeeded"
	TransactionStatusAuthorizationDeclined    TransactionStatus = "authorization_declined"
	TransactionStatusAuthorizationFailed      TransactionStatus = "authorization_failed"
	TransactionStatusAuthorizationVoided      TransactionStatus = "authorization_voided"
	TransactionStatusAuthorizationVoidPending TransactionStatus = "authorization_void_pending"
	TransactionStatusCaptureSucceeded         TransactionStatus = "capture_succeeded"
	TransactionStatusCapturePending           TransactionStatus = "capture_pending"
	TransactionStatusBuyerApprovalPending     TransactionStatus = "buyer_approval_pending"
)

func (e TransactionStatus) ToPointer() *TransactionStatus {
	return &e
}
