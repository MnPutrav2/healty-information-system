import type { SearchLimit } from "@/types/response";
import { formatDate } from "../formatDate";

interface RegisterData {
  care_number: string;
  register_number: string;
  register_date:string;
  medical_record: string;
  payment_method:string;
  policlinic: string;
  doctor: string;
}

export async function createRegister(register: RegisterData): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/registration`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(register)
  })

  return response
}

export async function getRegisters(date1: string, date2: string, param: SearchLimit): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/registration?date1=${date1}&date2=${date2}&limit=${param.limit}&search=${param.search}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function deleteRegister(careNum: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/registration?care-number=${careNum}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function getCurrentRegisterNumber(date: string, policlinic: string): Promise<Response> {
  const d = new Date(date)

  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/current-register-number?date=${formatDate(d)}&policlinic=${policlinic}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function getCurrentCareNumber(date: string): Promise<Response> {
  const d = new Date(date)

  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/current-care-number?date=${formatDate(d)}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function updateStatus(id: string, status: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/registration/update-status?care-number=${id}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify({
      status: status
    })
  })

  return response
}
