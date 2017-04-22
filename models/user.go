package models

import (
	"time"
	"secsys/db"
)

// CompanyBasic type represents basic company info
type CompanyBasic struct {
  ComName string `json:"comName" db:"com_name"`       
  ComField string `json:"comField" db:"com_field"`        
  ComMan string `json:"comMan" db:"com_man"`         
  ComPhone string `json:"comPhone" db:"com_phone"`        
  ComRegnum string `json:"comRegnum" db:"com_regnum"`      
  ComRegcap string `json:"comRegcap" db:"com_regcap"`       
  ComCapreport string `json:"comCapreport" db:"com_capreport"`    
  ComBatch string `json:"comBatch" db:"com_batch"`        
  ComLicense string `json:"comLicense" db:"com_license"`    
}

// User type represents the registered user.
type User struct {
  // user acount info
  ID string `json:"id"`
  IsAdmin bool `json:"isAdmin" db:"is_admin"`
  Phone string `json:"phone"`
  Email *string `json:"email"`
  NickName *string `json:"nickName" db:"nick_name"`    //make it pointer so that it can be null
  Avator *string `json:"avator"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
  Password string `json:"password,omitempty"`
  // company basic info
  CompanyBasic
  // TODO add this field
  /*
  com_level text,
  appli_date text,
  appli_level text,
  appli_result text,
  certf_date text,
  certf_num text,
  verif_date text,
  verif_result text,
  com_turnover decimal,
  com_area decimal,
  police_num integer,
  police_duty integer,
  police_cancel integer,
  police_dutycancel decimal,
  list_duty integer,
  list_dutycancel decimal,
  emp_num integer,
  emp_contract integer,
  emp_lccr decimal,
  cont_num integer,
  cont_vac decimal,
  cont_samptnum integer,
  cont_sampfnum integer,
  cont_sampvac decimal,
  emp_sep integer,
  emp_seprate decimal,
  list_certrate decimal,
  list_sampcertrate decimal,
  emp_ssemanum integer,
  emp_ssemarate decimal,
  emp_semanum integer,
  emp_semarate decimal,
  emp_jsenum integer,
  emp_jserate decimal,
  train_period integer,
  com_salary decimal,
  train_funds decimal,
  train_fundsrate decimal,
  com_comins integer,
  com_sosec integer,
  com_sosecrate decimal,
  com_party text,
  com_youth text,
  com_union text,
  com_crime text,
  com_acc text,
  com_mwgs text,
  */
}

// CreateUser create user and return id
func CreateUser(user User, hashedPassword string, isAdmin bool) (string, error) {
  // TODO more field need to be insert
  var id string
  createUserSQL := `INSERT INTO users (phone, password, is_admin) 
    VALUES($1, $2, $3) 
    RETURNING id;`
  err := db.Pool.Get(&id, createUserSQL, user.Phone, hashedPassword, isAdmin)
  return id, err
}

// GetUserByPhone get user by phone
func GetUserByPhone(phone string) (User, error) {
  var user User
  getUserSQL := `SELECT * FROM users WHERE phone=$1`
  err := db.Pool.Get(&user, getUserSQL, phone)
  return user, err
}

// GetUserByID get user by id
func GetUserByID(userID string) (User, error) {
  // TODO
}

// GetUsers return companybasic info list
func GetUsers() ([]CompanyBasic, error) {
  // TODO
}

// UpdateUser update user with whole user entity
func UpdateUser(user User) (error) {

}

// DeleteUserByID delete user by id
func DeleteUserByID(userID string) (error) {

}

