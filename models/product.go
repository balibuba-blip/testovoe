package models

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitCost  float64 `json:"unit_cost"`
	MeasureID int     `json:"measure_id"`
}
