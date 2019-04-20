package client

type UserInfo struct {
	Id       int64
	Username string
	Password string
}
type CreateRequest struct {
	Username string
	Password string
}

type CreateResponse struct {
	Username string
}

type DeleteRequest struct {
	UserID int64
}
type DeleteResponse struct {
}

type GetUserRequest struct {
	Username string
}
type GetUserResponse struct {
	UserInfo
}

type ListRequest struct {
	Username string
	Offset   int64
	Limit    int64
}

type ListResponse struct {
	TotalCount uint64
	UserList   []*UserInfo
}

type UpdateRequest struct {
	UserInfo
}

type UpdateResponse struct {
}
