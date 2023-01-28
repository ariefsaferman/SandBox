/* eslint-disable no-console */
/* eslint-disable prefer-template */
/* eslint-disable indent */
/* eslint-disable no-unused-vars */
fetch('https://api.github.com/users/ariefsaferman')
  .then((response) => {
    if (response.ok) {
      const data = response.json();
      return data;
    }
     throw new Error('Status code error :' + response.status);
  }).then((data) => {
    console.log('success');
    console.log('name: ' + data.name);
    console.log('image_url: ' + data.avatar_url);
    console.log('public_repos: ' + data.public_repos);
    // console.log(data);
  }).catch((error) => {
    console.log('error');
    console.error(error);
  });
