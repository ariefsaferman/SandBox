import { legacy_createStore as createStore } from "redux";

interface ICounterAction {
  type: 'INCREMENT' | 'DECREMENT';
}

interface ICounterState {
  value: number;
}

function counterReducer(state: ICounterState = { value: 0 }, action: ICounterAction) {
  switch (action.type) {
    case 'INCREMENT':
      return { value: state.value + 1 };
    case 'DECREMENT':
      return { value: state.value - 1 };
    default:
      return state;
  }
}

const store = createStore(counterReducer);

store.subscribe(() => console.log(store.getState()));

store.dispatch({ type: 'INCREMENT' });
store.dispatch({ type: 'DECREMENT' });
store.dispatch({ type: 'INCREMENT' });
