import { applyMiddleware } from 'redux';
// import logger from './logger';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
export default applyMiddleware(logger, thunk);
