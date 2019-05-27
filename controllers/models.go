package controllers

type FriendsList struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type Friends struct {
	Friends []string `json:"friends"`
}

type Friend struct {
	Email string `json:"email"`
}

type Subscribe struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
