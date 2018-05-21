package decode_mbus

import "time"

type Items1 struct {
	ActPowPos int `json:"act_pow_pos"` /* OBIS Code 1.0.1.7.0.255 - Active Power + (Q1+Q4) */
}

type Items9 struct {
	ObisListVersion string `json:"obis_list_version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string `json:"gs1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string `json:"meter_model"`       /* OBIS Code 0.0.96.1.7 - Meter Type */
	ActPowPos       int    `json:"act_pow_pos"`       /* Active Power + */
	ActPowNeg       int    `json:"act_pow_neg"`       /* Actve Power - */
	ReactPowPos     int    `json:"react_pow_pos"`     /* Reactive Power + */
	ReactPowNeg     int    `json:"react_pow_neg"`     /* Reactive Power - */
	CurrL1          int    `json:"curr_L1"`           /* Current Phase L1 */
	VoltL1          int    `json:"volt_L1"`           /* Voltage L1 */
}

type Items13 struct {
	ObisListVersion string `json:"obis_list_version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string `json:"gs1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string `json:"meter_model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int    `json:"act_pow_pos"`       /* Active Power + */
	ActPowNeg       int    `json:"act_pow_neg"`       /* Active Power - */
	ReactPowPos     int    `json:"react_pow_pos"`     /* Reactive Power + */
	ReactPowNeg     int    `json:"react_pow_neg"`     /* Reactive Power - */
	CurrL1          int    `json:"curr_L1"`           /* Current Phase L1 mA */
	CurrL2          int    `json:"curr_L2"`           /* Current Phase L2 mA */
	CurrL3          int    `json:"curr_L3"`           /* Current Phase L3 mA */
	VoltL1          int    `json:"volt_L1"`           /* Voltage L1 */
	VoltL2          int    `json:"volt_L2"`           /* Voltage L2 */
	VoltL3          int    `json:"volt_L3"`           /* Voltage L3 */
}

type Items14 struct {
	ObisListVersion string    `json:"obis_list_version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string    `json:"gs1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string    `json:"meter_model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int       `json:"act_pow_pos"`       /* Active Power + */
	ActPowNeg       int       `json:"act_pow_neg"`       /* Active Power - */
	ReactPowPos     int       `json:"react_pow_pos"`     /* Reactive Power + */
	ReactPowNeg     int       `json:"react_pow_neg"`     /* Reactive Power - */
	CurrL1          int       `json:"curr_L1"`
	VoltL1          int       `json:"volt_L1"`
	DateTime        time.Time `json:"date_time"` /* OBIS Code: 0.0.1.0.0.255 - Clock and Date in Meter */
	ActEnergyPos    int       `json:"act_energy_pos"`
	ActEnergyNeg    int       `json:"act_energy_neg"`
	ReactEnergyPos  int       `json:"react_energy_pos"`
	ReactEnergyNeg  int       `json:"react_energy_neg"`
}

type Items18 struct {
	ObisListVersion string    `json:"obis_list_version"` /* OBIS Code 1.1.0.2.129.255 - OBIS List Version Identifier */
	Gs1             string    `json:"gs1"`               /* OBIS Code 0.0.96.1.0.255 - Meter-ID(GIAI GS1 - 16 digits */
	MeterModel      string    `json:"meter_model"`       /* OBIS Code 0.0.96.1.7.255 - Meter Type */
	ActPowPos       int       `json:"act_pow_pos"`       /* Active Power + */
	ActPowNeg       int       `json:"act_pow_neg"`       /* Active Power - */
	ReactPowPos     int       `json:"react_pow_pos"`     /* Reactive Power + */
	ReactPowNeg     int       `json:"react_pow_neg"`     /* Reactive Power - */
	CurrL1          int       `json:"curr_L1"`           /* Current L1 */
	CurrL2          int       `json:"curr_L2"`           /* Current L2 */
	CurrL3          int       `json:"curr_L3"`           /* Current L3 */
	VoltL1          int       `json:"volt_L1"`           /* Voltage L1 */
	VoltL2          int       `json:"volt_L2"`           /* Voltage L2 */
	VoltL3          int       `json:"volt_L3"`           /* Voltage L3 */
	DateTime        time.Time `json:"date_time"`         /* OBIS Code: 0.0.1.0.0.255 - Clock and Date in Meter */
	ActEnergyPa     int       `json:"act_energy_pa"`     /* Active Energy +A */
	ActEnergyMa     int       `json:"act_energy_ma"`     /* Active Energy -A */
	ActEnergyPr     int       `json:"act_energy_pr"`     /* Active Energy +R */
	ActEnergyMr     int       `json:"act_energy_mr"`     /* Active Energy -R */
}
