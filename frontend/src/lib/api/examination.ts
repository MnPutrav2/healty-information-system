import type { ExaminationReq } from "@/types/examination"

export async function getExaminationData(careNumber: string | null): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/examination?care-number=${careNumber}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function addExamination(data: ExaminationReq): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/examination`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}

export async function deleteExaminationData(id: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/examination?id=${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}
