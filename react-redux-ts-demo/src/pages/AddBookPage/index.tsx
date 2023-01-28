import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IBook } from '../../interfaces/api';

const AddBookPage: React.FC = () => {
  const navigate = useNavigate();
  const [book, setBook] = useState<IBook>({
    title: '',
    author: '',
    image_url: '',
    synopsis: '',
    price: 0,
  });

  const onChangeHandler = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    setBook({ ...book, [event.target.name]: event.target.value });
  };

  const onSubmitHandler = (event: React.FormEvent) => {
    event.preventDefault();
    fetch('http://localhost:8080/books', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(book),
    })
      .then((response) => {
        if (!response.ok) throw new Error('cannot post book');
        return response.json();
      })
      .then(() => {
        navigate('/books');
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <>
      <h1>Add Book</h1>
      <form onSubmit={onSubmitHandler}>
        <label htmlFor="title">Title</label>
        <br />
        <input
          id="title"
          name="title"
          type="text"
          value={book.title}
          onChange={onChangeHandler}
        />
        <br />
        <label htmlFor="author">Author</label>
        <br />
        <input
          id="author"
          name="author"
          type="text"
          value={book.author}
          onChange={onChangeHandler}
        />
        <br />
        <label htmlFor="imageUrl">Image URL</label>
        <br />
        <input
          id="image_url"
          name="image_url"
          type="text"
          value={book.image_url}
          onChange={onChangeHandler}
        />
        <br />
        <label htmlFor="synopsis">Synopsis</label>
        <br />
        <textarea
          id="synopsis"
          name="synopsis"
          rows={4}
          cols={50}
          value={book.synopsis}
          onChange={onChangeHandler}
        />
        <br />
        <label htmlFor="price">Price</label> <br />
        <input
          id="price"
          name="price"
          type="number"
          value={book.price}
          onChange={onChangeHandler}
        />{' '}
        <br />
        <br />
        <input type="submit" value="Add" />
      </form>
    </>
  );
};

export default AddBookPage;
