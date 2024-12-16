package services

import (
	"gomod/database"
	"gomod/models"
	"errors"
)

func CreateUser(user *models.Users) error {
	db, err := database.GetDB()

	if err != nil {
		return err
	}
    //check if the user already exists.
	var existingUser models.Users
	err = db.Where("user_name = ?", user.UserName).First(&existingUser).Error
	//if exist user already exists..ie err is empty we did not recive any value in err
	//then no new user needs to be created
	if err == nil{
         return errors.New("user already exists")
	}
    
	// Check if email already exists
	var existsEmail models.Users
	err = db.Where("email = ?", user.Email).First(&existsEmail).Error
	
	if err == nil{
		return errors.New("email already exists")
    }
    
    return db.Create(user).Error
}

func FindbyUsername(username string) (*models.Users, error) {
     db, err:= database.GetDB()
	 if err != nil {
		return nil, err
	}
	//if username exists
	var existsUser models.Users
	err = db.Where("username = ?", username).First(&existsUser).Error
	// directly returning the user and error
	// if err != nil{
	// 	return fmt.Errorf("user not found")
	// }

	return &existsUser,err 
}

func FindbyEmail(email string) (*models.Users, error) {
	db, err:= database.GetDB()
	//db not found then return nil for db and err
	if err != nil {
		return nil, err
	}
	var existsemail models.Users
    err = db.Where("email= ?", email).First(&existsemail).Error

	return &existsemail, err
}

//why are we always taking the input in err: we  have err as this will return the error from gorm .
//when do we not use return..return 1) when there is a return type 
// 2) when the condition is satisfied , we can return