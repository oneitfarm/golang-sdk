package cienv

import "fmt"

type MysqlConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
}

func (c *MysqlConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Password)
}
