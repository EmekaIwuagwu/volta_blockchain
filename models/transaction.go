package models

type Transaction struct {
    AddressFrom       string
    AddressTo         string
    Amount            decimal.Decimal
    DateOfTransaction time.Time
    CreatedAt         time.Time
}
