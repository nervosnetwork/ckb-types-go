package types

// Basic
type Uint32 = string // 0x-prefixed uint32
type Uint128 = string
type H256 = string

// Enum
type ScriptHashType string
type DepType string

type JsonBytes = []byte
type ProposalShortId = string

const (
	Data ScriptHashType = "data"
	Type ScriptHashType = "type"

	Code     DepType = "code"
	DepGroup DepType = "dep_group"
)

type Script struct {
	Args     JsonBytes      `json:"args"`
	CodeHash H256           `json:"code_hash"`
	HashType ScriptHashType `json:"hash_type"`
}

type OutPoint struct {
	TxHash H256   `json:"tx_hash"`
	Index  Uint32 `json:"index"`
}

type CellInput struct {
	PreviousOutput OutPoint `json:"previous_output"`
	Since          uint64   `json:"since"`
}

type CellOutput struct {
	Capacity uint64  `json:"capacity"`
	Lock     Script  `json:"lock"`
	Type     *Script `json:"type"`
}

type CellDep struct {
	OutPoint OutPoint `json:"out_point"`
	DepType_ DepType  `json:"dep_type"`
}

type Transaction struct {
	Version     Uint32       `json:"version"`
	CellDeps    []CellDep    `json:"cell_deps"`
	HeaderDeps  []H256       `json:"header_deps"`
	Inputs      []CellInput  `json:"inputs"`
	Outputs     []CellOutput `json:"outputs"`
	Witnesses   []JsonBytes  `json:"witnesses"`
	OutputsData []JsonBytes  `json:"outputs_data"`
}

type Header struct {
	Version          Uint32   `json:"version"`
	CompactTarget    Uint32   `json:"compact_target"`
	ParentHash       H256     `json:"parent_hash"`
	Timestamp        uint64   `json:"timestamp"`
	Number           uint64   `json:"number"`
	Epoch            uint64   `json:"epoch"`
	TransactionsRoot H256     `json:"transactions_root"`
	ProposalsHash    H256     `json:"proposals_hash"`
	UnclesHash       H256     `json:"uncles_hash"`
	Dao              [32]byte `json:"dao"`
	Nonce            Uint128  `json:"nonce"`
}

type UncleBlock struct {
	Header    Header            `json:"header"`
	Proposals []ProposalShortId `json:"proposals"`
}
type Block struct {
	Header       Header            `json:"header"`
	Uncles       []UncleBlock      `json:"uncles"`
	Transactions []Transaction     `json:"transactions"`
	Proposals    []ProposalShortId `json:"proposals"`
}
