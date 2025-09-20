import type {AnimDir, Ghost, GhostAnimBase, Pacman} from "./models.ts";
import {
    getGameState,
    getSpriteID,
    getUsername,
    sendPacmanGhostMessage,
    sendPelletMessage,
    sendPosMessage,
    sendPowerUpMessage
} from "./connection.ts";
import {delay} from "./utils.ts";

import Phaser from 'phaser';

type CursorKeys = Phaser.Types.Input.Keyboard.CursorKeys;
type GameObjectWithBody = Phaser.Types.Physics.Arcade.GameObjectWithBody;

import TilemapLayer = Phaser.Tilemaps.TilemapLayer;
import Text = Phaser.GameObjects.Text;
import Tile = Phaser.Tilemaps.Tile;

const playerSpriteIdKey = "playerID";

type CollisionBodyType =
    Phaser.Types.Physics.Arcade.GameObjectWithBody
    | Phaser.Physics.Arcade.Body
    | Phaser.Physics.Arcade.StaticBody
    | Phaser.Tilemaps.Tile


export class GameScene extends Phaser.Scene {
    cursors!: CursorKeys

    mapLayer!: TilemapLayer
    pelletLayer!: Phaser.Tilemaps.TilemapLayer
    powerLayer!: Phaser.Tilemaps.TilemapLayer

    controllingSprite: Pacman | Ghost | undefined

    gameOverText!: Text;
    spectatingText!: Text

    gameOver = false
    isRedirecting = true

    allSprites: { [key: string]: Ghost | Pacman } = {
        "pacman": {
            playerInfo: null,
            curAnimDir: 'up',
            defaultAnim: 'default',
            animBase: 'pacman',
            startPos: [110, 220],
            userNameText: null,
            movementSpeed: -200
        },
        'gh0': {
            playerInfo: null,
            curAnimDir: 'right',
            defaultAnim: 'left',
            animBase: 'ghostred',
            startPos: [670.1666666666666, 424.49999999999994],
            userNameText: null,
            movementSpeed: -160
        },
        'gh1': {
            playerInfo: null,
            curAnimDir: 'right',
            defaultAnim: 'left',
            animBase: 'ghostblue',
            startPos: [723.4999999999992, 424.49999999999994],
            userNameText: null,
            movementSpeed: -160
        },
        'gh2': {
            playerInfo: null,
            userNameText: null,
            curAnimDir: 'right',
            defaultAnim: 'left',
            animBase: 'ghostpink',
            startPos: [776.8333333333, 424.49999999999994],
            movementSpeed: -160
        },
    }

    constructor() {
        super({key: 'GameScene'});
    }

    preload(): void {
        console.log('Preloading assets...');
        // sprite sheets
        this.loadAssets();
    }

    create(): void {
        console.log('Creating scene...');

        this.cursors = this.input.keyboard!.createCursorKeys();
        this.createMap()
        this.initPacmanAnim()
        this.initGhostsAnim()
        this.loadPlayers()

        // Set up the "Game Ended" text
        const textStyle = {fontFamily: 'Arial', fontSize: 48, color: '#00ffc7'};
        this.gameOverText = this.add.text(700, 450, 'Game Ended', textStyle);
        this.gameOverText.setOrigin(0.5);
        this.gameOverText.visible = false; // Initially hide the text

        // Set up the "Spectating" text
        const spectatingTextStyle = {fontFamily: 'Arial', fontSize: 48, color: '#00ffc7'};
        this.spectatingText = this.add.text(800, 30, 'You were eaten... spectating', spectatingTextStyle);
        this.spectatingText.setOrigin(0.5);
        this.spectatingText.visible = false; // Initially hide the text


        const spriteID = getSpriteID();
        const sprite = this.allSprites[spriteID]
        if (!sprite) {
            throw Error(`Sprite not found: ${spriteID}`);
        }
        this.controllingSprite = sprite
        this.controllingSprite.userNameText!.setText(getUsername())
        console.log(getUsername())

        console.log(`
            Assigning ${spriteID} to 
            ${this.controllingSprite!.userNameText} 
            with username ${this.controllingSprite.userNameText!.text}`
        )

        const gameState = getGameState()

        // assign username for existing players
        for (const [playerSprite, playerData] of Object.entries(gameState.activePlayers)) {
            let username = (playerData as any).username as string
            let x = (playerData as any).x as number
            let y = (playerData as any).y as number

            this.allSprites[playerSprite]!.userNameText!.setText(username as string)
            this.allSprites[playerSprite]!.playerInfo!.setPosition(x, y);
        }

        // update map to match game state till now
        for (const pellet of gameState.pelletsEaten) {
            this.pelletLayer.removeTileAt(pellet.X, pellet.Y)
        }

        for (const powerUpId of gameState.powerUpsEaten) {
            this.powerLayer.removeTileAt(powerUpId.X, powerUpId.Y)
        }

        for (const ghId of gameState.ghostsEaten) {
            console.log(`Eaten ghosts ${ghId}`)
            this.allSprites[ghId]!.playerInfo!.destroy()
        }
    }

