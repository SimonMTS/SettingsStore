package systest

import (
	"testing"

	"github.com/cucumber/godog"
)

func iSendTheDefaultGlobalSettingRequest() error {
	return godog.ErrPending
}

func thereAreSettingsInTheDatabase(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeGlobalSettingInTheDatabase(arg1 int) error {
	return godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			// https://golang.testcontainers.org/features/docker_compose/

			ctx.Step(`^I send the default global setting request$`, iSendTheDefaultGlobalSettingRequest)
			ctx.Step(`^there are (\d+) settings in the database$`, thereAreSettingsInTheDatabase)
			ctx.Step(`^there should be (\d+) global setting in the database$`, thereShouldBeGlobalSettingInTheDatabase)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"."},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
