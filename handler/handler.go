package handler

import (
	//"fmt"
	"bwa-golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)

		return
	}

	// data := map[string]interface{}{
	// 	"title": "I'm Learning Golang Web",
	// 	"content": "I'm Learning Golang Web with Panji Hadjarati",
	// }

	// data := entity.Product{
	// 	ID:    1,
	// 	Name:  "Macbook Pro M1 2021",
	// 	Price: 48000000,
	// 	Stock: 3,
	// }

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 5},
		{ID: 2, Name: "Xpander", Price: 300000000, Stock: 4},
		{ID: 3, Name: "Pajero Sport", Price: 1000000000, Stock: 3},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)

		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar Golng Web"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product Page : %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

}
