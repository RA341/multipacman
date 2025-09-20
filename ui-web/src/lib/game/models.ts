import Phaser from 'phaser';
import Text = Phaser.GameObjects.Text; // Make sure to import Phaser

export type AnimDir = "up" | "down" | "left" | "right" | "default";
export type GhostAnimBase = "ghostred" | "ghostblue" | "ghostpink";
export type PacmanAnimBase = "pacman";

export type GameMessage = "pos" | "pel" | "pow" | "kill"

export type PlayerInfo = {
    playerInfo: Phaser.Physics.Arcade.Sprite | null,
    curAnimDir: AnimDir,
    defaultAnim: AnimDir,
    startPos: [number, number],
    userNameText?: Text | null,
    movementSpeed: number,
}

export function getAnimStr(info: Ghost | Pacman) {
    return info.animBase + info.curAnimDir
}

export type Ghost = PlayerInfo & {
    animBase: GhostAnimBase,
}

export type Pacman = PlayerInfo & {
    animBase: PacmanAnimBase,
}


export type PrevGameState = {}