package entity

import "time"

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusRunning   TransactionStatus = "running"
	StatusCompleted TransactionStatus = "completed"
	StatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID        string
	Status    TransactionStatus
	Sender    string
	Receiver  string
	Amount    float64
	CreatedAt time.Time
}
