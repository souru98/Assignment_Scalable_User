package userManagement

//const DBFileName = "/ProgramData/LMS/UserDB/userDB.db"
const DBFileWithPath = "/ProgramData/LMS/UserDB/userDB.db" //"./" + DBFileName
const DBFilePath = "/ProgramData/LMS/UserDB"
const URI = "localhost:8081"
const DBDriver = "sqlite3"

// Response Codes
const Code_200 = "200"
const Code_301 = "301" //User not found
const Code_302 = "302" //record already exists
const Code_303 = "303" //row not exists
const Code_304 = "304" //Paramter ID and Request Body ID Mismatch
const Code_305 = "305" //Username Can not be updated

// Response constants
const SuccessMsg = "Success"
const UserNotFound = "User not found"
const InvalidRequest = "Invalid Request"
const IDMismatch = "Paramter ID and Request Body ID Mismatch"
const UserNameUpdate_305 = "Username Can not be updated"

// default users
var goSuperuser = User{
	Name:       "SuperUser",
	UserName:   "su",
	Password:   "su",
	IsActive:   true,
	IsInternal: true,
}

var goSystemUser = User{
	Name:       "SystemUser",
	UserName:   "sys",
	Password:   "sys",
	IsActive:   true,
	IsInternal: false,
}
