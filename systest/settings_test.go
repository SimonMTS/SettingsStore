package systest

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
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
			compose, err := tc.NewDockerCompose("testresources/docker-compose.yml")
			assert.NoError(t, err, "NewDockerComposeAPI()")

			t.Cleanup(func() {
				assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
			})

			ctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)

			assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

			// do some testing here

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
