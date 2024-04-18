package userManagement

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	fmt.Println("service1:Init")
	//dbReset() // if any modifcation to DB structure
	if _, err := os.Stat(DBFileWithPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not found....dbReset")
		dbReset()
	} else {
		fmt.Println("file found....dbINIT")
		dbInit()
	}
}

func GetUsers(c *gin.Context) Response {
	fmt.Println("Service:Getuser")
	al, err := GetAllUsers()
	if err != nil {
		fmt.Println("ServiceClass_GetUsers1_Error: " + err.Error())
		return Resp(nil, Code_301, err.Error())
	} else {
		return Resp(al, Code_200, SuccessMsg)
	}
}
func GetUserByID(c *gin.Context) Response {
	fmt.Println("Service:GetUserByID")
	if c.Param("id") == "1" {
		fmt.Println("ServiceClass_GetUserByID1_Error: user with ID-1 is kept internal")
		return Resp(nil, Code_301, InvalidRequest) // preventing internal user
	}
	al, err := GetUsrByID(c.Param("id"))
	if err != nil {
		fmt.Println("ServiceClass_GetUserByID1_Error: " + err.Error())
		return Resp(nil, Code_303, err.Error())
	} else {
		return Resp(al, Code_200, SuccessMsg)
	}
}

func CreateUser(c *gin.Context) Response {
	fmt.Println("Service:Createuser")
	var newUser User
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println("ServiceClass_CreateUser1_Error: " + err.Error())
		return Resp(nil, Code_301, err.Error())
	}
	us, err := CreateUsr(newUser)
	if err != nil {
		fmt.Println("ServiceClass_CreateUser2_Error: " + err.Error())
		return Resp(nil, Code_302, err.Error())
	} else {
		return Resp(us, Code_200, SuccessMsg)
	}

}

func UpdateUser(c *gin.Context) Response {
	fmt.Println("Service:UpdateUser")
	var uu User
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&uu); err != nil {
		fmt.Println("ServiceClass_UpdateUser_Error: " + err.Error())
		return Resp(nil, Code_301, err.Error())
	}
	fmt.Println("id:" + c.Param("id"))
	fmt.Println("user" + uu.Name)
	eu, er := GetUsrByID(c.Param("id"))
	if er != nil {
		fmt.Println("ServiceClass_UpdateUser1_Error: " + er.Error())
		return Resp(nil, Code_301, UserNotFound)
	} else if eu.ID != uu.ID {
		fmt.Println("ServiceClass_UpdateUser2_Error: Param Id dont match with Update deatils")
		return Resp(nil, Code_304, IDMismatch)
	} else if eu.UserName != uu.UserName {
		fmt.Println("ServiceClass_UpdateUser3_Error: Username Can not be updated")
		return Resp(nil, Code_305, UserNameUpdate_305)
	}
	us, err := UpdateUsr(c.Param("id"), uu)
	if err != nil {
		fmt.Println("ServiceClass_UpdateUser4_Error: " + err.Error())
		return Resp(nil, Code_301, UserNotFound)
	} else {
		return Resp(us, Code_200, SuccessMsg)
	}
}

func AuthenticateUser(c *gin.Context) Response {
	fmt.Println("Service:AuthenticateUser")
	var usr User
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&usr); err != nil {
		fmt.Println("ServiceClass_AuthenticateUser_Error: " + err.Error())
		return Resp(nil, Code_301, err.Error())
	}
	us, err := AuthenticateUsr(usr)
	if err != nil {
		fmt.Println("ServiceClass_AuthenticateUser2_Error: " + err.Error())
		return Resp(nil, Code_301, UserNotFound)
	} else {
		return Resp(us, Code_200, SuccessMsg)
	}
}
