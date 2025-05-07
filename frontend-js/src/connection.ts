import {getBaseUrl, showError} from "./utils.ts";
import {AnimDir, GameMessage} from "./models.ts";
import {GameScene} from "./game.ts";

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
    "join": handleNewPlayerJoin,
    "state": handleGameStateMessage,
    "dis": handleDisconnect,
    "pos": handlePosMessage,
    "pel": handlePellet,
    "pow": handlePowerPelletStart,
    "powend": handlePowerPelletEnd,
    "kill": handlePlayerKilled
}

export function connectToWebSocket() {
    let wssProtocol = `${window.location.protocol === 'https:' ? 'wss://' : 'ws://'}`

    const queryString = window.location.search;
    const params = new URLSearchParams(queryString);
    const lobbyId = params.get('lobby');
    if (lobbyId === null) {
        throw new Error("lobbyId must be provided");
    }

    const url = wssProtocol + getBaseUrl() + `/api/game?lobby=${lobbyId}`;
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

export function handleMessage(msg: MessageEvent): void {
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

        handler(json)
    } catch (e: any) {
        console.log("unable to handle message")
        console.log(e)
        console.log(msg.data)
        showError(e)
    }
}

export function handleError(ev: Event): void {
    console.log('Error: ', ev.type);
    showError(ev.type);
}

export function handleGameStateMessage(msg: any) {
    prevGameState = msg;
    console.log(prevGameState)
    if (resolvePromise) {
        resolvePromise("")
    }
}

// actual join with player info

export function handleNewPlayerJoin(json: any) {
    if (!json.spriteType) {
        throw new Error("No sprite id found")
    }

    const spriteId = json.spriteId;
    gameScene?.allSprites[spriteId]?.userNameText!.setText(json.username)
}


export function handlePosMessage(json: any) {
    let spriteId = json.spriteType
    let x = json.x as number;
    let y = json.y as number;
    let anim = json.dir;

    if (spriteId == gameScene?.controllingSprite?.playerInfo!.getData(spriteId)) {
        return // skip if self update
    }

    // update player
    gameScene?.allSprites[spriteId]?.playerInfo!.setPosition(x, y);

    // update other player username text
    gameScene?.setUserNameTextPos(spriteId)

    let baseAnimKey = gameScene?.allSprites[spriteId]!.animBase!
    try {
        gameScene?.allSprites[spriteId].playerInfo!.anims.play(baseAnimKey + anim, true)
    } catch (e) {
        let defaultAnim = gameScene?.allSprites[spriteId]!.defaultAnim!
        // if invalid anim default to neutral image
        gameScene?.allSprites[spriteId]!.playerInfo!.anims.play(baseAnimKey + defaultAnim)
        console.warn(e)
    }
}


export function handleDisconnect(json: any) {
    console.log('disconnect')
    console.log(json)

    let spriteId = json.spriteId;

    gameScene?.allSprites[spriteId]!.userNameText!.setText('')
}


export function handlePellet(json: any) {
    const x = json.x as number
    const y = json.y as number
    console.log(`Pellet eaten at x:${x}, y:${y}`)

    gameScene?.pelletLayer.removeTileAt(x, y)
}


export function handlePowerPelletStart(json: any) {
    const x = json.x as number
    const y = json.y as number
    console.log(`Power eaten at x:${x}, y:${y}`)

    gameScene?.powerLayer.removeTileAt(x, y)

    // give pacman power up
    gameScene?.allSprites['pacman']!.playerInfo!.setTint(0xff0000);
    gameScene!.allSprites['pacman']!.movementSpeed = -160

}


export function handlePowerPelletEnd(_json: any) {
    console.log(`power up ended`)

    // give pacman power up
    gameScene?.allSprites['pacman']!.playerInfo!.setTint(0xffffff);
    gameScene!.allSprites['pacman']!.movementSpeed = -200
}


export function handlePlayerKilled(json: any) {
    const spriteId = json.spriteId;
    console.log(`spriteId ${spriteId} killed`)
    gameScene?.allSprites[spriteId]!.playerInfo!.destroy()
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
// send functions

export function handlePacmanDead() {
    // if (playerSprites['pcm'].isPoweredUp === 0) {
    //     console.log('Received pacman is dead')
    //     playerSprites['pcm'].playerInfo.destroy()
    //     gameEnd = true
    //     return
    // }
    //
    // const spriteId = json.id
    // if (spriteId !== undefined) {
    //     playerSprites[spriteId].playerInfo.destroy()
    //     setUserNameText(spriteId, '')
    //
    //     if (!playerSprites['gh1'].playerInfo.active
    //         && !playerSprites['gh2'].playerInfo.active
    //         && !playerSprites['gh3'].playerInfo.active
    //     ) {   // all ghosts eaten
    //         gameEnd = true
    //     }
    //     return;
    // }
    // console.log('Json does not have id')
}


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
    ws.send(data) // do not use json stringify here adds latency when sending messages
}
