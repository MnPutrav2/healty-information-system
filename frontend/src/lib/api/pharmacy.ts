import type { DrugData, RequestBodyDrugDataUpdate, RecipeForRequest, RecipeCompoundForRequest, ValidateTypeRecipe, HandoverTypeRecipe } from "@/types/pharmacy"
import type { SearchLimit } from "@/types/response"

export async function createDrugData(drug: DrugData): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/drug-data`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(drug)
  })

  return response
}

export async function getDrugData(param: SearchLimit): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/drug-data?search=${param.search}&limit=${param.limit}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function updateDrugData(data: RequestBodyDrugDataUpdate): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/drug-data`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(data)
  })

  return response
}

export async function deleteDrugData(id: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/drug-data?drug_id=${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function getDistributor() {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/distributor`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function createRecipe(drug: RecipeForRequest): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/recipe`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(drug)
  })

  return response
}

export async function createRecipeCompound(drug: RecipeCompoundForRequest): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/recipe-compound`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(drug)
  })

  return response
}

export async function getCurrentRecipeNumber(date: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/current-recipe-number?date=${date.split("T")[0]}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function getRecipeNumber(care: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/recipe-number?care_number=${care}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function getRecipeData(date1: string, date2: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/recipe?date1=${date1.split("T")[0]}&date2=${date2.split("T")[0]}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function getDrugRecipe(recipe: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/recipe-compound?recipe=${recipe}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function deleteDrugRecipes(recipe: string, drugID: string, comname: string): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/recipe?recipe_id=${recipe}&drug_id=${drugID}&compound_name=${comname}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
  })

  return response
}

export async function validateRecipe(recipe: string, val: ValidateTypeRecipe): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/validate?recipe_id=${recipe}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(val)
  })

  return response
}

export async function handoverRecipe(recipe: string, val: HandoverTypeRecipe): Promise<Response> {
  const response = await fetch(`${import.meta.env.VITE_API_PATH}/pharmacy/handover?recipe-id=${recipe}`, {
    method: "PUT",
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    },
    body: JSON.stringify(val)
  })

  return response
}
