export async function getDoctors(id: string | null): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/doctor?specialis=${id}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}

export async function getDoctorsAll(): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/doctor`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}
