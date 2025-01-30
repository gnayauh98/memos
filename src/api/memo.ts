import { RequestCode } from "."

export async function createMemo(content: string) {

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch('/api/memo', {
        method: 'PUT',
        headers: {
            'content-type': 'application/json',
            'authorization': `bearer ${accessToken}`
        },
        body: JSON.stringify({
            content: content
        })
    })

    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR,
        }
    }
    const json = await response.json()

    return {
        code: RequestCode.REQUEST_SUCCESS,
        data: json
    }
}

export async function updateMemo(content: string, id: string) {

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch('/api/memo', {
        method: 'PATCH',
        headers: {
            'content-type': 'application/json',
            'authorization': `bearer ${accessToken}`
        },
        body: JSON.stringify({
            content: content,
            id
        })
    })

    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR,
        }
    }
    const json = await response.json()

    return {
        code: RequestCode.REQUEST_SUCCESS,
        data: json
    }
}

export async function getMemoById(id: string) {

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch(`/api/memo/${id}`, {
        method: "POST",
        headers: {
            'content-type': 'application/json',
            'authorization': `bearer ${accessToken}`
        }
    })


    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR,
        }
    }
    const json = await response.json()

    return {
        code: RequestCode.REQUEST_SUCCESS,
        data: json
    }
}

export async function queryMemos(pageNo: number = 1) {

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch('/api/memo/all', {
        method: 'POST',
        headers: {
            'content-type': 'application/json',
            'authorization': `bearer ${accessToken}`
        },
        body: JSON.stringify({
            pageNo,
            pageSize: 10
        })
    })

    if (!response.ok) {
        return {
            code: RequestCode.REQUEST_ERROR,
        }
    }
    const json = await response.json()

    return {
        code: RequestCode.REQUEST_SUCCESS,
        data: json
    }
}