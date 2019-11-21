package boldchat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetOperators is the 'getOperators' API Call
// returns the full list of operators found or an error
// reference: https://developer.bold360.com/help/EN/Bold360API/Bold360API/c_data_extraction_list_data_getOperators.html
func (c *Client) GetOperators() ([]*Operator, error) {
	url := c.url("getOperators")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ur := getOperatorsResponse{}
	json.Unmarshal(body, &ur)
	if ur.Good() {
		return ur.Data, nil
	}
	return nil, errors.New(ur.Message)
}

// GetOperator is the 'getOperator' API Call
// requires the OperatorID
// returns the Operator or an error
// reference: https://developer.bold360.com/help/EN/Bold360API/Bold360API/c_data_extraction_single_item_getOperator.html
func (c *Client) GetOperator(id string) (*Operator, error) {
	url := c.url("getOperator", qs("OperatorID", id))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ur := getOperatorResponse{}
	json.Unmarshal(body, &ur)
	if ur.Good() {
		return ur.Data, nil
	}
	return nil, errors.New(ur.Message)
}

// GetOperatorByEmail is a convenience function to get
// an operator by their email address
// Calls the 'getOperators' function call, and returns
// the matching user or an error
func (c *Client) GetOperatorByEmail(email string) (*Operator, error) {
	users, err := c.GetOperators()
	if err != nil {
		return nil, err
	}
	for i := range users {
		if users[i].Email == email {
			return users[i], nil
		}
	}
	return nil, fmt.Errorf("Unable to find user with email %s", email)
}

// GetOperatorByName is a convenience function to get
// an operator by their Name
// Calls the 'getOperators' function call, and returns
// the matching user or an error
func (c *Client) GetOperatorByName(name string) (*Operator, error) {
	users, err := c.GetOperators()
	if err != nil {
		return nil, err
	}
	for i := range users {
		if users[i].Name == name {
			return users[i], nil
		}
	}
	return nil, fmt.Errorf("Unable to find user with name %s", name)
}

// EditOperator is the 'editOperator' API Call
// requires an Operator object
// Any Operator value that can be changed will be sent if
// the value in the object is not empty
// The LoginID value MUST be set correctly or an error is thrown
// returns an error if applicable
// reference: https://developer.bold360.com/help/EN/Bold360API/Bold360API/c_provisioning_operator_editOperator.html
func (c *Client) EditOperator(op *Operator) error {
	s, err := op.editString()
	if err != nil {
		return err
	}
	url := c.url("editOperator", s)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Bad Edit Request on Operator %s", op.LoginID)
	}

	return nil
}

// getOperatorsResponse is the response from the 'getOperators' call
type getOperatorsResponse struct {
	Response
	Data []*Operator `json:"Data"`
}

// getOperatorResponse is the response from the 'getOperator' call
type getOperatorResponse struct {
	Response
	Data *Operator `json:"Data"`
}

// Operator is the data type matching the Operator JSON object
// according to the BoldChat API
type Operator struct {
	Name              string `json:"Name"`
	OperatorName      string `json:"OperatorName"`
	Email             string `json:"Email"`
	LoginID           string `json:"LoginID"`
	ChatName          string `json:"ChatName"`
	SSONameID         string `json:"SSONameID"`
	EmailName         string `json:"EmailName"`
	PermissionGroupID string `json:"PermissionGroupID"`
	Disabled          string `json:"Disabled"`

	Departments    []*OperatorDepartment `json:"Departments"`
	TwitterService *OperatorService      `json:"TwitterService"`
	EmailService   *OperatorService      `json:"EmailService"`
	ChatService    *OperatorService      `json:"ChatService"`
	SmsService     *OperatorService      `json:"SmsService"`
}

// editString creates the query string for the edit API Call
// returns the string or an error
func (u Operator) editString() (string, error) {
	str, err := json.Marshal(u.Departments)
	if err != nil {
		return "", nil
	}
	var out string
	if u.LoginID != "" {
		out += qs("OperatorID", u.LoginID)
	} else {
		return "", fmt.Errorf("LoginID is empty")
	}
	if u.Departments != nil {
		out += qs("Departments", string(str))
	}
	if u.OperatorName != "" {
		out += qs("OperatorName", u.OperatorName)
	}
	if u.EmailName != "" {
		out += qs("EmailName", u.EmailName)
	}
	if u.ChatName != "" {
		out += qs("ChatName", u.ChatName)
	}
	if u.Email != "" {
		out += qs("Email", u.Email)
	}
	return out, nil
}

// String is a helper function to print
func (u Operator) String() string {
	return fmt.Sprintf("\nName: %s\nOperatorName: %s\nEmail: %s\nLoginID: %s\nChatName: %s\nSSONameID: %s\nEmailName: %s\nPermissionGroupID: %s\nDisabled: %s\nDepartments: %s\nTwitterService: %s\nEmailService: %s\nChatService: %s\nSmsService: %s\n", u.Name, u.OperatorName, u.Email, u.LoginID, u.ChatName, u.SSONameID, u.EmailName, u.PermissionGroupID, u.Disabled, u.Departments, u.TwitterService, u.EmailService, u.ChatService, u.SmsService)
}

// OperatorService is the struct for the Services in the Operators JSON Object
type OperatorService struct {
	Available bool `json:"Available"`
	Capacity  int  `json:"Capacity"`
}

// String is a helper function to print
func (us OperatorService) String() string {
	return fmt.Sprintf("\n\tAvailable: %v\n\tCapacity: %d", us.Available, us.Capacity)
}

// OperatorDepartment is the struct for the Departments in the Operators JSON
// Object
type OperatorDepartment struct {
	AssignmentPriority int    `json:"AssignmentPriority"`
	Priority           int    `json:"Priority"`
	DepartmentID       string `json:"DepartmentID"`
}

// String is a helper function to print
func (ud OperatorDepartment) String() string {
	return fmt.Sprintf("\n\tAssignmentPriority: %d\n\tPriority: %d\n\tDepartmentID: %s", ud.AssignmentPriority, ud.Priority, ud.DepartmentID)
}
