package cmd

import (
	"github.com/sincerefly/capybara/base/log"
)

func Execute() {

	log.InitLogger() // init log

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
