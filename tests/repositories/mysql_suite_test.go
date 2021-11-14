package repositories

import (
	"database/sql"
	"log"
	"testing"

	"github.com/dwadp/todos-api/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const connStr = "root:@tcp(127.0.0.1:3306)/todos_test?charset=utf8mb4&parseTime=True&loc=Local"

type MysqlRepositoryTestSuite struct {
	gormDB *gorm.DB
	db     *sql.DB
	suite.Suite
}

func (m *MysqlRepositoryTestSuite) SetupSuite() {
	dbConn, err := db.Connect(connStr)
	if err != nil {
		log.Fatal(err)
	}
	m.gormDB = dbConn

	sqlDB, _ := dbConn.DB()
	m.db = sqlDB
}

func TestMysqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &MysqlRepositoryTestSuite{})
}

func (mr *MysqlRepositoryTestSuite) SetupTest() {
	driver, _ := mysql.WithInstance(mr.db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://../../db/migration", "mysql", driver)

	assert.NoError(mr.T(), err)

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			// just ignore
			return
		}

		panic(err)
	}
}

func (mr *MysqlRepositoryTestSuite) TearDownTest() {
	driver, _ := mysql.WithInstance(mr.db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://../../db/migration", "mysql", driver)

	assert.NoError(mr.T(), err)
	assert.NoError(mr.T(), m.Down())
}
