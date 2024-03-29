package boldchat

// Response is the base Response object returned from BoldChat
// All API Responses should contain the following
// Status is the status returned from the API Call
// Message contains the message returned
//		Only populated if status is an error
type Response struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

// Good checks if the response is success
// returns a bool type
func (r Response) Good() bool {
	if r.Status == "success" {
		return true
	}
	return false
}
