package sematext

type App struct {
	ajaxThreshold         int64
	appType               string
	appTypeId             int64
	creatorEmail          string
	creditCardExpiry      string
	creditCardNumber      string
	description           string
	displayStatus         string
	firstDataSavedDate    int64
	id                    int64
	integration           sematext.ServiceIntegration
	lastDataReceivedDate  int64
	lastDataSavedDate     int64
	loggedInUserAppRole   string
	monthlyInvoiceAccount bool
	name                  string
	ownerEmail            string
	owningOrganization    sematext.BasicOrganizationDto
	pageLoadThreshold     int64
	paymentMethodId       int64
	plan                  sematext.Plan
	prepaidAccount        bool
	status                string
	token                 string
	trialEndDate          int64
	urlGroupLimit         int32
	userRoles             []sematext.UserRole
}
