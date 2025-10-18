package unitEngine

import "XiaXiaoMan/core"

func ReleaseResource() {
	sql, _ := core.DB.DB()
	sql.Close()
}
