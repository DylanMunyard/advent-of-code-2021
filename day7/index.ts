import { Inputs } from "./inputs";

let cheapest = 9_999_999;
let horiz = 0;
for (let i = 0; i < Inputs.length; i++) {
    let h = Inputs[i];

    let cost = Inputs.reduce((p, c) => p + Math.abs(h - c));
    if (cost - h < cheapest) {
        cheapest = cost - h;
        horiz = h;
    }
}

console.log(`It's cheapest to travel to ${horiz}:${cheapest} fuel`);

let min = Inputs[0], max = Inputs[Inputs.length - 1];
console.log(`min=${min} max=${max}`);
let pos = Array.from({ length: max - min }, (_, i) => min + i);

cheapest = 9_999_999_999;
horiz = 0;
for (let i = 0; i < pos.length; i++) {
    let h = pos[i];
    
    let cost = Inputs.reduce((p, c) => p + (Math.abs(h - c) * (Math.abs(h - c) + 1)) / 2, 0); // n(n+1)/2 is additive factorial
    if (cost < cheapest) {
        cheapest = cost;
        horiz = h;
    }
}

console.log(`It's cheapest to travel to ${horiz}:${cheapest} fuel`);

