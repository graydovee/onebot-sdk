package model

type AccountResult struct {
	Data    *Account `json:"data"`
	RetCode int      `json:"retcode"`
	Status  string   `json:"status"`
}

type FriendListResult struct {
	Data    []*Account `json:"data"`
	RetCode int        `json:"retcode"`
	Status  string     `json:"status"`
}

type Anonymous struct {
	//匿名用户 ID
	ID int64 `json:"id"`
	//匿名用户 ID
	Name int64 `json:"name"`
	//匿名用户 flag，在调用禁言 API 时需要传入
	Flag int64 `json:"flag"`
}

type Account struct {
	UserID int `json:"user_id"`
	//性别，male 或 female 或 unknown
	Sex string `json:"sex"`
	//昵称
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
	Remark   string `json:"remark"`
}

type Cokies struct {
	Cookies string `json:"cookies"`
}

type CokiesResult struct {
	Data *Cokies `json:"data"`
}

type CSRFToken struct {
	Token string `json:"token"`
}

type CSRFTokenResult struct {
	Data *CSRFToken `json:"data"`
}

type Credentials struct {
	Cookies   string `json:"cookies"`
	CSRFToken int32  `json:"csrf_token"`
}

type CredentialsResult struct {
	Data    *Credentials `json:"data"`
	RetCode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type File struct {
	File string `json:"file"`
}

type FileResult struct {
	Data    *File  `json:"data"`
	RetCode int    `json:"retcode"`
	Status  string `json:"status"`
}
