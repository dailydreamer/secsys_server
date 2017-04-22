package models

// Contract type represents contract
type Contract struct {
  ID string `json:"id"`
	// TODO
}

// CreateContract create contract and return id
func CreateContract(contract Contract) (string, error) {
  // TODO
}


// GetContractByID get contract by id
func GetContractByID(contractID string) (Contract, error) {
  // TODO
}

// GetContracts return Contract list
func GetContracts() ([]Contract, error) {
  // TODO
}

// UpdateContract update contract with whole contract entity
func UpdateContract(contract Contract) (error) {

}

// DeleteContractByID delete contract by id
func DeleteContractByID(contractID string) (error) {

}
