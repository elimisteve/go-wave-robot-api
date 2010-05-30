package waveapi

const (
	E_WaveletBlipCreated         = "WAVELET_BLIP_CREATED"
	E_WaveletBlipRemoved         = "WAVELET_BLIP_REMOVED"
	E_WaveletParticipantsChanged = "WAVELET_PARTICIPANTS_CHANGED"
	E_WaveletSelfAdded           = "WAVELET_SELF_ADDED"
	E_WaveletSelfRemoved         = "WAVELET_SELF_REMOVED"
	E_WaveletTitleChanged        = "WAVELET_TITLE_CHANGED"
	E_BlipContributorsChanged    = "BLIP_CONTRIBUTORS_CHANGED"
	E_BlipSubmitted              = "BLIP_SUBMITTED"
	E_DocumentChanged            = "DOCUMENT_CHANGED"
	E_FormButtonClicked          = "FORM_BUTTON_CLICKED"
	E_GadgetStateChanged         = "GADGET_STATE_CHANGED"
	E_AnnotatedTextChanged       = "ANNOTATED_TEXT_CHANGED"
	E_OperationError             = "OPERATION_ERROR"
	E_WaveletCreated             = "WAVELET_CREATED"
	E_WaveletFetched             = "WAVELET_FETCHED"
)

type Event struct {
	ModifiedBy string
	Timestamp  int64
	Type       string
	Properties map[string]interface{}
}
