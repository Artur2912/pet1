package config

import (
	"fmt"
	
	"os"
	"github.com/Artur2912/pet1/models"
	
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Config struct{
	DB *gorm.DB
	AppConfig *AppConfig
}


type DBConfig struct{
	Host string
	Port string
	User string
	Name string
	Password string
}


type AppConfig struct{
	JWT string
	Port string
}


func LoadEnv()error{
	if err := godotenv.Load(); err != nil{
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func NewDBconfig () *DBConfig{
	return &DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Name: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func (v *DBConfig) DSN()string{
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
						v.Host, v.Port, v.User, v.Name, v.Password)						
}

func NewDB(cfg *DBConfig)(*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil{
		return  nil, fmt.Errorf("error: %w", err)
	}
	return  db, nil
}

func RunMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

func NewConfigApp()*AppConfig{
	return &AppConfig{
		JWT: os.Getenv("JWT_SECRET"),
		Port: os.Getenv("JWT_PORT"),
	}
}

func LoadConfig()(*Config,error){
	if err := LoadEnv(); err != nil {
    return nil, err
}

	DBConfig := NewDBconfig()

	ConfigApp := NewConfigApp()

	db, err := NewDB(DBConfig)
	if err != nil{
		return nil, err
	}

	if err := RunMigration(db); err != nil{
		return nil, err
	}

	return &Config{
		DB: db,
		AppConfig: ConfigApp,
	},nil
}