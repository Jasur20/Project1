package request

import (
	// "brt_adapter/settings"
	// "brt_ispc/models"
	// "fmt"
	// "io"
	// "net/http"
)

func FimiGetInfo() string{
	return "Fimi request"
} 
// func FimiGetTranInfo(tranNumber, Type, pan string) (*models.ResponceNPC, *models.ErrorModel) {
// 	var (
// 		err      error
// 		url      = fmt.Sprintf(settings.AppSettings.FIMI.TranInfoURL, tranNumber, Type, pan)
// 		client   = http.Client{Timeout: settings.AppSettings.AppParams.TimeoutReq2NPCsec * time.Second}
// 		errModel = &models.ErrorModel{}
// 		npc      = &models.ResponceNPC{}
// 	)
// 	errModel.ErrorType = "FimiGetTranInfoError"

// 	resp, err := client.Get(url)
// 	if err != nil {
// 		errModel.ErrorCode = 400
// 		errModel.ErrorText = err.Error()
// 		return nil, errModel
// 	}
// 	defer resp.Body.Close()

// 	resString, _ := io.ReadAll(resp.Body)
// 	if resp.StatusCode != 200 {
// 		errModel.ErrorCode = resp.StatusCode
// 		errModel.ErrorText = string(resString)
// 		return nil, errModel
// 	}

// 	log.Println("FimiRequest -> response string: ", string(resString))
// 	// remove xml header
// 	req := regexp.MustCompile("(?m)^.*xml version.*$").ReplaceAllString(string(resString), "")
// 	req = removeNonXMLCharacters(req)

// 	if err = xml.Unmarshal([]byte(req), &npc); err != nil {
// 		errModel.ErrorCode = 400
// 		errModel.ErrorText = err.Error()
// 		return nil, errModel
// 	}

// 	return npc, nil
// }
