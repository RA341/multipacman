document.addEventListener('DOMContentLoaded', (event) => {
    const userId = document.getElementById('userToken').value;
    const lobbyId = document.getElementById('lobbyToken').value;

    if (!lobbyId || !userId || lobbyId === '{LobbyToken}' || userId === '{UserToken}') {
        console.log('User token or the lobby token are invalid')
        showError()
        return
    }

    let wssProtocol = `${window.location.protocol === 'https:' ? 'wss://' : 'ws://'}`
    const url = wssProtocol + window.location.host + `/ws/game?user=${userId}&lobby=${lobbyId}`;
    let ws = new WebSocket(url);


    let allPlayers = {};

    let prevGameState;

    ws.onerror = function (event) {
        console.log('WebSocket Error: ', event);
        showError()
    }

    ws.onmessage = function (msg) {
        if (!msg.data) {
            console.log("No data received")
            return
        }
        try {
            const json = JSON.parse(msg.data);
            if (json.redirect){
                // if lobby is full redirect to home page
                window.location = json.redirect
                return;
            }
            if (json.type === 'join') {
                handleNewPlayerJoin(json)
            } else if (json.type === 'pos') {
                handlePosMessage(json)
            } else if (json.type === 'dis') {
                handleDisconnect(json)
            } else if (json.type === 'state') {
                handleGameStateMessage(json)
            } else if (json.type === 'pellet') {
                handlePellet(json)
            } else if (json.type === 'power') {
                handlePowerPellet(json)
            } else if (json.type === 'pacded') {
                handlePacmanDead(json)
            } else {
                console.log("Unknown info type: " + json.type)
            }
        } catch (e) {
            console.log("whoa json failed")
            console.log(e)
        }
    };

    function handleGameStateMessage(msg) {
        prevGameState = msg
        // prevGameState = JSON.parse(msg);
        console.log(prevGameState.ghostsEaten)
        // start once all state has been received
        const game = new Phaser.Game(config);
    }

    // actual join with player info
    function handleNewPlayerJoin(json) {
        if (!json.spriteType) {
            throw new Error("No sprite id found")
        }
        allPlayers[json.playerid] = json
        playerSprites[json.spriteType].username = json.user
    }

    function handlePosMessage(json) {
        let spriteId = json.spriteType
        // update player
        allPlayers[json.id] = json

        if (playerSprites[spriteId].playerInfo === null) {
            console.log('Client not ready')
            return
        }

        // update other player sprites
        playerSprites[spriteId].playerInfo.x = json.x
        playerSprites[spriteId].playerInfo.y = json.y
        // update other player username text
        setUserNameTextPos(spriteId)

        try {
            playerSprites[spriteId].playerInfo.anims.play(json.spriteAnim, true)
        } catch (e) {
            // if invalid anim default to neutral image
            playerSprites[spriteId][0].anims.play()
        }
    }

    function handleDisconnect(json) {
        console.log('disconnect')
        console.log(json.playerid)
        delete allPlayers[json.playerid]
        setUserNameText(json.spriteType, '')
    }

    function handlePellet(json) {
        const x = json.x
        const y = json.y
        console.log(`Pellet eaten at x:${x}, y:${y}`)
        pelletLayer.removeTileAt(x, y)
        if (pelletLayer.tilesDrawn === 1) {
            // all pellets eaten
            gameEnd = true
        }
    }

    function handlePowerPellet(json) {
        const x = json.x
        const y = json.y
        console.log(`Pellet eaten at x:${x}, y:${y}`)
        powerLayer.removeTileAt(x, y)
        // give pacman power up
        playerSprites['pcm'].isPoweredUp = 800
    }

    // pacman dead
    function handlePacmanDead(json) {
        if (playerSprites['pcm'].isPoweredUp === 0) {
            console.log('Received pacman is dead')
            playerSprites['pcm'].playerInfo.destroy()
            gameEnd = true
            return
        }

        const spriteId = json.id
        if (spriteId !== undefined) {
            playerSprites[spriteId].playerInfo.destroy()
            setUserNameText(spriteId, '')

            if (!playerSprites['gh1'].playerInfo.active
                && !playerSprites['gh2'].playerInfo.active
                && !playerSprites['gh3'].playerInfo.active
            ) {   // all ghosts eaten
                gameEnd = true
            }
            return;
        }
        console.log('Json does not have id')
    }


    function pelletCallBack(pacman, pellet) {
        sendWsMessage('pellet', {user: userId, x: pellet.x, y: pellet.y})
        // console.log(pelletLayer.tilesDrawn)
    }

    function powerUpCallBack(pacman, powerUp) {
        sendWsMessage('power', {user: userId, x: powerUp.x, y: powerUp.y})
        // console.log(pelletLayer.tilesDrawn)
    }

    function killPacmanCallBack(ghost, pacman) {
        // tell all clients pacman is dead
        sendWsMessage('pacded', {user: userId, id: ghost.id})
    }

    function sendPosMessage(player) {
        sendWsMessage('pos', player)
    }

    function sendWsMessage(messageType, data) {
        data.type = messageType
        ws.send(JSON.stringify(data))
    }

////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main game

// helper variables
    let cursors;
    let gameEnd = false
    let gameOverText
    let spectatingText
    let isRedirecting = true
// mapped sprite id with:
// playerInfo
// default anim
// anim base so that we can add direction
// start pos : {x, y}
    /**
     *
     * @type {{gh2: {playerInfo: null, animBase: string, userNameTextPos: null, movementSpeed: number, defaultAnim: string, startPos: number[]}, gh1: {playerInfo: null, animBase: string, userNameTextPos: null, movementSpeed: number, defaultAnim: string, startPos: number[]}, gh3: {playerInfo: null, animBase: string, userNameTextPos: null, movementSpeed: number, defaultAnim: string, startPos: number[]}, pcm: {isPoweredUp: number, playerInfo: null, animBase: string, userNameTextPos: null, movementSpeed: number, defaultAnim: string, startPos: number[]}}}
     */
    const playerSprites = {
        'pcm': {
            playerInfo: null,
            defaultAnim: 'pacmanDefault',
            animBase: 'pacman',
            startPos: [1050, 925],
            userNameTextPos: null,
            isPoweredUp: 0,
            movementSpeed: -200

        },
        'gh1': {
            playerInfo: null,
            defaultAnim: 'ghostRedNorth',
            animBase: 'ghostRed',
            startPos: [670.1666666666666, 424.49999999999994],
            userNameTextPos: null,
            movementSpeed: -160

        },
        'gh2': {
            playerInfo: null,
            defaultAnim: 'ghostBlueNorth',
            animBase: 'ghostBlue',
            startPos: [723.4999999999992, 424.49999999999994],
            userNameTextPos: null,
            movementSpeed: -160
        },
        'gh3': {
            playerInfo: null,
            defaultAnim: 'ghostPinkNorth',
            animBase: 'ghostPink',
            startPos: [776.8333333333, 424.49999999999994],
            userNameTextPos: null,
            movementSpeed: -160

        },
    }

    const userNameOffSetX = 15
    const userNameOffSetY = 35
// movement

    /**
     *  @type {?Phaser.Tilemaps.TilemapLayer} mapLayer
     */
    let mapLayer
    /**
     *  @type {?Phaser.Tilemaps.TilemapLayer} mapLayer
     */
    let pelletLayer
    /**
     *  @type {?Phaser.Tilemaps.TilemapLayer} mapLayer
     */
    let powerLayer

    const config = {
        type: Phaser.AUTO,
        width: window.innerWidth,
        height: window.innerHeight,
        backgroundColor: 0x000000,
        physics: {
            default: 'arcade',
            arcade: {
                gravity: 0,
            }
        },
        scale: {
            mode: Phaser.Scale.FIT,
            autoCenter: Phaser.Scale.CENTER_BOTH,
            parent: 'thegame'
        },
        scene: {
            preload: preload,
            create: create,
            update: update
        }
    };

/////////////////////////////////////////////
// main phaser stuff
    function preload() {
        const assetsPath = "static/game/assets"
        // sprite sheets
        this.load.spritesheet("pacman", `${assetsPath}/pacmanSpriteSheet.png`, {
            frameWidth: 50,
            frameHeight: 50,
        });
        this.load.spritesheet("ghosts", `${assetsPath}/ghosts.png`, {
            frameWidth: 50,
            frameHeight: 50,
        });
        this.load.spritesheet("fruits", `${assetsPath}/fruits.png`, {
            frameWidth: 50,
            frameHeight: 50,
        });
        this.load.image('secondTile', `${assetsPath}/secondTile.png`)
        this.load.image('forthTile', `${assetsPath}/forthTile.png`)
        this.load.image('centrepoint', `${assetsPath}/centrepoint.png`)
        this.load.image('power-up', `${assetsPath}/powercent.png`)
        this.load.tilemapTiledJSON('map', `${assetsPath}/map.json`);
    }

    function create() {
        // keyboard detect
        cursors = this.input.keyboard.createCursorKeys();

        // Set Map
        createMap(this)

        // create animations
        pacmanAnimInit(this)
        ghostsAnimInit(this)

        // load in initial sprites
        for (let playerType in playerSprites) {
            const sprite = playerSprites[playerType]
            const defaultAnim = sprite.defaultAnim
            const startPos = sprite.startPos


            let tmp = this.physics.add.sprite(
                startPos[0],
                startPos[1],
                playerType === 'pcm' ? 'pacman' : 'ghosts'
            );

            tmp.anims.play(defaultAnim);
            tmp.setOrigin(0.5);
            tmp.setCollideWorldBounds(true);
            tmp.setScale(.7)
            // Set the width and height of the player sprite
            // tmp.displayWidth = 48; // Set width to 64 pixels
            // tmp.displayHeight = 48; // Set height to 64 pixels
            playerSprites[playerType].userNameTextPos = this.add.text(
                tmp.x - userNameOffSetX,
                tmp.y - userNameOffSetY, '', {
                    fontFamily: 'Arial',
                    fontSize: 12,
                    color: '#ffffff'
                });

            if (playerType === 'pcm') {
                // if pacman
                this.physics.add.overlap(tmp, pelletLayer, pelletCallBack)
                this.physics.add.collider(tmp, powerLayer, powerUpCallBack)
            } else {
                // if ghosts collide with pacman
                this.physics.add.collider(tmp, playerSprites['pcm'].playerInfo, killPacmanCallBack)
            }

            this.physics.add.collider(tmp, mapLayer)

            tmp.id = playerType
            // attach player info
            playerSprites[playerType].playerInfo = tmp
        }

        // Set up the "Game Ended" text
        const textStyle = {fontFamily: 'Arial', fontSize: 48, color: '#00ffc7'};
        gameOverText = this.add.text(700, 450, 'Game Ended', textStyle);
        gameOverText.setOrigin(0.5);
        gameOverText.visible = false; // Initially hide the text

        // Set up the "Spectating" text
        const spectatingTextStyle = {fontFamily: 'Arial', fontSize: 48, color: '#00ffc7'};
        spectatingText = this.add.text(800, 30, 'You were eaten... spectating', spectatingTextStyle);
        spectatingText.setOrigin(0.5);
        spectatingText.visible = false; // Initially hide the text


        // update map to match game state till now
        for (const pelletId of prevGameState.pelletsEaten) {
            pelletLayer.removeTileAt(pelletId[0], pelletId[1])
        }

        for (const powerUpId of prevGameState.powerUpsEaten) {
            powerLayer.removeTileAt(powerUpId[0], powerUpId[1])
        }

        for (const ghId of prevGameState.ghostsEaten) {
            playerSprites[ghId].playerInfo.destroy()
        }
    }

    function update() {
        if (gameEnd) {
            // Show the "Game Ended" text
            gameOverText.visible = true;
            if (isRedirecting) {
                // here so the game doesn't repeatedly redirect
                isRedirecting = false
                delay(2000).then(value => {
                    window.location.href = '/lobby';
                })
            }
            // Prevent further updates
            this.input.off('pointerdown'); // Remove the event listener to avoid multiple redirects
            return;
        }

        const tmp = allPlayers[userId]
        // initial empty because websocket is requesting info
        if (tmp === undefined) {
            console.log(`waiting for player ${userId}`)
            return
        }

        powerUpCheck();

        let currentPlayerSpriteType = allPlayers[userId].spriteType

        if (!playerSprites[currentPlayerSpriteType].playerInfo.active) {
            spectatingText.visible = true
            console.log('You are killed')
            return;
        }

        // anims
        const neutralAnim = playerSprites[currentPlayerSpriteType].defaultAnim
        const anim = playerSprites[currentPlayerSpriteType].animBase
        let curAnim;

        if (cursors.left.isDown) {
            // remove vertical velocity
            movePlayer(currentPlayerSpriteType, playerSprites[currentPlayerSpriteType].movementSpeed, 0);
            // set anim
            curAnim = `${anim}West`
            setSpriteAnim(playerSprites[currentPlayerSpriteType].playerInfo, curAnim)

        } else if (cursors.right.isDown) {
            // remove vertical velocity
            movePlayer(currentPlayerSpriteType, -playerSprites[currentPlayerSpriteType].movementSpeed, 0);

            curAnim = `${anim}East`
            setSpriteAnim(playerSprites[currentPlayerSpriteType].playerInfo, curAnim)
        } else if (cursors.up.isDown) {
            // remove horizontal velocity
            movePlayer(currentPlayerSpriteType, 0, playerSprites[currentPlayerSpriteType].movementSpeed);

            curAnim = `${anim}North`
            setSpriteAnim(playerSprites[currentPlayerSpriteType].playerInfo, curAnim)
        } else if (cursors.down.isDown) {
            // remove horizontal velocity
            movePlayer(currentPlayerSpriteType, 0, -playerSprites[currentPlayerSpriteType].movementSpeed);

            curAnim = `${anim}South`
            setSpriteAnim(playerSprites[currentPlayerSpriteType].playerInfo, curAnim)
        } else {
            playerSprites[currentPlayerSpriteType].playerInfo.setVelocityY(0);
            playerSprites[currentPlayerSpriteType].playerInfo.setVelocityX(0);

            curAnim = neutralAnim
            setSpriteAnim(playerSprites[currentPlayerSpriteType].playerInfo, curAnim, false)
        }

        // move username text
        setUserNameTextPos(currentPlayerSpriteType)

        // setup info to send to server
        allPlayers[userId].x = playerSprites[currentPlayerSpriteType].playerInfo.x
        allPlayers[userId].y = playerSprites[currentPlayerSpriteType].playerInfo.y
        allPlayers[userId].spriteAnim = curAnim
        sendPosMessage(allPlayers[userId])
    }


/////////////////////////////////////////////
// utility functions

    function powerUpCheck() {
        if (playerSprites['pcm'].isPoweredUp > 0 || playerSprites['pcm'].isPoweredUp > 800) {
            // time left till powerup
            playerSprites['pcm'].isPoweredUp -= 1
            playerSprites['pcm'].playerInfo.tint = 0xff0000;
            playerSprites['pcm'].movementSpeed = -160
        } else {
            playerSprites['pcm'].playerInfo.tint = 0xffffff; // Reset tint to white (no tint)
            playerSprites['pcm'].movementSpeed = -200
        }
    }

    function movePlayer(currentPlayerSpriteType, x = 0, y = 0) {
        if (playerSprites[currentPlayerSpriteType].playerInfo.active) {
            playerSprites[currentPlayerSpriteType].playerInfo.setVelocityY(y);
            playerSprites[currentPlayerSpriteType].playerInfo.setVelocityX(x);
        }
    }

    function setSpriteAnim(sprite, anim, loop = true) {
        // convert loop to bool
        sprite.anims.play(anim, loop)
    }

    function setUserNameTextPos(spriteId) {
        playerSprites[spriteId].userNameTextPos.setText(playerSprites[spriteId]?.username);
        playerSprites[spriteId]
            .userNameTextPos
            .setPosition(
                playerSprites[spriteId].playerInfo.x - userNameOffSetX,
                playerSprites[spriteId].playerInfo.y - userNameOffSetY
            );
    }

    function setUserNameText(spriteId, text) {
        if (playerSprites[spriteId].userNameTextPos === null) {
            console.log('Client is not ready yet to set username text')
            return
        }
        playerSprites[spriteId].userNameTextPos.setText(text);
    }

    function pacmanAnimInit(ctx) {
        let namesArray = [
            'pacmanEast',
            'pacmanNorth',
            'pacmanSouth',
            'pacmanWest',
        ]
        let indexArray = 0
        for (let i = 0; i < 12; i += 3) {
            ctx.anims.create({
                frames: ctx.anims.generateFrameNumbers('pacman', {start: i + 1, end: i + 3}),
                key: namesArray[indexArray],
                frameRate: 10,
                repeat: -1
            });
            indexArray++
        }

        // neutral pacman
        ctx.anims.create({
            frames: ctx.anims.generateFrameNumbers('pacman', {start: 0, end: 0}),
            key: 'pacmanDefault',
            frameRate: 1,
            repeat: -1
        });

    }

    function createMap(ctx) {
        const map = ctx.make.tilemap({key: 'map'})

        const blueTile = map.addTilesetImage('secondTile', 'secondTile')
        const redTile = map.addTilesetImage('forthTile', 'forthTile')
        const pellets = map.addTilesetImage('centrepoint', 'centrepoint')
        const powerUp = map.addTilesetImage('power-up', 'power-up')

        mapLayer = map.createLayer('map', [blueTile, redTile])
        pelletLayer = map.createLayer('pellets', pellets)
        powerLayer = map.createLayer('powerup', powerUp)

        if (mapLayer === null || pelletLayer === null || powerLayer == null) {
            throw Error('Map layer or pellet layer failed to initlize')
        }

        mapLayer.setCollisionByExclusion([-1], true, true)
        pelletLayer.setCollisionByExclusion([-1])
        powerLayer.setCollisionByExclusion([-1])
    }

    function ghostsAnimInit(ctx) {
        let namesArray = [
            'ghostRedEast',
            'ghostRedNorth',
            'ghostRedSouth',
            'ghostRedWest',
            'ghostBlueEast',
            'ghostBlueNorth',
            'ghostBlueSouth',
            'ghostBlueWest',
            'ghostPinkEast',
            'ghostPinkNorth',
            'ghostPinkSouth',
            'ghostPinkWest',
            'frigthenedAnim',
        ]
        for (let i = 0; i < namesArray.length; i++) {
            ctx.anims.create({
                frames: ctx.anims.generateFrameNumbers('ghosts', {start: i, end: i}),
                key: namesArray[i],
            })
        }
    }
})

function showError() {
    window.location = "/static/game/error.html"
}

const delay = ms => new Promise(res => setTimeout(res, ms));
