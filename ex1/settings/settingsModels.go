package settings

type Settings struct{
	OracleCFTDbParams `json:"oracleCFTDbParams"`
}

type OracleCFTDbParams struct {
	Server   string `mapstructure:"SERVER" json:"server"`
	User     string `mapstructure:"USER" json:"user"`
	Password string `mapstructure:"PASSWORD" json:"password"`
}