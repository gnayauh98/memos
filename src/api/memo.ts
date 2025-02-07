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

async function wrapperFetch(
    input: RequestInfo | URL,
    { withToken = true, ...init }: RequestInit & { withToken?: boolean } = {
        withToken: true
    }) {
    let headers = (init?.headers as Record<string, string>) || {}

    if (withToken) {
        const accessToken = localStorage.getItem("access-token")
        if (accessToken === null) {
            return Promise.reject("没有登录")
        }
        headers["authorization"] =
            `bearer ${accessToken}`
    }

    return fetch(input, {
        ...init,
        headers
    })
}

export interface MemoQueryFilter {
    pageNo?: number
    pageSize?: number
    tags?: string[]
}

export async function queryMemosByFilter({ pageNo = 1, pageSize = 10, tags = [] }: MemoQueryFilter) {

    const response = await wrapperFetch('/api/memo/filter', {
        method: 'POST',
        headers: {
            'content-type': 'application/json',
        },
        body: JSON.stringify({
            pageNo,
            pageSize,
            tags,
        })
    })

    // console.log(response)

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

export async function queryTags() {

    const accessToken = localStorage.getItem("access-token") ?? ""

    const response = await fetch('/api/memo/tags', {
        method: 'POST',
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