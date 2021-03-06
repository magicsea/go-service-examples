package internal

/*
替换字符串
	${SERVICE_NAME}:
	${REG_MSG2ID_FUNS}
	${REG_MSG_CREATORS}
末尾增加函数
*/

const kTemplateRegMsg = `
package internal

// Code generated by gen. Do not edit!
// 代码由 gen 生成。不要手工编辑！

import (
	${S2C_MSG_IMPORTS}
	"zeus/net/server"
)

// 注册消息.
func RegMsg_${SERVICE_NAME}(svr *server.Server) {
	regMsg2ID_${SERVICE_NAME}(svr)
	regMsgCreators_${SERVICE_NAME}(svr)
}

// 注册消息ID, 用于发送消息.
func regMsg2ID_${SERVICE_NAME}(svr *server.Server) {
	// [ServerToClient] 注册发送的消息。需要从类型获取ID。
	${REG_MSG2ID_FUNS}
}

// 注册消息创建器，用于接收数据后创建消息。
func regMsgCreators_${SERVICE_NAME}(svr *server.Server) {
	// [ClientToServer] 注册接收的消息。需要从ID创建消息。
	${REG_MSG_CREATORS}
}
`
