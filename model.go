package uchatlib

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/lvzhihao/goutils"
)

type UchatMessage struct {
	MerchantNo       string
	LogSerialNo      string
	ChatRoomSerialNo string
	WxUserSerialNo   string
	MsgTime          time.Time
	MsgType          int32
	Content          string
	VoiceTime        int32
	ShareTitle       string
	ShareDesc        string
	ShareUrl         string
	ExtraData        interface{} //补充数据，并非接口返回
}

func ConvertUchatMessage(b []byte) ([]*UchatMessage, error) {
	var rst map[string]interface{}
	err := json.Unmarshal(b, &rst)
	if err != nil {
		return nil, err
	}
	merchantNo, ok := rst["vcMerchantNo"]
	if !ok {
		return nil, errors.New("empty merchantNo")
	}
	data, ok := rst["Data"]
	if !ok {
		return nil, errors.New("empty Data")
	}
	var list []map[string]interface{}
	err = json.Unmarshal([]byte(goutils.ToString(data)), &list)
	if err != nil {
		return nil, err
	}
	ret := make([]*UchatMessage, 0)
	for _, v := range list {
		msg := &UchatMessage{}
		msg.MerchantNo = goutils.ToString(merchantNo)
		msg.LogSerialNo = goutils.ToString(v["vcSerialNo"])
		msg.ChatRoomSerialNo = goutils.ToString(v["vcChatRoomSerialNo"])
		msg.WxUserSerialNo = goutils.ToString(v["vcFromWxUserSerialNo"])
		msg.MsgTime, _ = time.ParseInLocation("2006-01-02 15:04:05", goutils.ToString(v["dtMsgTime"]), UchatTimeLocation)
		msg.MsgType = goutils.ToInt32(v["nMsgType"])
		content, err := base64.StdEncoding.DecodeString(goutils.ToString(v["vcContent"]))
		if err != nil {
			msg.Content = goutils.ToString(v["vcContent"])
		} else {
			msg.Content = goutils.ToString(content)
		}
		msg.VoiceTime = goutils.ToInt32(v["nVoiceTime"])
		msg.ShareTitle = goutils.ToString(v["vcShareTitle"])
		msg.ShareDesc = goutils.ToString(v["vcShareDesc"])
		msg.ShareUrl = goutils.ToString(v["vcShareUrl"])
		ret = append(ret, msg)
	}
	return ret, nil
}

// keyword model
type UchatKeyword struct {
	MerchantNo         string
	LogSerialNo        string
	ChatRoomSerialNo   string
	FromWxUserSerialNo string
	ToWxUserSerialNo   string
	Content            string
	ExtraData          interface{} //补充数据，并非接口返回
}

func ConvertUchatKeyword(b []byte) ([]*UchatKeyword, error) {
	var rst map[string]interface{}
	err := json.Unmarshal(b, &rst)
	if err != nil {
		return nil, err
	}
	merchantNo, ok := rst["vcMerchantNo"]
	if !ok {
		return nil, errors.New("empty merchantNo")
	}
	data, ok := rst["Data"]
	if !ok {
		return nil, errors.New("empty Data")
	}
	var list []map[string]interface{}
	err = json.Unmarshal([]byte(goutils.ToString(data)), &list)
	if err != nil {
		return nil, err
	}
	ret := make([]*UchatKeyword, 0)
	for _, v := range list {
		key := &UchatKeyword{}
		key.MerchantNo = goutils.ToString(merchantNo)
		key.LogSerialNo = goutils.ToString(v["vcSerialNo"])
		key.ChatRoomSerialNo = goutils.ToString(v["vcChatRoomSerialNo"])
		key.FromWxUserSerialNo = goutils.ToString(v["vcFromWxUserSerialNo"])
		key.ToWxUserSerialNo = goutils.ToString(v["vcToWxUserSerialNo"])
		key.Content = goutils.ToString(v["vcContent"])
		ret = append(ret, key)
	}
	return ret, nil
}

