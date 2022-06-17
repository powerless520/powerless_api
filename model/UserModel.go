package model

import "time"

type SysUser struct {
	Id                int       `json:"id"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	Password          string    `json:"password"`
	Salt              string    `json:"salt"`
	Email             string    `json:"email"`
	Mobile            string    `json:"mobile"`
	Status            int       `json:"status"`
	DeptId            string    `json:"dept_id"`
	LastLoginTime     time.Time `json:"last_login_time"`
	LastLoginIp       string    `json:"last_login_ip"`
	DelFlag           int       `json:"del_flag"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PlatformAuthority string    `json:"platform_authority"`
}
