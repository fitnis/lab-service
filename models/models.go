package models

type Sample struct {
	SampleID  string `json:"sampleId"`
	PatientID string `json:"patientId"`
	Type      string `json:"type"`
}

type SampleEvaluation struct {
	SampleID string `json:"sampleId"`
	Result   string `json:"result"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
