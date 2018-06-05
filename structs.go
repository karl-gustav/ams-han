package ams

import "time"

type MessageType int

func (m *MessageType) GetInt() int {
	return int(*m)
}

func (m *MessageType) GetByte() byte {
	return byte(*m)
}

const (
	MessageType1           MessageType = 1
	TwoFasesMessageType2   MessageType = 9
	TwoFasesMessageType3   MessageType = 14
	ThreeFasesMessageType2 MessageType = 13
	ThreeFasesMessageType3 MessageType = 18
)

type BaseItem struct {
	MessageType MessageType `json:"Message_Type"`
	MeterTime   time.Time   `json:"Meter_Time"`
	HostTime    time.Time   `json:"Host_Time"`
}

type Items1 struct {
	BaseItem
	ActPowPos int `json:"Act_Pow_P_Q1_Q4"` /* OBIS Code 1.0.1.7.0.255 - Active Power + (Q1+Q4) */
}

type Items9 struct {
	BaseItem
	ObisListVersion string `json:"OBIS_List_Version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string `json:"GS1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string `json:"Meter_Model"`       /* OBIS Code 0.0.96.1.7 - Meter Type */
	ActPowPos       int    `json:"Act_Pow_P_Q1_Q4"`   /* Active Power + */
	ActPowNeg       int    `json:"Act_Pow_M_Q2_Q3"`   /* Actve Power - */
	ReactPowPos     int    `json:"React_Pow_P"`       /* Reactive Power + */
	ReactPowNeg     int    `json:"React_Pow_M"`       /* Reactive Power - */
	CurrL1          int    `json:"Curr_L1"`           /* Current Phase L1 */
	VoltL1          int    `json:"Volt_L1"`           /* Voltage L1 */
}

type Items13 struct {
	BaseItem
	ObisListVersion string `json:"OBIS_List_Version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string `json:"GS1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string `json:"Meter_Model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int    `json:"Act_Pow_P_Q1_Q4"`   /* Active Power + */
	ActPowNeg       int    `json:"Act_Pow_M_Q2_Q3"`   /* Active Power - */
	ReactPowPos     int    `json:"React_Pow_P_Q1_Q2"` /* Reactive Power + */
	ReactPowNeg     int    `json:"React_Pow_M_Q3_Q4"` /* Reactive Power - */
	CurrL1          int    `json:"Curr_L1"`           /* Current Phase L1 mA */
	CurrL2          int    `json:"Curr_L2"`           /* Current Phase L2 mA */
	CurrL3          int    `json:"Curr_L3"`           /* Current Phase L3 mA */
	VoltL1          int    `json:"Volt_L1"`           /* Voltage L1 */
	VoltL2          int    `json:"Volt_L2"`           /* Voltage L2 */
	VoltL3          int    `json:"Volt_L3"`           /* Voltage L3 */
}

type Items14 struct {
	BaseItem
	ObisListVersion string    `json:"OBIS_List_Version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string    `json:"GS1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string    `json:"Meter_Model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int       `json:"Act_Pow_P_Q1_Q4"`   /* Active Power + */
	ActPowNeg       int       `json:"Act_Pow_M_Q2_Q3"`   /* Active Power - */
	ReactPowPos     int       `json:"React_Pow_P_Q1_Q2"` /* Reactive Power + */
	ReactPowNeg     int       `json:"React_Pow_M_Q3_Q4"` /* Reactive Power - */
	CurrL1          int       `json:"Curr_L1"`
	VoltL1          int       `json:"Volt_L1"`
	DateTime        time.Time `json:"Date_Time2"` /* OBIS Code: 0.0.1.0.0.255 - Clock and Date in Meter */
	ActEnergyPos    int       `json:"Act_Energy_P"`
	ActEnergyNeg    int       `json:"Act_Energy_M"`
	ReactEnergyPos  int       `json:"React_Energy_P"`
	ReactEnergyNeg  int       `json:"React_Energy_M"`
}

type Items18 struct {
	BaseItem
	ObisListVersion string    `json:"OBIS_List_Version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string    `json:"GS1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string    `json:"Meter_Model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int       `json:"Act_Pow_P_Q1_Q4"`   /* Active Power + */
	ActPowNeg       int       `json:"Act_Pow_M_Q2_Q3"`   /* Active Power - */
	ReactPowPos     int       `json:"React_Pow_P_Q1_Q2"` /* Reactive Power + */
	ReactPowNeg     int       `json:"React_Pow_M_Q3_Q4"` /* Reactive Power - */
	CurrL1          int       `json:"Curr_L1"`           /* Current L1 */
	CurrL2          int       `json:"Curr_L2"`           /* Current L2 */
	CurrL3          int       `json:"Curr_L3"`           /* Current L3 */
	VoltL1          int       `json:"Volt_L1"`           /* Voltage L1 */
	VoltL2          int       `json:"Volt_L2"`           /* Voltage L2 */
	VoltL3          int       `json:"Volt_L3"`           /* Voltage L3 */
	DateTime        time.Time `json:"Date_Time2"`        /* OBIS Code: 0.0.1.0.0.255 - Clock and Date in Meter */
	ActEnergyPa     int       `json:"Act_Energy_P"`      /* Active Energy +A */
	ActEnergyMa     int       `json:"Act_Energy_M"`      /* Active Energy -A */
	ActEnergyPr     int       `json:"React_Energy_P"`    /* Active Energy +R */
	ActEnergyMr     int       `json:"React_Energy_M"`    /* Active Energy -R */
}
