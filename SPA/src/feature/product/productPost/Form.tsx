import React, { useEffect } from 'react';
import {
  Button, Typography, Box, TextField, Radio, RadioGroup, FormControl, FormControlLabel, Grid,
} from '@mui/material';
import { FormikProvider, useFormik } from 'formik';
import { Link } from 'react-router-dom';
import type { ProductFull } from '../types';
// import InputTags from './tags';
import { validationSchema, EMPTY_ITEM, priorities } from './Form.validations';

type Props = { product?: ProductFull, onSubmit: (data: ProductFull) => void };

function Form({ product, onSubmit }: Props) {
  const formik = useFormik<ProductFull>({
    initialValues: EMPTY_ITEM,
    validationSchema,
    onSubmit,
  });

  useEffect(() => {
    if (product) {
      Object.keys(product).forEach((i) => {
        const key = i as keyof ProductFull;
        formik.setFieldValue(i, product[key]);
      });
    }
  }, [product]);

  // const setInitialTags = useCallback((callback) => {
  //   callback(product?.tags);
  // }, [product]);

  return (
    <FormikProvider value={formik}>
      <Box
        component="form"
        onSubmit={formik.handleSubmit}
        style={{ maxWidth: '500px' }}
        noValidate
        autoComplete="off"
      >
        <Grid
          container
          alignItems="center"
          direction="column"
          spacing={2}
        >
          <Grid item>
            <FormControl>
              <TextField
                fullWidth
                name="name"
                label="Name"
                value={formik.values.name}
                onChange={formik.handleChange}
                error={formik.touched.name && Boolean(formik.errors.name)}
                helperText={formik.touched.name && formik.errors.name}
              />
            </FormControl>
          </Grid>
          <Grid item>
            <FormControl>
              <TextField
                fullWidth
                name="description"
                label="Description"
                value={formik.values.description}
                onChange={formik.handleChange}
                error={formik.touched.description && Boolean(formik.errors.description)}
                helperText={formik.touched.description && formik.errors.description}
              />
            </FormControl>
          </Grid>
          {/* <Grid item>
            <FormControl>
              <FieldArray
                name="tags"
                validateOnChange={false}
                render={(arrayHelpers) => (
                  <InputTags
                    component={(
                      <TextField
                        fullWidth
                        error={formik.touched.tags && Boolean(formik.errors.tags)}
                        helperText={formik.touched.tags && formik.errors.tags}
                      />
                    )}
                    initiate={setInitialTags}
                    insert={arrayHelpers.insert}
                    remove={arrayHelpers.remove}
                  />
                )}
              />
            </FormControl>
          </Grid> */}
          <Grid item>
            <FormControl>
              {Boolean(formik.touched.status)
                && (
                  <Typography color="red">
                    {formik.errors.status}
                  </Typography>
                )}
              <RadioGroup aria-label="anonymous" name="anonymous" value={false} row>
                {priorities.map((i) => (
                  <FormControlLabel
                    checked={formik.values.status === i}
                    control={<Radio />}
                    value={`${i}`}
                    name="status"
                    label={i}
                    onChange={formik.handleChange}
                    key={i}
                  />
                ))}
              </RadioGroup>
              {Boolean(formik.touched.status)
                && (
                  <Typography color="red">
                    {formik.errors.status}
                  </Typography>
                )}
            </FormControl>
          </Grid>
          <Grid item>
            <Button color="primary" variant="contained" fullWidth type="submit">
              {product?.id ? 'edit' : 'add'}
            </Button>
            <Button component={Link} to="/" variant="outlined">cancel</Button>
          </Grid>
        </Grid>
      </Box>
    </FormikProvider>
  );
}

Form.defaultProps = { product: undefined };

export default Form;
