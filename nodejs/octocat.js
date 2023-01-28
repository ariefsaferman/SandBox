fetch('https://api.github.com/users/octoca92929t')
  .then((response) => {
    if (response.ok) {
      const data = response.json();
      console.log(data);
      return data;
    // eslint-disable-next-line prefer-template
    } throw new Error('Status code error :' + response.status);
  }).then((data) => {
    console.log('success');
    console.log(data);
  }).catch((error) => {
    console.log('error');
    console.error(error);
  });
