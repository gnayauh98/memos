import { RequestCode } from "."

export async function verify(accessToken: string) {

    try {
        const response = await fetch("/api/user/verify", {
            method: 'POST',
            headers: {
                'authorization': `bearer ${accessToken}`
            }
        })
        if (response.ok) {
            return true
        }
    } catch {

    }

    return { name: 'signin' }
}

export async function signin(email: string, password: string) {
    try {
        const response = await fetch("/api/user/signin", {
            method: 'POST',
            headers: {
                'content-type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            })
        })

        if (!response.ok) {
            return Promise.resolve({
                code: RequestCode.REQUEST_ERROR,
                message: response.statusText
            })
        }

        const json = await response.json()
        return Promise.resolve({
            ...json,
            code: RequestCode.REQUEST_SUCCESS
        })
    } catch (err: any) {
        return Promise.resolve({
            code: RequestCode.PROGRESS_ERROR,
            message: err.toString()
        })
    }
}