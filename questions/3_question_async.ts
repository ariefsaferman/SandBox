const greet = new Promise((resolve: (value: string | Promise<string>) => void, reject) => {
  setTimeout(() => {
    resolve('Hello world!');
  }, 3000);
});

greet
  .then((data) => {
    console.log(data);
  })
  .catch((error) => {
    console.log(error);
  });
