package main

import (
	"html/template"
	"log"
	"net/http"
)

const templ = `<html><head><title>Hello</title></head><body>
こんにちは {{ .RemoteAddr }}
あなたは {{ .Method }} リクエストを {{ .URL }} へ送信しました。
</body></html>`

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	t := template.Must(template.New("html").Parse(templ))
	t.Execute(w, req)
}

func main() {
	log.Println("http://localhost:7777/ へ接続してください")
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
