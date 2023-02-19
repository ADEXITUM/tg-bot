package consts

var (
	API_KEY = "793379e5-4050-4e06-b44a-a5f67878024d"
	CS_URL  = "https://api.writesonic.com/v2/business/content/chatsonic?engine=premium&language=en"
)

type Response struct {
	Message string `json:"message"`
}
