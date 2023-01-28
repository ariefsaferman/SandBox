"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const redux_1 = require("redux");
function counterReducer(state = { value: 0 }, action) {
    switch (action.type) {
        case "INCREMENT":
            return { value: state.value + 1 };
        case "DECREMENT":
            return { value: state.value - 1 };
        default:
            return state;
    }
}
const store = (0, redux_1.legacy_createStore)(counterReducer);
store.subscribe(() => console.log(store.getState()));
store.dispatch({ type: "INCREMENT" });
