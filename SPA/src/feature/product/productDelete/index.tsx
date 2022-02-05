import React from 'react';
import { Button, Grid, Typography } from '@mui/material';
import { Link } from 'react-router-dom';

function Product() {
  return (
    <Grid container columns={{ xs: 4, md: 12 }}>
      <Grid item xs={2}>
        <Typography>delete product</Typography>
        <Typography>are you sure? do you want do delete the product?</Typography>
        <Button color="warning" variant="contained">delete</Button>
        <Button component={Link} to="../" variant="outlined">cancel</Button>
      </Grid>
    </Grid>
  );
}

export default Product;
