export interface ExaminationReq {
  care_number: string
  examination: string
  doctor_id: string
  nurse_id: string
  service_type: string
  date: string
}

export interface Examination {
  id: string
  examination: string
  doctor_id: string
  nurse_id: string
  service_type: string
  date: string
}

export interface ExaminationDetail {
  id: string
  examination_id: string
  examination: string
  doctor_id: string
  doctor_name: string
  nurse_id: string
  nurse_name: string
  service_type: string
  date: string
  total: number
}
