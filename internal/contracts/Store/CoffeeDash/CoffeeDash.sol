pragma solidity >=0.4.22<0.6.0;

contract CoffeeDash {
  struct Qlearning {
    string state;
    uint256 learningRate;
    uint256 gamma;
    uint256 epsilon;
    uint256 epsilonDecay;
    uint successratio;
    uint prediction;
    mapping(string => qvals) qtable;
  }

  struct qvals {
    string data;
    bool exists;
  }

  mapping(address => Qlearning) private users;
  string[280] public states;
  uint8 public stateCnt;

  constructor(
    string memory _st,
    uint256 _lr,
    uint256 _gm,
    uint256 _ep,
    uint256 _epd,
    uint _sr,
    uint _pd
  ) public payable {
    users[msg.sender] = Qlearning(_st, _lr, _gm, _ep, _epd, _sr, _pd);
    addState(_st, "0, 0");
  }

  function getQtableState(uint8 i)
    public
    view
    returns (string memory, string memory)
  {
    return (states[i], users[msg.sender].qtable[states[i]].data);
  }

  function setQValue(string memory _st, string memory qvalArr) public {
    if (!users[msg.sender].qtable[_st].exists) revert();
    users[msg.sender].qtable[_st].data = qvalArr;
  }

  function addState(string memory _st, string memory qvalArr) public {
    if (users[msg.sender].qtable[_st].exists) revert();
    users[msg.sender].qtable[_st].exists = true;
    users[msg.sender].qtable[_st].data = qvalArr;
    states[stateCnt] = _st;
    stateCnt++;
  }
}
