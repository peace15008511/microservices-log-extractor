package adminUsers

var response string

type Profile struct {
	Username        string
	Password        string
	Rights          []string
	AccountIsActive bool
}

func CreateUserProfile() string {
	response = "Please connect me to mongo now so I can do some magics"
	return response

}
