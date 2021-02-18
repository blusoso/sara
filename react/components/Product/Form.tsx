import React from 'react';
import { useForm } from 'react-hook-form';

interface Product {
  id: string;
  sku: string;
  title: string;
  brand: string;
  condition: string;
  meta_description: string;
  short_description: string;
  long_description: string;
  colors?: string;
  sizes?: string;
  price: number;
  qty: number;
  warranty_id: number;
  shop: string;
  is_active: boolean;
  publish_date: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

interface Props {
  submitProduct: (product: Product) => void;
  defaultProduct?: Product
}

const ProductForm: React.FC<Props> = ({ submitProduct, defaultProduct }) => {
  const { register, handleSubmit, errors } = useForm({
    defaultValues: {
        sku: defaultProduct?.sku,
        title: defaultProduct?.title,
        brand: defaultProduct?.brand,
        condition: defaultProduct?.condition,
        meta_description: defaultProduct?.meta_description,
        short_description: defaultProduct?.short_description,
        long_description: defaultProduct?.long_description,
        colors: defaultProduct?.colors,
        sizes: defaultProduct?.sizes,
        price: defaultProduct?.price,
        qty: defaultProduct?.qty,
        warranty_id: defaultProduct?.warranty_id,
        shop: defaultProduct?.shop,
        is_active: defaultProduct?.is_active,
        publish_date: defaultProduct?.publish_date,
    },
  });

  const onSubmit = (product: Product) => submitProduct(product);

  return (
    <React.Fragment>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="form-group">
          <label>SKU</label>
          <input
            type="text"
            name="sku"
            ref={register({ required: 'Please enter a SKU' })}
          />
          {errors.sku?.message}
        </div>

        <div className="form-group">
          <label>Title</label>
          <input
            type="text"
            name="title"
            ref={register({ required: 'Please enter a title' })}
          />
          {errors.title?.message}
        </div>

        <div className="form-group">
          <label>Brand</label>
          <input
            type="text"
            name="brand"
            ref={register({ required: 'Please enter a brand' })}
          />
          {errors.brand?.message}
        </div>

        <div className="form-group">
          <label>Condition</label>
          <input
            type="text"
            name="condition"
            ref={register({ required: 'Please enter a condition' })}
          />
          {errors.condition?.message}
        </div>

        <div className="form-group">
          <label>Meta description</label>
          <input
            type="text"
            name="meta_description"
            ref={register({ required: 'Please enter a meta description' })}
          />
          {errors.meta_description?.message}
        </div>
        <div className="form-group">
          <label>Description</label>
          <input
            type="text"
            name="description"
            ref={register({ required: 'Please enter a description' })}
          />
          {errors.description?.message}
        </div>
        <div className="form-group">
          <label>Colors</label>
          <input type="text" name="colors" ref={register} />
          {errors.colors?.message}
        </div>
        <div className="form-group">
          <label>Sizes</label>
          <input type="text" name="sizes" ref={register} />
          {errors.sizes?.message}
        </div>
        <div className="form-group">
          <label>Price</label>
          <input
            type="number"
            name="price"
            step="any"
            ref={register({
              required: 'Please enter a price',
              valueAsNumber: true,
            })}
          />
          {errors.price?.message}
        </div>
        <div className="form-group">
          <label>Quantity</label>
          <input
            type="number"
            name="qty"
            ref={register({
              required: 'Please enter a qty',
              valueAsNumber: true,
            })}
          />
          {errors.qty?.message}
        </div>
        <div className="form-group">
          <label>Warranty Id</label>
          <input
            type="number"
            name="warranty_id"
            ref={register({
              required: 'Please enter a warranty_id',
              valueAsNumber: true,
            })}
          />
          {errors.warranty_id?.message}
        </div>

        <div className="form-group">
          <label>Shop Id</label>
          <input
            type="text"
            name="shop"
            ref={register({
              required: 'Please enter a shop',
            })}
          />
          {errors.shop?.message}
        </div>

        <div className="form-group">
          <label>Publish Date</label>
          <input type="text" name="publish_date" ref={register} />
          {errors.publish_date?.message}
        </div>

        <div className="form-group">
          <label>
            <input
              type="checkbox"
              name="status"
              className="form-control"
              ref={register}
            />
            Active
          </label>

          <div>
            <button type="submit">Submit</button>
          </div>
        </div>
      </form>
    </React.Fragment>
  );
};

export default ProductForm;
