// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetReportGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *GetReportGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type GetReportRequest struct {
	// The ID of the report to retrieve details for.
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	// The ID of the merchant account to use for this request.
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *GetReportRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *GetReportRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}
