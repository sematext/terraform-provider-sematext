package sematext

type AppMetadata struct {
	awsCloudWatchAccessKey string
	awsCloudWatchSecretKey string
	awsFetchFrequency      string   // TODO Enum:[ MINUTE, FIVE_MINUTES, FIFTEEN_MINUTES ]
	awsRegion              string   // TODO Enum: [ US_EAST_1, US_WEST_1, EU_WEST_1, US_WEST_2, AP_SOUTHEAST_1, AP_SOUTHEAST_2, AP_NORTHEAST_1, SA_EAST_1, GovCloud, CN_NORTH_1, US_EAST_2, AP_SOUTH_1, AP_NORTHEAST_2, CA_CENTRAL_1, EU_CENTRAL_1, EU_WEST_2 ]
	subTypes               []string //TODO array of enum example: aws_ec2,aws_elb Comma separated list of AWS types monitored by created app Enum:[ aws_ec2, aws_elb, aws_ebs ]]
}
