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
    uint8 stateCnt;
    string[] states;
    mapping(string => qvals) qtable;
  }

  struct qvals {
    string data;
    bool exists;
  }

  mapping(address => Qlearning) private users;

  constructor(
    string memory _st,
    uint256 _lr,
    uint256 _gm,
    uint256 _ep,
    uint256 _epd,
    uint _sr,
    uint _pd
  ) public payable {
    string[] memory localArr;
    users[msg.sender] = Qlearning(
      _st,
      _lr,
      _gm,
      _ep,
      _epd,
      _sr,
      _pd,
      0,
      localArr
    );
    addState(_st, "0, 0");
  }

  function getQtableState(uint8 i)
    public
    view
    returns (string memory, string memory)
  {
    string memory _state = users[msg.sender].states[i];
    return (users[msg.sender].states[i], users[msg.sender].qtable[_state].data);
  }

  function getState(uint8 i) public view returns (string memory) {
    return (users[msg.sender].states[i]);
  }

  function getStateValues(uint8 i) public view returns (string memory) {
    string memory _state = users[msg.sender].states[i];
    return (users[msg.sender].qtable[_state].data);
  }

  function setQValue(string memory _st, string memory qvalArr) public {
    if (!users[msg.sender].qtable[_st].exists) revert();
    users[msg.sender].qtable[_st].data = qvalArr;
  }

  function addState(string memory _st, string memory qvalArr) public {
    if (users[msg.sender].qtable[_st].exists) revert();
    //uint8 stcnt = users[msg.sender].stateCnt;
    users[msg.sender].qtable[_st].exists = true;
    users[msg.sender].qtable[_st].data = qvalArr;
    // users[msg.sender].states[stcnt] = _st;
    users[msg.sender].states.push(_st);
    users[msg.sender].stateCnt++;
  }

  function getStateCnt() public view returns (uint8) {
    return (users[msg.sender].stateCnt);
  }

  function stateExists(string memory _st) public view returns (bool) {
    return (users[msg.sender].qtable[_st].exists);
  }
}
