package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/seijihg/api_template_mongodb/lib"
	"github.com/seijihg/api_template_mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	customvalidator "github.com/seijihg/api_template_mongodb/customvalidator"
)

// User struct definition. Uppercase for the JSON package to see their value.

// CreateUser POST request
func CreateUser(golangDB *mongo.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Decode ERROR:", err)
			lib.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Response if validation are not passed.
		validationError := customvalidator.CheckUserValid(user)
		if validationError != nil {
			fmt.Println("hitting here")
			lib.WriteResponse(w, http.StatusBadRequest, validationError)
			return
		}

		usersCollection := golangDB.Collection("users")

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	lib.WriteResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// err = json.Unmarshal(body, &user)
		// if err != nil {
		// 	lib.WriteResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// Bcrypting the password.
		bcryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Bcrypt Error:", err)
			lib.WriteResponse(w, http.StatusBadRequest, err.Error())
		}
		user.Password = string(bcryptedPass)

		// Adding Update and Create at.
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// Checking if email already exists.
		var result models.User
		userNotExist := usersCollection.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&result)

		// If user exist it will go to here.
		if userNotExist == nil {
			fmt.Println("User exists")
			lib.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "User already exists"})
			return
		}

		// If user does not exist then Insert to DB.
		res, err := usersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println("InsertOne ERROR:", err)
			lib.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("Response from DB:", res)

		lib.WriteResponse(w, http.StatusOK, res)
	}
}
