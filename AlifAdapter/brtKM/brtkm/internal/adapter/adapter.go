package adapter

import (
	"brtkm/internal/integration"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"
)

type adapter struct {
	cftUrl        string
	ispcUrl       string
	agentLogin    string
	agentPassword string
	pans          []string
	httpClient    *http.Client
}

func NewAdapter(cftUrl string, ispcUrl string, agentLogin string, agentPassword string, pans []string, httpClient *http.Client) integration.Adapter {
	return &adapter{cftUrl: cftUrl, ispcUrl: ispcUrl, agentLogin: agentLogin, agentPassword: agentPassword, pans: pans, httpClient: httpClient}
}

func (a adapter) PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error) {

	if slices.Contains(a.pans, account[0:9]) {

		status, description, rawInfo, err = cftCheck(account, a.cftUrl, a.httpClient)

	} else {
		//TODO KM check pan
	}

	return status, description, rawInfo, err
}

// Payment implements integration.Adapter.
func (a *adapter) Payment(account string, serviceID string, amount string, trnID string, notifyRoute string) (status int64, description string, paymentID string, err error) {

	data := url.Values{}
	data.Add("agentLogin", a.agentLogin)
	data.Add("agentPassword", a.agentPassword)
	data.Add("cardHash", "")
	data.Add("cardNumber", account)
	data.Add("clientCode", "0")
	data.Add("transID", trnID)
	data.Add("amount", amount)
	data.Add("reason", "")
	encodedBody := data.Encode()
	req, err := http.NewRequest(http.MethodPost, a.ispcUrl+"/brt-v1/npc/v2c", strings.NewReader(encodedBody))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		logrus.WithError(err).Info("on creating request")
		return 0, "", "0", err
	}
	res, err := a.httpClient.Do(req)
	if err != nil {
		logrus.WithError(err).Info("on doing request")
		return 0, "", "0", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.WithError(err).Warn("on reading response body")
		return 0, "", "0", err
	}

	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")

	resBody := &PaymentResp{}
	err = xml.Unmarshal(body, resBody)
	if err != nil {
		logrus.WithError(err).Info("on parsing response")
		return 0, "", "0", err
	}

	return int64(resBody.ResultCode), resBody.Result, resBody.TransID, nil
}

// PostCheck implements integration.Adapter.
func (a *adapter) PostCheck(trnID string) (status int64, description string, err error) {

	data := url.Values{}
	data.Add("agentLogin", a.agentLogin)
	data.Add("transID", trnID)
	encodedBody := data.Encode()
	req, err := http.NewRequest(http.MethodPost, a.ispcUrl+"/brt-v1/cards/checkStatus", strings.NewReader(encodedBody))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		logrus.WithError(err).Info("on creating request")
		return 0, "", err
	}
	res, err := a.httpClient.Do(req)
	if err != nil {
		logrus.WithError(err).Info("on doing request")
		return 0, "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.WithError(err).Warn("on reading response body")
		return 0, "", err
	}

	logrus.WithFields(logrus.Fields{
		"body":   string(body),
		"status": res.StatusCode,
	}).Info("response to adapter")

	if res.StatusCode == http.StatusNotFound {
		resBody := &PostCheckRespError{}
		err = xml.Unmarshal(body, resBody)
		if err != nil {
			logrus.WithError(err).Info("on parsing response error")
			return 0, "", err
		}
		return int64(resBody.ErrorCode), resBody.ErrorText, nil
	}
	resBody := &PaymentResp{}
	err = xml.Unmarshal(body, resBody)
	if err != nil {
		logrus.WithError(err).Info("on parsing response")
		return 0, "", err
	}
	return int64(resBody.ResultCode), resBody.Result, nil
}

func cftCheck(account string, url string, httpClient *http.Client) (status int64, description string, rawInfo map[string]string, err error) {
	req, err := http.NewRequest(http.MethodGet, url+"/card/holder/"+account, nil)
	if err != nil {
		logrus.WithError(err).Info("on creating request")
		return 0, "", nil, err
	}
	res, err := httpClient.Do(req)

	if err != nil {
		logrus.WithError(err).Warn("on doing request")
		return 0, "", nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.WithError(err).Warn("on reading response body")
		return 0, "", nil, err
	}

	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")
	resBody := &PreCheckResp{}
	err = json.Unmarshal(body, resBody)
	if err != nil {
		logrus.WithError(err).Info("on parsing response")
		return 0, "", nil, err
	}
	rawInfo = make(map[string]string)
	rawInfo["message"] = resBody.Message
	//rawInfo["currency"] = resBody.Currency

	return int64(res.StatusCode), res.Status, rawInfo, nil
}
