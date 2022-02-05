import React from 'react';
import {
  Button, Container, Grid, Typography,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { loadCurrent } from 'feature/product/slice';
import ProductPost from 'feature/product/productPost';
import ProductsList from 'feature/product/products';
import { Modal, useModal } from './modal';

function Home() {
  const dispatch = useDispatch();
  const { isOpen, handleOpen, handleClose } = useModal();
  const handleCreateProduct = () => {
    dispatch(loadCurrent());
    handleOpen();
  };
  return (
    <Container maxWidth="sm">
      <Grid container columns={{ xs: 4, md: 12 }}>
        <Grid item xs={2}>
          <Button onClick={handleCreateProduct} variant="outlined">add new product</Button>
          <Button component={Link} to="/product" variant="outlined">products list</Button>
        </Grid>
        <Grid item xs={2}>
          <Typography variant="h5">Latest added of products:</Typography>
          <ProductsList take={3} />
        </Grid>
      </Grid>
      <Modal open={isOpen} onClose={handleClose}>
        <ProductPost />
      </Modal>
    </Container>
  );
}

export default Home;
