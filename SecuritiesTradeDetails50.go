package iso20022

// Details of the securities trade.
type SecuritiesTradeDetails50 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails50) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails50) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails50) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails50) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails50) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails50) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails50) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails50) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails50) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails50) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails50) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails50) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails50) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails50) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails50) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails50) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails50) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails50) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails50) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
