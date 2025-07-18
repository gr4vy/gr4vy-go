// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ForterAntiFraudOptionsCartItemBeneficiary struct {
	PersonalDetails ForterAntiFraudOptionsCartItemBeneficiaryPersonalDetails `json:"personal_details"`
	// Address information of the beneficiary.
	Address *ForterAntiFraudOptionsCartItemBeneficiaryAddress `json:"address,omitempty"`
	// Phone numbers associated with the beneficiary.
	Phone []ForterAntiFraudOptionsCartItemBeneficiaryPhone `json:"phone,omitempty"`
	// Comments related to the beneficiary.
	Comments *ForterAntiFraudOptionsCartItemBeneficiaryComments `json:"comments,omitempty"`
}

func (o *ForterAntiFraudOptionsCartItemBeneficiary) GetPersonalDetails() ForterAntiFraudOptionsCartItemBeneficiaryPersonalDetails {
	if o == nil {
		return ForterAntiFraudOptionsCartItemBeneficiaryPersonalDetails{}
	}
	return o.PersonalDetails
}

func (o *ForterAntiFraudOptionsCartItemBeneficiary) GetAddress() *ForterAntiFraudOptionsCartItemBeneficiaryAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ForterAntiFraudOptionsCartItemBeneficiary) GetPhone() []ForterAntiFraudOptionsCartItemBeneficiaryPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *ForterAntiFraudOptionsCartItemBeneficiary) GetComments() *ForterAntiFraudOptionsCartItemBeneficiaryComments {
	if o == nil {
		return nil
	}
	return o.Comments
}
