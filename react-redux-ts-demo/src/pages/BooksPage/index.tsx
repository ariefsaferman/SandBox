import React, { useEffect } from 'react';
import Book from '../../components/Book';
import Loading from '../../components/Loading';
import Error from '../../components/Error';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../store';
import {
  BooksDispatch,
  deleteBook,
  fetchBooks,
} from '../../store/slices/booksSlice';

const BooksPage: React.FC = () => {
  const { books, booksError, booksLoading } = useSelector(
    (state: RootState) => state.books,
  );
  const dispatch: BooksDispatch = useDispatch();

  // const deleteBook = (bookId: number) => {
  //   return new Promise<void>((resolve, reject) => {
  //     fetch(`http://localhost:8080/books/${bookId}`, {
  //       method: 'DELETE',
  //     })
  //       .then((response) => {
  //         if (!response.ok) throw 'cannot delete book';
  //         return response.json();
  //       })
  //       .then(() => {
  //         resolve();
  //       })
  //       .catch((error) => {
  //         reject(error);
  //       });
  //   });
  // };

  useEffect(() => {
    dispatch(fetchBooks());
  }, []);

  const handleOnClickDelete = (bookId: number): void => {
    dispatch(deleteBook(bookId))
      .unwrap()
      .then((data) => {
        console.log('DELETE PROCESS SUCCEED:', data);
        dispatch(fetchBooks());
      })
      .catch((error) => {
        console.log('DELETE PROCESS FAILED:', error);
      });
  };

  return (
    <>
      <h1>Books</h1>
      {booksLoading && <Loading />}
      {!booksLoading && booksError && <Error />}
      {!booksLoading &&
        !booksError &&
        books.map((book, index) => {
          return (
            <Book
              key={index}
              book={book}
              onClickHandler={handleOnClickDelete}
            />
          );
        })}
    </>
  );
};

export default BooksPage;
