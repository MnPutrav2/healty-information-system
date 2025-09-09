export interface RegistrationData {
  care_number: string;
  register_number: string;
  register_date: string;
  medical_record: string;
  name: string;
  gender: string;
  payment_method: string;
  policlinic_id: string;
  policlinic_name: string;
  doctor_id: string;
  doctor_name: string;
  status: string;
}

export interface RegisterData {
  care_number: string;
  register_number: string;
  register_date:string;
  medical_record: string;
  payment_method:string;
  policlinic: string;
  doctor: string;
}

export interface Address {
  id: number;
  name: string;
}
