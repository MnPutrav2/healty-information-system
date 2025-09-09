export interface Patient {
  medical_record: string;
  name: string;
  gender: string;
  wedding: string;
  religion: string;
  education: string;
  birth_place: string;
  birth_date: string;
  work: string;
  address: string;
  village: number;
  district: number;
  regencie: number;
  province: number;
  nik: string;
  bpjs: string;
  phone_number: string;
  parent_name: string;
  relationship: string;
  parent_gender: string;
}

export interface Doctor {
  id: string;
  name:  string;
  specialist: boolean;
  specialist_name:  string;
}
