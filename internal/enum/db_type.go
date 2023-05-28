package enum

type DBType string

const (
	DBTypePostgres DBType = "postgres"
	DBTypeMariaDB  DBType = "mariadb"
)

func (t DBType) Validate() bool {
	switch t {
	case DBTypePostgres, DBTypeMariaDB:
		return true
	}

	return false
}

func (t DBType) String() string {
	return string(t)
}
