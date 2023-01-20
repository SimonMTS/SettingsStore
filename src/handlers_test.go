package src_test

import (
	"database/sql"
	"fmt"
	"settingsstore/src"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	sqlDB, err := sql.Open("ramsql", "Test")
	fmt.Println(err)
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	fmt.Println(err)
	// src.Migrate(db)
	// todo: close
	_ = db
}

func (suite *ExampleTestSuite) TestExample() {
	// expectedSetting := src.Setting{
	// 	ID:    42,
	// 	Type:  "default",
	// 	Value: "some value",
	// 	End:   time.Now(),
	// }
	// inputSetting := models.Setting{
	// 	ID:    &expectedSetting.ID,
	// 	Type:  &expectedSetting.Type,
	// 	Value: &expectedSetting.Value,
	// 	End:   &models.DateTime{Time: expectedSetting.End},
	// }

	// _, _ = expectedSetting, inputSetting
	// _ = suite.handler.AddSetting(operations.AddSettingParams{Setting: &inputSetting}, nil)

	// assert.Equal(suite.T(), operations.NewAddSettingCreated(), result)
	// assert.Equal(suite.T(), expectedSetting, suite.SettingInDB())
}

func (suite *ExampleTestSuite) SettingInDB() (s src.Setting) {
	suite.db.Take(&s)
	return
}
