package database

import (
	lg "bmcat/src/daemon/log"
)

type ConnData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	Network string `json:"network"`
	Server  string `json:"server"`
	Port    int    `json:"port"`
	DbName  string `json:"dbname"`
}

var MainSqlite Sqlite

func createTables() {
	sql := `CREATE TABLE "conn" (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"type" integer,
		"user" TEXT,
		"pwd" TEXT,
		"network" TEXT,
		"server" TEXT,
		"port" integer,
		"dbname" TEXT
	  );`

	_, err := MainSqlite.conn.Exec(sql)
	if err != nil {
		lg.ELogger.Println("create table ", err)
		return
	}

	lg.ELogger.Println("create table success")
}
func InitMainSqlite(path string) {
	MainSqlite.Path = path
	r := MainSqlite.Connect()
	if !r {
		lg.ELogger.Println("MainSqlite connect fail")
	}

	createTables()
}

func AddConnect(d ConnData) (int64, error) {
	res, err := MainSqlite.conn.Exec("INSERT INTO conn (name,type,user,pwd, network,server,port,dbname) VALUES(?,?,?,?,?,?,?,?)",
		d.Name, d.Type, d.User, d.Pwd, d.Network, d.Server, d.Port, d.DbName)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetConnects() ([]ConnData, error) {
	var cds []ConnData
	rows, err := MainSqlite.conn.Query("SELECT * FROM conn")
	if err != nil {
		return cds, err
	}

	for rows.Next() {
		var cd ConnData
		err = rows.Scan(&cd.Id, &cd.Name, &cd.Type, &cd.User, &cd.Pwd, &cd.Network, &cd.Server, &cd.Port, &cd.DbName)
		if err != nil {
			continue
		}

		cds = append(cds, cd)
	}

	return cds, nil
}
