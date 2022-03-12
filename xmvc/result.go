package xmvc

type Result[T any] struct {
	RspCd  string `json:"rspCd"`
	RspInf string `json:"rspInf"`
	Error  string `json:"error"`
	Data   T      `json:"data"`
}

func OkResult[T any](data T) *Result[T] {
	return &Result[T]{Data: data, RspInf: "Success", RspCd: "00000"}
}

func ErrResult[T any](data T, rspCd string, rspInf string) *Result[T] {
	return &Result[T]{Data: data, RspCd: rspCd, RspInf: rspInf}
}
