package valueobject

type EntrantStatus string

const (
	EntrantStatusRegistered EntrantStatus = "registered"
	EntrantStatusConfirmed  EntrantStatus = "confirmed"
	EntrantStatusStarted    EntrantStatus = "started"
	EntrantStatusFinished   EntrantStatus = "finished"
	EntrantStatusDNF        EntrantStatus = "dnf" // Did Not Finish
	EntrantStatusDQ         EntrantStatus = "dq"  // Disqualified
)