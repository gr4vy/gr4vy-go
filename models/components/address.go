// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Address struct {
	// The city for the address.
	City *string `json:"city,omitempty"`
	// The country for the address in ISO 3166 format.
	Country *string `json:"country,omitempty"`
	// The postal code or zip code for the address.
	PostalCode *string `json:"postal_code,omitempty"`
	// The state, county, or province for the address.
	State *string `json:"state,omitempty"`
	// The code of state, county, or province for the address in ISO 3166-2 format.
	StateCode *string `json:"state_code,omitempty"`
	// The house number or name for the address. Not all payment services use this field but some do.
	HouseNumberOrName *string `json:"house_number_or_name,omitempty"`
	// The first line of the address.
	Line1 *string `json:"line1,omitempty"`
	// The second line of the address.
	Line2 *string `json:"line2,omitempty"`
	// The optional name of the company or organisation to add to the address.
	Organization *string `json:"organization,omitempty"`
}

func (o *Address) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Address) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Address) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *Address) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Address) GetStateCode() *string {
	if o == nil {
		return nil
	}
	return o.StateCode
}

func (o *Address) GetHouseNumberOrName() *string {
	if o == nil {
		return nil
	}
	return o.HouseNumberOrName
}

func (o *Address) GetLine1() *string {
	if o == nil {
		return nil
	}
	return o.Line1
}

func (o *Address) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *Address) GetOrganization() *string {
	if o == nil {
		return nil
	}
	return o.Organization
}
