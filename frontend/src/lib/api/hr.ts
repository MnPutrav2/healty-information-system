import type { EmployeeData, UserReq } from "@/types/hr";
import type { SearchLimit } from "@/types/response";

export async function getEmployee(param: SearchLimit) {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/human-resource/employees?search=${param.search}&limit=${param.limit}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function addEmployee(data: EmployeeData) {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/human-resource/employees`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}

export async function deleteEmployee(id: string) {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/human-resource/employees?id=${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function updateEmployee(data: EmployeeData, id: string) {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/human-resource/employees?id=${id}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}

export async function addUser(data: UserReq) {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/human-resource/users`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}
