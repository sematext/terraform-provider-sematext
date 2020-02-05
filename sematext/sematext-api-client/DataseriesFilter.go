package sematext

type DataSeriesFilter struct {
	aggregation string // TODO Enum: [ NONE, SUM, AVG, MIN, MAX ]
	multiValue  bool
	values      []string
}
