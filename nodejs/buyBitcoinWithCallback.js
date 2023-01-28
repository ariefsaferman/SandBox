/* eslint-disable no-console */
// Get user id
// Check balance from user id
// Check current bitcoin price
// Calculate total price

const getUserId = (userId, next) => {
  console.log('1, Checking user id in localStorage');
  return next(userId);
};

const checkBalance = (userId, next) => {
  console.log('2. Check balance from userId');
  if (userId === '123') {
    return next(userId, 1000);
  }
  return next(userId, 0);
};

const checkBitcoinPrice = (userId, balance, next) => {
  console.log('3. Check current bitcoin price');
  return next(userId, balance, 20);
};

const buyBitCoin = (userId, balance, price, qty, next) => {
  console.log('4. Buy citcoin price based on the totla price');
  const totalPrice = qty * price;
  if (balance < totalPrice) {
    return next(`Error: not enogu balance to boy ${qty} bitcoin with the tptal price of ${totalPrice}\n Your balance is ${balance}`);
  }

  return next(`Success buy ${qty} bitcoin with the total price of ${totalPrice}`);
};

//callback hell 
const buyBitcoinFromStartToFinish = (qty) => {
  getUserId('123', (userId) => {
    checkBalance(userId, (userId, balance) => {
      checkBitcoinPrice(userId, balance, (userId, balance, price) => {
        buyBitCoin(userId, balance, price, qty, (msg) => {
          console.log(msg);
        });
      });
    });
  });
};

buyBitcoinFromStartToFinish(10);
