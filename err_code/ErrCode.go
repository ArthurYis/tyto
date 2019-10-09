/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-20 14:40:34
 * @LastEditTime: 2019-09-21 09:49:47
 * @LastEditors: Arthur
 */
package err_code

const (
	SUCCESS    = 200 //成功
	FAILURE    = 2   //失败
	SERVER_ERR = 500 //服务器错误

	//客户端错误
	BAD_REQUEST                  = 4000 // 数据格式错误
	LOGIN_USER_UN_EXTIST         = 4001 // 用户不存在
	LOGIN_PASSWORD_ACCOUNT_ERROR = 4002 // 密码或者账号错误
	LOGIN_NO_TOKEN               = 4003 // 未登录,请重新登陆
	LOGIN_NO_TOKEN_TIMEOUT       = 4005 //令牌已失效,请重新登陆
	DATA_EXIST                   = 4006 //数据已存在,重复插入
	JSON_MARSHAL_ERROR           = 4007 //JSON编码错误
	JSON_UNMARSHAL_ERROR         = 4008 //JSON解码错误
	HTTP_VERSION_ERROR           = 4009 //HTTP请求版本错误
	HTTP_POST_COMMAND_ERROR      = 4010 //HTTP请求命令错误
	USER_STATE_DISABLE           = 4011 //此用户不可用

	PAR_PARAMETER_NOT_LEGAL   = 4012 //参数不合法
	DATA_CONVERSION_EXCEPTION = 4013 //数据转换异常
	UNAUTHORIZED_OPERATION    = 4014 //无操作权限
	PAR_PARAMETER_IS_NULL     = 4015 //参数为空
	LOGIN_USER_NO_PERMISSION  = 4016 //用户没有登录权限
	HTTP_READ_ERROR           = 4017 //读取http body数据失败

	//服务端错误
	SERVICE_BUSSY     = 5000 //服务器繁忙
	TIME_OUT          = 5001 //服务器查询超时
	DATA_NULL         = 5002 //数据不存在
	MGROBJ_FIND_ERROR = 5010 //获取设备信息失败
	GRPC_CONN_ERR     = 5011 //grpc连接失败
	GRPC_METHOD_ERR   = 5012 //GRPC方法调用失败

	UNKOWN_ERROR = -1 // "未知错误

	//
	DATA_ERR = 502 //数据不存在
)
