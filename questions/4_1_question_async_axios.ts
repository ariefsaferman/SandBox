import axios from 'axios';

interface postType {
  userId: number;
  id: number;
  title: string;
  body: string
}

axios<postType>({
  url: 'https://jsonplaceholder.typicode.com/posts',
  method: 'GET'
})
  .then(({ data = [] }) => {
    console.log(data);
  })
  .catch((error) => {
    console.log(error)
  });
