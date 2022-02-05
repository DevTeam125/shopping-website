import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { initProducts } from './slice';
import type { Product, Rating } from './types';

function generateRandomBetween<T extends number>(min: number, max: number): T {
  return Math.floor(Math.random() * (max - min + 1)) + min as T;
}

function useProducts() {
  const dispatch = useDispatch();
  useEffect(() => {
    const products: Product[] = Array<number>(2).fill(1).map<Product>((x, i) => ({
      id: String(x + i),
      name: String(x + i),
      description: `hi this is description for ${i + x}`,
      price: x + i,
      rating: generateRandomBetween<Rating>(1, 5),
      status: 'available',
      photo: {
        title: 'photo title',
        url: 'https://picsum.photos/200/300',
      },
    }));
    dispatch(initProducts(products));
  }, []);
}

export default useProducts;
