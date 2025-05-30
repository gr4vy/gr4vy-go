// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CardScheme string

const (
	CardSchemeAccel           CardScheme = "accel"
	CardSchemeAmex            CardScheme = "amex"
	CardSchemeBancontact      CardScheme = "bancontact"
	CardSchemeCarteBancaire   CardScheme = "carte-bancaire"
	CardSchemeCirrus          CardScheme = "cirrus"
	CardSchemeCuliance        CardScheme = "culiance"
	CardSchemeDankort         CardScheme = "dankort"
	CardSchemeDinersClub      CardScheme = "diners-club"
	CardSchemeDiscover        CardScheme = "discover"
	CardSchemeEftposAustralia CardScheme = "eftpos-australia"
	CardSchemeElo             CardScheme = "elo"
	CardSchemeHipercard       CardScheme = "hipercard"
	CardSchemeJcb             CardScheme = "jcb"
	CardSchemeMaestro         CardScheme = "maestro"
	CardSchemeMastercard      CardScheme = "mastercard"
	CardSchemeMir             CardScheme = "mir"
	CardSchemeNyce            CardScheme = "nyce"
	CardSchemeOther           CardScheme = "other"
	CardSchemePulse           CardScheme = "pulse"
	CardSchemeRupay           CardScheme = "rupay"
	CardSchemeStar            CardScheme = "star"
	CardSchemeUatp            CardScheme = "uatp"
	CardSchemeUnionpay        CardScheme = "unionpay"
	CardSchemeVisa            CardScheme = "visa"
)

func (e CardScheme) ToPointer() *CardScheme {
	return &e
}
