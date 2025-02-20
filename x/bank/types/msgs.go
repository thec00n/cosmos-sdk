package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgSend{}
	_ sdk.Msg = &MsgMultiSend{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgSend - construct a msg to send coins from one account to another.
func NewMsgSend(fromAddr, toAddr string, amount sdk.Coins, lol string) *MsgSend {
	return &MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      amount,
	}
}

// NewMsgMultiSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgMultiSend(in Input, out []Output) *MsgMultiSend {
	return &MsgMultiSend{Inputs: []Input{in}, Outputs: out}
}

// NewMsgSetSendEnabled Construct a message to set one or more SendEnabled entries.
func NewMsgSetSendEnabled(authority string, sendEnabled []*SendEnabled, useDefaultFor []string) *MsgSetSendEnabled {
	return &MsgSetSendEnabled{
		Authority:     authority,
		SendEnabled:   sendEnabled,
		UseDefaultFor: useDefaultFor,
	}
}
