package models

import(
	"time"
)

type Users struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}

//create table name

func (users *Users) tableName() string {
	return "users";
}
type UserLogin struct {
    Email    string `form:"email" binding:"required"`
    Password string `form:"password" binding:"required"`
}

//UserRegister -> Request Binding for User Register
type UserRegister struct {
    Email     string `form:"email" json:"email" binding:"required"`
    Password  string `form:"password" json:"password" binding:"required"`
    FirstName string `form:"first_name"`
    LastName  string `form:"last_name"`
}
//create response

func (users Users) ResponseMap() map[string]interface{} {
  resp := make(map[string]interface{});
  resp["id"]= users.ID
  resp["first_name"]= users.FirstName
  resp["last_name"]= users.LastName
  resp["email"]= users.Email
  resp["password"]= users.Password
  resp["is_active"]=users.IsActive
  resp["created_at"]=users.CreatedAt
  resp["updated_at"]=users.UpdatedAt
  return resp

}


