import { Diagnostics } from './inputs';

let GammaRate: string = "", EpsilonRate: string = "";

interface CountItem {
    [key: string]: number
}
let Counts: Array<CountItem>;
Counts = Array.from({ length: 12 }, () => {
    return { "0": 0, "1": 0 };
});

Diagnostics.forEach((value) => {
    for (let i = 0; i < value.length; i++) {
        Counts[i][value[i]]++;
    }
});

Counts.forEach((val) => {
    GammaRate += (val["0"] >= val["1"] ? "0" : "1");
    EpsilonRate += (val["0"] < val["1"] ? "0" : "1");
})

console.log(GammaRate);
console.log(EpsilonRate);
console.log(parseInt(GammaRate, 2) * parseInt(EpsilonRate, 2));