package toy1

// beego session manager

var (
	sessionOn            bool   // 是否开启 session 模块
	sessionProvider      string // -
	sessionName          string // 客户端保存的 cookies 的名称
	sessionGCMaxLifeTime int64  // cookies有效期
)
