package fair

type Op int

const (
	Deploy Op = iota
	Mint
	Transfer
	UNKNOWN_OP
)

type DeployData struct {
	p    string
	op   string
	tick string
	dec  string
	max  string
	lim  string

	startBlock string
	duration   string
}

type MintOrTransferData struct {
	p    string
	op   string
	tick string
	amt  string
}
