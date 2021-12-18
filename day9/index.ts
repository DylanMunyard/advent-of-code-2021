import { Heights } from './input';


interface Point { r: number, c: number, height: number };
function Adjacent(rows: number, row: Array<number>, r: number, c: number) : Array<Point> {
    let result = new Array<Point>();

    if (c > 0) {
        result.push({ r: r, c: c - 1, height: Heights[r][c-1] });
    }
    if (c + 1 < row.length) {
        result.push({ r: r, c: c + 1, height: Heights[r][c+1]});
    }

    if (r > 0) {
        result.push({ r: r - 1, c: c, height: Heights[r-1][c]});   
    }
    if (r + 1 < rows) {
        result.push({ r: r + 1, c: c, height: Heights[r+1][c]});   
    }

    return result;
}


let rows = Heights.length;
let risks = new Array<Point>();
Heights.forEach((heights, r) => {
    heights.forEach((height, c) => {
        let lowest = true;
        Adjacent(rows, heights, r , c).forEach(h => {
            if (!lowest) return;
            
            if (Heights[h.r][h.c] <= height) {
                lowest = false;
            } 
        });
        if (lowest) risks.push({ r: r, c: c, height: height });
    });
});

console.log(risks.reduce((p, c) => p + c.height + 1, 0));

// #part2

function basin(p: Point) : Array<Point> {
    let result = new Array<Point>();

    let adjacents = Adjacent(rows, Heights[p.r], p.r , p.c);
    adjacents.forEach(adjp => {
        if (Heights[adjp.r][adjp.c] < 9 && Heights[adjp.r][adjp.c] > p.height) {
            result.push(adjp);
            result = result.concat(basin(adjp));
        }
    });

    return result;
}

let basin_sizes = new Array<number>();
risks.forEach(r => {
    let basins = basin(r);
    let basins_filtered : { [key: string]: number } = { [`${r.r}${r.c}`]: r.height };
    basins.forEach(b => {
        let key = `${b.r}${b.c}`;
        if (basins_filtered[key]) return;

        basins_filtered[key] = b.height;
    });

    basin_sizes.push(Object.keys(basins_filtered).length);
});

basin_sizes.sort((a,b) => b - a).forEach(a => console.log(a));
console.log(basin_sizes.sort((a,b) => b - a).slice(0, 3));
console.log(basin_sizes.sort((a,b) => b - a).slice(0, 3).reduce((p,c) => p * c));