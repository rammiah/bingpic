package lark

const (
	successCode = 0
)

/*
错误码列表：
https://open.feishu.cn/document/ukTMukTMukTM/ugjM14COyUjL4ITN
*/
func MapError(code int) string {
	if code >= 18000 && code <= 20000 {
		return "内部错误"
	}
	switch code {
	case 0:
		return "接口调用成功"
	case 10002:
		return "发送消息失败"
	case 10003:
		return "请求参数不合法"
	case 10004:
		return "获取用户信息失败或者用户 ID 不存在"
	case 10005:
		return "生成 token 的 app_id 和相关 chat、open_id 的 app_id 不一致"
	case 10009:
		return "获取 open_chat_id 失败"
	case 10010:
		return "禁止发送消息，请检查 scope 权限，机器人可见性范围"
	case 10012:
		return "获取 app access token 失败"
	case 10013, 10014:
		return "获取 tenant access token 失败，可能是appid、app secret不正确等，具体信息请查看msg字段"
	case 10015:
		return "app_secret 不正确"
	case 10016:
		return "发送 app_ticket 失败"
	case 10019:
		return "加急类型不支持"
	case 10020:
		return "消息 ID 不正确"
	case 10023:
		return "没有加急 scope 权限"
	case 10029:
		return "open_chat_id 不合法"
	case 10030:
		return "机器人不在群里"
	case 10032:
		return "所有 open_id 都不合法"
	case 10034:
		return "不支持跨企业操作"
	case 10037:
		return "获取 message_id 失败"
	case 11000:
		return "获取 sso_access_token 失败"
	case 11001:
		return "获取 CheckSecurityToken 失败"
	case 11100:
		return "open_chat_id不合法或者chat不存在"
	case 11101:
		return "open_id 不存在"
	case 11102:
		return "查询用户open_id 失败"
	case 11103:
		return "open_department_id 不存在"
	case 11104:
		return "查询用户 open_department_id 失败"
	case 11105:
		return "employee_id 不存在"
	case 11106:
		return "查询用户 employee_id 失败"
	case 11200:
		return "更新群名称失败"
	case 11201, 11208:
		return "机器人不是群主"
	case 11202:
		return "只有群主才能拉用户进群"
	case 11203:
		return "机器人没有给用户批量发送权限"
	case 11204:
		return "机器人没有给部门批量发送权限"
	case 11205:
		return "应用没有机器人"
	case 11206:
		return "用户不在群中不能被设置为群主"
	case 11207:
		return "app不可用"
	case 11209:
		return "app不存在"
	case 11210:
		return "AppUsageInfo不存在"
	case 11211:
		return "拉人进群参数错误"
	case 11212:
		return "踢人出群参数错误"
	case 11213:
		return "更新群参数错误"
	case 11214:
		return "上传图片参数错误"
	case 11215:
		return "chat_id为空"
	case 11216:
		return "获取chat_id失败"
	case 11217:
		return "拉机器人进群失败"
	case 11218:
		return "群机器人已满"
	case 11219:
		return "不支持chat跨租户"
	case 11220:
		return "禁止机器人解散群"
	case 11221:
		return "机器人不能获取不属于自己的图片"
	case 11222:
		return "机器人的Owner不在群里"
	case 11223:
		return "没有打开应用发消息权限"
	case 11224:
		return "message id参数错误"
	case 11225:
		return "你的应用对用户不可见"
	case 11226:
		return "app id参数不对或者没传"
	case 11227:
		return "image_key不存在"
	case 11228:
		return "chat_id不合法"
	case 11229:
		return "无user scope权限"
	case 11230:
		return "无发送系统消息权限"
	case 11231:
		return "chat_id不存在"
	case 11232, 11233:
		return "发消息触发限流"
	case 99991400:
		return "发消息频率超过频控限制，目前每个AppID每个接口50/s、1000/min的限制，超过这个频率的接口会返回99991400这个错误码。下一秒/下一分钟会自动恢复。"
	}
	return "未知错误"
}
