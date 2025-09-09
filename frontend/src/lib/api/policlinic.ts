export async function getPoliclinics(): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/policlinic`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}
