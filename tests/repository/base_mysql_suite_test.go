package repository

import (
	"database/sql"
	"log"

	"github.com/dwadp/todos-api/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const connStr = "root:@tcp(127.0.0.1:3306)/todos_test?charset=utf8mb4&parseTime=True&loc=Local"

type BaseMysqlTestSuite struct {
	gormDB *gorm.DB
	db     *sql.DB
	suite.Suite
}

func (m *BaseMysqlTestSuite) SetupSuite() {
	dbConn, err := db.Connect(connStr)
	if err != nil {
		log.Fatal(err)
	}
	m.gormDB = dbConn
	sqlDB, _ := dbConn.DB()
	m.db = sqlDB
}

func (m *BaseMysqlTestSuite) SetupTest() {
	migration, err := createMigration(m.db)

	assert.NoError(m.T(), err)
	if err := migration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		panic(err)
	}
}

func (m *BaseMysqlTestSuite) TearDownTest() {
	migration, err := createMigration(m.db)
	assert.NoError(m.T(), err)
	assert.NoError(m.T(), migration.Down())
}

func createMigration(db *sql.DB) (*migrate.Migrate, error) {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	return migrate.NewWithDatabaseInstance("file://../../db/migration", "mysql", driver)
}
