package models

import (
	"time"
	"secsys/db"
)

// CompanyBasic type represents basic company info
type CompanyBasic struct {
  ID string `json:"id"`
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
  IsAdmin bool `json:"isAdmin" db:"is_admin"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  NickName string `json:"nickName" db:"nick_name"`    //make it pointer so that it can be null
  Avator string `json:"avator"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
  Password string `json:"password,omitempty"`
  // company basic info
  CompanyBasic
  // company detailed info
  ComLevel string `json:"comLevel" db:"com_level"`
  AppliDate string `json:"appliDate" db:"appli_date"`
  AppliLevel string `json:"appliLevel" db:"appli_level"`
  AppliResult string `json:"appliResult" db:"appli_result"`
  CertfDate string `json:"certfDate" db:"certf_date"`
  CertfNum string `json:"certfNum" db:"certf_num"`
  VerifDate string `json:"verifDate" db:"verif_date"`
  VerifResult string `json:"verifResult" db:"verif_result"`
  ComTurnover float64 `json:"comTurnover" db:"com_turnover"`
  ComArea float64 `json:"comArea" db:"com_area"`
  PoliceNum int `json:"policeNum" db:"police_num"`
  PoliceDuty int `json:"policeDuty" db:"police_duty"`
  PoliceCancel int `json:"policeCancel" db:"police_cancel"`
  PoliceDutycancel float64 `json:"policeDutycancel" db:"police_dutycancel"`
  ListDuty int `json:"listDuty" db:"list_duty"`
  ListDutycancel float64 `json:"listDutycancel" db:"list_dutycancel"`
  EmpNum int `json:"empNum" db:"emp_num"`
  EmpContract int `json:"empContract" db:"emp_contract"`
  EmpLccr float64 `json:"empLccr" db:"emp_lccr"`
  ContNum int `json:"contNum" db:"cont_num"`
  ContVac float64 `json:"contVac" db:"cont_vac"`
  ContSamptnum int `json:"contSamptnum" db:"cont_samptnum"`
  ContSampfnum int `json:"contSampfnum" db:"cont_sampfnum"`
  ContSampvac float64 `json:"contSampvac" db:"cont_sampvac"`
  EmpSep int `json:"empSep" db:"emp_sep"`
  EmpSeprate float64 `json:"empSeprate" db:"emp_seprate"`
  ListCertrate float64 `json:"listCertrate" db:"list_certrate"`
  ListSampcertrate float64 `json:"listSampcertrate" db:"list_sampcertrate"`
  EmpSsemanum int `json:"empSsemanum" db:"emp_ssemanum"`
  EmpSsemarate float64 `json:"empSsemarate" db:"emp_ssemarate"`
  EmpSemanum int `json:"empSemanum" db:"emp_semanum"`
  EmpSemarate float64 `json:"empSemarate" db:"emp_semarate"`
  EmpJsenum int `json:"empJsenum" db:"emp_jsenum"`
  EmpJserate float64 `json:"empJserate" db:"emp_jserate"`
  TrainPeriod int `json:"trainPeriod" db:"train_period"`
  ComSalary float64 `json:"comSalary" db:"com_salary"`
  TrainFunds float64 `json:"trainFunds" db:"train_funds"`
  TrainFundsrate float64 `json:"trainFundsrate" db:"train_fundsrate"`
  ComComins int `json:"comComins" db:"com_comins"`
  ComSosec int `json:"comSosec" db:"com_sosec"`
  ComSosecrate float64 `json:"comSosecrate" db:"com_sosecrate"`
  ComParty string `json:"comParty" db:"com_party"`
  ComYouth string `json:"comYouth" db:"com_youth"`
  ComUnion string `json:"comUnion" db:"com_union"`
  ComCrime string `json:"comCrime" db:"com_crime"`
  ComAcc string `json:"comAcc" db:"com_acc"`
  ComMwgs string `json:"comMwgs" db:"com_mwgs"`
}

