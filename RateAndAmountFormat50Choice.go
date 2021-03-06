package iso20022

// Choice of format between a rate, an amount, index points or a unspecified rate.
type RateAndAmountFormat50Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

	// Value is expressed as a currency and amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Price expressed in index points.
	//
	IndexPoints *RestrictedFINDecimalNumber `xml:"IndxPts"`
}

func (r *RateAndAmountFormat50Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateAndAmountFormat50Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

func (r *RateAndAmountFormat50Choice) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat50Choice) SetIndexPoints(value string) {
	r.IndexPoints = (*RestrictedFINDecimalNumber)(&value)
}
