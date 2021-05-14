package database

import (
	"database/sql"

	lg "bmcat/src/daemon/log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Tp int
	// conn, err := sql.Open("sqlite3", "./foo.db")
	Path string
	conn *sql.DB
}

func (s *Sqlite) Connect() bool {
	conn, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		return false
	}
	if conn != nil {
		s.conn = conn
		lg.ILogger.Println("connect ", s.Path, " success")
		return true
	}

	return false
}

func (s *Sqlite) IsConnected() bool {
	return s.conn != nil
}
func (s *Sqlite) Tables() []string {
	var r []string
	rows, err := s.conn.Query("SELECT * FROM sqlite_master WHERE type='table'")
	if err != nil {
		return r
	}

	for rows.Next() {
		var tp, name, tblName, rootPage, sql string

		err = rows.Scan(&tp, &name, &tblName, &rootPage, &sql)
		if err != nil {
			continue
		}

		r = append(r, name)
	}

	return r
}
