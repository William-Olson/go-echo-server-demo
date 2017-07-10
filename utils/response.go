package utils

type Response struct {
	Payload map[string](interface{})
}

/*

	Response payload helpers

*/

func NewResponse(s string, v interface{}) Response {

	resp := Response{(map[string](interface{}){})}
	resp.Payload[s] = v
	return resp

}

func (r *Response) Set(s string, v interface{}) {

	r.Payload[s] = v

}
