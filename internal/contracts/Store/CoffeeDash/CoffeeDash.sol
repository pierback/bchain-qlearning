pragma solidity >=0.4.22<0.6.0;
pragma experimental ABIEncoderV2;

contract CoffeeDash {
  mapping(address => Qlearning) private users;
  struct Qlearning {
    string state;
    uint256 learningRate;
    uint256 gamma;
    uint256 epsilon;
    uint256 epsilonDecay;
    uint successratio;
    uint prediction;
    mapping(string => string[2]) qtable;
  }

  //mapping(uint => string[2]) public qtable;
  //mapping(uint => string) private stateMap;
  //mapping(address => Qlearning) public users;

  constructor(
    string memory _st,
    uint256 _lr,
    uint256 _gm,
    uint256 _ep,
    uint256 _epd,
    uint _sr,
    uint _pd
  ) public payable {
    Qlearning memory ql = Qlearning(_st, _lr, _gm, _ep, _epd, _sr, _pd);
    users[msg.sender] = ql;

    addState(_st, ["0", "0"]);
  }

  string[] public states;

  function getStateCount() public view returns (uint) {
    return states.length;
  }

  function getQtableState(uint i)
    public
    view
    returns (string memory, string[2] memory)
  {
    Qlearning storage ql = users[msg.sender];
    string storage state = states[i];
    return (states[i], ql.qtable[state]);
  }

  function setQValue(string memory _st, string[2] memory qvalArr) public {
    int stateIndex = stateExists(_st, states);
    require(stateIndex > -1);
    Qlearning storage ql = users[msg.sender];
    if (ql.qtable[_st].length > 0) revert();
    ql.qtable[_st] = qvalArr;

  }

  function addState(string memory _st, string[2] memory qvalArr) public {
    Qlearning storage ql = users[msg.sender];
    if (ql.qtable[_st].length < 1) revert();
    ql.qtable[_st] = qvalArr;
    states.push(_st);
  }

  function stateExists(string memory st, string[] memory stateArr)
    private
    pure
    returns (int)
  {
    for (uint i = 0; i < stateArr.length; i++) {
      //if (equals(stateArr[i], st)) {
      if (compareStrings(stateArr[i], st)) {
        return int(i);
      }
    }
    return -1;
  }

  function compareStrings(string memory a, string memory b)
    internal
    pure
    returns (bool)
  {
    return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
  }

  /*function equals(State memory st1, State memory st2)
    internal
    pure
    returns (bool)
  {
    return (keccak256(
      abi.encodePacked(
        st1.coffeeCount,
        st1.mateCount,
        st1.waterCount,
        st1.weekday,
        st1.timeslot
      )
    ) == keccak256(
      abi.encodePacked(
        st2.coffeeCount,
        st2.mateCount,
        st2.waterCount,
        st2.weekday,
        st2.timeslot
      )
    ));
  }*/

}
