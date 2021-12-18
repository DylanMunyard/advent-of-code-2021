import { BingoNumbers } from './inputs';
import { SumUnmarked, Winner, Hits, Hit, Inputs, Bingo, Cell } from './rawto';

let grids = Inputs();

grids.then((Grids: Array<Array<Cell>>) => {
    let bingo : Bingo = { Bingo: false, Grid: new Array<Cell>() };
    let last_number: number = 0;
    BingoNumbers.forEach((number, idx) => {
        if (bingo.Bingo) return;

        bingo = Hits(Grids, number);
        if (bingo.Bingo) {
            console.log(`BINGO! BINGOOO! Grid=${bingo.GridNumber} Row=${bingo.Row} Col=${bingo.Col}`);
            last_number = number; 
        }
    });

    Winner(bingo);
    console.log(`Last number: ${last_number}`);
    console.log(`Sum unmarked: ${SumUnmarked(bingo)}`);
    console.log(`${SumUnmarked(bingo) * last_number}`);

    let bingod = function(bingos: Array<number>, gridnumber: number) {
        let found = false;
        bingos.forEach(idx => {
            if (found) return;
            found = idx === gridnumber;
        });

        return found;
    }

    bingo = { Bingo :false, Grid: new Array<Cell>() };
    let bingos = new Array<number>();
    last_number = 0;
    BingoNumbers.forEach((number, idx) => {
        Grids.forEach((grid, gridIdx) => {
            if (bingod(bingos, gridIdx)) return;

            bingo = Hit(grid, number, gridIdx);
            if (bingo.Bingo) {
                bingos.push(gridIdx);
                last_number = number;
            }
        });
    });


    Winner(bingo);
    console.log(`Last number: ${last_number}`);
    console.log(`Sum unmarked: ${SumUnmarked(bingo)}`);
    console.log(`${SumUnmarked(bingo) * last_number}`);
});