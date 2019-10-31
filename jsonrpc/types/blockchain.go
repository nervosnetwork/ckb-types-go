package types

// Basic

// Uint32 ckb uint32, '0x' prefix hex number
type Uint32 string

// Uint64 ckb uint64, '0x' prefix hex number
type Uint64 string

// Uint128 ckb uint128, '0x' prefix hex number
type Uint128 string

// Hash ckb hash, '0x' prefix hex string
type Hash string

// Enum type

// ScriptHashType ckb script hash type
type ScriptHashType string

// DepType ckb dep type
type DepType string

// Bytes ckb json bytes
type Bytes string

// ProposalShortID ckb proposal short id
type ProposalShortID string

// Enum values
const (
	Data ScriptHashType = "data"
	Type ScriptHashType = "type"

	Code     DepType = "code"
	DepGroup DepType = "dep_group"
)

// Script ckb script
type Script struct {
	CodeHash Hash           `json:"code_hash"`
	HashType ScriptHashType `json:"hash_type"`
	Args     Bytes          `json:"args"`
}

// OutPoint ckb outpoint
type OutPoint struct {
	TxHash Hash   `json:"tx_hash"`
	Index  Uint32 `json:"index"`
}

// CellInput ckb cell input
type CellInput struct {
	Since          Uint64   `json:"since"`
	PreviousOutput OutPoint `json:"previous_output"`
}

// CellOutput ckb cell output
type CellOutput struct {
	Capacity Uint64  `json:"capacity"`
	Lock     Script  `json:"lock"`
	Type     *Script `json:"type"`
}

// CellDep ckb cell dep
type CellDep struct {
	OutPoint OutPoint `json:"out_point"`
	DepType  DepType  `json:"dep_type"`
}

// WitnessArgs ckb witness args
type WitnessArgs struct {
	Lock       *Bytes `json:"lock"`
	InputType  *Bytes `json:"input_type"`
	OutputType *Bytes `json:"output_type"`
}

// Transaction ckb transaction
type Transaction struct {
	Version     Uint32       `json:"version"`
	CellDeps    []CellDep    `json:"cell_deps"`
	HeaderDeps  []Hash       `json:"header_deps"`
	Inputs      []CellInput  `json:"inputs"`
	Outputs     []CellOutput `json:"outputs"`
	Witnesses   []Bytes      `json:"witnesses"`
	OutputsData []Bytes      `json:"outputs_data"`
}

// Header ckb header
type Header struct {
	Version          Uint32  `json:"version"`
	CompactTarget    Uint32  `json:"compact_target"`
	ParentHash       Hash    `json:"parent_hash"`
	Timestamp        Uint64  `json:"timestamp"`
	Number           Uint64  `json:"number"`
	Epoch            Uint64  `json:"epoch"`
	TransactionsRoot Hash    `json:"transactions_root"`
	ProposalsHash    Hash    `json:"proposals_hash"`
	UnclesHash       Hash    `json:"uncles_hash"`
	Dao              string  `json:"dao"`
	Nonce            Uint128 `json:"nonce"`
}

// UncleBlock ckb uncle block
type UncleBlock struct {
	Header    Header            `json:"header"`
	Proposals []ProposalShortID `json:"proposals"`
}

// Block ckb block
type Block struct {
	Header       Header            `json:"header"`
	Uncles       []UncleBlock      `json:"uncles"`
	Transactions []Transaction     `json:"transactions"`
	Proposals    []ProposalShortID `json:"proposals"`
}
