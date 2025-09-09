export interface LabsType {
  id: string
  name: number
  management: number
  referral_fee: number
  officer_fee: number
  total: number
}

export interface LabTypeRequest {
  care_number: string
  lab_id: string
  officer_id: string
  referral_id: string
  date: string
  data: LabsType[]
}

export interface LabolatoriumCreate {
  id: string
  name: string
  referral_fee: number
  officer_fee: number
  management: number
}

export interface LabolatoriumEdit {
  id: string
  data: LabolatoriumCreate
}

export interface LabolatoriumDatas {
  id: string
  name: string
  referral_fee: number
  officer_fee: number
  management: number
  total: number
}

export interface LabolatoriumTemplateCreate {
  id: string
  value: LabolatoriumTemplate
}

export interface LabolatoriumTemplate {
  template_id: string
  id: string
  name: string
  unit: string
  normal_value: string
}

export interface LabolatoriumRequestData {
  lab_id: string;
  care_number: string;
  patient_name: string;
  referral_name: string;
  officer_id: string;
  request_date: string;
  sample: string;
  validate: boolean;
  validate_date: string;
}

export interface LabUpdateSample {
  status: boolean;
  value: string;
}

export interface LabolatoriumDetailData {
  id: string
  name: string
  template: LabolatoriumTemplate[]
}

export interface LabolatoriumValue {
  date: string
  value: LabReq[]
}

export interface LabReq {
  template_id: string
  id: string
  value: string
}
