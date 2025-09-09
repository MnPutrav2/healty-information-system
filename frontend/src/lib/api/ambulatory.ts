interface AmbulatoryCare {
  care_number: string;
  medical_record: string;
  date: string;
  body_temperature: number;
  tension: string;
  pulse: number;
  respiration: number;
  height: number;
  weight: number;
  spo2: number;
  gcs: number;
  awareness: string;
  complaint: string;
  examination: string;
  allergy: string;
  followup: string;
  assessment: string;
  instructions: string;
  evaluation: string;
  officer: string;
}

interface AmbulatoryCareUpdate {
  care_number: string;
  date: string;
  data: AmbulatoryCare
}

export async function createAmbulatoryCare(ambulatory: AmbulatoryCare): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/ambulatory`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(ambulatory)
  })

  return response
}

export async function getAmbulatoryCare(careNumber: string | null, date1: string, date2: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/ambulatory?care-number=${careNumber}&date1=${date1}&date2=${date2}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function deleteAmbulatoryCare(careNumber: string | null, date: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/ambulatory?care-number=${careNumber}&date=${date}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function updateAmbulatoryCare(data: AmbulatoryCareUpdate): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/ambulatory`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}
