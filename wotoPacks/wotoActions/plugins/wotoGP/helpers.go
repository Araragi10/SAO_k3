package wotoGP

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	"github.com/Araragi10/SAO_k3/wotoPacks/appSettings"
	"github.com/Araragi10/SAO_k3/wotoPacks/interfaces"
	"github.com/Araragi10/SAO_k3/wotoPacks/wotoMD"
	ws "github.com/Araragi10/SAO_k3/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/Araragi10/SAO_k3/wotoPacks/wotoValues"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gotd/td/tg"
)

func IsAdmin(adminList []tgbot.ChatMember, id int64) bool {

	for _, ad := range adminList {
		if ad.HasLeft() || ad.WasKicked() || ad.IsAnonymous {
			continue
		}

		if ad.User.ID == id {
			return true
		}

	}

	return false
}

// IsValidUsername will check if `text` is a valid username or not.
// please fix the string before using this function,
// for example you have to do TrimPrefix, TrimSuffix, etc...
func IsValidUsername(text string) bool {
	l := len(text)
	if l <= wv.Base4Bit || l > wv.Base32Bit {
		return false
	}

	for i, c := range text {
		if i == wv.BaseIndex {
			if !wotoMD.IsEnglish(c) {
				return false
			}

			if !wotoMD.IsNumOrEng(c) {
				return false
			}
		}
	}

	return true
}

func GetMeInGp(api *tgbot.BotAPI, chat *tgbot.ChatConfig) *tgbot.ChatMember {
	return GetChatMember(api, chat, api.Self.ID)
}

func GetMember(api *tgbot.BotAPI, chat *tgbot.ChatConfig, user *tgbot.User) *tgbot.ChatMember {
	return GetChatMember(api, chat, user.ID)
}

