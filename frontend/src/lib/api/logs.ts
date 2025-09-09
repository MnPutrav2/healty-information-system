export async function logsApi(date1: string, date2: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/logs?date1=${date1}&date2=${date2}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  })

  return response
}
