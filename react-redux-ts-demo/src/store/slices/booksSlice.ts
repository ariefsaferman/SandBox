import { AnyAction } from 'redux';
import { ThunkDispatch } from 'redux-thunk';
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { IBook } from '../../interfaces/api';

export interface IBooksState {
  books: IBook[];
  booksLoading: boolean;
  booksError: string | null;
}

const initialState: IBooksState = {
  books: [],
  booksLoading: false,
  booksError: null,
};

export const fetchBooks = createAsyncThunk<
  IBook[],
  void,
  { rejectValue: string }
>('FETCH_BOOKS', (_: any, { rejectWithValue }: any) => {
  return fetch('http://localhost:8080/books')
    .then((response) => {
      if (!response.ok) throw new Error('failed to fetch books');
      return response.json();
    })
    .then((data) => {
      return data;
    })
    .catch((error) => {
      return rejectWithValue(error.message);
    });
});

export const deleteBook = createAsyncThunk<IBook, number>(
  'book/delete',
  (bookId, { rejectWithValue }) => {
    return fetch(`http://localhost:8080/books/${bookId}`, {
      method: 'DELETE',
    })
      .then((response) => {
        if (!response.ok) throw new Error('failed to delete book');
        return response.json();
      })
      .then((data) => {
        return data;
      })
      .catch((error) => {
        return rejectWithValue(error.message);
      });
  },
);

export const booksSlice = createSlice({
  name: 'books',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchBooks.fulfilled, (state, action) => {
      return { ...state, books: action.payload, booksLoading: false };
    });
    builder.addCase(fetchBooks.pending, (state) => {
      return { ...state, booksError: null, booksLoading: true };
    });
    builder.addCase(fetchBooks.rejected, (state, action) => {
      return action.payload
        ? { ...state, booksError: action.payload, booksLoading: false }
        : { ...state, booksError: 'unknown error', booksLoading: false };
    });
  },
});

export default booksSlice.reducer;
export type BooksDispatch = ThunkDispatch<IBooksState, any, AnyAction>;
