package user

import (
	"entity-call-entity/src/services/servicetype"
	"entity-call-entity/src/services/team/teamdata"

	"github.com/giant-tech/go-service/framework/idata"

	"github.com/giant-tech/go-service/framework/iserver"

	log "github.com/cihub/seelog"
	"github.com/globalsign/mgo/bson"
	dbservice "github.com/giant-tech/go-service/base/mongodbservice"

	"entity-call-entity/src/entitydef"
)

// FriendsInfo 好友信息 临时数据结构
type FriendsInfo struct {
	MyFriendsDbid    []uint64 `bson:"MyFriendsDbid"`
	ApplyFriendsDbid []uint64 `bson:"ApplyFriendsDbid"`
}

// RPCHello hello
func (gu *GatewayUser) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)

	err := gu.AsyncCall(servicetype.ServiceTypeClient, "Hello", name, id)
	if err != nil {
		log.Error("err: = ", err)
	}
}

// RPCModifyAttr 修改rpc attr
func (gu *GatewayUser) RPCModifyAttr(name string, index uint32, val uint32) {
	log.Debug("RPCModifyAttr, name: ", name, ", index: ", index, " val: ", val)

	gu.SetLevel(val)

	selectProps := bson.M{}
	selectProps["Friends"] = 1

	ret := bson.M{}
	dbservice.MongoDBQueryOneWithSelect("game", "player", bson.M{"dbid": 1}, selectProps, ret)

	friends := gu.GetFriends()

	//friends.MyFriendsName = "yekoufeng"
	gu.SetFriends(friends)

	//modify heros

	var info entitydef.HEROINFO
	info.HeroName = "yekoufeng"
	info.HeroID = 2

	heros := gu.GetHero()
	//(*heros)["yekoufeng"] = info
	//gu.SetHero(heros)

	log.Debug("RPCModifyAttr, name: ", name, ", index: ", index, " val: ", val, " ,selectProps: ", selectProps, " ,ret: ", ret, " ,friends: ", friends, " ,heros: ", heros)

}

// RPCCreateTeam 创建队伍
func (gu *GatewayUser) RPCCreateTeam(name string) {
	log.Debug("RPCCreateTeam, name: ", name)

	//通过自定义函数删选所需的服务
	proxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeTeam)
	if proxy.IsValid() {
		teamData := &teamdata.CreateTeamData{}
		teamData.TeamName = "test"
		teamData.PlayerInfo = &teamdata.TeamPlayerInfo{PlayerID: gu.GetEntityID(), PlayerName: "Jim"}
		proxy.AsyncCall("CreateTeam", teamData)
	}
}

// RPCCreateTeamResult 创建结果队伍
func (gu *GatewayUser) RPCCreateTeamResult(teamID uint64) {
	log.Debug("RPCCreateTeamResult, teamID: ", teamID)

	gu.AsyncCall(idata.ServiceClient, "CreateTeamResult", teamID)
}
