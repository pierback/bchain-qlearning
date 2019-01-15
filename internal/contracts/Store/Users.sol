pragma solidity >=0.4.22<0.6.0;
pragma experimental ABIEncoderV2;
import "./Qlearning.sol";

contract Users {
  struct User {
    string firstName;
    string lastName;
  }

  mapping(address => Qlearning) private users;

  constructor(
    string memory _st,
    bytes32 _lr,
    bytes32 _gm,
    bytes32 _ep,
    bytes32 _epd,
    uint _sr,
    uint _pd
  ) public {
    Qlearning ql = new Qlearning(_st, _lr, _gm, _ep, _epd, _sr, _pd);
    users[msg.sender] = ql;
  }

  function getQtable() public returns (string[] memory, bytes32[2][] memory) {
    Qlearning ql = users[msg.sender];
    return ql.getQtable();
  }

  function setQValue(string memory _st, bytes32[2] memory qvalArr) public {
    Qlearning ql = users[msg.sender];
    ql.setQValue(_st, qvalArr);
  }

  function addState(string memory st, bytes32[2] memory qvalArr) public {
    Qlearning ql = users[msg.sender];
    ql.addState(st, qvalArr);
  }
}
