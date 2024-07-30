package aws

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	region       string
	startTime    string
	endTime      string
	startTimeIso string
	endTimeIso   string
)

// CloudTrail represents the CloudTrail command
var CloudTrail = &cobra.Command{
	Use:     "cloudtrail",
	Short:   "Subcommand for terraform scope",
	Aliases: []string{"ct", "events"},
	Run: func(cmd *cobra.Command, args []string) {
		startTimeIso = parseTime(startTime)
		endTimeIso = parseTime(endTime)
		command := fmt.Sprintf("aws cloudtrail lookup-events --output text --region %s --start-time %s --end-time %s --query 'Events[].CloudTrailEvent' | jq -r ' . | select(.errorCode != null) | [.eventTime,.eventID,.eventName,.errorCode,.errorMessage] | @csv'", region, startTimeIso, endTimeIso)
		fmt.Println(111, region)

		process := exec.Command("bash", "-c", command)
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr

		_ = process.Run()
	},
}

func init() {
	CloudTrail.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "Region for which cloudtrail command is executed")
	CloudTrail.PersistentFlags().StringVarP(&startTime, "start", "s", "now", "Start date required to show range")
	CloudTrail.PersistentFlags().StringVarP(&endTime, "end", "e", "now", "End date required to show range")
}

func parseTime(s string) string {
	switch {
	case strings.ToLower(s) == "now":
		return time.Now().Format(time.RFC3339)
	case strings.HasSuffix(s, "d"):
		days, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
		if err != nil {
			log.Fatalf("%s is incorrect value", s)
		}
		return time.Now().Add(time.Duration(-24*days) * time.Hour).Format(time.RFC3339)
	case strings.HasSuffix(s, "h"):
		hours, err := strconv.Atoi(strings.TrimSuffix(s, "h"))
		if err != nil {
			log.Fatalf("%s is incorrect value", s)
		}
		return time.Now().Add(time.Duration(-hours) * time.Hour).Format(time.RFC3339)
	case strings.HasSuffix(s, "m"):
		minutes, err := strconv.Atoi(strings.TrimSuffix(s, "m"))
		if err != nil {
			log.Fatalf("%s is incorrect value", s)
		}
		return time.Now().Add(time.Duration(-minutes) * time.Minute).Format(time.RFC3339)
	default:
		log.Fatal("Incorrect input")
	}
	return ""
}
