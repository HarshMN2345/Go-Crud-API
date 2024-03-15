package main
import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"strconv"
	"math/rand"
)
type Movie struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}
var movies []Movie
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main(){
	r:=mux.NewRouter()
	movies = append(movies,Movie{ID:"1",Name:"The Shawshank Redemption",Year:"1994",Director:&Director{FirstName:"Frank",LastName:"Darabont"}})
	movies = append(movies,Movie{ID:"2",Name:"The Godfather",Year:"1972",Director:&Director{FirstName:"Francis",LastName:"Ford Coppola"}})
	movies = append(movies,Movie{ID:"3",Name:"The Dark Knight",Year:"2008",Director:&Director{FirstName:"Christopher",LastName:"Nolan"}})
	movies = append(movies,Movie{ID:"4",Name:"The Godfather: Part II",Year:"1974",Director:&Director{FirstName:"Francis",LastName:"Ford Coppola"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}