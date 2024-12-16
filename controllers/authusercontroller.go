package controller

import (
	"gomod/models"
	"gomod/services"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//in the controller we will create 1) Registation request 2)login request for authorization

type RegistrationRequest struct{
	UserName string `json:"username`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string  `json:"email"`
	Password string `jsom:"password"`
}

// steps:Request
// step1: decode/marshal
// step2: hashpassword
// step3: create user of type Users with hashed password
// step4:return hw.header with success response if no err



func Register(w http.ResponseWriter, r *http.Request){
    var req RegistrationRequest
	err:= json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//hash password
    //we need to hash password to bring in security
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}


	//Create user

	// type UsersMock struct{
	// need to create new user instance with hassed password
	user:= models.Users{
		UserName: req.UserName,
		Email: req.Email,
		Password: string(hashedPassword), 
		IsActive: true,
	}
  
	//create --> user
    err = services.CreateUser(&user)

	if err!= nil{
		http.Error(w,"user not created", http.StatusBadRequest)
	}

	//if no errors
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode( "user registerd successfully")

}

//steps:Login
//step1: check if user correct or not
//step2:check password correctness

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
    err:= json.NewDecoder(r.Body).Decode(&req)
	if err!= nil{
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	//check user correct (by mail id...login takes email and it will be unique)
	user, err:= services.FindbyEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}
    //check for password is correct or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)) //where is actually the password stored
    if err != nil{
		http.Error(w,"user or password is invalid", http.StatusBadRequest)
		return
	}
    //if correct
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
	// return user,nil we havent menstioned any return type



}