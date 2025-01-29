import * as user from './user'

export enum RequestCode {
    REQUEST_ERROR = 1001,
    REQUEST_SUCCESS = 1000,
    PROGRESS_ERROR = 1002
}

const Api = {
    user
}

export default Api