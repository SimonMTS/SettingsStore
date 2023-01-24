package src_test

import (
	"fmt"
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations"
	"settingsstore/src"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

type ExampleTestSuite struct {
	suite.Suite
	db      *gorm.DB
	handler src.Handler
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (suite *ExampleTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	fmt.Println(err)
	err = src.Migrate(db)
	fmt.Println(err)

	suite.handler = src.Handler{
		Database: db,
	}
	suite.db = db
}

func (suite *ExampleTestSuite) TearDownTest() {
	// suite.db.Close()
}

func (suite *ExampleTestSuite) TestExample() {
	expectedSetting := src.Setting{
		ID:    42,
		Type:  "default",
		Value: "some value",
		End:   time.Now().UTC(),
	}
	inputSetting := models.Setting{
		ID:    &expectedSetting.ID,
		Type:  &expectedSetting.Type,
		Value: &expectedSetting.Value,
		End:   &models.DateTime{Time: expectedSetting.End},
	}

	// _, _ = expectedSetting, inputSetting
	result := suite.handler.AddSetting(operations.AddSettingParams{Setting: &inputSetting}, nil)

	assert.Equal(suite.T(), operations.NewAddSettingCreated(), result)
	assert.Equal(suite.T(), expectedSetting, suite.SettingInDB())
}

func (suite *ExampleTestSuite) SettingInDB() (s src.Setting) {
	suite.db.Take(&s)
	return
}
