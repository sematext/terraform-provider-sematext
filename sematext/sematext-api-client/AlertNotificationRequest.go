package sematext

type AlertNotificationRequest struct {
	defaultInterval int64
	end             string // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
	interval        string
	start           string // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
}
