import { SampleInput, WhatDigit, SegmentLocation } from "./sampleinput";

let UniqueSegments : { [key: number]: number } = {
    2: 1,
    4: 4,
    3: 7,
    7: 8
}

let tally = 0;
SampleInput.forEach(value => {
    let output = value.split(' | ')[1].split(/\s+/);
    tally += output.filter(d => UniqueSegments[d.length]).length;
});

console.log(tally);

// #part2

type Signal = {
    [key: string]: string
};

let digits : { [key: number]: string } = {};

tally = 0;
SampleInput.forEach(sample => {
    let signal = sample.split(" | ")[0];
    let segments : { [key: number]: string } = {};
    signal.split(/\s+/).forEach(digit => {
        digits[digit.length] = digit;
        let signals = digit.split("");
        signals.forEach(s => {
            segments[(signal.match(new RegExp(s, "g")) || []).length] = s;
        });
    });

    // Unique segment overlaps:
    // top left: 6
    // bottom left: 4
    // bottom right: 9

    let segmentmap : Signal = {};
    segmentmap[SegmentLocation.TopLeft] = segments[6];
    segmentmap[SegmentLocation.BottomLeft] = segments[4];
    segmentmap[SegmentLocation.BottomRight] = segments[9];

    // Can figure out remaining segments through elimination of digits with known segments
    segmentmap[SegmentLocation.TopRight] = digits[2].replace(segmentmap["br"], ""); // 1 has 2 segments
    segmentmap[SegmentLocation.Top] = digits[3].replace(segmentmap["br"], "").replace(segmentmap[SegmentLocation.TopRight], ""); // 7 has 3 segments
    segmentmap[SegmentLocation.Middle] = digits[4].replace(segmentmap["br"], "").replace(segmentmap[SegmentLocation.TopRight], "").replace(segmentmap["tl"], ""); // 4 has 4 segments
    segmentmap[SegmentLocation.Bottom] =  digits[7].replace(segmentmap["tl"], "").replace(segmentmap[SegmentLocation.Top], "").replace(segmentmap[SegmentLocation.TopRight], "").replace(segmentmap[SegmentLocation.Middle], "").replace(segmentmap[SegmentLocation.BottomLeft], "").replace(segmentmap[SegmentLocation.BottomRight], ""); // 8 has 7 segments

    //console.log(segmentmap);
    //console.log(digits);

    let output = sample.split(" | ")[1].split(/\s+/);
    let display = "";
    output.forEach(number => {
        display += WhatDigit(number, segmentmap);
    });

    tally += parseInt(display, 10);
});

console.log(tally);
