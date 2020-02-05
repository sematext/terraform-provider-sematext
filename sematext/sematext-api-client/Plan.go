package sematext

type Plan struct {
	appType            string
	custom             bool
	dataRetentionHours int64 // TODO Check datatype
	defaultTrialPlan   bool
	free               bool
	freeTrialDays      int64
	id                 int64
	maxAlerts          int64
	maxDailyEvents     int64
	name               string
	planScheme         string // TODO Enum: [ SPM_1_0, SPM_2_0, SA_1_0, SEARCHENE_1_0, LOGSENE_1_0, LOGSENE_2_0, RUM_1_0, RUM_EA, SYNTHETICS_1_0 ]
	sematextService    string // TODO Enum: [ LOGSENE, SPM, RUM, SYNTHETICS ]
	trialPlan          bool
}
