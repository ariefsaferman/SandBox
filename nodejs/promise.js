/* eslint-disable no-console */
const succeed = false;

// Produce
const myPromise = new Promise((resolve, reject) => {
  setTimeout(() => {
    if (succeed) {
      resolve('Succeed');
    } else {
      reject(new Error('Failed'));
    }
  }, 1000);
});

myPromise.then((res) => {
  console.log(res);
  return 'Response from the first .then() call';
}).then((res) => {
  console.log(res);
}).catch((err) => {
  console.log(err);
}).finally(() => {
  console.log('Finally');
});