/*
// CreateUser create user and return id
func CreateUser(user User, hashedPassword string, isAdmin bool) (string, error) {
  // more field need to be insert
  var id string
  createUserSQL := `INSERT INTO users (phone, password, is_admin)
    VALUES($1, $2, $3)
    RETURNING id;`
  err := db.Pool.Get(&id, createUserSQL, user.Phone, hashedPassword, isAdmin)
  return id, err
}
*/

// CreateUser create user and return id
func CreateUser(user User, hashedPassword string, isAdmin bool) (string, error) {
  var id string
  createUserSQL := `INSERT INTO users(phone, password, is_admin, email, nick_name, avator, com_name, com_field, com_man, com_phone, com_regnum, com_regcap, com_capreport, com_batch, com_license, com_level, appli_date, appli_level, appli_result, certf_date, certf_num, verif_date, verif_result, com_turnover, com_area, police_num, police_duty, police_cancel, police_dutycancel, list_duty, list_dutycancel, emp_num, emp_contract, emp_lccr, cont_num, cont_vac, cont_samptnum, cont_sampfnum, cont_sampvac, emp_sep, emp_seprate, list_certrate, list_sampcertrate, emp_ssemanum, emp_ssemarate, emp_semanum, emp_semarate, emp_jsenum, emp_jserate, train_period, com_salary, train_funds, train_fundsrate, com_comins, com_sosec, com_sosecrate, com_party, com_youth, com_union, com_crime, com_acc, com_mwgs)
  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62)
  RETURNING id;`
  err := db.Pool.Get(&id, createUserSQL, user.Phone, hashedPassword, isAdmin, user.Email, user.NickName, user.Avator, user.ComName, user.ComField, user.ComMan, user.ComPhone, user.ComRegnum, user.ComRegcap, user.ComCapreport, user.ComBatch, user.ComLicense, user.ComLevel, user.AppliDate, user.AppliLevel, user.AppliResult, user.CertfDate, user.CertfNum, user.VerifDate, user.VerifResult, user.ComTurnover, user.ComArea, user.PoliceNum, user.PoliceDuty, user.PoliceCancel, user.PoliceDutycancel, user.ListDuty, user.ListDutycancel, user.EmpNum, user.EmpContract, user.EmpLccr, user.ContNum, user.ContVac, user.ContSamptnum, user.ContSampfnum, user.ContSampvac, user.EmpSep, user.EmpSeprate, user.ListCertrate, user.ListSampcertrate, user.EmpSsemanum, user.EmpSsemarate, user.EmpSemanum, user.EmpSemarate, user.EmpJsenum, user.EmpJserate, user.TrainPeriod, user.ComSalary, user.TrainFunds, user.TrainFundsrate, user.ComComins, user.ComSosec, user.ComSosecrate, user.ComParty, user.ComYouth, user.ComUnion, user.ComCrime, user.ComAcc, user.ComMwgs)
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
  var user User
  getUserSQL := `SELECT * FROM users WHERE id=$1`
  err := db.Pool.Get(&user, getUserSQL, userID)
  return user, err
}

// GetUserByComName get user by comName
func GetUserByComName(comName string) (User, error) {
  var user User
  getUserSQL := `SELECT * FROM users WHERE com_name=$1`
  err := db.Pool.Get(&user, getUserSQL, comName)
  return user, err
}

// GetUsers return companybasic info list
func GetUsers() ([]CompanyBasic, error) {
  companyBasicList := []CompanyBasic{}
  getCompanyBasicSQL := `SELECT id, com_name, com_field, com_man, com_phone, com_regnum, com_regcap, com_capreport, com_batch, com_license FROM users`
  err := db.Pool.Select(&companyBasicList, getCompanyBasicSQL)
  return companyBasicList, err
}

