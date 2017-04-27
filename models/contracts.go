package models

import (
	"time"
	"secsys/db"
)

// Contract type represents contract
type Contract struct {
  ID string `json:"id" db:"id"`
  UserID string `json:"userID" db:"user_id"`
  ComName string `json:"comName" db:"com_name"`
  ContractNo string `json:"contractNo" db:"contract_no"`
  ProjectName string `json:"projectName" db:"project_name"`
  ComField string `json:"comField" db:"com_field"`
  CustomerName string `json:"customerName" db:"customer_name"`
  CustomerType string `json:"customerType" db:"customer_type"`
  PeopleNum int `json:"peopleNum" db:"people_num"`
  StartTime time.Time `json:"startTime" db:"start_time"`
  EndTime time.Time `json:"endTime" db:"end_time"`
  UnitPrice float64 `json:"unitPrice" db:"unit_price"`
  TotalPrice float64 `json:"totalPrice" db:"total_price"`
  Income string `json:"income" db:"income"`
  Created time.Time `db:"created"`
  Modified time.Time `db:"modified"`
}

// CreateContract create contract and return id
func CreateContract(contract Contract) (string, error) {
  var id string
  createContractSQL := `INSERT INTO contracts (
    user_id,
    com_name,
    contract_no,
    project_name,
    com_field,
    customer_name,
    customer_type,
    people_num,
    start_time,
    end_time,
    unit_price,
    total_price,
    income
  ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    RETURNING id;`
  err := db.Pool.Get(&id, createContractSQL,
    contract.UserID,
    contract.ComName,
    contract.ContractNo,
    contract.ProjectName,
    contract.ComField,
    contract.CustomerName,
    contract.CustomerType,
    contract.PeopleNum,
    contract.StartTime,
    contract.EndTime,
    contract.UnitPrice,
    contract.TotalPrice,
    contract.Income)
  return id, err
}

// GetContractByID get contract by id
func GetContractByID(contractID string) (Contract, error) {
  var contract Contract
  getContractSQL := `SELECT * FROM contracts WHERE id=$1`
  err := db.Pool.Get(&contract, getContractSQL, contractID)
  return contract, err
}

// GetContracts return Contract list
func GetContracts() ([]Contract, error) {
  contracts := []Contract{}
  getContractsSQL := `SELECT * FROM contracts`
  err := db.Pool.Select(&contracts, getContractsSQL)
  return contracts, err
}

// UpdateContract update contract with whole contract entity
func UpdateContract(contract Contract) (error) {
  updateContractSQL := `UPDATE contracts SET
    user_id=$1,
    com_name=$2,
    contract_no=$3,
    project_name=$4,
    com_field=$5,
    customer_name=$6,
    customer_type=$7,
    people_num=$8,
    start_time=$9,
    end_time=$10,
    unit_price=$11,
    total_price=$12,
    income=$13
  WHERE id=$14`
  _, err := db.Pool.Exec(updateContractSQL,
    contract.UserID,
    contract.ComName,
    contract.ContractNo,
    contract.ProjectName,
    contract.ComField,
    contract.CustomerName,
    contract.CustomerType,
    contract.PeopleNum,
    contract.StartTime,
    contract.EndTime,
    contract.UnitPrice,
    contract.TotalPrice,
    contract.Income,
    contract.ID)
  return err
}

// DeleteContractByID delete contract by id
func DeleteContractByID(contractID string) (error) {
  deleteContractSQL := `DELETE FROM contracts WHERE id=$1`
  _, err := db.Pool.Exec(deleteContractSQL, contractID)
  return err
}
