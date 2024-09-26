package adapter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type adapter struct {
	alifUrl    string
	userid     string
	httpClient *http.Client
	Password   string
}

func NewAdapter(alifUrl string, userid string, httpClient *http.Client, password string) integration.Adapter {
	return &adapter{alifUrl: alifUrl, userid: userid, httpClient: httpClient, Password: password}
}


func GetSha256(text string, secret []byte) string {

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(text))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash	
}

func (a adapter) PreCheck(req *ReqForCheck) (status int64, description string, rawInfo map[string]string, err error) {
	
	hash:=GetSha256((a.userid+"."+req.Datatime),[]byte(a.Password))
	req.Hash=hash
	req.UserID=a.userid

	var resp *RespForPreCheck
	err=a.DoRequest(req,resp)
	if err!=nil{
		return 0,"0",rawInfo,nil
	}

	rawInfo = make(map[string]string)
	rawInfo["message"] = resp.Message
	return int64(resp.Code),resp.Message,resp.AccountInfo,nil

}

func (a adapter) Payment(req *Req) (status int64, description string, paymentID string, err error){

	hash:=GetSha256(a.userid+req.Account+req.TxnID+"0",[]byte(a.userid))
	req.Hash=hash
	req.UserID=a.userid

	var resp Resp
	err=a.DoRequest(req,resp)
	if err!=nil{
		return  0,"","0",err
	}	
	
	return int64(resp.StatusCode),resp.Message,string(resp.ID),nil
}

func (a adapter) PostCheck(req *Req) (status int64, description string,err error){
	
	hash:=GetSha256(a.userid+req.Account+req.TxnID+"0",[]byte(a.userid))
	req.Hash=hash
	req.UserID=hash

	var resp Resp
	err=a.DoRequest(req,resp)
	if err!=nil{
		return 0,"",nil
	}
	return int64(resp.Code),resp.Message,nil
}	



func (a adapter) DoRequest(req any,resp any)(err error){
    
	convertReq,err:=json.Marshal(req)
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return err	
	}
	reqBody,err:=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(convertReq))
	if err!=nil{
		logrus.WithError(err).Info("on req")
		return nil	
	}
	reqBody.Header.Add("Content-Type", "application/json; charset=utf8")

	res,err:=a.httpClient.Do(reqBody)
	if err!=nil{
		logrus.WithError(err).Info("on response")
		return err
	}
	defer res.Body.Close()
	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading responce body")
		return err
	}
	logrus.WithFields(
		logrus.Fields{
			"body": string(body),
			"status": res.StatusCode,
		},
	).Info("response to adapter")

	err=json.Unmarshal(body,resp)
	if err!=nil{
		logrus.WithError(err).Info("on parsing response")
		return err
	}
	return nil
}