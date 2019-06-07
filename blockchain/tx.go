package blockchain

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock check if input is equal to sig
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked check if output is equal to pubkey
// if true, this means the account own data in output
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
