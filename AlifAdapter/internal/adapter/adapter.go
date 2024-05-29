package adapter

import (
	"alif/internal/integration"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)


type adapter struct {
	alifUrl       string
	userid        string
	httpClient    *http.Client
}

func NewAdapter(alifUrl string ,userid string,httpClient *http.Client) integration.Adapter {
	return &adapter{alifUrl: alifUrl, userid: userid, httpClient: httpClient}
}

func GetSha256(text string, secret []byte) string {

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(text))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func (a adapter) PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error) {
 
	body,err:=SendRequest(account,serviceID,"0","0","0")
	if err!=nil{
		return 0,"0",rawInfo,nil
	}
	resBody:=Resp{}
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on parsing responce")
		return 0,"0",rawInfo,nil
	}
	rawInfo = make(map[string]string)
	rawInfo["message"] = resBody.Message
	return int64(resBody.Code),resBody.Message,resBody.AccountInfo,nil

}

func (a adapter) Payment(account string,serviceID string,amount string,trnID string, notifyRoute string) (status int64, description string, paymentID string, err error){

	// hash:=GetSha256(a.userid+account+trnID+"0",[]byte(a.userid))
	// data,err:=json.Marshal(Req{Service:serviceID,UserID: a.userid,Hash: hash,Account: account,Amount: amount,Currency: "TJS",TxnID: "",Phone: "",
	// Fee: "",Providerid: "0",Last_Name: "",First_Name: "",Middle_Name: "",Sender_Birthday: "",Address: "",Resident_Country: "",Postal_Code: "",Recipient_Name: ""})
	body,err:=SendRequest(account,serviceID,amount,trnID,notifyRoute)
	if err!=nil{
		return  0,"","0",err
	}	
	
	resBody:=&RespForPayment{}
	err=json.Unmarshal(body,resBody)
	if err!=nil{
		logrus.WithError(err).Info("on parsing response")
		return 0, "","0",err
	}
	return int64(resBody.StatusCode),resBody.Message,string(resBody.ID),nil
}

func (a adapter) PostCheck(trnID string,account string,serviceID string) (status int64, description string,err error){

	res,err:=a.httpClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on response")
		return nil,err
	}
	// body,err:=SendRequest(account,serviceID,"0",trnID,"0")
	// if err!=nil{
	// 	return 0,"",nil
	// }
	
	resBody:=Resp{}
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on unmarshaling resBoby")
		return 0,"0",nil
	}
	return int64(resBody.Code),resBody.Message,nil
}	

func SendRequest(account,serviceid,amount,trnID,notifyRoute string,konstruct )(body []byte,err error){
	var a adapter

	time:=time.Now()
	hash:=GetSha256(a.userid+account+"0"+"0",[]byte(a.userid))
    data,err:=json.Marshal(ReqFor{Account: account,Currency: "TJS",Hash: hash,ProviderID: 0,Service: serviceID,UserID: a.userid,Datatime:time})
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return 0,"0",rawInfo,nil	
	}
	req,err:=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(reqBody))
	if err!=nil{
		logrus.WithError(err).Info("on req")
		return nil,nil	
	}
	req.Header.Add("Content-Type", "application/json; charset=utf8")

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