package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCloudName() string {
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		return val
	} else {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_CLOUD_NAME")
	}
}

func EnvCloudAPIKey() string {
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		return val
	} else {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_KEY")
	}
}

func EnvCloudAPISecret() string {
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		return val
	} else {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_SECRET")
	}
}

func EnvCloudUploadFolder() string {
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		return val
	} else {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	}
}
