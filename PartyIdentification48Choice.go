package iso20022

// Choice of identification of a party.
type PartyIdentification48Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`
}

func (p *PartyIdentification48Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification48Choice) AddProprietaryIdentification() *GenericIdentification19 {
	p.ProprietaryIdentification = new(GenericIdentification19)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification48Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}
