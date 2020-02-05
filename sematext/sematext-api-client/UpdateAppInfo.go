package sematext

type UpdateAppInfo struct {
	description        string
	ignorePercentage   int32
	maxEvents          int64
	maxLimitMB         int64
	name               string
	sampling           bool
	samplingPercentage int32
	staggering         bool
	status             string // TODO example: ACTIVE  Enum: [ ACTIVE, DISABLED ]
}
