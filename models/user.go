package models

type User struct {
    Address   string
    Passkey   string
    UUID      string
    Balance   decimal.Decimal
    CreatedAt time.Time
}
