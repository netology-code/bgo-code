package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getCards", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("getCards")) // FIXME: для краткости
	})
	mux.HandleFunc("/addCard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("addCard")) // FIXME: для краткости
	})
	mux.HandleFunc("/editCard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("editCard")) // FIXME: для краткости
	})
	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		log.Println(r.Header) // можем доставать все заголовки
		log.Println(r.URL)
		log.Println(r.URL.Query()) // а это уже Values, который мы проходили

		// можно "вытаскивать" данные по ключу (Query + POST Form)
		// возвращает только первое значение или ""
		log.Println(r.FormValue("key"))
		// можно "вытаскивать" данные по ключу (POST Form)
		// возвращает только первое значение или ""
		log.Println(r.PostFormValue("key"))

		// только после ParseForm в r.Form и r.PostForm будут данные
		r.ParseForm() // FIXME: для краткости
		// Values, собираются из Query URL + POST, содержит все значения
		log.Println(r.Form)
		// Values, собираются из POST, содержит все значения
		log.Println(r.PostForm)

		// r.Body -> io.ReadCloser, можем целиком вычитать
	})
	server := &http.Server{
		Addr: "0.0.0.0:9999",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}


