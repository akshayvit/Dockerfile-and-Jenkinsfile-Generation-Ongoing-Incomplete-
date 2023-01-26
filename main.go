package main

import (
	"context",
	"log",
	"strings"
	"net/http",
	"go.mongodb.org/mongo-driver/mongo",
	"go.mongodb.org/mongo-driver/mongo/options",
	"go.mongodb.org/mongo-driver/bson/primitive",
	"github.com/gorilla/mux"
)

type dfile struct {
	Container_Name string         
	commands []string   
	src string     
	dest string
	port string
	wdir string
}

var coll *mongo.Collecton
var ctx=context.TODO()

func init() {
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client,err:=mongo.Connect(ctx,clientoptions)
	if err!=nil {
		log.Fatal(err)
	} else {
		err:=client.Ping(ctx,nil)
		if err!=nil {
			log.Fatal(err)
		} else {
			coll=client.Database("")
			return coll
		}
	}
	return nil
}

func home(w http.ResponseWriter,r *http.Request) {
	tec:=string.Split(r.URL.Path,"/")[-2]
	ver:=string.Split(r.URL.Path,"/")[-1]
	dict:=make(map[string]string)
	dict["net6.0"]="microsoft/net/et6.0"
	dict["python3"]="python/pypy3"
	dict["nodejs18"]="node/node18"
	var dfileobj dfile
	fmt.Fprintf(w,`
	   FROM %s
	   WORKDIR .
	   COPY build/app .
	   RUN 
	`)
}