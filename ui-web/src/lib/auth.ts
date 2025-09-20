import {authClient, callRPC} from "./api.ts";

export const login = async (username: string, password: string) => {
    return await callRPC(() => authClient.login({
        username,
        password,
    }))
}

export const register = async (username: string, password: string, passwordVerify: string) => {
    return await callRPC(() => authClient.register({
        username,
        password,
        passwordVerify,
    }))
}

export const guestLogin = async () => {
    return await callRPC(() => authClient.guestLogin({}))
}

export const logout = async () => {
    return await callRPC(() => authClient.logout({}))
}

export const getUserInfo = async () => {
    const {val} = await callRPC(() => authClient.test({}))
    return val?.username ?? ""
}
