import { configureStore } from '@reduxjs/toolkit';
import logger from 'redux-logger';
import thunk from 'redux-thunk';
import counterReducer from './slices/counterSlice';
import booksReducer from './slices/booksSlice';

export const store = configureStore({
  reducer: {
    counter: counterReducer,
    books: booksReducer,
  },
  middleware: [logger, thunk],
});

export type RootState = ReturnType<typeof store.getState>;
