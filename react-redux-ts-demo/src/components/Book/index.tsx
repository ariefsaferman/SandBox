import React from 'react';
import { IBook } from '../../interfaces/api';

type BookProps = {
  book: IBook;
  onClickHandler: (bookId: number) => void;
};

const Book: React.FC<BookProps> = (props: BookProps) => {
  const { book, onClickHandler } = props;
  return (
    <div>
      <h2>{book.title}</h2>
      <img src={book.image_url} alt={book.title} height="300" />
      <p>ID: {book.id}</p>
      <p>Author: {book.author}</p>
      <p>Synopsis: {book.synopsis}</p>
      <button onClick={() => onClickHandler(book.id!)}>Delete</button>
    </div>
  );
};

export default Book;
