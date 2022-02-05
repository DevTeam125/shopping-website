import * as yup from 'yup';
import type { ProductFull } from '../types';

export const EMPTY_ITEM: ProductFull = {
  id: '',
  name: '',
  description: '',
  status: 'available',
  price: 0,
  rating: 5,
  // TODO : add quantity
  // quantity: 0,
  photos: [],
  features: [],
};

export const priorities = ['available', 'pending', 'sold'] as ProductFull['status'][];

export const validationSchema = yup.object({
  name: yup
    .string()
    .trim()
    .min(1)
    .required('Email is required'),
  description: yup
    .string()
    .min(8, 'description should be of minimum 8 characters length'),
  tags: yup
    .array()
    .of(yup.string())
    .nullable(),
  status: yup
    .mixed()
    .oneOf(priorities)
    .defined(),
});
