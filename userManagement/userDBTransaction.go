package userManagement

import (
	"database/sql"
	"fmt"
	"os"
)

var userRepository *SQLiteRepository

func dbInit() {
	fmt.Println("UserDBTransaction:dbINIT!")
	db, err := sql.Open(DBDriver, DBFileWithPath)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:dbInit" + err.Error())
	} else {
		userRepository = NewSQLiteRepository(db)
	}
}
func dbReset() {
	fmt.Println("UserDBTransaction:dbReset!")
	os.Remove(DBFileWithPath)
	er := os.MkdirAll(DBFilePath, 0777)
	if er != nil {
		fmt.Println("UserDBTransaction:Remove!- Unable to create directory")
	}
	fmt.Println("UserDBTransaction:Remove!")
	db, err := sql.Open(DBDriver, DBFileWithPath)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:dbReset1" + err.Error())
	}

	userRepository = NewSQLiteRepository(db)

	if err := userRepository.Migrate(); err != nil {
		fmt.Println("Error:UserDBTransaction:dbReset2" + err.Error())
	}

	createdGoSuperuser, err := userRepository.Create(goSuperuser)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:dbReset3" + err.Error())
	} else {
		fmt.Println("SU recorde created" + createdGoSuperuser.Name)
	}

	createdGoSystemUser, err := userRepository.Create(goSystemUser)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:dbReset4" + err.Error())
	} else {
		fmt.Println("test recorde created" + createdGoSystemUser.Name)
	}
}
func GetAllUsers() ([]UserResponse, error) {
	all, err := userRepository.All()
	if err != nil {
		fmt.Println("Error:UserDBTransaction:GetAllUsers1" + err.Error())
		return nil, err
	}

	fmt.Printf("\nAll Users:\n")
	for _, user := range all {
		fmt.Printf("User: %+v\n", user)
	}
	return all, nil
}

func GetUserByName(name string) (*UserResponse, error) {
	usr, err := userRepository.GetByName(name)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:GetUserByName1" + err.Error())
		return nil, err
	} else {
		fmt.Printf("get by name: %+v\n", usr)
		return usr, nil
	}
}
func GetUsrByID(id string) (*UserResponse, error) {
	usr, err := userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:GetUsrByID1" + err.Error())
		return nil, err
	} else {
		fmt.Printf("get by ID: %+v\n", usr)
		return usr, nil
	}
}

func CreateUsr(user User) (*User, error) {
	createdUser, err := userRepository.Create(user)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:CreateUsr1" + err.Error())
		return nil, err
	}
	return createdUser, nil
}

func AuthenticateUsr(user User) (*UserResponse, error) {
	validatedUser, err := userRepository.Validate(user.UserName, user.Password)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:AuthenticateUsr1" + err.Error())
		return nil, err
	}
	return validatedUser, nil
}

func UpdateUsr(id string, updatedUser User) (*User, error) {
	al, err := userRepository.Update(id, updatedUser)
	if err != nil {
		fmt.Println("Error:UserDBTransaction:UpdateUsr1" + err.Error())
		return nil, err
	}
	return al, nil
}
