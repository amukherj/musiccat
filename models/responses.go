package models

type APISuccess struct {
	BaseModel
	Message string
	ID      int64
	Time    int64
}

type APIFailure struct {
	BaseModel
	Message   string
	ErrorCode int64
	HTTPCode  int
	Time      int64
}
