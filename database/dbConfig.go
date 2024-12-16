package database
import (
	"gomod/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	
dbUser := os.Getenv("DB_USER")
dbPass := os.Getenv("DB_PASS")
dbHost := os.Getenv("DB_HOST")
dbPort := os.Getenv("DB_PORT")
dbName := os.Getenv("DB_NAME")

connectionString:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",dbUser, dbPass, dbHost, dbPort, dbName)
   // Open db connection
   
   db, err:= gorm.Open(mysql.Open(connectionString), &gorm.Config{})
   if err != nil{
	  return nil,fmt.Errorf("DB Connection unestablished,failed to connect to db %v", err)
   }


   // Auto migrate models
   err = db.AutoMigrate(&models.Users{},)

   if err!= nil {
	 return nil, fmt.Errorf("database migration failed %v", err)
   }
    
   return db, nil
}

func GetDB() (*gorm.DB, error) {
    db, err := InitDB()
	return db , err 
}


