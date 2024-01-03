package main

import (
	//"encoding/json"
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	//"log"
	"net/http"
	"strings"

	//"reflect"
	"github.com/go-chi/chi"
	//"github.com/go-chi/chi/v5"
)

type api_json struct {
    Word string
    Synonyms []string
    Antonyms []string
}

func parseJson(body string)  {
    s := body[1 : len(body)-1]
    parts := strings.Split(s, ", ")
    fmt.Println(len(parts))
    fmt.Println(parts)
}

func sereveRoot(w http.ResponseWriter, r *http.Request)  {
    //w.Write([]byte("Server is Up!, Yet to implement index HTML"))
    fmt.Println(r.URL)
    q := r.URL.Query()
    if q.Get("search") == "" {
        fmt.Println("Empty")
    } else {
        fmt.Println(q.Get("search"))
    }
    tmp, _ := template.ParseFiles("index.html")
    tmp.ExecuteTemplate(w, "index.html", nil)
}

func sereveWord(w http.ResponseWriter, r *http.Request)  {
    fmt.Println(r.URL)
    tmp := template.Must(template.New("main").Parse(fmt.Sprintf("%d",2543)))
    tmp.ExecuteTemplate(w, "main", nil)
}

func serveSearch(w http.ResponseWriter, r *http.Request)  {
    q := r.URL.Query()
    //fmt.Println(r.URL)
    //fmt.Fprintf(w, "%s", q.Get("search"))
    //url_ar := []string{"https://api.api-ninjas.com/v1/thesaurus?word="}
    //fmt.Println(url_ar)
    //url_ar = append(url_ar, q.Get("search"))
    //url := strings.Join(url_ar,"")
    //req, err := http.NewRequest("Get", url, nil)
    //if err != nil {
    //    log.Fatalln(err)
    //}
    //req.Header.Set("X-Api-Key", "ACBq6BHaf8YlgW5OJDEgZw==9jTghNt9a2bmfagu")
    //client := &http.Client{}
    //resp, err := client.Do(req)
    //if err != nil {
    //    log.Fatalln(err)
    //}
    url := []string{"https://api.dictionaryapi.dev/api/v2/entries/en/", q.Get("search")}
    url_s := strings.Join(url, "")
    resp, err := http.Get(url_s)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    body_s := string(body)
    fmt.Println(body_s)
    //var res map[string]interface{}
    var res api_json
    err = json.Unmarshal([]byte(body), &res)
    //syn := res["synonyms"]
    //fmt.Println(reflect.TypeOf(syn))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(res.Synonyms)
    //parseJson(string(body))
}

func main()  {
    fmt.Println("Main started..., http://localhost:3434")
    r := chi.NewRouter()
    r.Get("/", sereveRoot)
    r.Get("/search", serveSearch)
    r.Post("/word", sereveWord)
    http.ListenAndServe("localhost:3434", r)
}
