package ams

import "time"

type messageTypes int

func (m *messageTypes) GetInt() int {
	return int(*m)
}

func (m *messageTypes) GetByte() byte {
	return byte(*m)
}

const (
	messageType1           messageTypes = 1
	twoFasesMessageType2   messageTypes = 9
	twoFasesMessageType3   messageTypes = 14
	threeFasesMessageType2 messageTypes = 13
	threeFasesMessageType3 messageTypes = 18
)

type BaseItem struct {
	MessageType messageTypes `json:"Message_Type"`
	MeterTime   time.Time    `json:"Meter_Time"`
	HostTime    time.Time    `json:"Host_Time"`
}

type MessageType1 struct { // Also known as aka Items1
	BaseItem
	ActPowPos int `json:"Act_Pow_P_Q1_Q4"` /* OBIS Code 1.0.1.7.0.255 - Active Power + (Q1+Q4) */
}

type TwoFasesMessageType2 struct { // Also known as aka Items9
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

type ThreeFasesMessageType2 struct { // Also known as Items13
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

type TwoFasesMessageType3 struct { // Also known as Items14
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

type ThreeFasesMessageType3 struct { // Also known as Items18
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
