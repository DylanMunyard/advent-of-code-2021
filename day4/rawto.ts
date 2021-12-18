import { createReadStream } from "fs";
import { createInterface } from "readline";

export type Cell = {
    Value: number;
    hit: boolean;
}

export type Bingo = {
    Bingo: boolean;
    Grid: Array<Cell>;
    GridNumber?: number;
    Row?: number;
    Col?: number;
}

export function SumUnmarked(bingo: Bingo) : number {
    let total = 0;
    bingo.Grid.forEach(cell => {
        if (!cell.hit) { total += cell.Value; }
    })

    return total;
}

export function Winner(bingo: Bingo) {
    if (bingo.Row !== undefined) {
        let cell = bingo.Row * 5;
        console.log(`${bingo.Grid[cell].Value},${bingo.Grid[cell+1].Value},${bingo.Grid[cell+2].Value},${bingo.Grid[cell+3].Value},${bingo.Grid[cell+4].Value}`)
    }

    if (bingo.Col !== undefined) {
        let cell = bingo.Col;
        for (let i = 0; i < 5; i++) {
            console.log(`${bingo.Grid[cell + (i * 5)].Value}`)
        }
    }
}

export function Inputs() : Promise<Array<Array<Cell>>> { 
    let promise = new Promise<Array<Array<Cell>>>((resolve, reject) => {
        let raw = createInterface({ input: createReadStream('raw.txt'), output: process.stdout, terminal: false });
        let Grids : Array<Array<Cell>> = [new Array<Cell>()];
        let GridNumber = 0;
        let line_number = 1;

        raw.on('line', (input: string) => {
            let row = input.trim().split(/\s+/);
            if (row.length < 4) return;

            row.forEach((value) => { Grids[GridNumber].push({ Value: parseInt(value.trim(), 10), hit: false })});

            if (line_number % 5 === 0) {
                Grids.push(new Array<Cell>());
                GridNumber++;
            }
            line_number++;
        });

        raw.on('close', () => {
            resolve(Grids.slice(0, Grids.length - 1));
        });
    });

    return promise;
}

function IsBingo(grid: Array<Cell>) : Bingo {
    for (let row = 0; row < 5; row++) {
        let cell = row * 5;
        if (grid[cell].hit && grid[cell+1].hit && grid[cell+2].hit && grid[cell+3].hit && grid[cell+4].hit) {
            return { Bingo: true, Row: row, Grid: grid };
        }
    }

    for (let column = 0; column < 5; column++) {
        if (grid[column].hit && grid[column + 5].hit && grid[column + 10].hit && grid[column + 15].hit && grid[column + 20].hit) {
            return { Bingo: true, Col: column, Grid: grid };
        }
    }

    return { Bingo: false, Grid: grid };
}

export function Hits(grids: Array<Array<Cell>>, input: number) : Bingo {
    let bingo : Bingo = { Bingo: false, Grid: new Array<Cell>() };
    grids.forEach((grid, idx) => {
        if (bingo.Bingo) return; 
        bingo = Hit(grid, input, idx);
    });

    return bingo
}

export function Hit(grid: Array<Cell>, input: number, gridNumber: number) : Bingo {
    let bingo : Bingo = { Bingo: false, Grid: new Array<Cell>() };
    grid.forEach(cell => {
        if (cell.Value === input) { cell.hit = true; }
    })

    let isbingo = IsBingo(grid);
    if (isbingo.Bingo) { bingo = isbingo; bingo.GridNumber = gridNumber; bingo.Grid = grid; }

    return bingo;
}
