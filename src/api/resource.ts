import { RequestCode } from "."

export async function createResource(file: File) {
    const formData = new FormData()

    formData.append("type", "image")
    formData.append("files", file)

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch("/api/resource", {
        method: 'POST',
        headers: {
            'authorization': `bearer ${accessToken}`
        },
        body: formData
    })

    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR
        }
    }

    const data = await response.json()

    return data
}

export async function getResources() {
    const accessToken = localStorage.getItem("access-token") ?? ""
    const response = await fetch("/api/resource/all", {
        method: 'POST',
        headers: {
            'authorization': `bearer ${accessToken}`
        },
    })

    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR
        }
    }

    const data = await response.json()

    return data
}