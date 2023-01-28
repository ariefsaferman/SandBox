import { legacy_createStore as createStore } from "redux";

type CounterActionType = "INCREMENT" | "DECREMENT";

interface ICounterAction {
  type: CounterActionType;
  payload: number;
}

const incrementCounter = (payload: number): ICounterAction => ({
  type: "INCREMENT",
  payload,
});

const decrementCounter = (payload: number): ICounterAction => ({
  type: "DECREMENT",
  payload,
});

interface ICounterState {
  value: number;
}

function counterReducer(
  state: ICounterState = { value: 0 },
  action: ICounterAction
) {
  switch (action.type) {
    case "INCREMENT":
      return { value: state.value + action.payload };
    case "DECREMENT":
      return { value: state.value - action.payload };
    default:
      return state;
  }
}

const store = createStore(counterReducer);

store.subscribe(() => console.log(store.getState()));

store.dispatch(incrementCounter(3));
store.dispatch(decrementCounter(6));
store.dispatch(incrementCounter(3));
