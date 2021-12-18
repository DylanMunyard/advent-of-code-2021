import { Max, StraightLines } from './inputs';

let straightLines = StraightLines();
let max = Max(straightLines);

let Grid : Array<Array<number>> = new Array<Array<number>>();
for (let x = 0; x <= max.x; x++) {
    Grid[x] = new Array<number>();
    for (let y = 0; y <= max.y; y++) {
        Grid[x][y] = 0;
    }
}

straightLines.forEach(line => {
    if (line.x0 === line.x1) {
        let y1 = Math.max(line.y1, line.y0),
            y0 = Math.min(line.y1, line.y0);

        for (let y = y0; y <= y1; y++) {
            Grid[line.x0][y]++;
        }
    } else if (line.y0 === line.y1) {
        let x1 = Math.max(line.x1, line.x0),
            x0 = Math.min(line.x1, line.x0);

        for (let x = x0; x <= x1; x++) {
            Grid[x][line.y0]++;
        }
    } else {
        let xcrement = 1;
        if (line.x0 > line.x1) {
            xcrement = -1;
        }

        let ycrement = 1;
        if (line.y0 > line.y1) {
            ycrement = -1;
        }

        let x = line.x0, y = line.y0, i = 0, length = Math.abs(line.x1 - line.x0);
        while (i <= length) {   
            Grid[x][y]++;
            x += xcrement;
            y += ycrement;
            i++;
        }
    }
});

let dangerzone = 0;
Grid.forEach((ys, x) => {
    ys.forEach((lines, y) => {
        if (lines > 1) {
            dangerzone++;
        }
    })
})

console.log(dangerzone);