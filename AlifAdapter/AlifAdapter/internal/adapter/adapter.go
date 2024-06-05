package adapter

import (
	"alif/internal/integration"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)


type adapter struct {
	alifUrl       string
	userid        string
	httpClient    *http.Client
	secret        string
}

func NewAdapter(alifUrl string ,userid string,httpClient *http.Client,secret string) integration.Adapter {
	return &adapter{alifUrl: alifUrl, userid: userid, httpClient: httpClient,secret: secret}
}

func GetSha256(text string, secret []byte) string {

	h := hmac.New(sha256.New, secret)

	h.Write([]byte(text))

	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func (a adapter) PreCheck(data *CheckAccountReq) (status int64, description string, rawInfo map[string]string, err error) {
	
	datatime:=time.Now()
	hash:=GetSha256(a.userid+"."+datatime.String(),[]byte(a.secret))
	data.UserID=a.userid
	data.Hash=hash
	
    temp,err:=json.Marshal(data)
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return 0,"0",rawInfo,nil
	}

	var resBody CheckAccountResp
	body,err:=SendRequest(temp,a.alifUrl)
	if err!=nil{
		logrus.WithError(err).Info("on sendRequest")
		return 0,"",rawInfo,err
	}

	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on unmarshaling")
		return 0,"",rawInfo,nil
	}
	rawInfo = make(map[string]string)
	rawInfo["message"] = resBody.Message
	Code, err := strconv.ParseFloat(resBody.Code, 64)

	return int64(Code),resBody.Message,rawInfo,nil

}

func (a adapter) Payment(data *Req) (status int64, description string, paymentID string, err error){

	hash:=GetSha256(a.userid+data.Account+data.TxnID+"0",[]byte(a.userid))
	
	data.Hash=hash
	data.UserID=a.userid

	temp,err:=json.Marshal(data)
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return 0,"","0",err 
	}

	body,err:=SendRequest(temp,a.alifUrl)
	if err!=nil{
		logrus.WithError(err).Info("on sendRequest")
		return 0,"","0",err
	}
	var resBody *Resp
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on unmarshaling")
		return 0,"","",nil
	}

	return int64(resBody.StatusCode),resBody.Message,string(resBody.ID),nil
	
}

func (a adapter) PostCheck(data *Req) (status int64, description string,err error){
	
	hash:=GetSha256(a.userid+data.Account+data.TxnID+"0",[]byte(a.userid))
	data.Hash=hash
	data.UserID=a.alifUrl
	reqbody,err:=json.Marshal(data)
	if err!=nil{
		logrus.WithError(err).Info("on marshaling data")
		return 0,"",nil
	}

	body,err:=SendRequest(reqbody,a.alifUrl)
	if err!=nil{
		logrus.WithError(err).Info("on respBody")
		return 0,"",nil
	}
	var resBody *Resp
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on unmarshaling")
		return 0,"",nil
	}
	return int64(resBody.StatusCode),resBody.Message,nil
}	

func SendRequest(temp []byte,alifUrl string/*,construct any*/)(body []byte ,err error){

	req,err:=http.NewRequest(http.MethodPost,alifUrl,bytes.NewBuffer(temp))
	if err!=nil{
		logrus.WithError(err).Warn("on sendRequest")
		return nil,err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf8")
	var a adapter
	res,err:=a.httpClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on response")
		return nil,err
	}
	defer res.Body.Close()
	body,err=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading responce body")
		return nil,err
	}
	logrus.WithFields(
		logrus.Fields{
			"body": string(body),
			"status": res.StatusCode,
		},
	).Info("response to adapter")

	return body,nil
}

