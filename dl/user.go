package dl

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
	Base
}

type DeleteRequest struct {
	UserID int64
}
type DeleteResponse struct {
	Base
}

type GetUserRequest struct {
	Username string
}
type GetUserResponse struct {
	UserInfo
	Base
}

type ListRequest struct {
	Username string
	Offset   int64
	Limit    int64
}

type ListResponse struct {
	TotalCount uint64
	UserList   []*UserInfo
	Base
}

type UpdateRequest struct {
	UserInfo
}

type UpdateResponse struct {
	Base
}
