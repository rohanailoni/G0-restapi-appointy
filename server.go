package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//this connector has the clien object which can bee access to all the
type connector struct {
	sync.Mutex
	client  *mongo.Client
	context *context.Context
}

//structure that contains usermodel
type Usermodel struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Post     []Postmodel `json:"post"`
}
type Postmodel struct {
	UseriD  string `json:"userId"`
	Url     string `json:"url"`
	Caption string `json:"caption"`
	Time    string `json:"Time"`
}

//this recives function http.handlerfunc and deports it
func (h1 *connector) userread(h http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h1.userget(h, r)
		return
	case "POST":
		h1.userpost(h, r)
	default:
		h.WriteHeader(http.StatusMethodNotAllowed)
		h.Write([]byte("method not allowed"))
		return
	}
}

//this function is used to get the values from the user to be used
func (h1 *connector) userget(h http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	var result []bson.M
	coll := h1.client.Database("appointy").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	if !ok || len(keys[0]) < 1 {
		cursor, err := coll.Find(ctx, bson.M{})
		if err != nil {

			log.Fatal(err)
		}
		if err = cursor.All(ctx, &result); err != nil {

			log.Fatal(err)
		}
		jsonbytes, err := json.Marshal(result)
		if err != nil {
			h.WriteHeader(http.StatusInternalServerError)

		}
		h.Header().Add("content-type", "application-json")
		h.WriteHeader(http.StatusOK)
		h.Write(jsonbytes)
		return
	}
	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]
	fmt.Print(key)
	h1.Lock()
	objectId, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		log.Println("Invalid id")
	}
	var ans bson.M

	fmt.Print(ans)
	if err = coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&ans); err != nil {
		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte("NO document in the result\n"))

	}
	h1.Unlock()
	jsonbytes, err := json.Marshal(ans)
	if err != nil {
		h.WriteHeader(http.StatusInternalServerError)

	}
	h.Header().Add("content-type", "application-json")
	h.WriteHeader(http.StatusOK)
	h.Write(jsonbytes)
	return

}

//this values gets the parms from the User and make changes in USER collections in mongodb
func (h1 *connector) userpost(h http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.WriteHeader(http.StatusInternalServerError)
		h.Write([]byte("cannt read the body"))
		return
	}
	var userm Usermodel
	err = json.Unmarshal(bodybytes, &userm)
	if err != nil {
		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte(err.Error()))
		return
	}
	data := md5.Sum([]byte(userm.Password))
	userm.Password = hex.EncodeToString(data[:])
	h1.Lock()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	col := h1.client.Database("appointy").Collection("users")
	insertResult, err := col.InsertOne(ctx, userm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
	h.WriteHeader(http.StatusOK)
	h1.Unlock()
}

func userreader(c mongo.Client, d context.Context) *connector {
	return &connector{client: &c, context: &d}
}

//***********************************************************************now post method functions********************
//this function is also method of connector structure but used for post
func (h1 *connector) postread(h http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h1.postget(h, r) //for sendint the value
		return
	case "POST":
		h1.postsend(h, r) //for getting an saving in database
	default:
		h.WriteHeader(http.StatusMethodNotAllowed)
		h.Write([]byte("method not allowed"))
		return
	}
}

func (h1 *connector) postget(h http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["userid"]
	coll := h1.client.Database("appointy").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	if !ok || len(keys[0]) < 1 {

		h.Header().Add("content-type", "application-json")
		h.WriteHeader(http.StatusOK)
		h.Write([]byte("you cannt get the value userId requried"))
		return
	}
	key := keys[0]
	h1.Lock()
	objectId, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		log.Println("Invalid id")
	}
	var ans Usermodel
	if err = coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&ans); err != nil {
		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte("NO document in the result\n"))
	}
	h1.Unlock()
	jsonbytes, err := json.Marshal(ans.Post)
	if err != nil {
		h.WriteHeader(http.StatusInternalServerError)

	}
	h.Header().Add("content-type", "application-json")
	h.WriteHeader(http.StatusOK)
	h.Write(jsonbytes)
	return

}
func (h1 *connector) postsend(h http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.WriteHeader(http.StatusInternalServerError)
		h.Write([]byte("cannt read the body"))
		return
	}
	var userm Postmodel
	err = json.Unmarshal(bodybytes, &userm)

	if err != nil {

		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte(err.Error()))
		return
	}
	h1.Lock()

	coll := h1.client.Database("appointy").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	objectId, err := primitive.ObjectIDFromHex(userm.UseriD)
	if err != nil {

		h.WriteHeader(http.StatusInternalServerError)
		h.Write([]byte("Invalid ID"))
		return
	}

	var user Usermodel
	if err = coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {

		h.WriteHeader(http.StatusBadRequest)
		h.Write([]byte("NO document in the result\n"))
		return
	}
	var array []Postmodel
	array = append(array, userm)

	result, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.D{
			{"$set", bson.D{{"post", array}}},
		},
	)
	fmt.Print(result)
	h1.Unlock()
	fmt.Print(user)
	h.Header().Add("content-type", "application-json")
	h.WriteHeader(http.StatusOK)
	return

}

//function to connect to mongodb cluster
func mongoconnect() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rohan:Z9RicwzgtBzh4hzI@cluster0.29lux.mongodb.net/Cluster0?retryWrites=true&w=majority"))

	if err != nil {

		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {

		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {

		log.Fatalf("Monogb is not reachable")
	}

	return client, ctx
}
func main() {

	client, context1 := mongoconnect()

	uss := userreader(*client, context1)
	http.HandleFunc("/users", uss.userread)
	http.HandleFunc("/post", uss.postread)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