func GetChatMember(api *tgbot.BotAPI, chat *tgbot.ChatConfig, id int64) *tgbot.ChatMember {
	mem, err := api.GetChatMember(tgbot.GetChatMemberConfig{
		ChatConfigWithUser: tgbot.ChatConfigWithUser{
			ChatID:             chat.ChatID,
			SuperGroupUsername: chat.SuperGroupUsername,
			UserID:             id,
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return &mem
}

func GetAllMembers(api *tgbot.BotAPI, chat *tgbot.ChatConfig) {

}

func GetUserInfo(username string) (u *tg.User, err error) {
	if ws.IsEmpty(&username) {
		return nil, errors.New(usernameEmptyString)
	}

	s := appSettings.GetExisting()
	if s == nil {
		return nil, errors.New(settingsNilString)
	}

	g, ctx := s.GetGClient()
	if g == nil {
		return nil, errors.New(gClientNilString)
	}

	r, err := g.ContactsResolveUsername(*ctx, username)

	if err != nil {
		return
	}

	if r == nil {
		return nil, errors.New(resolvedPeerNilString)
	}

	if len(r.Users) == wv.BaseIndex {
		return nil, errors.New(userZeroString)
	}

	ur := r.Users[wv.BaseIndex]
	u, ok := ur.(*tg.User)
	if !ok {
		return nil, errors.New(convertFailedString)
	}

	return
}

func GetTotalInfo(username string, md bool) string {
	if ws.IsEmpty(&username) {
		return wv.EMPTY
	}

	s := appSettings.GetExisting()
	if s == nil {
		return wv.EMPTY
	}

	g, ctx := s.GetGClient()
	if g == nil {
		return wv.EMPTY
	}

	retMe := func(v string) string {
		if md {
			return wotoMD.GetNormal(v).ToString()
		}
		return v
	}
	retErr := func(e error) string {
		return retMe(e.Error())
	}
	r, err := g.ContactsResolveUsername(*ctx, username)

	if err != nil {
		return retErr(err)
	}

	if r == nil {
		return retMe(resolvedPeerNilString)
	}

	if len(r.Chats) != wv.BaseIndex {
		myChat := r.Chats[wv.BaseIndex]
		ch, ok := myChat.(*tg.Channel)
		if !ok {
			ct, ok := myChat.(*tg.Chat)
			if !ok {
				return retMe(unkownChatType +
					reflect.TypeOf(myChat).String())
			}

			return parseChatInfo(ct, username)
		}

		chFull, err := g.ChannelsGetFullChannel(*ctx, ch.AsInput())
		if err != nil {
			return retErr(err)
		}

		return parseChannelInfo(ch, chFull, md)
	} else if len(r.Users) != wv.BaseIndex {
		myUser := r.Users[wv.BaseIndex]
		ur, ok := myUser.(*tg.User)

		if !ok {
			return retMe(unkownUserType +
				reflect.TypeOf(myUser).String())
		}

		return parseUserInfo(ur, md)
	} else {
		return retMe(userChZeroString)
	}
}

func parseChatInfo(chat *tg.Chat, uname string) string {
	appSettings.GetExisting().SendSudo("In the first of parse chat info.")
	yn := ws.YesOrNo
	str := groupInfoStr
	str += groupNameStr + chat.Title + wv.LineEscape
	if !ws.IsEmpty(&uname) {
		str += groupUnameStr + uname + wv.LineEscape
	}
	str += groupIDStr + strconv.Itoa(chat.ID) + wv.LineEscape
	str += membersCountStr + strconv.Itoa(chat.ParticipantsCount) + wv.LineEscape
	str += groupVersionStr + strconv.Itoa(chat.Version) + wv.LineEscape
	str += leftChatStr + yn(chat.Deactivated) + wv.LineEscape
	if chat.Deactivated {
		return str
	}

	str += activeCallStr + yn(chat.CallActive) + wv.LineEscape
	if chat.CallActive {
		str += emptyCallStr + yn(chat.CallNotEmpty) + wv.LineEscape
	}
	str += kickedChatStr + yn(chat.Kicked) + wv.LineEscape
	str += leftChatStr + yn(chat.Left) + wv.LineEscape

	return str
}

func parseChannelInfo(ch *tg.Channel, chF *tg.MessagesChatFull, md bool) string {
	isChannel := !ch.Gigagroup && !ch.Megagroup
	yn := ws.YesOrNo
	str := wv.EMPTY
	var tmp, tmp2 interfaces.WMarkDown
	genLine := func(first, second string) string {
		if !md {
			return first + second + wv.LineEscape
		}

		if ws.IsEmpty(&second) {
			return wotoMD.GetBold(first + wv.LineEscape).ToString()
		}

		tmp = wotoMD.GetBold(first)
		tmp2 = wotoMD.GetMono(second + wv.LineEscape)
		return tmp.Append(tmp2).ToString()
	}
	genLineI := func(first string, second int) string {
		return genLine(first, strconv.Itoa(second))
	}
	genLineB := func(first string, second bool) string {
		return genLine(first, yn(second))
	}
	f := chF.FullChat.(*tg.ChannelFull)
	if isChannel {
		str += genLine(channelInfoStr, str)
		str += genLine(chNameStr, ch.Title)
		str += genLine(chUnameStr, ch.Username)
		str += genLineI(channelIDStr, ch.ID)
		if f.ParticipantsCount != wv.BaseIndex {
			str += genLineI(membersCountStr, f.ParticipantsCount)
		}
		if f.KickedCount != wv.BaseIndex {
			str += genLineI(kickedCountStr, f.KickedCount)
		}
		if f.BannedCount != wv.BaseIndex {
			str += genLineI(bannedCountStr, f.BannedCount)
		}
		str += genLineB(isVerifiedStr, ch.Verified)
		str += genLineB(hasBroadcastStr, ch.Broadcast)
		str += genLineB(hasSignaturesStr, ch.Signatures)
		str += genLineB(hasLinkStr, ch.HasLink)
		str += genLineB(isScamStr, ch.Scam)
		str += genLineB(leftChatStr, !ch.Left)
	} else {
		str += genLine(groupInfoStr, str)
		str += genLine(groupNameStr, ch.Title)
		str += genLine(groupUnameStr, ch.Username)
		str += genLineI(groupIDStr, ch.ID)
		str += genLineI(membersCountStr, f.ParticipantsCount)
		if f.KickedCount != wv.BaseIndex {
			str += genLineI(kickedCountStr, f.KickedCount)
		}
		if f.BannedCount != wv.BaseIndex {
			str += genLineI(bannedCountStr, f.BannedCount)
		}
		if f.AdminsCount != wv.BaseIndex {
			str += genLineI(adminsCountStr, f.AdminsCount)
		}
		if f.OnlineCount != wv.BaseIndex {
			str += genLineI(onlineCountStr, f.OnlineCount)
		}
		str += genLineB(issupergpStr, ch.Megagroup)
		str += genLineB(isVerifiedStr, ch.Verified)
		str += genLineB(isSlowModeStr, ch.SlowmodeEnabled)
		if ch.Photo != nil {
			ph, ok := ch.Photo.(*tg.ChatPhoto)
			if ok && ph != nil {
				str += genLineB(hasVideoStr, ph.HasVideo)
				str += genLineI(uDCIDStr, ph.DCID)
			}
		}

		str += genLineB(hasBroadcastStr, ch.Broadcast)
		str += genLineB(hasLinkStr, ch.HasLink)
		str += genLineB(activeCallStr, ch.CallActive)
		if ch.CallActive {
			str += genLineB(emptyCallStr, !ch.CallNotEmpty)
		}
		str += genLineB(isScamStr, ch.Scam)
		str += genLineB(leftChatStr, !ch.Left)
	}

	return str
}

func parseUserInfo(user *tg.User, md bool) string {
	yn := ws.YesOrNo
	str := wv.EMPTY
	var tmp, tmp2 interfaces.WMarkDown
	genLine := func(first, second string) string {
		if !md {
			return first + second + wv.LineEscape
		}
		if ws.IsEmpty(&second) {
			return wotoMD.GetBold(first + wv.LineEscape).ToString()
		}

		tmp = wotoMD.GetBold(first)
		tmp2 = wotoMD.GetMono(second + wv.LineEscape)
		return tmp.Append(tmp2).ToString()
	}
	genLineI := func(first string, second int) string {
		return genLine(first, strconv.Itoa(second))
	}
	genLineB := func(first string, second bool) string {
		return genLine(first, yn(second))
	}

	str += genLine(userInfoStr, str)
	if !ws.IsEmpty(&user.LastName) {
		str += genLine(uFirstNameStr, user.FirstName)
		str += genLine(uLastNameStr, user.FirstName)
	} else {
		str += genLine(uNameStr, user.FirstName)
	}

	str += genLine(uUsernameStr, user.Username)
	str += genLineI(userIDStr, user.ID)
	if user.Photo != nil {
		ph, ok := user.Photo.(*tg.UserProfilePhoto)
		if ok {
			if ph != nil {
				str += genLineI(uDCIDStr, ph.DCID)
				str += genLineB(hasVideoStr, ph.HasVideo)
			}
		}
	}

	str += genLineB(isScamStr, user.Scam)
	str += genLineB(isTgSupportStr, user.Support)
	str += genLineB(isVerifiedStr, user.Verified)
	tmp = wotoMD.GetBold(uMentionStr)
	tmp2 = wotoMD.GetUserMention(linkStr, int64(user.ID))
	str += tmp.Append(tmp2).ToString()

	tmp = wotoMD.GetHyperLink("\nThis is a test blah blah blah: "+
		user.FirstName, "https://t.me/"+user.Username)

	str += tmp.ToString()

	return str
}