// MemberJoin model
// todo add more chatroom info
type UchatMemberJoin struct {
	MerchantNo           string
	ChatRoomSerialNo     string
	WxUserSerialNo       string
	FatherWxUserSerialNo string
	WxId                 string
	NickName             string
	HeadImages           string
	JoinChatRoomType     int32
	JoinDate             time.Time
	ExtraData            interface{} //补充数据，并非接口返回
}

func ConvertUchatMemberJoin(b []byte) ([]*UchatMemberJoin, error) {
	var rst map[string]interface{}
	err := json.Unmarshal(b, &rst)
	if err != nil {
		return nil, err
	}
	merchantNo, ok := rst["vcMerchantNo"]
	if !ok {
		return nil, errors.New("empty merchantNo")
	}
	data, ok := rst["Data"]
	if !ok {
		return nil, errors.New("empty Data")
	}
	var list []map[string]interface{}
	err = json.Unmarshal([]byte(goutils.ToString(data)), &list)
	if err != nil {
		return nil, err
	}
	ret := make([]*UchatMemberJoin, 0)
	for _, v := range list {
		key := &UchatMemberJoin{}
		key.MerchantNo = goutils.ToString(merchantNo)
		key.ChatRoomSerialNo = goutils.ToString(v["vcChatRoomSerialNo"])
		key.WxUserSerialNo = goutils.ToString(v["vcWxUserSerialNo"])
		key.FatherWxUserSerialNo = goutils.ToString(v["vcFatherWxUserSerialNo"])
		key.WxId = goutils.ToString(v["vcWxId"])
		nickName, err := base64.StdEncoding.DecodeString(goutils.ToString(v["vcBase64NickName"]))
		if err != nil {
			key.NickName = goutils.ToString(v["vcNickName"])
		} else {
			key.NickName = goutils.ToString(nickName)
		}
		key.HeadImages = goutils.ToString(v["vcHeadImages"])
		key.JoinChatRoomType = goutils.ToInt32(v["nJoinChatRoomType"])
		key.JoinDate, _ = time.ParseInLocation("2006-01-02T15:04:05.999", goutils.ToString(v["dtCreateDate"]), UchatTimeLocation)
		ret = append(ret, key)
	}
	return ret, nil
}

// MemberQuit model
type UchatMemberQuit struct {
	MerchantNo       string
	ChatRoomSerialNo string
	WxUserSerialNo   string
	WxId             string
	NickName         string
	QuitDate         time.Time
	ExtraData        interface{} //补充数据，并非接口返回
}

func ConvertUchatMemberQuit(b []byte) ([]*UchatMemberQuit, error) {
	var rst map[string]interface{}
	err := json.Unmarshal(b, &rst)
	if err != nil {
		return nil, err
	}
	merchantNo, ok := rst["vcMerchantNo"]
	if !ok {
		return nil, errors.New("empty merchantNo")
	}
	data, ok := rst["Data"]
	if !ok {
		return nil, errors.New("empty Data")
	}
	var list []map[string]interface{}
	err = json.Unmarshal([]byte(goutils.ToString(data)), &list)
	if err != nil {
		return nil, err
	}
	ret := make([]*UchatMemberQuit, 0)
	for _, v := range list {
		key := &UchatMemberQuit{}
		key.MerchantNo = goutils.ToString(merchantNo)
		key.ChatRoomSerialNo = goutils.ToString(v["vcChatRoomSerialNo"])
		key.WxUserSerialNo = goutils.ToString(v["vcWxUserSerialNo"])
		key.WxId = goutils.ToString(v["vcWxId"])
		key.NickName = goutils.ToString(v["vcNickName"])
		key.QuitDate, _ = time.ParseInLocation("2006-01-02T15:04:05", goutils.ToString(v["dtCreateDate"]), UchatTimeLocation)
		ret = append(ret, key)
	}
	return ret, nil
}
