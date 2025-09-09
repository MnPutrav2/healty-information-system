import type { ExaminationCost, ExaminationCostRequest } from "@/types/finance"
import type { SearchLimit } from "@/types/response"

export async function getExamination(param: SearchLimit): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/finance/examination-cost?search=${param.search}&limit=${param.limit}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function createExamination(data: ExaminationCostRequest): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/finance/examination-cost`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body:  JSON.stringify(data)
  })

  return response
}

export async function deleteExamination(id: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/finance/examination-cost?id=${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function updateExamination(data: ExaminationCost, id: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/finance/examination-cost?id=${id}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}
