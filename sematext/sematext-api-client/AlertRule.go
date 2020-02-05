package sematext

type AlertRule struct {
	accountEmail                          string // TODO readOnly: true
	allowedAppTypes                       []int64
	analyzingTime                         string
	appDisplayState                       string // TODO readOnly: true
	appId                                 int64
	appName                               string // TODO readOnly: true
	appState                              string // TODO readOnly: true
	appToken                              string // TODO readOnly: true
	appType                               string // TODO readOnly: true
	backToNormalNeeded                    bool
	chartKey                              string
	color                                 string
	creatorEmail                          string // TODO readOnly: true
	defaultAggType                        string // TODO readOnly: true
	description                           string
	disallowedAppTypes                    []int64
	enabled                               bool
	estimateOperation                     string // TODO Enum: [] LESS, MORE, EQUAL, UN_EQUAL, LESS_OR_EQUAL, MORE_OR_EQUAL ]
	estimateValue                         float64
	filterValues                          string
	filterValuesObj                       []FilterValue
	ignoreRegularEventsEnabled            bool
	integrations                          string            // TODO readOnly: true
	lastDataReceivedDate                  int64             // TODO readOnly: true
	lastSent                              int64             // TODO readOnly: true
	lastTriggered                         int64             // TODO readOnly: true
	metadata                              map[string]string // TODO mixed freeform in API?
	metricKey                             string            // TODO readOnly: true
	metricLabel                           string
	minDelayBetweenNotificationsInMinutes string
	name                                  string
	notificationEmails                    []string
	notificationIntegrations              []NotificationIntegration
	notificationsEnabled                  bool
	query                                 string
	reportName                            string
	ruleKey                               int64 // TODO readOnly: true
	ruleType                              string
	runbook                               string
	savedQueryId                          int64 // TODO readOnly: true
	schedule                              []AlertRuleScheduleWeekdayDto
	sematextService                       string // TODO readOnly: true
	sendToEmail                           string
	timezone                              string
	useOnlyAlertRuleIntegrations          bool
	userPermissions                       UserPermissions
	valueColumnName                       string // TODO readOnly: true
	valueName                             string // TODO readOnly: true
}
