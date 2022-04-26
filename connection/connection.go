package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetConnection(host, username, password, dbname, port string) (*gorm.DB, error) {
	const dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	return gorm.Open(postgres.Open(fmt.Sprintf(dsn, host, username, password, dbname, port)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //Para que no use el nombre de las tablas en plural
			},
		})
}
