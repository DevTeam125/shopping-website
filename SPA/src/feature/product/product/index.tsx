import React, { useEffect } from 'react';
import { Grid, Typography } from '@mui/material';
import { useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { loadCurrent, selectCurrentProduct } from '../slice';
import ProductItem from './Item';

function Product() {
  const { productId } = useParams();
  const dispatch = useDispatch();
  useEffect(() => {
    if (productId) {
      dispatch(loadCurrent(productId));
    }
  }, [productId]);
  const product = useSelector(selectCurrentProduct);
  return (
    <Grid container columns={{ xs: 4, md: 12 }}>
      <Grid item xs={2}>
        <Typography>product view</Typography>
      </Grid>
      <Grid item xs={4}>
        {product && <ProductItem product={product} />}
        {!product && <Typography>product not found</Typography>}
      </Grid>
    </Grid>
  );
}

export default Product;