// UpdateUser update user with whole user entity
func UpdateUser(user User) (error) {
  updateUserSQL := `UPDATE users SET
    is_admin=$1,
    phone=$2,
    email=$3,
    nick_name=$4,
    avator=$5,

    com_name=$6,
    com_field=$7,
    com_man=$8,
    com_phone=$9,
    com_regnum=$10,
    com_regcap=$11,
    com_capreport=$12,
    com_batch=$13,
    com_license=$14,

    com_level=$15,
    appli_date=$16,
    appli_level=$17,
    appli_result=$18,
    certf_date=$19,
    certf_num=$20,
    verif_date=$21,
    verif_result=$22,
    com_turnover=$23,
    com_area=$24,
    police_num=$25,
    police_duty=$26,
    police_cancel=$27,
    police_dutycancel=$28,
    list_duty=$29,
    list_dutycancel=$30,
    emp_num=$31,
    emp_contract=$32,
    emp_lccr=$33,
    cont_num=$34,
    cont_vac=$35,
    cont_samptnum=$36,
    cont_sampfnum=$37,
    cont_sampvac=$38,
    emp_sep=$39,
    emp_seprate=$40,
    list_certrate=$41,
    list_sampcertrate=$42,
    emp_ssemanum=$43,
    emp_ssemarate=$44,
    emp_semanum=$45,
    emp_semarate=$46,
    emp_jsenum=$47,
    emp_jserate=$48,
    train_period=$49,
    com_salary=$50,
    train_funds=$51,
    train_fundsrate=$52,
    com_comins=$53,
    com_sosec=$54,
    com_sosecrate=$55,
    com_party=$56,
    com_youth=$57,
    com_union=$58,
    com_crime=$59,
    com_acc=$60,
    com_mwgs=$61
  WHERE id=$62`
  _, err := db.Pool.Exec(updateUserSQL,
    user.IsAdmin,
    user.Phone,
    user.Email,
    user.NickName,
    user.Avator,
    user.ComName,
    user.ComField,
    user.ComMan,
    user.ComPhone,
    user.ComRegnum,
    user.ComRegcap,
    user.ComCapreport,
    user.ComBatch,
    user.ComLicense,
    user.ComLevel,
    user.AppliDate,
    user.AppliLevel,
    user.AppliResult,
    user.CertfDate,
    user.CertfNum,
    user.VerifDate,
    user.VerifResult,
    user.ComTurnover,
    user.ComArea,
    user.PoliceNum,
    user.PoliceDuty,
    user.PoliceCancel,
    user.PoliceDutycancel,
    user.ListDuty,
    user.ListDutycancel,
    user.EmpNum,
    user.EmpContract,
    user.EmpLccr,
    user.ContNum,
    user.ContVac,
    user.ContSamptnum,
    user.ContSampfnum,
    user.ContSampvac,
    user.EmpSep,
    user.EmpSeprate,
    user.ListCertrate,
    user.ListSampcertrate,
    user.EmpSsemanum,
    user.EmpSsemarate,
    user.EmpSemanum,
    user.EmpSemarate,
    user.EmpJsenum,
    user.EmpJserate,
    user.TrainPeriod,
    user.ComSalary,
    user.TrainFunds,
    user.TrainFundsrate,
    user.ComComins,
    user.ComSosec,
    user.ComSosecrate,
    user.ComParty,
    user.ComYouth,
    user.ComUnion,
    user.ComCrime,
    user.ComAcc,
    user.ComMwgs,
    user.ID)
  return err
}

// DeleteUserByID delete user by id
func DeleteUserByID(userID string) (error) {
  deleteUserSQL := `DELETE FROM users WHERE id=$1`
  _, err := db.Pool.Exec(deleteUserSQL, userID)
  return err
}

// GetUserContracts return Contract list of specific user
func GetUserContracts(userID string) ([]Contract, error) {
  contracts := []Contract{}
  getContractsSQL := `SELECT * FROM contracts WHERE user_id=$1`
  err := db.Pool.Select(&contracts, getContractsSQL, userID)
  return contracts, err
}

// GetUserScores return Score list of specific user
func GetUserScores(userID string) ([]Score, error) {
  scores := []Score{}
  getScoresSQL := `SELECT * FROM scores WHERE user_id=$1`
  err := db.Pool.Select(&scores, getScoresSQL, userID)
  return scores, err
}
