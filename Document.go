package idenfy

type Document string

const (
	ID_CARD                Document = "ID_CARD"
	PASSPORT               Document = "PASSPORT"
	RESIDENCE_PERMIT       Document = "RESIDENCE_PERMIT"
	DRIVER_LICENSE         Document = "DRIVER_LICENSE"
	AADHAAR                Document = "AADHAAR"
	PAN_CARD               Document = "PAN_CARD"
	VISA                   Document = "VISA"
	BORDER_CROSSING        Document = "BORDER_CROSSING"
	ASYLUM                 Document = "ASYLUM"
	NATIONAL_PASSPORT      Document = "NATIONAL_PASSPORT"
	INTERNATIONAL_PASSPORT Document = "INTERNATIONAL_PASSPORT"
	VOTER_CARD             Document = "VOTER_CARD"
	OLD_ID_CARD            Document = "OLD_ID_CARD"
	TRAVEL_CARD            Document = "TRAVEL_CARD"
	PHOTO_CARD             Document = "PHOTO_CARD"
	MILITARY_CARD          Document = "MILITARY_CARD"
	PROOF_OF_AGE_CARD      Document = "PROOF_OF_AGE_CARD"
	DIPLOMATIC_ID          Document = "DIPLOMATIC_ID"
)

var validDocuments = []Document{ID_CARD, PASSPORT, RESIDENCE_PERMIT, DRIVER_LICENSE, AADHAAR, PAN_CARD, VISA,
	BORDER_CROSSING, ASYLUM, NATIONAL_PASSPORT, INTERNATIONAL_PASSPORT, VOTER_CARD, OLD_ID_CARD, TRAVEL_CARD,
	PHOTO_CARD, MILITARY_CARD, PROOF_OF_AGE_CARD, DIPLOMATIC_ID}
