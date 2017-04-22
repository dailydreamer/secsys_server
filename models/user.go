package models

import (
	"time"
	"secsys/db"
)

// User type represents the registered user.
type User struct {
  ID string `json:"id"`
  IsAdmin bool `json:"isAdmin" db:"is_admin"`
  Phone string `json:"phone"`
  Email *string `json:"email"`
  NickName *string `json:"nickName" db:"nick_name"`    //make it pointer so that it can be null
  Avator *string `json:"avator"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
  Password string `json:"password,omitempty"`
  ComName *string `json:"comName" db:"com_name"`
  // TODO add this field
  /*
  com_field text,
  com_man text,
  com_phone text,
  com_regnum text,
  com_regcap decimal,
  com_capreport decimal,
  com_batch text,
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
  com_license text
  */
}

// CreateUser create user and return id
func CreateUser(phone string, hashedPassword string, isAdmin bool) (string, error) {
  var id string
  createUserSQL := `INSERT INTO users (phone, password, is_admin) 
    VALUES($1, $2, $3) 
    RETURNING id;`
  err := db.Pool.Get(&id, createUserSQL, phone, hashedPassword, isAdmin)
  return id, err
}

// FindUserByPhone find user by phone
func FindUserByPhone(phone string) (User, error) {
  var user User
  findUserSQL := `SELECT * FROM users WHERE phone=$1`
  err := db.Pool.Get(&user, findUserSQL, phone)
  return user, err
}

