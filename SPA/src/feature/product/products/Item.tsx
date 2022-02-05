import React from 'react';
import {
  Avatar, Box, Button, ButtonGroup, Grid, Paper, Typography,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { Product } from '../types';

function ItemComponent({ product }: { product: Product }) {
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
            <Typography variant="h4" component={Link} to={`${product.id}`}>{product.name}</Typography>
            {/* <Typography noWrap>{product.description}</Typography> */}
          </Grid>
          <Grid item>
            <ButtonGroup variant="outlined" aria-label="outlined secondary button group" size="small">
              <Button component={Link} to={`/product/edit/${product.id}`}>🖊</Button>
              <Button component={Link} to={`/product/delete/${product.id}`}>🗑</Button>
            </ButtonGroup>
            <div>
              {/* {product.tags.join(', ')} */}
              {product.status}
            </div>
          </Grid>
        </Grid>
      </Paper>
    </Box>

  );
}

export default ItemComponent;
