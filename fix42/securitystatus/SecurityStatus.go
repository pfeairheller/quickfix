//Package securitystatus msg type = f.
package securitystatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a SecurityStatus wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityStatusReqID is a non-required field for SecurityStatus.
func (m Message) SecurityStatusReqID() (*field.SecurityStatusReqIDField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatusReqID reads a SecurityStatusReqID from SecurityStatus.
func (m Message) GetSecurityStatusReqID(f *field.SecurityStatusReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for SecurityStatus.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityStatus.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityStatus.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityStatus.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityStatus.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityStatus.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for SecurityStatus.
func (m Message) IDSource() (*field.IDSourceField, quickfix.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from SecurityStatus.
func (m Message) GetIDSource(f *field.IDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityStatus.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityStatus.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityStatus.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityStatus.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for SecurityStatus.
func (m Message) MaturityDay() (*field.MaturityDayField, quickfix.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from SecurityStatus.
func (m Message) GetMaturityDay(f *field.MaturityDayField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityStatus.
func (m Message) PutOrCall() (*field.PutOrCallField, quickfix.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityStatus.
func (m Message) GetPutOrCall(f *field.PutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityStatus.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityStatus.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityStatus.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityStatus.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityStatus.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityStatus.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityStatus.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityStatus.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityStatus.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityStatus.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityStatus.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityStatus.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityStatus.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityStatus.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityStatus.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityStatus.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityStatus.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityStatus.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityStatus.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityStatus.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityStatus.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityStatus.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityStatus.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityStatus.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityStatus.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityStatus.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for SecurityStatus.
func (m Message) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, quickfix.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from SecurityStatus.
func (m Message) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityTradingStatus is a non-required field for SecurityStatus.
func (m Message) SecurityTradingStatus() (*field.SecurityTradingStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityTradingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityTradingStatus reads a SecurityTradingStatus from SecurityStatus.
func (m Message) GetSecurityTradingStatus(f *field.SecurityTradingStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for SecurityStatus.
func (m Message) FinancialStatus() (*field.FinancialStatusField, quickfix.MessageRejectError) {
	f := &field.FinancialStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from SecurityStatus.
func (m Message) GetFinancialStatus(f *field.FinancialStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityStatus.
func (m Message) CorporateAction() (*field.CorporateActionField, quickfix.MessageRejectError) {
	f := &field.CorporateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityStatus.
func (m Message) GetCorporateAction(f *field.CorporateActionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HaltReasonChar is a non-required field for SecurityStatus.
func (m Message) HaltReasonChar() (*field.HaltReasonCharField, quickfix.MessageRejectError) {
	f := &field.HaltReasonCharField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHaltReasonChar reads a HaltReasonChar from SecurityStatus.
func (m Message) GetHaltReasonChar(f *field.HaltReasonCharField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InViewOfCommon is a non-required field for SecurityStatus.
func (m Message) InViewOfCommon() (*field.InViewOfCommonField, quickfix.MessageRejectError) {
	f := &field.InViewOfCommonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInViewOfCommon reads a InViewOfCommon from SecurityStatus.
func (m Message) GetInViewOfCommon(f *field.InViewOfCommonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DueToRelated is a non-required field for SecurityStatus.
func (m Message) DueToRelated() (*field.DueToRelatedField, quickfix.MessageRejectError) {
	f := &field.DueToRelatedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDueToRelated reads a DueToRelated from SecurityStatus.
func (m Message) GetDueToRelated(f *field.DueToRelatedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BuyVolume is a non-required field for SecurityStatus.
func (m Message) BuyVolume() (*field.BuyVolumeField, quickfix.MessageRejectError) {
	f := &field.BuyVolumeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBuyVolume reads a BuyVolume from SecurityStatus.
func (m Message) GetBuyVolume(f *field.BuyVolumeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SellVolume is a non-required field for SecurityStatus.
func (m Message) SellVolume() (*field.SellVolumeField, quickfix.MessageRejectError) {
	f := &field.SellVolumeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSellVolume reads a SellVolume from SecurityStatus.
func (m Message) GetSellVolume(f *field.SellVolumeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HighPx is a non-required field for SecurityStatus.
func (m Message) HighPx() (*field.HighPxField, quickfix.MessageRejectError) {
	f := &field.HighPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHighPx reads a HighPx from SecurityStatus.
func (m Message) GetHighPx(f *field.HighPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LowPx is a non-required field for SecurityStatus.
func (m Message) LowPx() (*field.LowPxField, quickfix.MessageRejectError) {
	f := &field.LowPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLowPx reads a LowPx from SecurityStatus.
func (m Message) GetLowPx(f *field.LowPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for SecurityStatus.
func (m Message) LastPx() (*field.LastPxField, quickfix.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from SecurityStatus.
func (m Message) GetLastPx(f *field.LastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityStatus.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityStatus.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Adjustment is a non-required field for SecurityStatus.
func (m Message) Adjustment() (*field.AdjustmentField, quickfix.MessageRejectError) {
	f := &field.AdjustmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustment reads a Adjustment from SecurityStatus.
func (m Message) GetAdjustment(f *field.AdjustmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SecurityStatus messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SecurityStatus.
func Builder(
	symbol *field.SymbolField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("f"))
	builder.Body().Set(symbol)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "f", r
}
