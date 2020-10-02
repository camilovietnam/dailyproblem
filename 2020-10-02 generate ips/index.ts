/**
* Given a string of digits, generate all possible valid IP address combinations.
* IP addresses must follow the format A.B.C.D, where A, B, C, and D are numbers
* between 0 and 255. Zero-prefixed numbers, such as 01 and 065, are not allowed,
* except for 0 itself.
*
* For example, given "2542540123", you should return ['254.25.40.123', '254.254.0.123'].
*
* The code implements this solution https://www.youtube.com/watch?v=KU7Ae2513h0
* "The IP Address Decomposition Problem - Compute All Valid IP Addresses From Raw IP String"
* by youtube user "Back To Back SWE"
*
*/

function split (
    ipAddresses: string[],
    source: string,
    builderIndex: number,
    segments: string[],
    segmentNumber: number
) {
    if (segmentNumber === 4 && builderIndex === source.length) {
        ipAddresses.push(`${segments[0]}.${segments[1]}.${segments[2]}.${segments[3]}`);
        return;
    } else if (segmentNumber === 4 || builderIndex === source.length) {
        return;
    }

    for (let length = 1; length <= 3 && (builderIndex+length) <= source.length; length++ ) {
        const candidateSegment = source.substr(builderIndex, length);
        const intValue = parseInt(candidateSegment);

        const segmentStartsWithZero = length >= 2 && candidateSegment[0] === '0';
        const numberTooBig = intValue > 255;

        if (segmentStartsWithZero || numberTooBig) {
            break;
        }

        segments[segmentNumber] = candidateSegment;
        split(ipAddresses, source, builderIndex+length, segments, segmentNumber+1);
    }
}

const myString: string = '2542540123';
const allIpAddresses: string[] = [];
const path : string[] = [];

split(allIpAddresses, myString, 0, path, 0);

console.log(`\nFor string '${myString}', here are the possible IPs:`);
console.log(allIpAddresses);