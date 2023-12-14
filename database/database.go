package database

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	host              = os.Getenv("host")
	password          = os.Getenv("password")
	user              = os.Getenv("user")
	port              = os.Getenv("port")
	catalog           = os.Getenv("catalog")
	serviceAccountKey = os.Getenv("service_account_key")
)

func GetConnection() (*sql.DB, error) {

	server := fmt.Sprintf("server=%s,%s;user id=%s;password=%s;database=%s;trustservercertificate=true;encrypt=DISABLE",
		host, port, user, password, catalog)

	return sql.Open("mssql", server)
}

func GetDecodedFireBaseKey() ([]byte, error) {

	decodedKey, err := base64.StdEncoding.DecodeString(serviceAccountKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

// func SetupFirebase() (*firebase.App, error) {
// 	decodedKey, err := GetDecodedFireBaseKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	opt := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}

// 	// serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
// 	// if err != nil {
// 	// 	// panic("Unable to load serviceAccountKeys.json file")
// 	// 	return nil, fmt.Errorf("error initializing app load serviceAccountKeys: %v", err)
// 	// }

// 	// opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

// 	app, err := firebase.NewApp(context.Background(), nil, opt...)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}
// 	return app, nil
// }