    update(_time: number, _delta: number): void {
        if (this.gameOver) {
            // Show the "Game Ended" text
            this.gameOverText.visible = true;
            if (this.isRedirecting) {
                // here so the game doesn't repeatedly redirect
                this.isRedirecting = false
                delay(2000).then(_ => {
                    window.location.assign('/');
                })
            }

            // Prevent further updates
            this.input.off('pointerdown'); // Remove the event listener to avoid multiple redirects
            return;
        }

        if (!(this.controllingSprite!.playerInfo!.active)) {
            this.spectatingText.visible = true
            // console.log('You are killed')
            return;
        }

        // anims
        let curAnim: AnimDir;
        let spriteID = this.controllingSprite!.playerInfo!.getData(playerSpriteIdKey) as string;

        if (this.cursors.left.isDown) {
            this.movePlayer(spriteID, this.controllingSprite!.movementSpeed, 0);
            curAnim = `left`
            this.setSpriteAnim(spriteID, curAnim)
        } else if (this.cursors.right.isDown) {
            this.movePlayer(spriteID, -this.controllingSprite!.movementSpeed, 0);

            curAnim = `right`
            this.setSpriteAnim(spriteID, curAnim)
        } else if (this.cursors.up.isDown) {
            this.movePlayer(spriteID, 0, this.controllingSprite!.movementSpeed);
            curAnim = `up`
            this.setSpriteAnim(spriteID, curAnim)
        } else if (this.cursors.down.isDown) {
            this.movePlayer(spriteID, 0, -this.controllingSprite!.movementSpeed);
            curAnim = `down`
            this.setSpriteAnim(spriteID, curAnim)
        } else {
            this.movePlayer(spriteID, 0, 0);
            curAnim = this.controllingSprite!.defaultAnim
            this.setSpriteAnim(spriteID, curAnim, false)
        }

        this.controllingSprite!.curAnimDir = curAnim;
        // move username text
        this.setUserNameTextPos(spriteID)
        sendPosMessage(
            this.controllingSprite?.playerInfo!.x!,
            this.controllingSprite?.playerInfo!.y!,
            this.controllingSprite?.curAnimDir!
        )
    }

    setSpriteAnim(spriteId: string, anim: string, loop = true) {
        const finalAnim = this.allSprites[spriteId]!.animBase + anim
        this.allSprites[spriteId]!.playerInfo!.anims.play(finalAnim, loop)
    }

    movePlayer(spriteID: string, x: number, y: number) {
        if (this.allSprites[spriteID]!.playerInfo!.active) {
            this.allSprites[spriteID]!.playerInfo!.setVelocityY(y);
            this.allSprites[spriteID]!.playerInfo!.setVelocityX(x);
        }
    }

    setUserNameTextPos(spriteID: string) {
        this.allSprites[spriteID]!.userNameText!.setPosition(
            this.allSprites[spriteID]!.playerInfo!.x - userNameOffSetX,
            this.allSprites[spriteID]!.playerInfo!.y - userNameOffSetY
        );
    }

