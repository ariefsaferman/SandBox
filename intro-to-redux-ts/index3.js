"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const redux_1 = require("redux");
const incrementCounter = (payload) => {
    return { type: 'INCREMENT', payload };
};
const decrementCounter = (payload) => {
    return { type: 'DECREMENT', payload };
};
function counterReducer(state = { value: 0 }, action) {
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
const store = (0, redux_1.legacy_createStore)(counterReducer);
store.subscribe(() => console.log(store.getState()));
store.dispatch(incrementCounter(3));
store.dispatch(decrementCounter(3));
store.dispatch(incrementCounter(3));
