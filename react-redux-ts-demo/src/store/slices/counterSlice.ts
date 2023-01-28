import { createSlice } from '@reduxjs/toolkit';

interface ICounterState {
  value: number;
}

const initialState: ICounterState = {
  value: 0,
};

export const counterSlice = createSlice({
  name: 'counter',
  initialState,
  reducers: {
    increment: (state, action) => {
      return { ...state, value: state.value + action.payload };
    },
    decrement: (state, action) => {
      return { ...state, value: state.value - action.payload };
    },
  },
});

export const { increment, decrement } = counterSlice.actions;
export default counterSlice.reducer;
