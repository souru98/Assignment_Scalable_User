package userManagement

type User struct {
	ID         int64
	Name       string
	UserName   string
	Password   string
	IsActive   bool
	IsInternal bool
}

type UserResponse struct {
	ID       int64
	Name     string
	UserName string
	IsActive bool
}

type Response struct {
	Data         any    `json:"data"`
	ResponseCode string `json:"code"`
	Message      string `json:"message"`
}

func Resp(data any, code string, msg string) Response {
	var rp Response
	rp.Data = data
	rp.ResponseCode = code
	rp.Message = msg
	return rp
}
