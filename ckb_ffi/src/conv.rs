use crate::error::Error;

use ckb_types::prelude::*;

// json <=> mol
macro_rules! convert {
    (try($target_type:ident, $source:expr, $from_json:expr)) => {{
        if $from_json {
            convert!(from_json($target_type, $source))
        } else {
            convert!(to_json($target_type, $source))
        }
    }};
    (from_json($target_type:ident, $json:expr)) => {{
        let json_struct = serde_json::from_slice::<ckb_jsonrpc_types::$target_type>($json)?;
        let packed_struct: ckb_types::packed::$target_type = json_struct.into();

        Ok(packed_struct.as_slice().to_owned())
    }};
    (to_json($target_type:ident, $mol:expr)) => {{
        let packed_struct = ckb_types::packed::$target_type::from_slice($mol)?;
        let json_struct: ckb_jsonrpc_types::$target_type = packed_struct.into();

        Ok(serde_json::to_vec::<_>(&json_struct)?)
    }};
}

fn convert_tx(source: &[u8], from_json: bool) -> Result<Vec<u8>, Error> {
    if from_json {
        encode_tx(source)
    } else {
        decode_raw_tx(source)
    }
}

// First turn Transaction into RawTransaction, then return its molecule bytes
fn encode_tx(source: &[u8]) -> Result<Vec<u8>, Error> {
    let json_struct = serde_json::from_slice::<ckb_jsonrpc_types::Transaction>(source)?;
    let packed_struct: ckb_types::packed::Transaction = json_struct.into();

    Ok(packed_struct.as_reader().raw().as_slice().to_owned())
}

// Decode molecule RawTransaction bytes, turn into Transaction, then return its
// json bytes
fn decode_raw_tx(source: &[u8]) -> Result<Vec<u8>, Error> {
    let raw_tx = ckb_types::packed::RawTransaction::from_slice(source)?;
    let tx = ckb_types::packed::Transaction::new_builder()
        .raw(raw_tx)
        .build();
    let json_struct: ckb_jsonrpc_types::Transaction = tx.into();

    Ok(serde_json::to_vec::<_>(&json_struct)?)
}

pub fn try_convert(ty: &str, source: &[u8], from_json: bool) -> Result<Vec<u8>, Error> {
    match ty {
        // Chain Types
        "ProposalShortId" => convert!(try(ProposalShortId, source, from_json)),
        "Script" => convert!(try(Script, source, from_json)),
        "OutPoint" => convert!(try(OutPoint, source, from_json)),
        "CellInput" => convert!(try(CellInput, source, from_json)),
        "CellOutput" => convert!(try(CellOutput, source, from_json)),
        "CellDep" => convert!(try(CellDep, source, from_json)),
        "Transaction" => convert_tx(source, from_json),
        "Header" => convert!(try(Header, source, from_json)),
        "UncleBlock" => convert!(try(UncleBlock, source, from_json)),
        "Block" => convert!(try(Block, source, from_json)),
        // Unimplement in ckb_jsonrpc_types
        // "RawTransaction" => convert!(try(RawTransaction, source, from_json)),
        // "RawHeader" => convert!(try(RawHeader, source, from_json)),
        // "ScriptOpt" => convert!(try(ScriptOpt, source, from_json)),
        // "UncleBlockVec" => convert!(try(UncleBlockVec, source, from_json)),
        // "TransactionVec" => convert!(try(TransactionVec, source, from_json)),
        // "ProposalShortIdVec" => convert!(try(ProposalShortIdVec, source, from_json)),
        // "CellDepVec" => convert!(try(CellDepVec, source, from_json)),
        // "CellInputVec" => convert!(try(CellInputVec, source, from_json)),
        // "CellOutputVec" => convert!(try(CellOutputVec, source, from_json)),
        // "CellbaseWitness" => convert!(try(CellbaseWitness, source, from_json)),
        _ => Err(Error::UnsupportedType),
    }
}

#[cfg(test)]
mod tests {
    use super::try_convert;

    #[test]
    fn should_decode_encode_transaction() {
        let tx_json = r#"{
            "version": "0x0",
            "cell_deps": [],
            "header_deps": [],
            "inputs": [],
            "outputs": [],
            "witnesses": [],
            "outputs_data": []
        }"#;

        let mol = try_convert("Transaction", tx_json.as_bytes(), true).expect("try from json fail");
        let json = try_convert("Transaction", mol.as_slice(), false).expect("try to json fail");

        let expect_tx: ckb_jsonrpc_types::Transaction =
            serde_json::from_str(tx_json).expect("tx_json from str fail");
        let tx: ckb_jsonrpc_types::Transaction =
            serde_json::from_slice(json.as_slice()).expect("json from slice fail");

        assert_eq!(tx, expect_tx);
    }

    #[test]
    fn should_decode_encode_block() {
        let blk_json = r#"{
            "header": {
                "version": "0x0",
                "parent_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "timestamp": "0x0",
                "number": "0x0",
                "epoch": "0x0",
                "transactions_root": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "proposals_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "compact_target": "0x0",
                "uncles_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "dao": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "nonce": "0x0"
            },
            "uncles": [],
            "transactions": [],
            "proposals": []
        }"#;

        let mol = try_convert("Block", blk_json.as_bytes(), true).expect("try from json fail");
        let json = try_convert("Block", mol.as_slice(), false).expect("try to json fail");

        let expect_blk: ckb_jsonrpc_types::Block =
            serde_json::from_str(blk_json).expect("blk_json from str fail");
        let blk: ckb_jsonrpc_types::Block =
            serde_json::from_slice(json.as_slice()).expect("json from slice fail");

        assert_eq!(blk, expect_blk);
    }
}
