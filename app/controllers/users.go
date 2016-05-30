package controllers

import (
	"github.com/revel/revel"
	"chat-go/app/chatroom"
//	"github.com/boltdb/bolt"
//	"fmt"
//	"log"
//	"encoding/json"

//    "crypto/rand"
//     "encoding/base64"
//	"time"
//	"fmt"
//	"unicode/utf8"
	"encoding/binary"
//	"unicode"

)
type Users struct {
	*revel.Controller
}
 

func itob(v uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return
}

func btoi(v []byte) (i uint64) {
	if len(v) == 8 {
		i = binary.BigEndian.Uint64(v)
	} else {
		i = 0
	}
	return
}


func ReturnAppropriateResultOnline(IsModerator bool, lol map[string]*chatroom.Session) interface{} {
	if IsModerator {
		var retValue = make([]*chatroom.OnlineEarlData, len(lol))
		i := 0
		for k, e := range lol {
			if /*e.LastSeenTimestamp + 3*60 > time.Now().Unix() ||*/ e.IsHere {
				retValue[i] = &chatroom.OnlineEarlData{e.Ip, e.Id, e.Nickname, e.IndexNickname, e.IsModerator, e.IsLoggedIn, e.IsBanned, e.LastMessageTimestamp, e.LastSeenTimestamp, k, e.IsHere}
				i += 1
			}
		}
		return retValue
	} else {
		var retValue = make([]*chatroom.OnlinePeonData, len(lol))

		i := 0
		for _, e := range lol {
			if /*e.LastSeenTimestamp + 3*60 > time.Now().Unix() || */e.IsHere{
				nick := e.Nickname
				if !e.IsLoggedIn {
					nick = ""
				}
				retValue[i] = &chatroom.OnlinePeonData{nick, e.IsModerator}
				i += 1
			}
		}
		return retValue
	}
}


func (c Users) Login(name string, password string) revel.Result {

	status := chatroom.Send_request_login(c.Session.Id(), name, password)
	return c.RenderJson(status)

};


func (c Users) Register(name string, password string) revel.Result {
	
	status := chatroom.Send_request_register(c.Session.Id(), name, password)
	return c.RenderJson(status)
}



func (c Users) State() revel.Result {


	status := chatroom.Send_request_state(c.Session.Id())
	return c.RenderJson(status)
}



func (c Users) Online() revel.Result {


	allshit := chatroom.Send_request_online(c.Session.Id())
	return c.RenderJson(ReturnAppropriateResultOnline(allshit.IsModerator, allshit.AllSessions))
}



func (c Users) Logout() revel.Result {
	status := chatroom.Send_request_logout(c.Session.Id())
	return c.RenderJson(status)
}

