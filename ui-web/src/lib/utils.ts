
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
