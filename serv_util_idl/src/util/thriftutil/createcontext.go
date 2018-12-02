package thriftutil


func CreateContextByUid(uid int64) *Context {
	return &Context{
		Head:&Head{
			Uid:uid,
		},
	}
}

func CreateContext(uid int64, source, dt int32, unionid, ip,region string) *Context {
	return &Context{
		Head:&Head{
			Uid:uid,
			Ip:ip,
			Region:region,
			Source:source,
			Dt:dt,
			Unionid:unionid,
		},
	}
}
