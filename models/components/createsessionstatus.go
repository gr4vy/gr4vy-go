// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateSessionStatus string

const (
	CreateSessionStatusSucceeded CreateSessionStatus = "succeeded"
	CreateSessionStatusFailed    CreateSessionStatus = "failed"
)

func (e CreateSessionStatus) ToPointer() *CreateSessionStatus {
	return &e
}
