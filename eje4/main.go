package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// referiacia a la carpéta viwe
var plantilla = template.Must(template.ParseGlob("view/*"))

type Emplado struct {
	Id    int
	Name  string
	Email string
}

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := "root"
	Name := "sistema"
	Host := "localhost"

	conexion, err := sql.Open(Driver, User+":"+Password+"@tcp("+Host+")/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return conexion

}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delecte", Delecte)
	http.HandleFunc("/showUpdate", ShowUpdate)
	http.HandleFunc("/update", Update)

	log.Println("Server corriendo...")
	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {

	conex := conexionDB()

	registo, err := conex.Query("select * from empleado;")

	if err != nil {
		panic(err.Error())
	}

	empleado := Emplado{}

	aregloEm := []Emplado{}

	for registo.Next() {
		var id int
		var name string
		var email string
		err = registo.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Name = name
		empleado.Email = email

		aregloEm = append(aregloEm, empleado)
	}

	plantilla.ExecuteTemplate(w, "inicio", aregloEm)

}

func Crear(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hola weii")
	plantilla.ExecuteTemplate(w, "create", nil)

}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		nama := r.FormValue("name")
		emal := r.FormValue("email")

		conex := conexionDB()
		inser, err := conex.Prepare("INSERT into empleado (`name`,email) VALUE (?,?);")
		if err != nil {
			panic(err.Error())
		}
		inser.Exec(nama, emal)
		http.Redirect(w, r, "/", 301)
	}
}

func Delecte(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")

	conex := conexionDB()
	delecte, err := conex.Prepare("DELETE FROM empleado WHERE id=?;")
	if err != nil {
		panic(err.Error())
	}
	delecte.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)

}
func ShowUpdate(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	conex := conexionDB()
	registro, err := conex.Query("Select * FROM empleado WHERE id=?;", idEmpleado)

	if err != nil {
		panic(err.Error())
	}

	empleado := Emplado{}

	for registro.Next() {
		var id int
		var name string
		var email string
		err = registro.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Name = name
		empleado.Email = email
	}

	plantilla.ExecuteTemplate(w, "edit", empleado)

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		conex := conexionDB()
		update, err := conex.Prepare("Update empleado set name=?,email=? where id=?;")
		if err != nil {
			panic(err.Error())
		}
		update.Exec(name,email, id)
		http.Redirect(w, r, "/", 301)
	}

}
