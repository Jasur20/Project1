package adapter

import (
	"crypto/md5"
	"dc_adapter/internal/usecase"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type adapter struct {
	dcUrl      string
	login      string
	httpClient *http.Client
	password   string
}

func NewAdapter(dcUrl string, login string, httpClient *http.Client, password string) usecase.Adapter {
	return &adapter{dcUrl: dcUrl, login: login, httpClient: httpClient, password: password}
}

// func GetHmac(text string, secret []byte) string {

// 	h := hmac.New(md5.New, secret)
// 	h.Write([]byte(text))
// 	hash := hex.EncodeToString(h.Sum(nil))
// 	return hash
// }

func (a adapter) PreCheck(reqBody *usecase.ReqForCheck) (status int, description string, rawInfo map[string]string, err error) {

	hashText := fmt.Sprintf("%s%s", a.login, a.password)
	hash := md5.Sum([]byte(hashText))
	sign := hex.EncodeToString(hash[:])

	url := fmt.Sprintf("%s?command=%s&login=%s&txn_id=%s&account=%s&prvid=%s&ccy=%s&sum=%.2f&txn_date=%s&sign=%s",
		a.dcUrl, reqBody.Command, a.login, "1", reqBody.Account, reqBody.Prv_ID, reqBody.Ccy, reqBody.Sum, reqBody.Txn_Date, sign)

	logrus.WithFields(logrus.Fields{"url": url}).Info("URL PreCheck")
	
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("DC preCheck request error: %v", err.Error())
		return
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logrus.Errorf("Couldn't read response body: %v", err.Error())
		return
	}

	logrus.Infof("Response from adapter(PreCheck): %s", string(body))

	var response PreCheckResp

	err = xml.Unmarshal(body, &response)
	if err != nil {
		logrus.Errorf("Error during unmarshaling: %v", err.Error())
		return
	}
	fmt.Println(response.Comment)
	fmt.Println(response.Info)
	fmt.Println(response.Txn_Date)
	status = response.Result
	description = response.Comment
	rawInfo = make(map[string]string)
	rawInfo["name"] = fmt.Sprint(response.Info) //response.PreCheckInfo.RawInfo
	return

}

func (a adapter) Payment(reqBody *usecase.ReqForPayment) (status int, description string, paymentID string, err error) {

	// serviceId, err := strconv.Atoi(reqBody.ServiceID)
	// if err != nil {
	// 	logrus.Errorf("err with convert servise_id ot int:  %v", err.Error())
	// 	return
	// }

	hashText := fmt.Sprintf("%s%s%s%s", a.login, reqBody.Txn_ID, reqBody.Account, a.password)

	hash := md5.Sum([]byte(hashText))
	sign := hex.EncodeToString(hash[:])
	url := fmt.Sprintf("%s?command=%s&login=%s&txn_id=%s&prvid=%s&account=%s&ccy=%s&sum=%.2f&sign=%s&cr_amount=%s&txn_date=%s&wallet=",
		a.dcUrl, reqBody.Command, a.login, reqBody.Txn_ID, reqBody.Prv_ID, reqBody.Account, reqBody.Ccy, reqBody.Sum, sign, "true", reqBody.Txn_Date)

	logrus.WithFields(logrus.Fields{"url": url}).Info("URL Payment")

	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("DC payment request error: %v", err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logrus.Errorf("Couldn't read response body: %v", err.Error())
		return
	}

	logrus.Infof("Response from adapter(Payment): %s", string(body))

	var response PaymentResp

	err = xml.Unmarshal(body, &response)

	if err != nil {
		logrus.Errorf("Couldn't unmarshal body: %v", err.Error())
		return
	}

	fmt.Println(response.Comment)
	fmt.Println(response.TxnDate)

	status = response.Result
	description = response.Comment
	paymentID = fmt.Sprint(response.PrvTxn)

	return
}

func (a adapter) PostCheck(req *usecase.ReqForPostCheck) (status int, description string, err error) {

	hashText := fmt.Sprintf("%s%s%s", a.login, req.TxnId, a.password)

	hash := md5.Sum([]byte(hashText))
	sign := hex.EncodeToString(hash[:])
	url := fmt.Sprintf("%s?command=%s&login=%s&txn_id=%s&ccy=%s&sign=%s",
		a.dcUrl, req.Command, a.login, req.TxnId,"TJS",sign)

	logrus.WithFields(logrus.Fields{"url": url}).Info("URL PostCheck")

	resp, err := http.Get(url)

	if err != nil {
		logrus.Errorf("DC postCheck request error: %v", err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logrus.Errorf("Couldn't read response body: %v", err.Error())
		return
	}

	logrus.Infof("Response from adapter(PostCheck): %s", string(body))

	var response PostCheckResp
	fmt.Println((body))

	err = xml.Unmarshal(body, &response)

	if err != nil {
		logrus.Errorf("Couldn't unmarshal body: %v", err.Error())
		return
	}

	status = response.Result
	description = response.Comment

	return
}
