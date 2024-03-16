package model

type PoolInfo struct {
	PoolId         int32  `json:"pool_id"`         // 池id
	Cid            string `json:"cid"`             // ipfs上图片存储地址
	Name           string `json:"name"`            // 名字
	Description    string `json:"description"`     // 描述
	Price          int32  `json:"price"`           // 价格
	Amount         int32  `json:"amount"`          // 数量
	Left           int32  `json:"left"`            // 剩余数量
	LimitAmount    int32  `json:"limit_amount"`    // 限制数量
	CreatorName    string `json:"creator_name"`    // 创建者名字
	CreatorAddress string `json:"creator_address"` // 创建者地址
	Status         string `json:"status"`          // 状态（true=>新品，false=>非卖）
}

type DcInfo struct {
	Id             int32  `json:"id"`              // 唯一标识ID
	Hash           string `json:"hash"`            // 藏品hash
	Cid            string `json:"cid"`             // ipfs上图片存储地址
	Name           string `json:"name"`            // 名字
	Description    string `json:"description"`     // 描述
	Price          int32  `json:"price"`           // 价格
	OwnerName      string `json:"owner_name"`      // 拥有者名字
	OwnerAddress   string `json:"owner_address"`   // 拥有者地址
	CreatorName    string `json:"creator_name"`    // 创建者名字
	CreatorAddress string `json:"creator_address"` // 创建者地址
}

type ActivityInfo struct {
	Id            int32  `json:"id"`             // 链中返回的id
	Name          string `json:"name"`           // 活动名
	Description   string `json:"description"`    // 活动描述
	DcDescription string `json:"dc_description"` // nft描述
	Cid           string `json:"cid"`            // 图片存储地址
	HostName      string `json:"host_name"`      // 举办人名字
	HostAddress   string `json:"host_address"`   // 举办人地址
	Amount        int32  `json:"amount"`         // 总数量
	Remainder     int32  `json:"remainder"`      // 剩余数量
	Status        string `json:"status"`         // 状态(0=>展示，1=>不展示)
}
