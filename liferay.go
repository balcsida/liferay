package liferay

import (
	"fmt"
)

// JDBC contains the portal-ext.properties for the created database
// It contains the full properties, meaning, key-value pairs.
type JDBC struct {
	Driver   string
	URL      string
	User     string
	Password string
}

// OracleJDBC returns a JDBC struct which has values corresponding to that
// of Liferay's portal.properties entry for connecting the portal to an
// Oracle database (with the thin driver).
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func OracleJDBC(address, port, sid, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=oracle.jdbc.driver.OracleDriver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:oracle:thin:@%s:%s:%s", address, port, sid),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}

// PostgreJDBC returns a JDBC struct which has values corresponding to that
// of Liferay's portal.properties entry for connecting the portal to a
// PostgreSQL database.
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func PostgreJDBC(address, port, database, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=org.postgresql.Driver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:postgresql://%s:%s/%s", address, port, database),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}

// MysqlJDBC returns a JDBC struct which has values corresponding to that
// of Liferay's portal.properties entry for connecting the portal to a
// Mysql database.
//
// Creates the entry that works with Liferay Portal 6.2 and earlier.
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func MysqlJDBC(address, port, database, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=com.mysql.jdbc.Driver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:mysql://%s:%s/%s?useUnicode=true&characterEncoding=UTF-8&useFastDateParsing=false&useSSL=false", address, port, database),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}

// MariaDBJDBC returns a JDBC struct which has values corresponding to that
// of Liferay's portal.properties entry for connecting the portal to a
// MariaDB database.
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func MariaDBJDBC(address, port, database, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=org.mariadb.jdbc.Driver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:mariadb://%s:%s/%s?useUnicode=true&characterEncoding=UTF-8&useFastDateParsing=false&useSSL=false", address, port, database),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}

// MysqlJDBCDXP returns a JDBC struct which has values corresponding to that
// of Liferay's portal.properties entry for connecting the portal to a
// Mysql database.
//
// Creates the entry that works with Liferay DXP.
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func MysqlJDBCDXP(address, port, database, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=com.mysql.jdbc.Driver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:mysql://%s:%s/%s?characterEncoding=UTF-8&dontTrackOpenResources=true&holdResultsOpenOverStatementClose=true&useFastDateParsing=false&useUnicode=true&useSSL=false", address, port, database),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}

// MSSQLJDBC returns a JDBC struct which has values corresponding to those
// of Liferay's portal.properties entries for connecting the portal to an
// SQl Server database.
//
// Each of its fields corresponds to the full key=value pair from the
// properties file.
func MSSQLJDBC(address, port, database, user, password string) JDBC {
	return JDBC{
		Driver:   "jdbc.default.driverClassName=net.sourceforge.jtds.jdbc.Driver",
		URL:      fmt.Sprintf("jdbc.default.url=jdbc:jtds:sqlserver//%s:%s/%s", address, port, database),
		User:     fmt.Sprintf("jdbc.default.username=%s", user),
		Password: fmt.Sprintf("jdbc.default.password=%s", password),
	}
}
