package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Mahasiswa struct {
	Id   int
	Nama string
	Kota string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "tb_mahasiswa"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl, err = template.ParseGlob("views/*")

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Mahasiswa ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Mahasiswa{}
	res := []Mahasiswa{}
	for selDB.Next() {
		var id int
		var nama, kota string
		err = selDB.Scan(&id, &nama, &kota)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Kota = kota
		res = append(res, emp)
	}
	err = tmpl.ExecuteTemplate(w, "Index", res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Mahasiswa WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Mahasiswa{}
	for selDB.Next() {
		var id int
		var nama, kota string
		err = selDB.Scan(&id, &nama, &kota)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Kota = kota
	}
	err = tmpl.ExecuteTemplate(w, "Show", emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	err = tmpl.ExecuteTemplate(w, "New", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Mahasiswa WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	emp := Mahasiswa{}
	for selDB.Next() {
		var id int
		var nama, kota string
		err = selDB.Scan(&id, &nama, &kota)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Kota = kota
	}
	err = tmpl.ExecuteTemplate(w, "Edit", emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		kota := r.FormValue("kota")
		insForm, err := db.Prepare("INSERT INTO Mahasiswa(nama, kota) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, kota)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		kota := r.FormValue("kota")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Mahasiswa SET nama=?, kota=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, kota, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Mahasiswa WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	var address = ":9000"
	fmt.Println("Started at ", address)
	http.ListenAndServe(address, nil)
}
