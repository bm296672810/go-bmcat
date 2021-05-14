package httpServer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"bmcat/src/daemon/database"
	lg "bmcat/src/daemon/log"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)
var conns map[int]database.Db

// mux = make(map[string]func(http.ResponseWriter, *http.Request))

type MyHandler struct {
	// mux := make(map[string]func(http.ResponseWriter, *http.Request))
}

func initHandler() {
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	// mux["/user/login"] = Login
	// curl localhost:2200/add/connect -X POST -d '{"name": "database", "type": 0, "server": "./database.db"}' --header "Content-Type: application/json"
	mux["/add/connect"] = AddConnect
	// curl localhost:2200/get/connects -X GET --header "Content-Type: application/json"
	mux["/get/connects"] = GetConnects
	// curl localhost:2200/get/tables&id=1 -X GET --header "Content-Type: application/json"
	mux["/get/tables"] = GetTables

	conns = make(map[int]database.Db)
}

func (lh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		body, err := getRequestBody(w, r)
		if err != nil {
			lg.ELogger.Println("getRequestBody:", err)
			return
		}
		// w.Write(w, string(body))
		w.WriteHeader(http.StatusForbidden)
		w.Write(body)
		return
	}
	url := r.URL
	lg.ILogger.Println("url:", url.String())

	if h, ok := mux[url.Path]; ok {
		h(w, r)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	// io.WriteString(w, "URL:"+r.URL.String())
}

func AddConnect(w http.ResponseWriter, r *http.Request) {
	body, err := getRequestBody(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var d database.ConnData
	err = parseJson(body, &d)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := database.AddConnect(d)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	d.Id = int(id)
	js, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
	// w.Write("add connect success")
}

func GetConnects(w http.ResponseWriter, r *http.Request) {
	cds, err := database.GetConnects()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	go func(cds []database.ConnData) {
		for _, d := range cds {
			td := conns[d.Id]
			if td != nil {
				continue
			}

			switch d.Type {
			case 0:
				var db database.Db
				db = &database.Sqlite{Path: d.Server, Tp: d.Type}
				if db.Connect() {
					conns[d.Id] = db
				}
			}
		}
	}(cds)

	js, err := json.Marshal(cds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetTables(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	lg.ILogger.Println("values:", values)
	v := values.Get("id")
	lg.ILogger.Println("id:", v)
	// database.GetTables()
	id, err := strconv.Atoi(v)
	if err != nil {
		lg.ELogger.Println("id to int error:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	db := conns[id]
	if db != nil {
		tbs := db.Tables()
		if len(tbs) > 0 {
			js, err := json.Marshal(tbs)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(js)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func getRequestBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// ud.ELogger.Println("read body error!")
		lg.ELogger.Println("read body error!")
		return body, err
	}
	return body, err
}

func parseJson(d []byte, v interface{}) error {
	err := json.Unmarshal(d, v)
	if err != nil {
		lg.ELogger.Println("parse json error:", err)
		lg.ELogger.Println("data:", string(d))
		return err
	}
	return nil
}
