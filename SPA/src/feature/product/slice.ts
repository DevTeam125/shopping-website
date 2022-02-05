/* eslint-disable no-param-reassign */
import { createDraftSafeSelector, createSlice, PayloadAction } from '@reduxjs/toolkit';
import type { Product, ProductFull } from './types';

export interface PartialState {
  products: Product[];
  currentProductId?: string;
  searchText?: string;
  range?: {
    skip: number;
    take: number;
  }
}

const initialState: PartialState = {
  products: [],
  currentProductId: '1',
  range: { take: 50, skip: 0 },
};

function castProductFullToProduct({
  name, id, price, rating, status, photos,
}: ProductFull): Product {
  return {
    id, name, price, rating, status, photo: photos[0],
  };
}

export const productSlice = createSlice({
  name: 'counter',
  initialState,
  reducers: {
    initProducts: (state, action: PayloadAction<Product[]>) => {
      state.products = action.payload;
    },
    addProduct: (state, action: PayloadAction<ProductFull>) => {
      state.products.push(castProductFullToProduct(action.payload));
    },
    editProduct: (state, action: PayloadAction<ProductFull>) => {
      const idx = state.products.findIndex((i) => i.id === action.payload.id);
      state.products[idx] = castProductFullToProduct(action.payload);
    },
    removeProduct: (state, action: PayloadAction<string>) => {
      state.products = state.products.filter((i) => i.id !== action.payload);
    },
    loadProducts: (state, action: PayloadAction<Product[]>) => {
      state.products = action.payload;
      delete state.currentProductId;
    },
    loadCurrent: (state, action: PayloadAction<string | undefined>) => {
      state.currentProductId = action.payload;
    },
    loadSearchedItems: (state, action: PayloadAction<string>) => {
      state.searchText = action.payload;
    },
  },
});

export const {
  initProducts,
  addProduct,
  editProduct,
  removeProduct,
  loadCurrent,
  loadSearchedItems,
  loadProducts,
} = productSlice.actions;

const selectSelf = (state: { product: PartialState }) => state;

export const selectProducts = createDraftSafeSelector(
  selectSelf,
  (state) => state.product.products.slice(state.product.range?.skip, state.product.range?.take),
);

export const selectCurrentProduct = createDraftSafeSelector(
  selectSelf,
  (state) => state.product.products.find((i) => i.id === state.product.currentProductId),
);

export const selectCurrentProductFull = createDraftSafeSelector(
  selectSelf,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  (state) => ({
    id: '1', // state.product.currentProductId,
    name: '',
    description: '',
    rating: 5,
    price: 0,
    status: 'available',
    photos: [],
    features: [],
  } as ProductFull),
);

export default productSlice.reducer;
