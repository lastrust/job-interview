// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

contract FizzBuzz {
    uint256 counter;

    constructor() {
        counter = 0;
    }

    function increment() public {
        counter++;
        emit Increment(msg.sender, counter);
    }

    function getFizzBuzz() public view returns (string memory) {
        if (counter == 0) return "0";
        if (counter % 5 == 0 && counter % 3 == 0) return "FizzBuzz";
        if (counter % 3 == 0) return "Fizz";
        if (counter % 5 == 0) return "Buzz";
        return uint2str(counter);
    }

    function uint2str(uint256 _i)
        internal
        pure
        returns (string memory _uintAsString)
    {
        if (_i == 0) {
            return "0";
        }
        uint256 j = _i;
        uint256 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint256 k = len - 1;
        while (_i != 0) {
            bstr[k--] = bytes1(uint8(48 + (_i % 10)));
            _i /= 10;
        }
        return string(bstr);
    }

    event Increment(address from, uint counter);
}
