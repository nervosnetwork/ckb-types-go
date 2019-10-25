package types

// Basic

// Uint32 ckb uint32, '0x' prefix nubmer
type Uint32 = string

// Uint64 ckb uint64, '0x' prefix number
type Uint64 = string

// Uint128 ckb uint128, '0x' prefix number
type Uint128 = string

// H256 ckb H256
type H256 = string

// Enum type

// ScriptHashType ckb script hash type
type ScriptHashType string

// DepType ckb dep type
type DepType string

// JSONBytes ckb json bytes
type JSONBytes = string

// ProposalShortID ckb proposal short id
type ProposalShortID = string

// Enum values
const (
	Data ScriptHashType = "data"
	Type ScriptHashType = "type"

	Code     DepType = "code"
	DepGroup DepType = "dep_group"
)

// Script ckb script
type Script struct {
	Args     JSONBytes      `json:"args"`
	CodeHash H256           `json:"code_hash"`
	HashType ScriptHashType `json:"hash_type"`
}

// OutPoint ckb outpoint
type OutPoint struct {
	TxHash H256   `json:"tx_hash"`
	Index  Uint32 `json:"index"`
}

// CellInput ckb cell input
type CellInput struct {
	PreviousOutput OutPoint `json:"previous_output"`
	Since          Uint64   `json:"since"`
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

// Transaction ckb transaction
type Transaction struct {
	Version     Uint32       `json:"version"`
	CellDeps    []CellDep    `json:"cell_deps"`
	HeaderDeps  []H256       `json:"header_deps"`
	Inputs      []CellInput  `json:"inputs"`
	Outputs     []CellOutput `json:"outputs"`
	Witnesses   []JSONBytes  `json:"witnesses"`
	OutputsData []JSONBytes  `json:"outputs_data"`
}

// Header ckb header
type Header struct {
	Version          Uint32   `json:"version"`
	CompactTarget    Uint32   `json:"compact_target"`
	ParentHash       H256     `json:"parent_hash"`
	Timestamp        Uint64   `json:"timestamp"`
	Number           Uint64   `json:"number"`
	Epoch            Uint64   `json:"epoch"`
	TransactionsRoot H256     `json:"transactions_root"`
	ProposalsHash    H256     `json:"proposals_hash"`
	UnclesHash       H256     `json:"uncles_hash"`
	Dao              [32]byte `json:"dao"`
	Nonce            Uint128  `json:"nonce"`
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
