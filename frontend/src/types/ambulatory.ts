export interface AmbulatoryCare {
  care_number: string
  medical_record: string
  name: string
  date: string
  body_temperature: number
  tension: string
  pulse: number
  respiration: number
  height: number
  weight: number
  spo2: number
  gcs: number
  awareness: string
  complaint: string
  examination: string
  allergy: string
  followup: string
  assessment: string
  instructions: string
  evaluation: string
  officer: string
  officer_name: string
}

export interface AmbulatoryCareRequest {
  care_number: string
  medical_record: string
  date: string
  body_temperature: number
  tension: string
  pulse: number
  respiration: number
  height: number
  weight: number
  spo2: number
  gcs: number
  awareness: string
  complaint: string
  examination: string
  allergy: string
  followup: string
  assessment: string
  instructions: string
  evaluation: string
  officer: string
}

export interface AmbulatoryCareRequestUpdate {
  care_number: string;
  date: string;
  data: AmbulatoryCareRequest
}
