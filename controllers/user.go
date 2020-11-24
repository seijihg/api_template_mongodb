package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/seijihg/api_template_mongodb/lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User struct definition. Uppercase for the JSON package to see their value.
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Name     string             `json:"name" bson:"name"`
	Surname  string             `json:"surname" bson:"surname"`
	Password string             `json:"password" bson:"password" validate:"required"`
	Dob      primitive.DateTime `json:"dob" bson:"dob"`
}

// CreateUser POST request
func CreateUser(golangDB *mongo.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Decode ERROR:", err)
			lib.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Validation User Struct at this point.
		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")

		validate := validator.New()
		en_translations.RegisterDefaultTranslations(validate, trans)
		err = validate.Struct(user)

		if err != nil {
			errs := err.(validator.ValidationErrors)
			lib.WriteResponse(w, http.StatusBadRequest, errs.Translate(trans))
			return
		}

		// usersCollection := golangDB.Collection("users")

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

		// res, err := usersCollection.InsertOne(context.TODO(), user)
		// if err != nil {
		// 	fmt.Println("InsertOne ERROR:", err)
		// 	lib.WriteResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// fmt.Println("Response from DB:", res)

		// lib.WriteResponse(w, http.StatusOK, res)
	}
}
