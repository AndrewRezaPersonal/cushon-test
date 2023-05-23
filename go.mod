module cushonTest

go 1.20

require github.com/go-sql-driver/mysql v1.7.1

replace testDatabase => \testDatabase
require testDatabase v0.0.1

replace types => \types
require types v0.0.1
