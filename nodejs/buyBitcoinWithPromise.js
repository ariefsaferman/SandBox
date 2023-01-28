const getUserId = (userId = '123') => {
  console.log('Step 1: Checking userId in localstorage');
  return userId;
};
const checkBalance = (userId) => {
  console.log('Step 2: Check balance from userId');
  if (userId === '123') {
    return (userId, 10000);
  }
  return 0;
};

const checkBitcoinPrice = () => {
  console.log('Step 3: Check current bitcoin price');
  return 20;
};

const buyBitcoin = (qty, price, balance) => {
  console.log('Step 4: Buy bitcoin based on total price');
  const totalPrice = qty * price;
  if (balance < totalPrice) {
    return `Error: not enough balance to buy ${qty} bitcoin with total price of ${totalPrice}! \nYour balance is ${balance}`;
  }
  return `Success buy ${qty} bitcoin with total price of ${totalPrice}`;
};

const myPromise = new Promise((resolve, reject) => {
  console.log('Start promise');
  resolve('foo');
});

myPromise
  .then(() => getUserId())
  .then((userId) => checkBalance(userId))
  .then((balance) => checkBitcoinPrice())
  .then((price) => buyBitcoin(1000, price))
  .then((message) => console.log(message))
  .catch((err) => { console.log(err); });