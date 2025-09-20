import {showError} from "./utils.ts";
import type {AnimDir, GameMessage} from "./models.ts";
import {GameScene} from "./game.ts";
import {getBaseUrl} from "../api.ts";

let ws: WebSocket;
let resolvePromise: ((value: string | PromiseLike<string>) => void) | undefined;
// let _rejectPromise: ((reason?: any) => void | undefined) | undefined;
let gameStatePromise: Promise<string>

let prevGameState: any = {}

let gameScene: GameScene | null = null;

export function setGame(state: GameScene) {
    gameScene = state;
}

export function getSpriteID() {
    return (prevGameState.spriteId ?? "") as string
}

export function getUsername() {
    return (prevGameState.username ?? "") as string
}

let messageHandlers: { [key: string]: (json: any) => void } = {
    "active": handleNewPlayerJoin,
    "state": handleGameStateMessage,
    "dis": handleDisconnect,
    "pos": handlePosMessage,
    "pel": handlePellet,
    "pow": handlePowerPelletStart,
    "powend": handlePowerPelletEnd,
    "kill": handlePlayerKilled,
    "gameover": handleGameOver
}

export function connectToWebSocket() {
    let wssProtocol = `${window.location.protocol === 'https:' ? 'wss://' : 'ws://'}`

    const queryString = window.location.search;
    const params = new URLSearchParams(queryString);
    const lobbyId = params.get('lobby');
    if (lobbyId === null) {
        throw new Error("lobbyId must be provided");
    }

    const base = new URL(getBaseUrl())
    const url = wssProtocol + base.host + `/api/game?lobby=${lobbyId}`;
    ws = new WebSocket(url);

    ws.onmessage = (ev) => handleMessage(ev);
    ws.onerror = (ev) => handleError(ev);

    gameStatePromise = new Promise<string>((resolve, _) => {
        resolvePromise = resolve;
        // _rejectPromise = reject;
    });
}

export async function waitForGameState() {
    await gameStatePromise;
}


export function getGameState() {
    return prevGameState;
}

function handleMessage(msg: MessageEvent): void {
    if (!msg.data) {
        console.log("No data received")
        return
    }

    try {
        const json = JSON.parse(msg.data);
        const mType = json["type"] as string

        const handler = messageHandlers[mType]
        if (!handler) {
            console.warn(`No handler found: ${mType}`);
            return;
        }

        if (Object.keys(prevGameState).length === 0 && mType !== "state") {
            // state has not been received
            // game has not started, ignore all messages
            return;
        }

        handler(json)
    } catch (e: any) {
        console.log("unable to handle message")
        console.log(e)
        console.log(msg.data)
        // showError(e)
    }
}

function handleError(ev: Event): void {
    console.log('Error: ', ev.type);
    showError(ev.type);
}

function handleGameStateMessage(msg: any) {
    prevGameState = msg;
    console.log(prevGameState)
    if (resolvePromise) {
        resolvePromise("")
    }
}

function setPlayerJoinedUsername(spriteId: string, username: string) {
    console.log(`New player joined ${spriteId}: ${username}`);
    gameScene?.allSprites[spriteId]?.userNameText!.setText(username)
}

// actual join with player info
function handleNewPlayerJoin(json: any) {
    if (!json.spriteType) {
        throw new Error("No sprite id found")
    }

    const spriteId = json.spriteType;
    setPlayerJoinedUsername(spriteId, json.user);
}

function handlePosMessage(json: any) {
    let spriteId = json.spriteType
    let x = json.x as number;
    let y = json.y as number;
    let anim = json.dir as string;

    // update player
    gameScene?.allSprites[spriteId]?.playerInfo!.setPosition(x, y);
    // update other player username text
    gameScene?.setUserNameTextPos(spriteId)

    try {
        gameScene?.setSpriteAnim(spriteId, anim)
    } catch (e) {
        // if invalid anim default to neutral image
        let defaultAnim = gameScene?.allSprites[spriteId]!.defaultAnim!
        gameScene?.setSpriteAnim(spriteId, defaultAnim)
        console.warn(e)
    }
}


function handleDisconnect(json: any) {
    console.log('disconnect')
    console.log(json)

    let spriteId = json.spriteType;

    gameScene?.allSprites[spriteId]!.userNameText!.setText('')
}


function handlePellet(json: any) {
    const x = json.x as number
    const y = json.y as number
    console.log(`Pellet eaten at x:${x}, y:${y}`)

    gameScene?.pelletLayer.removeTileAt(x, y)
    console.log(gameScene?.pelletLayer.tilesDrawn)

    if (gameScene?.pelletLayer.tilesDrawn === 0) {
        gameScene!.gameOver = true
        gameScene!.gameOverText.setText("Pacman wins")
    }
}


function handlePowerPelletStart(json: any) {
    const x = json.x as number
    const y = json.y as number
    console.log(`Power eaten at x:${x}, y:${y}`)

    gameScene?.powerLayer.removeTileAt(x, y)

    // give pacman power up
    gameScene?.allSprites['pacman']!.playerInfo!.setTint(0xff0000);
    gameScene!.allSprites['pacman']!.movementSpeed = -160
    // gameScene!.allSprites['pacman']!.movementSpeed = -300
}


function handlePowerPelletEnd(_json: any) {
    console.log(`power up ended`)

    // give pacman power up
    gameScene?.allSprites['pacman']!.playerInfo!.setTint(0xffffff);
    gameScene!.allSprites['pacman']!.movementSpeed = -200
    // gameScene!.allSprites['pacman']!.movementSpeed = -300
}


function handlePlayerKilled(json: any) {
    let spriteId = json.spriteId;
    console.log(`spriteId ${spriteId} killed`)
    gameScene?.allSprites[spriteId]!.playerInfo!.destroy()
    gameScene?.allSprites[spriteId]!.userNameText!.setText("") // username empty
}


function handleGameOver(msg: any) {
    console.log(`game over: ${msg.reason}`)
    gameScene!.gameOver = true
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
// send functions

export function sendPelletMessage(x: number, y: number) {
    console.log('eating pellet')
    sendWsMessage('pel', {x: x, y: y})
}


export function sendPowerUpMessage(x: number, y: number) {
    console.log('eating power')
    sendWsMessage('pow', {x: x, y: y})
}


export function sendPacmanGhostMessage(ghostSpriteId: string) {
    console.log('ghost and pacman collided')
    // tell all clients pacman is dead
    sendWsMessage('kill', {id: ghostSpriteId})
}


export function sendPosMessage(x: number, y: number, dir: AnimDir) {
    // console.log('position')
    sendWsMessage('pos', {x: x, y: y, dir: dir})
}


export function sendWsMessage(messageType: GameMessage, data: any) {
    data.type = messageType
    data.secretToken = prevGameState.secretToken
    const mess = JSON.stringify(data)
    ws.send(mess) // do not use json stringify here adds latency when sending messages
}
