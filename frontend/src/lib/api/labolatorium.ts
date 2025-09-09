import type { LabolatoriumCreate, LabolatoriumEdit, LabolatoriumValue, LabolatoriumTemplateCreate, LabTypeRequest, LabUpdateSample } from "@/types/labolatorium"

export async function getLabolatoriumData(limit: number, search: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/data?limit=${limit}&${search}`, {
        method: "GET",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function createLabolatoriumData(data: LabolatoriumCreate): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/data`, {
        method: "POST",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(data)
    })

    return response
}

export async function deleteLabolatoriumData(id: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/data?id=${id}`, {
        method: "DELETE",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function updateLabolatoriumData(id: string, data: LabolatoriumCreate): Promise<Response> {
    const req: LabolatoriumEdit = {
        id: id,
        data
    }

    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/data`, {
        method: "PUT",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(req)
    })

    return response
}

export async function createLabolatoriumTemplateData(data: LabolatoriumTemplateCreate): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/template`, {
        method: "POST",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(data)
    })

    return response
}

export async function getLabolatoriumTemplate(id: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/template?id=${id}`, {
        method: "GET",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function createRequestLabolatorium(data: LabTypeRequest): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/request`, {
        method: "POST",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(data)
    })

    return response
}

export async function deleteLabolatoriumTemplate(id: string, unit: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/template?id=${id}&unit=${unit}`, {
        method: "DELETE",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function getLabolatoriumNumber(date: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/common/current-laboratorium-number?date=${date}`, {
        method: "GET",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function getLabolatoriumRequest(date1: string, date2: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/request?date1=${date1}&date2=${date2}`, {
        method: "GET",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
    })

    return response
}

export async function updateLabolatoriumSample(data: LabUpdateSample, id: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/request?id=${id}`, {
        method: "PUT",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(data)
    })

    return response
}

export async function getLabolatoriumDetailData(id: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/detail?id=${id}`, {
        method: "GET",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        }
    })

    return response
}

export async function updateLabolatoriumValue(data: LabolatoriumValue, id: string): Promise<Response> {
    const response = await fetch(`${import.meta.env.VITE_API_PATH}/laboratorium/value?lab_id=${id}`, {
        method: "PUT",
        headers: {
        "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(data)
    })

    return response
}
