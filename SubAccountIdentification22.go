package iso20022

// Account to or from which a securities entry is made.
type SubAccountIdentification22 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation16 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification22) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification22) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification22) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification22) AddBalanceForSubAccount() *AggregateBalanceInformation16 {
	newValue := new(AggregateBalanceInformation16)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
