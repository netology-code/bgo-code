package conversations

type ResponseDTO struct {
	Error *Error `json:"error"`
	Response *Response `json:"response"`
}

type Error struct {
	Code int `json:"error_code"`
	Message string `json:"error_msg"`
}

type Response struct {
	Count int `json:"count"`
	Items []Items `json:"items"`
}

type Items struct {
	Conversation Conversation `json:"conversation"`
}

type Conversation struct {
	Peer Peer `json:"peer"`
}

type Peer struct {
	Id int64 `json:"id"`
}

