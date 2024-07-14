package models

type Block struct {
    BlockID      string
    Timestamp    time.Time
    Hash         string
    PreviousHash string
    Address      string
}
