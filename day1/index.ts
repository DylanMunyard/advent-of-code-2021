import { Input } from './input';

let measurement = 0;
Input.forEach((_, idx) => {
    if (idx === 0) return;

    if (Input[idx-1] < Input[idx]) { measurement++; }
});

console.log(`The answer to part 1 is ${measurement}`);

let SlidingInputs = Input.map((value, idx, array) => {
   if (idx < array.length - 2) {
       return value + array[idx+1] + array[idx+2];
   }
   return value;
}).slice(0, Input.length - 2);

let slidingMeasurement = 0;
SlidingInputs.forEach((value, idx) => {
    if (idx === 0) return;

    if (SlidingInputs[idx-1] < SlidingInputs[idx]) { slidingMeasurement++; }
});

console.log(`The answer to part 2 is ${slidingMeasurement}`);