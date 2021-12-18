import { tCountoFush, CountoFush, LanternFish } from "./inputs";

let oCountOfFush = {...CountoFush};

console.log(`Initial state: ${PrintCountoFush(CountoFush)}`);
for (let i = 1; i <= 256; i++) {
    let cof : tCountoFush = {
        "8": oCountOfFush["0"],
        "7": oCountOfFush["8"],
        "6": oCountOfFush["7"] + oCountOfFush["0"],
        "5": oCountOfFush["6"],
        "4": oCountOfFush["5"],
        "3": oCountOfFush["4"],
        "2": oCountOfFush["3"],
        "1": oCountOfFush["2"],
        "0": oCountOfFush["1"]
    };
    oCountOfFush = cof;
}

console.log(`Total fush: ${TallyFush(oCountOfFush)}`);

function TallyFush(count: tCountoFush) : number {
    return count["0"] + count["1"] + count["2"] + count["3"] + count["4"] + count["5"] + count["6"] + count["7"] + count["8"];
}

function PrintCountoFush(count: tCountoFush) {
    return `${count["0"]},${count["1"]},${count["2"]},${count["3"]},${count["4"]},${count["5"]},${count["6"]},${count["7"]},${count["8"]}`;
}

