import { legacy_createStore as createStore } from "redux";

// ===== action type file =====
type CounterActionType = 'INCREMENT' | 'DECREMENT';
// ===== action type file =====

// ===== action creator file =====
interface ICounterAction {
  type: CounterActionType;
  payload: number;
}

const incrementCounter = (payload: number): ICounterAction => {
  return { type: 'INCREMENT', payload };
};

const decrementCounter = (payload: number): ICounterAction => {
  return { type: 'DECREMENT', payload };
}
// ===== action creator file =====

// ===== reducer file =====
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
// ===== reducer file =====

const store = createStore(counterReducer);

store.subscribe(() => console.log(store.getState()));

store.dispatch(incrementCounter(3));
store.dispatch(decrementCounter(3));
store.dispatch(incrementCounter(3));