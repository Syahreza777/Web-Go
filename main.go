package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type DaftarRoutePost struct{
	Post string
	Link string
}

type Index struct {
	RoutePost []DaftarRoutePost
	Text string
}

type DaftarKomen struct{
	Nama string
	Komen string
}

type Post struct{
	Deskripsi string
	Message []DaftarKomen
	ImgSrc string
	Title string
}



func main() {


	// Index
	var tmpl, err = template.ParseGlob("views/*")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var index = Index{
			RoutePost: []DaftarRoutePost{
				{Post: "Judul Posting 1", Link: "/post/1"},
				{Post: "Judul Posting 2", Link: "/post/2"},
				{Post: "Judul Posting 3", Link: "/post/3"},
			},
			Text: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias rem dolore minima, ipsum voluptates, dolorem, deserunt doloremque mollitia harum commodi ut assumenda reprehenderit! Accusamus dignissimos officia unde harum animi veritatis.",
		}

		err = tmpl.ExecuteTemplate(w, "index", index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	// Post 1
	http.HandleFunc("/post/1", func(w http.ResponseWriter, r *http.Request) {
		var post_1 = Post{
			ImgSrc: "/img/hei.jpg",
			Deskripsi: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Message: []DaftarKomen{
				{Nama: "Yaza", Komen: "Bagus"},
				{Nama: "Apis", Komen: "Ajarin dong mas"},
				{Nama: "Irvan", Komen: "Malah lebih pro aku"},
			},
			Title: "Judul Postingan 1",
		}

		err = tmpl.ExecuteTemplate(w, "post", post_1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	// Post 2
	http.HandleFunc("/post/2", func(w http.ResponseWriter, r *http.Request) {
		var post_2 = Post{
			ImgSrc: "/img/irvanCupu.PNG",
			Deskripsi: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Message: []DaftarKomen{
				{Nama: "Deddy", Komen: "Keren ja"},
				{Nama: "Erick", Komen: "Ok lah"},
				{Nama: "Bocil", Komen: "Sipp"},
			},
			Title: "Judul Postingan 2",
		}

		err = tmpl.ExecuteTemplate(w, "post", post_2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	// Post 3
	http.HandleFunc("/post/3", func(w http.ResponseWriter, r *http.Request) {
		var post_3 = Post{
			ImgSrc: "/img/irvanCupu1.PNG",
			Deskripsi: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Message: []DaftarKomen{
				{Nama: "Ridho", Komen: "Good job"},
				{Nama: "Samsul", Komen: "Bagus Juga"},
				{Nama: "Arip", Komen: "Josss"},
			},
			Title: "Judul Postingan 3",
		}

		err = tmpl.ExecuteTemplate(w, "post", post_3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})



	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.Handle("/img/",
		http.StripPrefix("/img/",
			http.FileServer(http.Dir("image"))))



	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)

}
