package iso20022

// ATM information.
type AutomatedTellerMachine9 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities7 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine9) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine9) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine9) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine9) AddCapabilities() *PointOfInteractionCapabilities7 {
	a.Capabilities = new(PointOfInteractionCapabilities7)
	return a.Capabilities
}

func (a *AutomatedTellerMachine9) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}
