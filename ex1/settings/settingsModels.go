package settings

type Settings struct{
	AppParams `json:"appparams"`
	OracleCFTDbParams `json:"oracleCFTDbParams"`
	FimiService `json:"fimiservice"`
}

type AppParams struct{
	TimeoutReq2NPCSec string `json:"timeoutreq2npcsec"`
}
type OracleCFTDbParams struct {
	Server   string `mapstructure:"SERVER" json:"server"`
	User     string `mapstructure:"USER" json:"user"`
	Password string `mapstructure:"PASSWORD" json:"password"`
}

type FimiService struct{
	FimiUrl string `json:"fimiurl"`
	FimiSetPinURL string `json:"fimisetpinurl"`
	FimiTranInfoURL string `json:"fimitraninfourl"`
	FimiReverseTranUrl string `json:"fimireversetranurl"`
}