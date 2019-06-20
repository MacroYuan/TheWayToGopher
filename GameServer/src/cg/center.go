package cg

import (
	"GameServer/src/ipc"
	"encoding/json"
	"errors"
	"sync"
)

type Message struct {
	From string "from"
	To string "to"
	Connect string "connect"
}

var _ ipc.Server = &CenterServer{} //确认实现了Server的接口

type CenterServer struct {
	servers map[string] ipc.Server
	players []*Player
	rooms []*Room//目前无服务器管理
	mutex sync.RWMutex //学习使用锁
}

func NewCenterServer() *CenterServer {
	servers := make(map[string] ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers:servers, players:players}
}

func (server *CenterServer)addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player) //获取json里的error

	if err != nil {
		return err
	}

	server.mutex.Lock() //锁住写入
	defer server.mutex.Unlock()

	server.players = append(server.players, player)

	return nil
}

func (server *CenterServer)removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[:i+1]...)
			}
			return nil
		}
	}
	return errors.New("Player not found.")
}

func (server *CenterServer)listPlayer(params string) error {

}

func (server *CenterServer) broadcast(params string) error {

}

func (server *CenterServer) Handle(method, pramas string) *ipc.Response {

}