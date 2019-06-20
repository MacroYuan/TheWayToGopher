package cg

import "GameServer/src/ipc"

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {

}

func (client *CenterClient) RemovePlayer(name string) error {

}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {

}

func (client *CenterClient) Broadcast(message string) error {

}
