interface Data {
  Response?: Response; 
}

interface Response {
  postId: number; 
  id: number; 
  name: string;
  email: string;
  body: string
}

fetch('https://jsonplaceholder.typicode.com/posts', {
  method: 'GET'
})
  .then((response) => {
    if (!response.ok) throw new Error('cannot fetch');
    return response.json();
  })
  .then((data: Data) => {
    console.log(data);
  })
  .catch((error) => {
    console.log(error);
  });

fetch('https://jsonplaceholder.typicode.com/comments', {
  method: 'GET'
})
  .then((response) => {
    if (!response.ok) throw new Error('cannot fetch');
    return response.json();
  })
  .then((data) => {
    console.log(data);
  })
  .catch((error) => {
    console.log(error);
  });
