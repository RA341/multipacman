const wsErrorKey = 'websocketGameError';

export function showError(message: string) {
    try {
        sessionStorage.setItem(wsErrorKey, message);
    } catch (e) {
        console.error("Failed to set sessionStorage:", e);
    }
    // window.location.assign("/error.html")
}


export function getBaseUrl() {
    if (import.meta.env.DEV) {
        const devUrl = "localhost:11200"
        console.log(`Application is running in Debug mode using ${devUrl}`);
        return devUrl
    } else {
        const url = window.location.host
        console.log(`Application is running in Production mode using ${url}`);
        return url
    }
}

export const delay = (ms: number) => new Promise(res => setTimeout(res, ms));

