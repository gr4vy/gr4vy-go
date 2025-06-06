// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GooglePaySessionRequest struct {
	// The domain on which Google Pay is being loaded.
	OriginDomain string `json:"origin_domain"`
}

func (o *GooglePaySessionRequest) GetOriginDomain() string {
	if o == nil {
		return ""
	}
	return o.OriginDomain
}
