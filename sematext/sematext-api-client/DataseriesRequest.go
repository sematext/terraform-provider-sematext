package sematext

type DataSeriesRequest struct {
	defaultInterval int64
	end             string // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
	filters         map[int]DataSeriesFilter
	granularity     string // TODO - Data points interval granularity between two data points.Default value is “AUTO” - calculated based on selected time span. Not required while getting filters. Enum:[ AUTO, ONE_MINUTE, FIVE_MINUTES, HOUR, DAY, WEEK, MONTH ]
	interval        string
	metric          string
	start           string // TODO Start time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format

}
