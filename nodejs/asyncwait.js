/* eslint-disable no-console */
// const promise = new Promise((resolve, reject) => {
//   setTimeout(() => resolve('I love you'), 3000);
// });

// async function call() {
//   const result = await promise;
//   console.log(result);
// }

// call();

// with arrow function
const fetchApi = async (url) => {
  try {
    const response = await fetch(url);
    const data = await response.json();
    console.log('success');
    console.log(data);
  } catch (error) {
    console.log('error');
    console.error(error);
  }
};

fetchApi('https://api.github.com/users/octocats');
