package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Person struct {
	ID       primitive.ObjectID `json"_id,omitempty" bson:"_id,omitempty"`
	Nombre   string             `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Apellido string             `json:"apellido,omitempty" bson:"apellido,omitempty"`
}

var participante *mongo.Client

func CrearPersonaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	var persona Person
	_ = json.NewDecoder(req.Body).Decode(&persona)
	lista := participante.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	resultado, _ := lista.InserOne(ctx, persona)
	json.NewDecoder(w).Encode(resultado)
}

func main() {
	fmt.Println("comenzando")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	participante, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	router.HandleFunc("/persona", CrearPersonaEndpoint).Methods("POST")
	http.ListenAndServe(":12345", router)
}
