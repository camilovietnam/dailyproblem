// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Parenthesis {
    // accepts a string, returns the count of wrong parentheses
    function countWrong(string memory text) public pure returns (int) {
        bytes memory b = bytes(text);

        bytes1 open = '(';
        bytes1 close = ')';

        int accum = 0;
        int badCount = 0;

        for (uint i=0; i<b.length;i++) {
            // if open a parenthesis
            if (b[i] == open) {
                accum++;
                continue;
            }
            // if close a parenthesis with accum
            if (b[i] == close && accum > 0) {
                accum--;
                continue;
            }

            // if close a parenthesis without accum (wrong)
            if (b[i] == close && accum == 0) {
                badCount++;
                continue;
            }
        }

        // return the wrongs I found plus whatever opened parenthesises I have
        return badCount + accum;
    }
}