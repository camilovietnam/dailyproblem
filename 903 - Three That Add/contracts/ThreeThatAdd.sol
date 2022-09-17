// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.16;

contract ThreeThatAdd {
    uint [] public items;

    function addItem(uint item) public {
        items.push(item);
    }

    function getAllItems() public view returns(uint[] memory) {
        return items;
    }

    function resetItems() public {
        delete items ;
    }

    function threeThatAdd(uint sum)public view returns (string memory) {
        for (uint i = 0; i < items.length; i++) {
            for (uint j = i+1; j < items.length; j++) {
                for (uint k = j+1; k < items.length; k++) {
                    if (items[i]+items[j]+items[k] == sum) {
                        string memory first = uint2str(items[i]);
                        string memory second = uint2str(items[j]);
                        string memory third = uint2str(items[k]);

                        return string.concat("found 3 elements: ", first, ", ", second, ", ", third);
                    }
                }
            }
        }
        return "no items found";
    }

    function uint2str(uint256 _i)internal pure returns (string memory str){
        if (_i == 0) {
            return "0";
        }
        uint256 j = _i;
        uint256 length;
        while (j != 0){
            length++;
            j /= 10;
        }
        bytes memory bstr = new bytes(length);
        uint256 k = length;
        j = _i;
        while (j != 0){
            bstr[--k] = bytes1(uint8(48 + j % 10));
            j /= 10;
        }
        str = string(bstr);
    }
}
