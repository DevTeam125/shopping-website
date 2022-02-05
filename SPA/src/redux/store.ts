import { configureStore } from '@reduxjs/toolkit';
import productReducer from '../feature/product/slice';

export default configureStore({
  reducer: {
    product: productReducer,
  },
});
