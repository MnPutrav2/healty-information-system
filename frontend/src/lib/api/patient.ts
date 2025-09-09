import type { SearchLimit } from "@/types/response";

interface PatientData {
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

export async function createPatient(patient: PatientData): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/patient`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(patient)
  })

  return response
}

export async function getPatient(param:  SearchLimit): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/patient?limit=${param.limit}&search=${param.search}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function deletePatient(mr: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/patient?mr=${mr}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function updatePatient(patient: PatientData, mr: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/patient`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify({
      medical_record: mr,
      update: patient
    })
  })

  return response
}

export async function getCurrentMedicalRecord(): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/current-medical-record`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}
