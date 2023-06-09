package main

import (
	"html/template"
	// "log"
	"net/http"
	"os"
)

// func main() {
// 	templateIfAndRange()
// 	http.HandleFunc("/guestbook", viewHandlertest)
// 	err := http.ListenAndServe("localhost:8080", nil)
// 	log.Fatal(err)
// }

func templateTest() {
	templateText := "template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}

func executeTemplate(text string, data interface{}) {
	temp, err := template.New("test").Parse(text)
	check(err)
	err = temp.Execute(os.Stdout, data)
	check(err)
}

func templateIfAndRange() {
	executeTemplate("start {{if.}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if.}} Dot is true!{{end}} finish\n", false)
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []int{1, 2, 3})
	type Subscriber struct {
		Name   string
		Rate   float64
		Active bool
	}
	templateText = "Name: {{.Name}}\n{{if.Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Anne Yang", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Anne wang", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)

}

func viewHandlertest(writer http.ResponseWriter, request *http.Request) {
	// placeholder := []byte("signature list goes here")
	// _, err := writer.Write(placeholder)
	// check(err)
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
