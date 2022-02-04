/* eslint-disable react/require-default-props */
import React, { useEffect } from 'react';
import { Container, Grid, Typography } from '@mui/material';
import { useDispatch, useSelector } from 'react-redux';
import ProductComponent from './Item';
import { loadProducts, selectProducts } from '../slice';
import { Product } from '../types';

interface Props {
  take?: number;
  skip?: number;
}

function ProductsList({ take = 10, skip = 0 }: Props) {
  const Products = useSelector(selectProducts);
  const dispatch = useDispatch();
  useEffect(() => {
    fetch(`/api/products?take=${take}&skip=${skip}`)
      .then((res) => res.json())
      .then((data: Product[]) => {
        dispatch(loadProducts(data));
      });
  }, [dispatch, take, skip]);
  return (
    <Container maxWidth="sm">
      <Grid container columns={{ xs: 4, md: 12 }}>
        <Grid item xs={4} md={12}>
          <Typography>list of Products</Typography>
        </Grid>
        {Products.map((p) => <ProductComponent product={p} />)}
      </Grid>
    </Container>
  );
}

export default ProductsList;
