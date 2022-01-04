package response

var zhCNText = map[int]string{
	ServerError:        "内部服务器错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数解析失败",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",

	AuthorizedCreateError:    "创建调用方失败",
	AuthorizedListError:      "获取调用方列表失败",
	AuthorizedDeleteError:    "删除调用方失败",
	AuthorizedUpdateError:    "更新调用方失败",
	AuthorizedDetailError:    "获取调用方详情失败",
	AuthorizedCreateAPIError: "创建调用方 API 地址失败",
	AuthorizedListAPIError:   "获取调用方 API 地址列表失败",
	AuthorizedDeleteAPIError: "删除调用方 API 地址失败",
}
