import React, { useEffect, useState } from 'react';
import { Grid, Typography } from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import {
  addProduct, editProduct, loadCurrent, selectCurrentProductFull,
} from '../slice';
import type { ProductFull } from '../types';
import ProductForm from './Form';

function ProductPost() {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const { productId } = useParams();
  const product = useSelector(selectCurrentProductFull);
  const [formState, setFormState] = useState<ProductFull>();

  useEffect(() => {
    if (productId) {
      dispatch(loadCurrent(productId));
    }
  }, [productId]);

  useEffect(() => {
    setFormState(product);
  }, [productId, product]);

  const submitData = (data: ProductFull) => {
    if (data.id) {
      dispatch(editProduct(data));
    } else {
      // eslint-disable-next-line no-param-reassign
      data.id = String(Math.random());
      dispatch(addProduct(data));
    }
    navigate('../');
  };
  return (
    <Grid container columns={{ xs: 4, md: 12 }}>
      <Grid item xs={2}>
        {productId
          && (
            <Typography>
              edit product :
              {productId}
            </Typography>
          )}
        {
          !productId
          && <Typography>add a new product</Typography>
        }
      </Grid>
      <Grid item xs={4}>
        <ProductForm product={formState} onSubmit={submitData} />
      </Grid>
    </Grid>
  );
}

export default ProductPost;
