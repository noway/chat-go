package controllers

import (
	"github.com/revel/revel"
	"chat-go/app/chatroom"
	"unicode/utf8"
	"log"
)

type LongPolling struct {
	*revel.Controller
}

const cooldownTimeSec = 15
const messageMaxLength = 140
const messageMinLength = 1
const perPage = 30

func ReturnAppropriateResult(IsModerator bool, actualEvents []chatroom.Event) interface{} {
	if IsModerator {
		var retValue = make([]*chatroom.EarlEvent, len(actualEvents))

		for i, e := range actualEvents {
			retValue[i] = &chatroom.EarlEvent{e.Index, e.Id, e.Ip, e.Nickname, e.IndexNickname, e.IsModerator, e.IsLoggedIn, e.Message, e.Timestamp}
		}
		return retValue
	} else {
		var retValue = make([]*chatroom.PeonEvent, len(actualEvents))

		for i, e := range actualEvents {
			retValue[i] = &chatroom.PeonEvent{e.Index, e.Nickname, e.IsModerator, e.IsLoggedIn, e.Message, e.Timestamp}
		}
		return retValue
	}
}


func (c LongPolling) Room(user string) revel.Result {
	
	return c.Render(user)
}

func (c LongPolling) Say(user, message string) revel.Result {

	
	if utf8.RuneCountInString(message) > messageMaxLength{
		return nil
	}
	if utf8.RuneCountInString(message) < messageMinLength{
		return nil
	}
	
	chatroom.Say(c.Session.Id(), user, message)
	return nil
}


func (c LongPolling) PrevPage() revel.Result {

	var page chatroom.Page

	page.PrevPage = chatroom.Send_request_for_page_count();

	return c.RenderJson(page)
}

func (c LongPolling) WaitMessages(last uint64) revel.Result {
	subscription := chatroom.Subscribe()

	defer func() {
		subscription.Cancel()
		chatroom.AllFuckingSessions[c.Session.Id()].IsHere = false
		log.Println("Ishere unset")
	}()
	var events []chatroom.Event
	for _, event := range subscription.Archive {	
		if last < event.Index {
			events = append(events, event)
		}
	}
			
	if len(events) > 0 {
		return c.RenderJson(ReturnAppropriateResult(chatroom.AllFuckingSessions[c.Session.Id()].IsModerator, events))
	}

	chatroom.AllFuckingSessions[c.Session.Id()].IsHere = true
		log.Println("Ishere set")

	event := <-subscription.New
	return c.RenderJson(ReturnAppropriateResult(chatroom.AllFuckingSessions[c.Session.Id()].IsModerator, []chatroom.Event{event}))
}

func (c LongPolling) LoadPage(last uint64, page uint64) revel.Result {
	var events []chatroom.Event

	if page > 0 {
		events = chatroom.Send_request_for_page(page);
		if len(events) != perPage {
			return c.RenderJson(&chatroom.Status{1, "Wrong page",0})
		}
		return c.RenderJson(ReturnAppropriateResult(chatroom.AllFuckingSessions[c.Session.Id()].IsModerator, events))
	} else {
		events = chatroom.Send_request_for_page(0);
		return c.RenderJson(ReturnAppropriateResult(chatroom.AllFuckingSessions[c.Session.Id()].IsModerator, events))
	}
}

func (c LongPolling) Leave(user string) revel.Result {
	//chatroom.Leave(user)
	return c.Redirect(Application.Index)
}