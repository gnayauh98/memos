import { RequestCode } from "."

export async function createResource(file: File) {
    const formData = new FormData()

    formData.append("type", "image")
    formData.append("files", file)

    const response = await fetch("/api/resource", {
        method: 'POST',
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