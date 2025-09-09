interface AmbulatoryCare {
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

export function resetForm(form: AmbulatoryCare, officer: string) {
  form.care_number = ''
  form.medical_record = ''
  form.date = ''
  form.body_temperature = 0
  form.tension = ''
  form.pulse = 0
  form.respiration = 0
  form.height = 0
  form.weight = 0
  form.spo2 = 0
  form.gcs = 0
  form.awareness = ''
  form.complaint = ''
  form.examination = ''
  form.allergy = ''
  form.followup = ''
  form.assessment = ''
  form.instructions = ''
  form.evaluation = ''
  form.officer = officer
}
