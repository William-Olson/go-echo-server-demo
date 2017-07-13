package utils

// Response : holds a payload for json rest responses
type Response struct {
	Payload map[string](interface{})
}

// NewResponse : create a new response
func NewResponse(s string, v interface{}) Response {

	resp := Response{(map[string](interface{}){})}
	resp.Payload[s] = v
	return resp

}

// Set : add a new entry to the reponses payload
func (r *Response) Set(s string, v interface{}) {

	r.Payload[s] = v

}
