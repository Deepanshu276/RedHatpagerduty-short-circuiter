package ui

import "github.com/gdamore/tcell/v2"

const (

	// Text Format
	TitleFmt = " [lightcyan::b]%s "

	// Table Titles
	AlertsTableTitle          = "[ ALERTS ]"
	TrigerredAlertsTableTitle = "[ TRIGERRED ALERTS ]"
	HighAlertsTableTitle      = "[ TRIGERRED ALERTS - HIGH ]"
	LowAlertsTableTitle       = "[ TRIGERRED ALERTS - LOW ]"
	AlertMetadataViewTitle    = "[ ALERT DATA ]"
	IncidentsTableTitle       = "[ TRIGERRED INCIDENTS ]"
	AckIncidentsTableTitle    = "[ ACKNOWLEDGED INCIDENTS ]"
	OncallTableTitle          = "ONCALL"
	NextOncallTableTitle      = "[ NEXT ONCALL ]"
	AllTeamsOncallTableTitle  = "[ ALL TEAMS ONCALL ]"

	// Page Titles
	AlertsPageTitle          = "Alerts"
	AlertDataPageTitle       = "Metadata"
	AlertMetadata            = "AlertData"
	TrigerredAlertsPageTitle = "Trigerred"
	HighAlertsPageTitle      = "High Alerts"
	LowAlertsPageTitle       = "Low Alerts"
	IncidentsPageTitle       = "Incidents"
	AckIncidentsPageTitle    = "AckIncidents"
	OncallPageTitle          = "Oncall Layer"
	NextOncallPageTitle      = "Next Oncall"
	AllTeamsOncallPageTitle  = "All Teams Oncall"
	ServiceLogsPageTitle     = "Service Logs"

	//Footer
	FooterText                = "[Esc] Go Back"
	FooterTextAlerts          = "[R] Refresh Alerts | [1] Acknowledged Incidents | [2] Trigerred Incidents\n" + FooterText
	FooterTextTrigerredAlerts = "[1] Acknowledged Incidents | [2] Trigerred Incidents\n" + FooterText
	FooterTextIncidents       = "[ENTER] Select Incident | [CTRL+A] Acknowledge Incidents\n" + FooterText
	FooterTextOncall          = "[N] Your Next Oncall Schedule | [A] All Teams Oncall | [<-] Previous Layer Oncall | [->] Next Layer Oncall \n" + FooterText
	TerminalFooterText        = "[CTRL + N] Next Slide | [CTRL + P] Previous Slide | [CTRL + A] Add Slide | [CTRL + E] Exit Slide | [CTRL + B] + [Num] Change to Slide with [Num]  | [CTRL + Q] Quit "
	TerminalFooterEscapeState = "Enter the Slide Number to Switch To : "

	// Colors
	TableTitleColor                = tcell.ColorLightCyan
	BorderColor                    = tcell.ColorLightGray
	FooterTextColor                = tcell.ColorGray
	InfoTextColor                  = tcell.ColorLightSlateGray
	ErrorTextColor                 = tcell.ColorRed
	PromptTextColor                = tcell.ColorLightGreen
	LoggerTextColor                = tcell.ColorGreen
	TerminalFooterTextColor        = tcell.ColorGreen
	TerminalFooterEscapeStateColor = tcell.ColorDarkGreen
)
