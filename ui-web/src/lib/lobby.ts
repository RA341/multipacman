import {callRPC, lobbyClient} from "./api.ts";
import type {Lobby} from "./generated/lobby/v1/lobby_pb.ts";
import type {Component} from "solid-js";

export const listLobbies = async () => {
    return await callRPC(() => lobbyClient.listLobbies({}))
}


export const addLobby = async (name: string) => {
    return await callRPC(() => lobbyClient.addLobby({lobbyName: name}))
}

export const deleteLobby = async (lobby: Lobby) => {
    return await callRPC(() => lobbyClient.deleteLobby({lobby}))
}

export const getRelativeTime = (dateString: string): string => {
    const now = new Date();
    const date = new Date(dateString);
    const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000);

    if (diffInSeconds < 60) {
        return 'Just now';
    } else if (diffInSeconds < 3600) {
        const minutes = Math.floor(diffInSeconds / 60);
        return `${minutes}m ago`;
    } else if (diffInSeconds < 86400) {
        const hours = Math.floor(diffInSeconds / 3600);
        return `${hours}h ago`;
    } else if (diffInSeconds < 604800) {
        const days = Math.floor(diffInSeconds / 86400);
        return `${days}d ago`;
    } else {
        return date.toLocaleDateString();
    }
};
