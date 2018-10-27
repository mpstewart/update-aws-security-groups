package main

import (
	"github.com/mpstewart/update-aws-security-groups/utils"
	"log"
	"os"
)

// Prepare the environment, handle whether or not we should actually run
func main() {
	setAWSProfileEnvar()

	if ipChanged() {
		log.Println("Updating AWS Security Groups...")
		updateIP()
		updatePorts()
		log.Println("Finished updating AWS Security Groups")
	}

	os.Exit(0)
}

// Update and persist the IP address in the config
func updateIP() {
	config := utils.GetConfig()
	realIP := utils.GetHomeIP()

	config.HomeIP = realIP
	config.Write()

	return
}

// Range through ports found in config, and update the security rules
func updatePorts() {
	config := utils.GetConfig()

	for _, port := range config.Ports {
		utils.UpdateRuleForPort(port)
	}
}

// Whether or not the hostname points at an IP other than what is stored
func ipChanged() (b bool) {
	config := utils.GetConfig()
	storedIP := config.HomeIP
	realIP := utils.GetHomeIP()

	if realIP != storedIP {
		b = true
	} else {
		b = false
	}

	return
}

// AWS expects an AWS_PROFILE to be set; this does that
func setAWSProfileEnvar() {
	config := utils.GetConfig()

	profile := config.AWSProfile

	err := os.Setenv("AWS_PROFILE", profile)

	if err != nil {
		utils.Logger.Fatalf("Unable to set AWS_PROFILE:\n%s", err)
	}

	return
}
