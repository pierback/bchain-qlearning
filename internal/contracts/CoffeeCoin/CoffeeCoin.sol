pragma solidity >=0.4.22<0.6.0;

contract CoffeeCoin {
  struct Account {
    uint256 balance;
    bool initialCredit;
  }

  address private chairAddress;

  uint256 private coffeePrice;
  uint256 private matePrice;
  uint256 private waterPrice;

  mapping(address => Account) public accounts;
  mapping(address => mapping(address => uint256)) public allowed;

  constructor(address _chairAddress, uint256 cp, uint256 mp, uint256 wp)
    public
    payable
  {
    chairAddress = _chairAddress;
    setWaterPrice(cp);
    setCoffeePrice(mp);
    setWaterPrice(wp);
  }

  function payCoffee() public returns (bool success) {
    return transferFrom(msg.sender, chairAddress, coffeePrice);
  }

  function payWater() public returns (bool success) {
    return transferFrom(msg.sender, chairAddress, waterPrice);
  }

  function payMate() public returns (bool success) {
    return transferFrom(msg.sender, chairAddress, matePrice);
  }

  function transfer(address to, uint256 tokens) public returns (bool success) {
    if (!accounts[msg.sender].initialCredit) {
      setInitialCredit();
      approve(msg.sender, 20);
    }

    require(accounts[msg.sender].balance >= tokens);

    accounts[msg.sender].balance -= tokens;
    accounts[to].balance += tokens;

    return true;
  }

  function approve(address spender, uint256 tokens)
    public
    returns (bool success)
  {
    allowed[msg.sender][spender] = tokens;
    return true;
  }

  function transferFrom(address from, address to, uint tokens)
    public
    returns (bool success)
  {
    if (!accounts[msg.sender].initialCredit) {
      setInitialCredit();
      approve(msg.sender, 20);
    }

    require(tokens <= accounts[from].balance);
    require(tokens <= allowed[from][msg.sender]);

    accounts[from].balance -= tokens;
    accounts[to].balance += tokens;

    allowed[from][msg.sender] -= tokens;

    return true;
  }

  function setInitialCredit() private {
    accounts[msg.sender].balance = 1000000000000000;
    accounts[msg.sender].initialCredit = true;
  }

  function balanceOf(address tokenOwner) public view returns (uint balance) {
    return accounts[tokenOwner].balance;
  }

  function getOwnBalance() public view returns (uint balance) {
    return accounts[msg.sender].balance;
  }

  function getChairBalance() public view returns (uint balance) {
    return balanceOf(chairAddress);
  }

  function setWaterPrice(uint256 price) public {
    waterPrice = price;
  }

  function getWaterPrice() public view returns (uint256) {
    return waterPrice;
  }

  function setMatePrice(uint256 price) public {
    matePrice = price;
  }

  function getMatePrice() public view returns (uint256) {
    return matePrice;
  }

  function setCoffeePrice(uint256 price) public {
    coffeePrice = price;
  }

  function getCoffeePrice() public view returns (uint256) {
    return coffeePrice;
  }
}
