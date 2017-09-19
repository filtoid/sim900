package sim900

// AT commands
const (
	CMD_AT             string = "AT"
	CMD_OK             string = "(^OK$)"
	CMD_ERROR          string = "(^ERROR$)"
	CMD_CMGF           string = "AT+CMGF?"
	CMD_CMGF_SET       string = "AT+CMGF=%s"
	CMD_CMGF_REGEXP    string = "(^[+]CMGF[:] [0-9]+$)"
	CMD_CMGF_RX        string = "+CMGF: "
	CMD_CTRL_Z         string = "\x1A"
	CMD_CMGS           string = "AT+CMGS=\"%s\""
	CMD_CMGS_RX_REGEXP string = "(^[+]CMGS[:] [0-9]+$)"
	CMD_CMGD           string = "AT+CMGD=%s"
	CMD_CMGR           string = "AT+CMGR=%s"
	CMD_CMGR_REGEXP    string = "(^[+]CMGR[:] .*)"
	CMD_CMGR_RX        string = "+CMGR: "
	CMD_CMTI_REGEXP    string = "(^[+]CMTI[:] \"SM\",[0-9]+$)"
	CMD_CMTI_RX        string = "+CMTI: \"SM\","
	CMD_SAPBR          string = "AT+SAPBR="
	CMD_HTTPINIT       string = "AT+HTTPINIT"
	CMD_HTTPPARA       string = "AT+HTTPPARA="
	CMD_HTTPDATA       string = "AT+HTTPDATA="
	CMD_HTTPACTION     string = "AT+HTTPACTION="
	CMD_HTTPREAD       string = "AT+HTTPREAD"
	CMD_HTTPTERM       string = "AT+HTTPTERM"
)

// SMS Message Format
const (
	PDU_MODE  string = "0"
	TEXT_MODE string = "1"
)
