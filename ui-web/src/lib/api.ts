import {createConnectTransport} from "@connectrpc/connect-web";
import {ConnectError, createClient} from "@connectrpc/connect";
import {AuthService} from "./generated/auth/v1/auth_pb.ts";
import {LobbyService} from "./generated/lobby/v1/lobby_pb.ts";

export function getBaseUrl() {
    if (import.meta.env.DEV) {
        const devUrl = "http://localhost:11300"
        console.log(`Application is running in Debug mode using ${devUrl}`);
        return devUrl
    } else {
        const url = typeof window !== 'undefined' ?
            window.location.protocol + "//" + window.location.host.toString() :
            "/"

        console.log(`Application is running in Production mode using ${url}`);
        return url
    }
}

const transport = createConnectTransport({
    baseUrl: getBaseUrl(),
});

export const authClient = createClient(AuthService, transport);
export const lobbyClient = createClient(LobbyService, transport);

export async function callRPC<T>(exec: () => Promise<T>): Promise<{ val: T | null; err: string; }> {
    try {
        const val = await exec()
        return {val, err: ""}
    } catch (error: unknown) {
        if (error instanceof ConnectError) {
            console.error(`Error: ${error.message}`);
            return {val: null, err: `${error.rawMessage}`};
        }

        return {val: null, err: `Unknown error while calling api: ${(error as Error).toString()}`};
    }
}
