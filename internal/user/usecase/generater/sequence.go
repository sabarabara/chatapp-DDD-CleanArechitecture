package id



type Sequence struct {
	SequenceName string `gorm:"primaryKey"`
	CurrentValue uint64
}
