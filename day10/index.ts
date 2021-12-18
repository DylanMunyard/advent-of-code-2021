import { Navigation } from "./input";

let Open : { [key: string]: string } = {
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">"
};
let Close : { [key: string]: string } = {
    ")": "(",
    "]": "[",
    "}": "{",
    ">": "<"
};

let Points : { [key: string]: number } = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137   
}

let invalids : Array<string> = [];
let incomplete_lines : Array<Array<string>> = [];
Navigation.forEach((line, idx) => {
    if (!Open[line[0]]) {
        console.log(`${line[0]} is not a way to open line`);
        return;
    }
    let opens : Array<string> = [];
    let valid = true;
    line.split("").forEach((c,i) => {
        if (!valid) return;

        if (Open[c]) {
            opens.push(c);
            return;
        }

        let last_open = opens.pop() || "";
        if (Close[c] !== last_open) {
            console.warn(`Expected ${Open[last_open]}, but found ${c}`);
            invalids.push(c);
            valid = false;
        }
    });

    if (valid) incomplete_lines.push(opens); 
});

let total = 0;
invalids.forEach(i => {
    total += Points[i];
})
console.log(total);

// #part2
let AcPoints : { [key: string]: number } = {
    ")": 1,
    "]": 2,
    "}": 3,
    ">": 4
};

let autocompletes : Array<Array<string>> = [];
incomplete_lines.forEach(line => {
    let closes : Array<string> = [];
    for (let i = line.length - 1; i >= 0; i--) {
        closes.push(Open[line[i]])
    }
    autocompletes.push(closes);
});

console.log(autocompletes);

let scores : Array<number> = [];
autocompletes.forEach(autocomplete => {
    total = 0;
    autocomplete.forEach(c => {
        console.log(`${total} * 5 + ${AcPoints[c]}`);
        total = total * 5 + AcPoints[c];
    });
    scores.push(total);
});

let middle = Math.floor(scores.length / 2);
console.log(scores.sort((a, b) => b - a)[middle]);