import { Diagnostics } from './inputs';

let MostBits: string = "";

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

for (let i = 0; i < 12; i++) {
    let CountOxygen = CommonBitCounter(OxygenDiagnostics, i);
    let CountScrubber = CountOxygen;

    if (i > 0) {
        // optimisation: the first iteration is the same array
        CountScrubber = CommonBitCounter(ScrubberDiagnostics, i);
    }

    console.log(i);
    console.log(CountOxygen);
    console.log(CountScrubber);

    if (OxygenDiagnostics.length > 1) OxygenDiagnostics = OxygenDiagnostics.filter((val) => { return val[i] === (CountOxygen["0"] > CountOxygen["1"] ? "0" : "1") });
    if (ScrubberDiagnostics.length > 1) ScrubberDiagnostics = ScrubberDiagnostics.filter((val) => { return val[i] === (CountScrubber["0"] <= CountScrubber["1"] ? "0" : "1") });
}

console.log(OxygenDiagnostics);
console.log(ScrubberDiagnostics);
console.log(parseInt(OxygenDiagnostics.toString(), 2) * parseInt(ScrubberDiagnostics.toString(), 2));