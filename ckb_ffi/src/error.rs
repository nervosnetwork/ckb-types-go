use derive_more::From;

#[derive(Debug, From)]
pub enum Error {
    SerdeJson(serde_json::error::Error),
    Molecule(ckb_types::error::VerificationError),
    UnsupportedType,
}
