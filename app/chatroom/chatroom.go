package chatroom

import (
	"github.com/revel/revel"
//	"github.com/cevaris/ordered_map"
	"github.com/boltdb/bolt"

	"container/list"
	"container/heap"
	"time"
	"runtime"
	"encoding/binary"
	"fmt"
	"log"
	"encoding/json"
//    "crypto/rand"
//    "encoding/base64"
	"crypto/sha1"
	"bytes"
	"unicode"
	"unicode/utf8"
	"net"
	"strings"
	"github.com/mtibben/confusables"
	"golang.org/x/text/unicode/norm" 
    "golang.org/x/crypto/bcrypt"
    "golang.org/x/text/secure/precis"
)

type Event struct {
	Index uint64
	Id uint64
	Ip []byte
	Nickname string
	IndexNickname string
	IsModerator bool
	IsLoggedIn bool
	Message string
	Timestamp int64
}

type PeonEvent struct {
	X uint64 // Index
	N string // Nickname
	A bool //  IsModerator
	L bool // IsLogggedIn
	M string // Message
	T int64 // Timestamp
}

type EarlEvent struct {
	X uint64 // Index
	D uint64 // Id
	I []byte // Ip
	N string // Nickname
	O string // IndexNickname
	A bool //  IsModerator
	L bool // IsLogggedIn
	M string // Message
	T int64 // Timestamp
}

type Session struct {
	Ip []byte
	Id uint64
	Nickname string
	IndexNickname string
	IsModerator bool
	IsLoggedIn bool
	IsBanned bool
	LastMessageTimestamp int64
	LastSeenTimestamp int64
	IsHere bool
}

type IpData struct {
	Ip    []byte
	Id uint64
	Index int
	Nickname string
	IndexNickname string
	IsModerator bool
	IsLoggedIn bool
	IsBanned bool
	LastMessageTimestamp int64
	LastSeenTimestamp int64
}

type OnlinePeonData struct {
	N string // Nickname
	A bool // IsModerator
}

type OnlineEarlData struct {
	I []byte // Ip
	D uint64 // Id
	N string // Nickname
	O string // IndexNickname
	A bool // IsModerator
	L bool // IsLoggedIn
	B bool // IsBanned
	T int64 // LastMessageTimestamp
	S int64 // LastSeenTimestamp
	K string // c.Session.Id()
	H bool // IsHere
}


type User struct {
	Ip []byte
	Id uint64
	RegistredTimestamp int64
	Nickname string
	IndexNickname string
	PasswordHash string
	IsModerator bool
	IsBanned bool   
}

type Status struct {
	Code int
	Message string
	RSS uint64
}

type Page struct {
	PrevPage      uint64 
}

// A PriorityQueue implements heap.Interface and holds Items.
type IpHeap []*IpData

func (k IpHeap) Len() int { return len(k) }


