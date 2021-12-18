import { Diagnostics } from './inputs';

interface CountItem {
    [key: string]: number
}
let OxygenDiagnostics = Diagnostics, ScrubberDiagnostics = Diagnostics;
function CommonBitCounter(arr: string[], position: number) : CountItem {
    let count : CountItem = { "0": 0, "1": 0 };
    arr.forEach((value) => {
        count[value.charAt(position)]++;
        // Optimisation, only looking for most common bit, not a total of each bit
        if (count[value.charAt(position)] * 2 > arr.length) return;
        if (count[value.charAt(position)] * 2 > arr.length) return; 
    });

    return count;
}

function reduceArray(arr: string[], position: number, fn: (count: CountItem, value: string, position: number) => boolean) : Array<string> {
    if (arr.length === 1) return arr;
    let count = CommonBitCounter(arr, position);

    return reduceArray(arr.filter((val) => { return fn(count, val, position) }), position + 1, fn);
}

OxygenDiagnostics = reduceArray(Diagnostics, 0, (count: CountItem, value: string, position: number) => {
    return value[position] === (count["0"] > count["1"] ? "0" : "1") }
);
ScrubberDiagnostics = reduceArray(Diagnostics, 0, (count: CountItem, value: string, position: number) => {
    return value[position] === (count["0"] <= count["1"] ? "0" : "1") }
);

console.log(OxygenDiagnostics);
console.log(ScrubberDiagnostics);
console.log(parseInt(OxygenDiagnostics.toString(), 2) * parseInt(ScrubberDiagnostics.toString(), 2));