    private loadAssets() {
        const assetsPath = "/gassets"
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

    private createMap() {
        const map = this.make.tilemap({key: 'map'})
        const blueTile = map.addTilesetImage('secondTile', 'secondTile')!
        const redTile = map.addTilesetImage('forthTile', 'forthTile')!
        const pellets = map.addTilesetImage('centrepoint', 'centrepoint')!
        const powerUp = map.addTilesetImage('power-up', 'power-up')!

        this.mapLayer = map.createLayer('map', [blueTile, redTile])!
        this.pelletLayer = map.createLayer('pellets', pellets)!
        this.powerLayer = map.createLayer('powerup', powerUp)!

        this.mapLayer.setCollisionByExclusion([-1], true, true)
        this.pelletLayer.setCollisionByExclusion([-1])
        this.powerLayer.setCollisionByExclusion([-1])
    }

    private initGhostsAnim() {
        // remember to maintain this order right, up, down, left (order in sprite sheet)
        const animDirections: AnimDir[] = ["right", "up", "down", "left"];
        const ghostColors: GhostAnimBase[] = ["ghostred", "ghostblue", "ghostpink"];

        let i = 0;
        for (let color of ghostColors) {
            for (let dir of animDirections) {
                this.anims.create({
                    frames: this.anims.generateFrameNumbers('ghosts', {start: i, end: i}),
                    key: color + dir,
                })
                i++
            }
        }
    }

    private initPacmanAnim() {
        let namesArray: AnimDir[] = [
            'right',
            'up',
            'down',
            'left',
        ]
        let ind = 0
        for (let i = 0; i < 12; i += 3) {
            this.anims.create({
                frames: this.anims.generateFrameNumbers('pacman', {start: i + 1, end: i + 3}),
                key: 'pacman' + namesArray[ind],
                frameRate: 10,
                repeat: -1
            });
            ind++
        }

        let defaultAnim: AnimDir = "default"
        // neutral pacman
        this.anims.create({
            frames: this.anims.generateFrameNumbers('pacman', {start: 0, end: 0}),
            key: 'pacman' + defaultAnim,
            frameRate: 1,
            repeat: -1
        });
    }

    private loadPlayers() {
        for (const [spriteId, spriteData] of Object.entries(this.allSprites)) {
            const defaultAnim = spriteData.animBase + spriteData.defaultAnim
            const startPos = spriteData.startPos

            let sprite = this.physics.add.sprite(
                startPos[0],
                startPos[1],
                spriteId
            );

            sprite.anims.play(defaultAnim);
            sprite.setOrigin(0.5);
            sprite.setCollideWorldBounds(true);
            sprite.setScale(.7)


            // hit box
            let desiredRadius = 20; // 50/2 - 5
            if (spriteId === "pacman") {
                sprite.body.setCircle(desiredRadius);
                const offsetX = (sprite.width / 2) - desiredRadius;
                const offsetY = (sprite.height / 2) - desiredRadius;
                sprite.body.setCircle(desiredRadius, offsetX, offsetY);
            } else {
                // ghosts have a sightly offset and smaller hitbox because of sprite size
                const offset = 8;
                sprite.body.setCircle(desiredRadius - 2, offset, offset);
            }

            spriteData.userNameText = this.addUsernameText(sprite);

            // collisions
            this.physics.add.collider(sprite, this.mapLayer)

            if (spriteId == "pacman") {
                this.physics.add.overlap(sprite, this.pelletLayer, this.pelletCallBack)
                this.physics.add.collider(sprite, this.powerLayer, this.powerUpCallBack)
            } else {
                this.physics.add.collider(
                    sprite,
                    this.allSprites["pacman"].playerInfo!,
                    this.pacmanGhostCollision
                )
            }

            sprite.setData(playerSpriteIdKey, spriteId)
            // attach player info
            spriteData.playerInfo = sprite

            this.allSprites[spriteId] = spriteData
        }
    }

    pelletCallBack(
        _pacman: CollisionBodyType,
        pellet: CollisionBodyType
    ) {
        let pel = pellet as Tile
        sendPelletMessage(pel.x, pel.y)
    }

    powerUpCallBack(
        _pacman: CollisionBodyType,
        power: CollisionBodyType
    ) {
        let pow = power as Tile
        console.log(pow.x, pow.y)

        sendPowerUpMessage(pow.x, pow.y)
    }

    pacmanGhostCollision(
        ghost: CollisionBodyType,
        _pacman: CollisionBodyType,
    ): void {
        const gho = ghost as GameObjectWithBody
        const ghostId = gho.getData(playerSpriteIdKey) as string
        console.log('pacmanGhostCollision', ghostId)

        sendPacmanGhostMessage(ghostId)
    }

    addUsernameText(sprite: Phaser.Physics.Arcade.Sprite) {
        return this.add.text(
            sprite.x - userNameOffSetX,
            sprite.y - userNameOffSetY,
            '',
            {
                fontFamily: 'Arial',
                fontSize: 12,
                color: '#ffffff'
            });
    }
}

const userNameOffSetX = 15
const userNameOffSetY = 35