func (pq IpHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (k IpHeap) Less(i, j int) bool {
	return k[i].LastSeenTimestamp > k[j].LastSeenTimestamp
}

func (k *IpHeap) Push(x interface{}) {
	n := len(*k)
	item := x.(*IpData)
	item.Index = n
	*k = append(*k, item)
}

func (k *IpHeap) Pop() interface{} {
	old := *k
	n := len(old)
	item := old[n-1]
	item = nil;
	*k = old[0 : n-1]
	old[n-1] = nil;
	return item
}


func (k *IpHeap) Free() {
	old := *k
	//n := len(old)
	var i int

	for i = len(old); i > 0; i-- {
		if old[i-1].LastSeenTimestamp + 15 * 60 > time.Now().Unix() {
			break;
		} else {
			
			delete(Ip_hash_map, string(old[i-1].Ip))
			old[i-1] = nil;
		}
	}

	*k = old[0 : i]
	
}

// update modifies the priority and value of an Item in the queue.
func (k *IpHeap) Update(item *IpData) {

	item.LastSeenTimestamp = time.Now().Unix()
	heap.Fix(k, item.Index)
}


func ifalright(k []byte, max uint64) bool {
	i := btoi(k)
	if i == 0 || i > max  {
		return false
	} else {
		return true
	}
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



type Subscription struct {
	Archive []Event      // All the events from the archive.
	New     <-chan Event // New events coming in.
}

// Owner of a subscription must cancel it when they stop listening to events.
func (s Subscription) Cancel() {
	unsubscribe <- s.New // Unsubscribe the channel.
	drain(s.New)         // Drain it, just in case there was a pending publish.
}

func newEvent(index uint64, id uint64, ip []byte, nickname string, indexnickname string, isModerator bool, isLoggedIn bool, msg string) Event {
	return Event{0, id, ip, nickname, indexnickname, isModerator, isLoggedIn, msg, time.Now().Unix()}
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}




type PageRequest struct {
	Page uint64      // All the events from the archive.
	Responce chan []Event // New events coming in.
}

func Send_request_for_page(page uint64) []Event {
	resp := PageRequest{page, make(chan []Event)}
	give_me_messages <- resp
	return <-resp.Responce
}

type PageCountRequest struct {
	Responce chan uint64
}

func Send_request_for_page_count() uint64 {
	resp := PageCountRequest{make(chan uint64)}
	give_me_pages <- resp
	return <-resp.Responce
}


type TrySay struct {
	Sid string
	Nickname string
	Message string
	
}

type TryRegister struct {
	Sid string
	Nickname string
	Password string
	Responce chan Status
}
type TryLogin struct {
	Sid string
	Nickname string
	Password string
	Responce chan Status
}

type TryLogout struct {
	Sid string
	Responce chan Status
}

type TryState struct {
	Sid string
	Responce chan Status
}

type FuckingCombinedFuckstery struct {
	AllSessions map[string]*Session
	IsModerator bool
}

type TryOnline struct {
	Sid string
	Responce chan FuckingCombinedFuckstery
}

func Send_request_register(sid, nickname, password string) Status {
	resp := TryRegister{sid, nickname, password, make(chan Status)}
	try_register <- resp
	return <-resp.Responce
}

func Send_request_login(sid, nickname, password string) Status {
	resp := TryLogin{sid, nickname, password, make(chan Status)}
	try_login <- resp
	return <-resp.Responce
}

func Send_request_logout(sid string) Status {
	resp := TryLogout{sid, make(chan Status)}
	try_logout <- resp
	return <-resp.Responce
}

func Send_request_state(sid string) Status {
	resp := TryState{sid, make(chan Status)}
	try_state <- resp
	return <-resp.Responce
}

func Send_request_online(sid string) FuckingCombinedFuckstery {
	resp := TryOnline{sid, make(chan FuckingCombinedFuckstery)}
	try_online <- resp
	return <-resp.Responce
}



func Say(sid string, user string, msg string)  {
	resp := TrySay{sid, user, msg}
	try_post_message <- resp
	//return <-resp.Responce
}


/*
func Say(sid string, user string, msg string) {
	publish <- newEvent(0, session.Id, session.Ip, session.Nickname, session.IsModerator, session.IsLoggedIn, msg)
}
*/
const archiveSize = 10

var (
	// Send a channel here to get room events back.  It will send the entire
	// archive initially, and then new messages as they come in.
	give_me_messages = make(chan PageRequest)
	give_me_pages = make(chan PageCountRequest)
	try_post_message = make(chan TrySay)

	try_login = make(chan TryLogin)
	try_register = make(chan TryRegister)
	try_logout = make(chan TryLogout)
	try_state = make(chan TryState)
	try_online = make(chan TryOnline)
	
	subscribe = make(chan (chan<- Subscription), 10)
	// Send a channel here to unsubscribe.
	unsubscribe = make(chan (<-chan Event), 10)
	// Send events here to publish them.
	publish = make(chan Event, 10)

	AllFuckingSessions map[string]*Session = make(map[string]*Session)
	Db bolt.DB
	Ip_heap IpHeap = make(IpHeap, 0)
	Ip_hash_map map[string]*IpData = make(map[string]*IpData)
) 

func only_alphanum(name string) (buffer string)  {
	for _, c := range name {
		if unicode.IsGraphic(c) && unicode.IsPrint(c) {
			Categories := unicode.Categories
			if unicode.In(c, Categories["Ll"],Categories["Lo"],Categories["Lt"],Categories["Lu"], Categories["Sc"], Categories["Sm"], Categories["So"], Categories["N"], Categories["Po"]) {
				//if unicode.IsLetter(c) || unicode.IsSymbol(c) || unicode.IsNumber(c) || unicode.IsDigit(c) {
		        buffer = buffer + string(c)
			}		
		}		
	}
	return
}

func strip_bidi_the_fuck_out(name string) (buffer string) {
	for _, c := range name {
		if string(c) == "\u200E" {
			continue
		}
		if string(c) == "\u200F" {
			continue
		}
		if string(c) == "\u061C" {
			continue
		}
		if string(c) == "\u202A" {
			continue
		}
		if string(c) == "\u202D" {
			continue
		}
		if string(c) == "\u202B" {
			continue
		}
		if string(c) == "\u202E" {
			continue
		}
		if string(c) == "\u202C" {
			continue
		}
		if string(c) == "\u2066" {
			continue
		}
		if string(c) == "\u2067" {
			continue
		}
		if string(c) == "\u2068" {
			continue
		}
		if string(c) == "\u2069" {
			continue
		}
        buffer = buffer + string(c)
	}
	return
}
const minEffectiveNicknameLength = 3;
const maxNicknameLength = 70;
func InitiateVisitor(c *revel.Controller) revel.Result {

	addr := strings.Split(c.Request.RemoteAddr, ":")
	ip := net.ParseIP(strings.Replace(strings.Replace(strings.Join(addr[:len(addr)-1],":" ), "[", "", -1), "]", "", -1))
	cur_time := time.Now().Unix()

	//fmt.Printf("addr %s ip %s ", addr, ip);
	if _, ok := AllFuckingSessions[c.Session.Id()]; ! ok {
		AllFuckingSessions[c.Session.Id()] = &Session{ip, 0, "", "", false, false, false, 0, cur_time, false};
	}

	if Ip_heap.Len() > 0 {
		Ip_heap.Free()
	}
	
	if _, ok := Ip_hash_map[string(ip)]; ! ok {
		Ip_hash_map[string(ip)] = &IpData{
			ip, 
			AllFuckingSessions[c.Session.Id()].Id, 
			0, 
			AllFuckingSessions[c.Session.Id()].Nickname, 
			AllFuckingSessions[c.Session.Id()].IndexNickname, 
			AllFuckingSessions[c.Session.Id()].IsModerator, 
			AllFuckingSessions[c.Session.Id()].IsLoggedIn, 
			AllFuckingSessions[c.Session.Id()].IsBanned, 
			AllFuckingSessions[c.Session.Id()].LastMessageTimestamp,
			time.Now().Unix()};
		heap.Push(&Ip_heap, Ip_hash_map[string(ip)])
	} else {
		Ip_heap.Update(Ip_hash_map[string(ip)])
		
	}

    AllFuckingSessions[c.Session.Id()].LastSeenTimestamp = cur_time
	return nil
}
func KekOnPanic(c *revel.Controller) revel.Result {
	//fmt.Println("panic was held, panic everybody, %#v", c)
	return nil
}

const cooldownTimeSec = 15
const messageMaxLength = 140
const messageMinLength = 1
const perPage = 30

// This function loops forever, handling the chat room pubsub

func prepare_nickname(name_orig string) (name, confusable_name, index_name string, err *Status) {

	// BASIS nickname: NORMALIZED+STRUPPED BIDI
	// separate index for BASIS->PRECIS+ONLY ALPHASYMNUM+LOWERCASE // also fail on normalize err
	// separate index for BASIS->PRECIS+ONLY ALPHASYMNUM+SKELETON and only then LOWERCASE

	// First we normalize the shit.	
	name = name_orig
	name = norm.NFC.String(name)
	name = strip_bidi_the_fuck_out(name)
	
	if utf8.RuneCountInString(name) > maxNicknameLength {
		err = &Status{};
		err.Code = 1
		err.Message = "Nick too long"

		return "", "", "", err
	}

	// make precis
	new_name, precis_err := precis.UsernameCasePreserved.String(name)
	if precis_err != nil {
		err = &Status{};
		err.Code = -1
		err.Message = "Your nick has unacceptable characters in it"
		fmt.Printf("Preciss precis_err on %s nickname: %#v\n", name, precis_err)

		return "", "", "", err
	}
	
	// check for idempotency
	tmp_name, precis_err2 := precis.UsernameCasePreserved.String(new_name)
	if precis_err2 != nil || tmp_name != new_name {
		err = &Status{};
		err.Code = -1
		err.Message = "This nick you got there is weird as fuck, go change it champ."

		return "", "", "", err
	}

	new_name = only_alphanum(new_name)
	if utf8.RuneCountInString(new_name) < minEffectiveNicknameLength {
		err = &Status{};
		err.Code = -1
		err.Message = "Bad nickname."

		return "", "", "", err
	}

	confusable_name = strings.ToLower(confusables.Skeleton(new_name))
	index_name = strings.ToLower(new_name)

	fmt.Printf("confusibe: %s, index: %s\n", confusable_name, index_name)
	return
}

func chatroom() {
	archive := list.New()
	subscribers := list.New()

    revel.InterceptFunc(InitiateVisitor, revel.BEFORE, &revel.Controller{})
    //revel.InterceptFunc(KekOnPanic, revel.PANIC, &revel.Controller{})
    revel.InterceptFunc(KekOnPanic, revel.AFTER, &revel.Controller{})
	
	err := Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("bucket-msgs"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("bucket-users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("bucket-users-index"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("bucket-users-confusable"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ch := <-try_register:

			session := AllFuckingSessions[ch.Sid]
			name := ch.Nickname
			password := ch.Password

			var status Status
			var confusable_name, index_name string

			if session.IsLoggedIn {
				status.Code = 1
				status.Message = "You are signed in"

				ch.Responce <- status
				break
			}

			if utf8.RuneCountInString(password) < 7 {
				status.Code = utf8.RuneCountInString(password)
				status.Message = "Password too short"

				ch.Responce <- status
				break
			}

			if len([]byte(password)) > 1024 {
				status.Code = 1
				status.Message = "Password too long"

				ch.Responce <- status
				break
			}
			
			var status_err *Status
			name, confusable_name, index_name, status_err = prepare_nickname(name)
			if status_err != nil {
				ch.Responce <- *status_err
				break
			}


			err := Db.Batch(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("bucket-users"))
				b2 := tx.Bucket([]byte("bucket-users-index"))
				b3 := tx.Bucket([]byte("bucket-users-confusable"))

					
		    	var id uint64
		    	id = btoi(b2.Get([]byte("idx")))
		    	id += 1

		    	v := b.Get([]byte(index_name))
		    	v3 := b3.Get([]byte(confusable_name))

		    	if v != nil || v3 != nil {
					status.Code = 1
					status.Message = "Already registred"
			    	return nil
		    	}

				hash,_ := bcrypt.GenerateFromPassword([]byte(password), 0);
				
				var user User
				user.Id = id
				user.Ip = session.Ip
				user.RegistredTimestamp = time.Now().Unix()
				user.Nickname = name
				user.IndexNickname = index_name
				user.PasswordHash = string(hash)
				user.IsModerator = false
				user.IsBanned = false
				
				buf, err := json.Marshal(user)
				if err != nil {
					return err
				}

				status.Code = 0
				status.Message = "OK"

				b.Put([]byte(index_name), buf)
				b2.Put([]byte("idx"), itob(id))
				b3.Put([]byte(confusable_name), itob(1)) // present
				return nil
			})

			if err != nil {
				log.Panic(err)
			}

			ch.Responce <- status

		case ch := <-try_login:

			session := AllFuckingSessions[ch.Sid]
			name := ch.Nickname
			password := ch.Password
			var status Status
			var index_name string
			var user User

			
			if session.IsLoggedIn {
				status.Code = 1
				status.Message = "Already signed im"

				ch.Responce <- status
				break
			}

			if utf8.RuneCountInString(password) < 7 {
				status.Code = utf8.RuneCountInString(password)
				status.Message = "Password too short"

				ch.Responce <- status
				break
			}

			if len([]byte(password)) > 1024 {
				status.Code = 1
				status.Message = "Password too long"

				ch.Responce <- status
				break
			}

			var status_err *Status
			name, _, index_name, status_err = prepare_nickname(name)
			if status_err != nil {
				ch.Responce <- *status_err
				break
			}


			err := Db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("bucket-users"))
		    	v := b.Get([]byte(index_name))
		    	
				json.Unmarshal(v, &user)
				if []byte(user.PasswordHash)[0] == []byte("$")[0] {	
					err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
					if err == nil {			
						status.Code = 0
						status.Message = "OK"
					} else {
						status.Code = 1
						status.Message = "Wrong name/password"
					}
				} else {
					heh := sha1.Sum([]byte(password))
					pass_hash_bytes := []byte(user.PasswordHash)

					if bytes.Equal(heh[:], pass_hash_bytes[:]){
						hash,_ := bcrypt.GenerateFromPassword([]byte(password), 0);
						user.PasswordHash = string(hash)
						
						buf, err := json.Marshal(user)
								
						if err != nil {
							return err
						}
						b.Put([]byte(index_name), buf)
						
						status.Code = 0
						status.Message = "OK"
					} else {
						status.Code = 1
						status.Message = "Wrong name/password"
					}
				}
				
				return nil
			})


			if err != nil {
				log.Panic(err)
			}

			if status.Message == "OK" {
				AllFuckingSessions[ch.Sid].Id = user.Id
				AllFuckingSessions[ch.Sid].Nickname = user.Nickname
				AllFuckingSessions[ch.Sid].IndexNickname = user.IndexNickname
				AllFuckingSessions[ch.Sid].IsModerator = user.IsModerator
				AllFuckingSessions[ch.Sid].IsLoggedIn = true
				AllFuckingSessions[ch.Sid].IsBanned = user.IsBanned
				AllFuckingSessions[ch.Sid].LastMessageTimestamp = 0
				AllFuckingSessions[ch.Sid].LastSeenTimestamp = time.Now().Unix()
				
				Ip_hash_map[string(session.Ip)].Id = AllFuckingSessions[ch.Sid].Id
				Ip_hash_map[string(session.Ip)].Nickname = AllFuckingSessions[ch.Sid].Nickname
				Ip_hash_map[string(session.Ip)].IndexNickname = AllFuckingSessions[ch.Sid].IndexNickname
				Ip_hash_map[string(session.Ip)].IsModerator = AllFuckingSessions[ch.Sid].IsModerator
				Ip_hash_map[string(session.Ip)].IsLoggedIn = AllFuckingSessions[ch.Sid].IsLoggedIn
				Ip_hash_map[string(session.Ip)].IsBanned = AllFuckingSessions[ch.Sid].IsBanned
				Ip_hash_map[string(session.Ip)].LastMessageTimestamp = AllFuckingSessions[ch.Sid].LastMessageTimestamp
				Ip_hash_map[string(session.Ip)].LastSeenTimestamp = AllFuckingSessions[ch.Sid].LastSeenTimestamp
			}
			ch.Responce <- status
		case ch := <-try_logout:
			session := AllFuckingSessions[ch.Sid]
			var status Status

			if session.IsLoggedIn {

				AllFuckingSessions[ch.Sid].IsLoggedIn = false
				AllFuckingSessions[ch.Sid].Nickname = ""
				Ip_hash_map[string(session.Ip)].IsLoggedIn = false
				Ip_hash_map[string(session.Ip)].Nickname = ""

				status.Code = 0
				status.Message = "OK"

				ch.Responce <- status
				break
			}

			status.Message = "You are not logged in"
			status.Code = 1

			ch.Responce <- status

		case ch := <-try_state:
					
			session := AllFuckingSessions[ch.Sid]
			var status Status
		    var mem runtime.MemStats

		    runtime.ReadMemStats(&mem)
		    status.RSS = mem.Sys
			if session.IsLoggedIn {
				status.Message = session.Nickname
				status.Code = 0
			} else {
				status.Message = ""
				status.Code = 0
			}

			ch.Responce <- status
		case ch := <-try_online:
			copy_sessions := make(map[string]*Session)
			for k,v := range AllFuckingSessions {
			  copy_sessions[k] = v
			}

			//copy_sessions := AllFuckingSessions//makes copy
			session := AllFuckingSessions[ch.Sid]
			ch.Responce <- FuckingCombinedFuckstery{copy_sessions, session.IsModerator}

		case ch := <-try_post_message:

			session := AllFuckingSessions[ch.Sid]
			ip := string(session.Ip)
			cur_time := time.Now().Unix()
			
			if !session.IsLoggedIn {
				AllFuckingSessions[ch.Sid].Nickname = ch.Nickname;
			}

			if Ip_hash_map[ip].LastMessageTimestamp + cooldownTimeSec < cur_time  {
				AllFuckingSessions[ch.Sid].LastMessageTimestamp = cur_time
				Ip_hash_map[ip].LastMessageTimestamp = cur_time

				publish <- newEvent(0, session.Id, session.Ip, session.Nickname, session.IndexNickname, session.IsModerator, session.IsLoggedIn, ch.Message)
			}

		case ch := <-give_me_pages:

			var prev_page uint64

			Db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("bucket-msgs"))
				c := b.Cursor()
				k, _ := c.Last()
				count := btoi(k)
	
				no_full_pages := count / perPage
				if no_full_pages > 0 {
					prev_page = no_full_pages - 1
				} else {
					prev_page = 0
				}
				return nil
			})

			ch.Responce <- prev_page

		case ch := <-give_me_messages:

			var events []Event
			page := ch.Page

			err := Db.View(func(tx *bolt.Tx) error {
				var start_from uint64
				var end_on uint64

				b := tx.Bucket([]byte("bucket-msgs"))
				c := b.Cursor()
				k, _ := c.Last()
				count := btoi(k)


				if page > 0 {
					start_from = (page - 1) * perPage + 1;
					end_on = page * perPage
				} else {
					full_pages := count / perPage	
					prev_page := full_pages - 1

					if full_pages == 0 {
						prev_page = 0
					}
					
					start_from = prev_page * perPage
					end_on = count
				}

				if count == 0 || count < end_on {
					return nil
				}

				for k, v := c.Seek(itob(start_from)); ifalright(k, end_on); k, v = c.Next() {
					var event Event
					err := json.Unmarshal(v, &event)

					if err != nil {
						return err;
					}
					events = append(events, event)
				}
				return nil
			})

			if err != nil {
				log.Panic(err)
			}

			ch.Responce <- events

		case ch := <-subscribe:
			var events []Event
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(Event))
			}
			subscriber := make(chan Event, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{events, subscriber}

		case event := <-publish:
					
			Db.Update(func(tx *bolt.Tx) error {
				var id uint64
				b := tx.Bucket([]byte("bucket-msgs"))
				id, _ = b.NextSequence()
				fmt.Println("heres the id")
				fmt.Println(id)
				event.Index = id
				buf, err := json.Marshal(event)
				if err != nil {
					return err
				}
				fmt.Println("going on")

				return b.Put(itob(id), buf)
			})

			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan Event) <- event
			}

			if archive.Len() >= archiveSize {
				archive.Remove(archive.Front())
			}
			archive.PushBack(event)

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan Event) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		}
	}
}

func init() {
	//var err error
	db, err := bolt.Open("my.db", 0600, nil)
	Db = *db
	if err != nil {
		log.Fatal(err)
	}

	heap.Init(&Ip_heap)
	
	go chatroom()
}

// Helpers

// Drains a given channel of any messages.
func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
