package models

import "time"

// BaseData 用户基础数据表
type BaseData struct {
	ID         int64  `xorm:"INT(11) notnull pk autoincr 'id'" json:"-"`
	OpenID     string `xorm:"CHAR(28) null unique 'openid'" json:"-"`  // wechat openid
	SessionKey string `xorm:"varchar(24) null 'session_key'" json:"-"` // wechat session key
	// register
	ShareType int `xorm:"INT(11) default(0) 'share_type'" json:"shareType"` // 注册分享类型
	Scene     int `xorm:"INT(11) default(0) 'scene'" json:"scene"`          // 注册场景id
	Cover     int `xorm:"INT(11) default(0) 'cover'" json:"cover"`          // 注册分享图片id
	// base
	Gate    uint32 `xorm:"INT(11) default(0) 'gate'" json:"gate"`
	Gold    uint32 `xorm:"INT(11) default(0) 'gold'" json:"gold"`
	Diamond uint32 `xorm:"INT(11) default(0) 'diamond'" json:"diamond"`
	// time
	CreatedAt time.Time `xorm:"created" json:"-"` // 创建时间
	UpdatedAt time.Time `xorm:"updated" json:"-"` // 更新时间
}

//func (u *BaseData) TableName() string {
//	return "ic_base_data"
//}
