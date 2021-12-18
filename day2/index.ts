import { Directions } from "./input";

let x : number = 0,y: number = 0, aim: number = 0;

interface MoveCommand {
    command: string,
    units: number
}

Directions.forEach((value, idx) => {
    let moveCommand: MoveCommand = value;
    switch (moveCommand.command) {
        case "forward": {
            x += moveCommand.units;
            y += moveCommand.units * aim;
            break;
        }
        case "up": {
            aim -= moveCommand.units;
            break;
        }
        case "down": {
            aim += moveCommand.units;
        }
    }
});


console.log(`Resting position (${x}x${y})=${x*y}`);