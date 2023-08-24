package template

import (
	"os"
	"html/template"
)

type Persona struct{
	Name string
	Edad int
	Hob[]string
}
var funcs =	template.FuncMap{
	"incre":func (num int) int  {
		
		return num +1
	},
}
func Init() {

	// Golang tiene dos template
	// text/template
		// html/template



		fabricio := &Persona{"fabricio",25,[]string{">","dormir","programar"}}
	
	loadTemplate("index.html",fabricio)
}
func loadTemplate(fileName string,data interface{}){

	tem,err := template.New(fileName).Funcs(funcs).ParseFiles("template/"+fileName)

	if err != nil {
			panic(err)
	}

	err = tem.Execute(os.Stdout,data)

	if err != nil {
		panic(err)
}

}