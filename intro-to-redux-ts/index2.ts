import { legacy_createStore as createStore } from "redux";

interface ICounterAction {
  type: 'INCREMENT' | 'DECREMENT';
  payload: number;
}

interface ICounterState {
  value: number;
}

function counterReducer(state: ICounterState = { value: 0 }, action: ICounterAction) {
  switch (action.type) {
    case 'INCREMENT':
      return { value: state.value + action.payload };
    case 'DECREMENT':
      return { value: state.value - action.payload };
    default:
      return state;
  }
}

const store = createStore(counterReducer);

store.subscribe(() => console.log(store.getState()));

store.dispatch({ type: 'INCREMENT', payload: 3 });
store.dispatch({ type: 'DECREMENT', payload: 3 });
store.dispatch({ type: 'INCREMENT', payload: 3 });
