import React from 'react';
import {
  Grid, Typography, Avatar, Box, Paper,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { Product } from '../types';

function ProductItem({ product }: { product: Product }) {
  return (
    <Box sx={{ flexGrow: 1, overflow: 'hidden', px: 3 }} key={product.id}>
      <Paper sx={{
        maxWidth: 400, my: 1, mx: 'auto', p: 2,
      }}
      >
        <Grid container wrap="nowrap" spacing={2}>
          <Grid item>
            <Avatar>{product.name}</Avatar>
          </Grid>
          <Grid item xs zeroMinWidth>
            <Typography variant="h4">{product.name}</Typography>
          </Grid>
        </Grid>
        <Grid item xs zeroMinWidth>
          {/* <Typography>{product.description}</Typography> */}
        </Grid>
        <Grid item>
          {/* {product.tags.map((tag) => (<Typography variant="caption">{tag}</Typography>))} */}
        </Grid>
        <Grid item>
          {['edit', 'delete'].map((i) => (
            <Typography variant="caption" component={Link} to={`../${i}/${product.id}`}>{i}</Typography>
          ))}
        </Grid>
      </Paper>
    </Box>
  );
}

export default ProductItem;
